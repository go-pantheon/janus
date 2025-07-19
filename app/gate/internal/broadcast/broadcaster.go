package broadcast

import (
	"context"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-kit/profile"
	kcp "github.com/go-pantheon/fabrica-net/kcp/server"
	tcp "github.com/go-pantheon/fabrica-net/tcp/server"
	ws "github.com/go-pantheon/fabrica-net/websocket/server"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/fabrica-util/xsync"
	"github.com/go-pantheon/janus/app/gate/internal/data"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/pool"
	v1 "github.com/go-pantheon/janus/gen/api/server/broadcaster/service/push/v1"
	"github.com/google/wire"
	"google.golang.org/protobuf/proto"
)

const (
	broadcastTopic = "gate.broadcast.message"
)

var ProviderSet = wire.NewSet(NewBroadcaster)

type Broadcaster struct {
	xsync.Stoppable

	log *log.Helper
	d   *data.Data
	svr []xnet.Server

	msgChan chan *v1.PushMessage
	subs    []*Subscriber
}

func NewBroadcaster(kcp *kcp.Server, ws *ws.Server, tcp *tcp.Server, d *data.Data) (b *Broadcaster, cleanup func()) {
	b = &Broadcaster{
		Stoppable: xsync.NewStopper(time.Second * 10),
		log:       log.NewHelper(log.With(log.DefaultLogger, "module", "broadcast")),
		d:         d,
		svr:       make([]xnet.Server, 0, 3),
	}

	b.svr = append(b.svr, tcp, ws, kcp)

	cleanup = func() {
		if err := b.Stop(context.Background()); err != nil {
			b.log.Errorf("stop broadcast: %v", err)
		}
	}

	return b, cleanup
}

func (b *Broadcaster) Start() error {
	if err := b.subscribe(); err != nil {
		return err
	}

	xsync.Go("broadcast.run", func() error {
		return b.run()
	})

	return nil
}

func (b *Broadcaster) subscribe() error {
	for _, topic := range []string{
		broadcastTopic,
		profile.GRPCEndpoint(),
	} {
		sub, err := newSubscriber(b.d.Rdb, topic, func(msg *v1.PushMessage) {
			b.msgChan <- msg
		})
		if err != nil {
			return err
		}

		b.subs = append(b.subs, sub)
	}

	return nil
}

func (b *Broadcaster) run() error {
	for {
		select {
		case <-b.StopTriggered():
			return xsync.ErrStopByTrigger
		case msg := <-b.msgChan:
			if err := b.send(msg); err != nil {
				b.log.Errorf("push message: %v", err)
			}
		}
	}
}

func (b *Broadcaster) send(msg *v1.PushMessage) (err error) {
	for _, body := range msg.Bodies {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		oup := pool.GetPacket()
		defer pool.PutPacket(oup)

		oup.Mod = body.Mod
		oup.Seq = body.Seq
		oup.Obj = body.Obj
		oup.Data = body.Data
		oup.DataVersion = body.DataVersion

		bytes, marshalErr := proto.Marshal(oup)
		if marshalErr != nil {
			err = errors.JoinUnsimilar(err, errors.Wrapf(marshalErr, "packet marshal failed. %d-%d", body.Mod, body.Seq))
			continue
		}

		if msg.Broadcast {
			if err = b.broadcast(ctx, msg, bytes); err != nil {
				err = errors.JoinUnsimilar(err, errors.Wrapf(err, "broadcast failed. %d-%d", body.Mod, body.Seq))
			}
			continue
		}

		if len(msg.Uids) == 0 {
			log.Warnf("broadcast message has no uids: %d-%d", body.Mod, body.Seq)
			continue
		}

		if len(msg.Uids) == 1 {
			if err = b.push(ctx, msg.Uids[0], bytes); err != nil {
				err = errors.JoinUnsimilar(err, errors.Wrapf(err, "push failed. %d-%d", body.Mod, body.Seq))
			}
		} else {
			if err = b.multicast(ctx, msg.Uids, bytes); err != nil {
				err = errors.JoinUnsimilar(err, errors.Wrapf(err, "multicast failed. %d-%d", body.Mod, body.Seq))
			}
		}
	}

	return err
}

func (b *Broadcaster) push(ctx context.Context, uid int64, pack xnet.Pack) (err error) {
	for _, srv := range b.svr {
		if sendErr := srv.Push(ctx, uid, pack); sendErr != nil {
			if errors.Is(sendErr, xnet.ErrWorkerNotFound) {
				continue
			}

			err = errors.JoinUnsimilar(err, errors.Wrapf(sendErr, "push failed. uid=%d", uid))
		}
	}

	return err
}

func (b *Broadcaster) multicast(ctx context.Context, uids []int64, pack xnet.Pack) (err error) {
	for _, srv := range b.svr {
		if sendErr := srv.Multicast(ctx, uids, pack); sendErr != nil {
			if errors.Is(sendErr, xnet.ErrWorkerNotFound) {
				continue
			}

			err = errors.JoinUnsimilar(err, errors.Wrapf(sendErr, "multicast failed. uids=%v", uids))
		}
	}

	return err
}

func (b *Broadcaster) broadcast(ctx context.Context, msg *v1.PushMessage, pack xnet.Pack) (err error) {
	for _, srv := range b.svr {
		if sendErr := srv.Broadcast(ctx, msg.Color, msg.Sid, pack); sendErr != nil {
			if errors.Is(sendErr, xnet.ErrWorkerNotFound) {
				continue
			}

			err = errors.JoinUnsimilar(err, errors.Wrapf(sendErr, "broadcast failed. color=%s, sid=%d", msg.Color, msg.Sid))
		}
	}

	return err
}

func (b *Broadcaster) Stop(ctx context.Context) error {
	return b.TurnOff(func() error {
		return b.stop(ctx)
	})
}

func (b *Broadcaster) stop(ctx context.Context) error {
	var (
		err = errors.NewSafeJoinError()
		wg  sync.WaitGroup
	)

	for _, sub := range b.subs {
		wg.Add(1)

		xsync.Go("stop.subscriber", func() error {
			defer wg.Done()
			if stopErr := sub.Stop(ctx); stopErr != nil {
				err.Join(stopErr)
			}

			return nil
		})
	}

	wg.Wait()

	close(b.msgChan)

	if err.HasError() {
		return err
	}

	return nil
}
