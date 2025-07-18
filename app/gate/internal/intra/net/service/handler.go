package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-kit/profile"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/compress"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/pool"
	"google.golang.org/protobuf/proto"
)

func (s *Service) Handle(ctx context.Context, ss xnet.Session, th xnet.TunnelManager, in xnet.Pack) (err error) {
	p := pool.GetPacket()
	defer pool.PutPacket(p)

	if err = proto.Unmarshal(in, p); err != nil {
		return errors.Wrap(err, "packet unmarshal failed")
	}

	defer func() {
		if err != nil {
			err = errors.WithMessagef(err, "mod=%d seq=%d obj=%d", p.Mod, p.Seq, p.Obj)
		}
	}()

	if ss.IsCrypto() {
		index := ss.CSIndex()
		if index != p.Index {
			return errors.Errorf("csindex validation failed. want=%d give=%d", index, p.Index)
		}

		ss.IncreaseCSIndex()
	}

	if p.Compress {
		p.Data, err = compress.Decompress(p.Data)
		if err != nil {
			return err
		}
	}

	// the p.Obj is 0 when the mod is belong to player tunnel
	if p.Obj == 0 {
		p.Obj = ss.UID()
	}

	t, err := th.Tunnel(ctx, p.Mod, p.Obj)
	if err != nil {
		return err
	}

	return t.Forward(ctx, t.PacketToTunnelMsg(p))
}

func (s *Service) Tick(ctx context.Context, ss xnet.Session) (err error) {
	if renewErr := s.renewRouteTable(ctx, ss); renewErr != nil {
		err = errors.Join(err, renewErr)
	}

	// TODO: check black list

	return err
}

func (s *Service) renewRouteTable(ctx context.Context, ss xnet.Session) (err error) {
	nt := ss.NextRenewTime()
	if nt.IsZero() {
		ss.UpdateNextRenewTime(time.Now().Add(s.gateRT.TTL() / time.Duration(2)))
		return nil
	}

	now := time.Now()

	if now.Before(nt) {
		return nil
	}

	ss.UpdateNextRenewTime(now.Add(s.gateRT.TTL() / time.Duration(2)))

	return s.gateRT.RenewSelf(ctx, ss.Color(), ss.UID(), profile.GRPCEndpoint())
}

func (s *Service) OnConnected(ctx context.Context, ss xnet.Session) (err error) {
	log.Debugf("client connected. %s", ss.LogInfo())
	return nil
}

func (s *Service) OnDisconnect(ctx context.Context, ss xnet.Session) (err error) {
	log.Debugf("client disconnected. %s", ss.LogInfo())
	return nil
}
