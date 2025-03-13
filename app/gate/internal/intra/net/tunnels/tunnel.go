package tunnels

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	verrors "github.com/go-pantheon/vulcan-kit/errors"
	"github.com/go-pantheon/vulcan-util/compress"
	"github.com/go-pantheon/vulcan-util/sync"
	"github.com/pkg/errors"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/pkg/pool"
	clipkt "github.com/vulcan-frame/vulcan-gate/gen/api/client/packet"
	"github.com/vulcan-frame/vulcan-gate/pkg/net"
	"github.com/vulcan-frame/vulcan-gate/pkg/net/tunnel"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
)

type AppTunnel interface {
	AppTunnelBase

	CSHandle(msg tunnel.ForwardMessage) error
	SCHandle() (tunnel.ForwardMessage, error)
	TransformMessage(from *clipkt.Packet) (to tunnel.ForwardMessage)
	OnStop()
	OnGroupStop(ctx context.Context, err error)
}

type AppTunnelBase interface {
	Log() *log.Helper
	Type() int32
	UID() int64
	Color() string
	OID() int64
	Session() net.Session
}

var _ tunnel.Tunnel = (*Tunnel)(nil)

type Tunnel struct {
	sync.Stoppable
	tunnel.Pusher

	app AppTunnel

	csChan chan tunnel.ForwardMessage
}

func NewTunnel(ctx context.Context, pusher tunnel.Pusher, app AppTunnel) *Tunnel {
	t := &Tunnel{
		Stoppable: sync.NewStopper(time.Second * 10),
		app:       app,
		Pusher:    pusher,
		csChan:    make(chan tunnel.ForwardMessage, 1024),
	}

	t.start(ctx)
	return t
}

func (t *Tunnel) Type() int32 {
	return t.app.Type()
}

func (t *Tunnel) Forward(ctx context.Context, p tunnel.ForwardMessage) error {
	if t.IsStopping() {
		return verrors.ErrTunnelStopped
	}

	msg, err := t.transform(p)
	if err != nil {
		return err
	}

	t.csChan <- msg
	return nil
}

func (t *Tunnel) transform(from tunnel.ForwardMessage) (to tunnel.ForwardMessage, err error) {
	p, ok := from.(*clipkt.Packet)
	if !ok {
		return nil, errors.New("invalid packet type")
	}
	to = t.app.TransformMessage(p)
	return
}

func (t *Tunnel) Push(ctx context.Context, pack []byte) error {
	return t.Pusher.Push(ctx, pack)
}

func (t *Tunnel) start(ctx context.Context) {
	sync.GoSafe(fmt.Sprintf("gate.Tunnel-%d-%d-%d", t.app.UID(), t.app.Type(), t.app.OID()), func() error {
		return t.run(ctx)
	})
}

func (t *Tunnel) run(ctx context.Context) error {
	defer t.stop()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.StopTriggered():
			return sync.GroupStopping
		}
	})
	eg.Go(func() error {
		return sync.RunSafe(func() error {
			return t.csLoop(ctx)
		})
	})
	eg.Go(func() error {
		return sync.RunSafe(func() error {
			return t.scLoop(ctx)
		})
	})
	if err := eg.Wait(); err != nil {
		t.app.OnGroupStop(ctx, err)
	}
	return nil
}

func (t *Tunnel) csLoop(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.StopTriggered():
			return sync.GroupStopping
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
			t.app.Log().WithContext(ctx).Errorf("[gate.Tunnel] uid=%d color=%s oid=%d push failed. %+v", t.app.UID(), t.app.Color(), t.app.OID(), err)
		}
	}
}

func (t *Tunnel) push(ctx context.Context, sc tunnel.ForwardMessage) error {
	p := pool.GetPacket()
	defer pool.PutPacket(p)

	p.Mod = sc.GetMod()
	p.Seq = sc.GetSeq()
	p.Obj = sc.GetObj()

	if newData, compressed, err := compress.Compress(p.Data); err != nil {
		return err
	} else {
		p.Data = newData
		p.Compress = compressed
	}

	p.Index = int32(t.app.Session().IncreaseSCIndex())

	bytes, err := proto.Marshal(p)
	if err != nil {
		return errors.Wrapf(err, "packet marshal failed")
	}
	if err = t.Push(ctx, bytes); err != nil {
		return err
	}

	return nil
}

func (t *Tunnel) stop() {
	t.DoStop(func() {
		close(t.csChan)
		t.app.OnStop()
	})
}
