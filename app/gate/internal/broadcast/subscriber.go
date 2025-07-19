package broadcast

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-util/xsync"
	v1 "github.com/go-pantheon/janus/gen/api/server/broadcaster/service/push/v1"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

type HandleFunc func(msg *v1.PushMessage)

type Subscriber struct {
	xsync.Stoppable

	log    *log.Helper
	rdb    redis.UniversalClient
	topic  string
	handle HandleFunc
	sub    *redis.PubSub
}

func newSubscriber(rdb redis.UniversalClient, topic string, handle HandleFunc) (*Subscriber, error) {
	s := &Subscriber{
		Stoppable: xsync.NewStopper(time.Second * 10),
		log:       log.NewHelper(log.With(log.DefaultLogger, "module", "broadcaster")),
		rdb:       rdb,
		topic:     topic,
		handle:    handle,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	s.sub = s.rdb.Subscribe(ctx, topic)

	s.GoAndStop("subscribe."+topic, func() error {
		return s.channel()
	}, func() error {
		return s.Stop(context.Background())
	})

	s.log.Infof("subscribe topic: %s", topic)

	return s, nil
}

func (s *Subscriber) channel() error {
	for {
		select {
		case <-s.StopTriggered():
			return xsync.ErrIsStopped
		case carrier := <-s.sub.Channel():
			var msg v1.PushMessage
			if err := proto.Unmarshal([]byte(carrier.Payload), &msg); err != nil {
				s.log.Errorf("unmarshal message: %v", err)
				continue
			}

			s.handle(&msg)
		}
	}
}

func (s *Subscriber) Stop(ctx context.Context) error {
	return s.TurnOff(func() error {
		return s.sub.Close()
	})
}
