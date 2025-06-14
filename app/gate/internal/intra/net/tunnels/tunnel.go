package tunnels

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-kit/xerrors"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/compress"
	"github.com/go-pantheon/fabrica-util/xsync"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/pool"
	clipkt "github.com/go-pantheon/janus/gen/api/client/packet"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
)

type AppTunnel interface {
	AppTunnelBase

	CSHandle(msg xnet.ForwardMessage) error
	SCHandle() (xnet.ForwardMessage, error)
	TransformMessage(from *clipkt.Packet) (to xnet.ForwardMessage)
	OnClose(err error)
}

type AppTunnelBase interface {
	Log() *log.Helper
	Type() int32
	UID() int64
	Color() string
	OID() int64
	Session() xnet.Session
}

const (
	recordPushErrStartAt = time.Second * 10
	recordPushErrCount   = 10
)

var _ xnet.Tunnel = (*Tunnel)(nil)

type Tunnel struct {
	xsync.Closable
	xnet.Pusher

	app    AppTunnel
	csChan chan xnet.ForwardMessage

	recordPushErrStartAt time.Time // the start time of the last 10 seconds
	recordPushErrCount   int       // the count of continuous push error in the last 10 seconds
}

func NewTunnel(ctx context.Context, pusher xnet.Pusher, app AppTunnel) *Tunnel {
	t := &Tunnel{
		Closable: xsync.NewClosure(time.Second * 10),
		app:      app,
		Pusher:   pusher,
		csChan:   make(chan xnet.ForwardMessage, 1024),
	}

	t.start(ctx)

	return t
}

func (t *Tunnel) Type() int32 {
	return t.app.Type()
}

func (t *Tunnel) Forward(ctx context.Context, p xnet.ForwardMessage) error {
	if t.OnClosing() {
		return xerrors.ErrTunnelStopped
	}

	t.csChan <- p

	return nil
}

func (t *Tunnel) TransformMessage(from xnet.PacketMessage) (to xnet.ForwardMessage, err error) {
	p, ok := from.(*clipkt.Packet)
	if !ok {
		return nil, errors.New("invalid packet type")
	}

	return t.app.TransformMessage(p), nil
}

func (t *Tunnel) Push(ctx context.Context, pack []byte) error {
	return t.Pusher.Push(ctx, pack)
}

func (t *Tunnel) start(ctx context.Context) {
	xsync.GoSafe(fmt.Sprintf("gate.Tunnel-%d-%d-%d", t.app.UID(), t.app.Type(), t.app.OID()), func() error {
		return t.run(ctx)
	})
}

func (t *Tunnel) run(ctx context.Context) (err error) {
	defer t.close(err)

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.CloseTriggered():
			return xsync.ErrGroupIsClosing
		}
	})
	eg.Go(func() error {
		return xsync.RunSafe(func() error {
			return t.csLoop(ctx)
		})
	})
	eg.Go(func() error {
		return xsync.RunSafe(func() error {
			return t.scLoop(ctx)
		})
	})

	return eg.Wait()
}

func (t *Tunnel) csLoop(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.CloseTriggered():
			return xsync.ErrGroupIsClosing
		case cs := <-t.csChan:
			if err := t.app.CSHandle(cs); err != nil {
				return err
			}
		}
	}
}

func (t *Tunnel) scLoop(ctx context.Context) error {
	for {
		msg, err := t.app.SCHandle()
		if err != nil {
			return err
		}

		if err = t.push(ctx, msg); err != nil {
			if t.recordPushErr(ctx, err) {
				return err
			}
		} else {
			t.recordPushErrCount = 0 // reset the count when push success
		}
	}
}

func (t *Tunnel) recordPushErr(ctx context.Context, err error) (overload bool) {
	t.app.Log().WithContext(ctx).Errorf("push failed uid=%d sid=%d color=%s status=%d oid=%d. %+v", t.app.UID(), t.app.Session().SID(), t.app.Color(), t.app.Session().Status(), t.app.OID(), err)

	now := time.Now()

	if now.Sub(t.recordPushErrStartAt) > recordPushErrStartAt {
		t.recordPushErrStartAt = now
		t.recordPushErrCount = 0
	}

	t.recordPushErrCount++

	return t.recordPushErrCount > recordPushErrCount
}

func (t *Tunnel) push(ctx context.Context, sc xnet.ForwardMessage) (err error) {
	p := pool.GetPacket()
	defer pool.PutPacket(p)

	p.Mod = sc.GetMod()
	p.Seq = sc.GetSeq()
	p.Obj = sc.GetObj()
	p.Index = sc.GetIndex()

	p.Data, p.Compress, err = compress.Compress(sc.GetData())
	if err != nil {
		return err
	}

	bytes, err := proto.Marshal(p)
	if err != nil {
		return errors.Wrapf(err, "packet marshal failed")
	}

	if err = t.Push(ctx, bytes); err != nil {
		return err
	}

	return nil
}

func (t *Tunnel) close(err error) {
	if closeErr := t.DoClose(func() {
		close(t.csChan)
		t.app.OnClose(err)
	}); closeErr != nil {
		t.app.Log().Errorf("close tunnel failed. %+v", closeErr)
	}
}
