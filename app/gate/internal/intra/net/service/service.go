package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/janus/app/gate/internal/client/player"
	"github.com/go-pantheon/janus/app/gate/internal/client/room"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/pool"
	playerv1 "github.com/go-pantheon/janus/gen/api/server/player/intra/v1"
	roomv1 "github.com/go-pantheon/janus/gen/api/server/room/intra/v1"
	"github.com/go-pantheon/vulcan-kit/tunnel"
	xnet "github.com/go-pantheon/vulcan-net"
	rctx "github.com/go-pantheon/vulcan-net/context"
	"github.com/go-pantheon/vulcan-util/compress"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

var ProviderSet = wire.NewSet(NewTCPService)

var _ xnet.Service = (*Service)(nil)

type Service struct {
	logger    log.Logger
	encrypted bool

	playerClient playerv1.TunnelServiceClient
	playerRT     *player.RouteTable

	roomClient roomv1.TunnelServiceClient
	roomRT     *room.RouteTable
}

func NewTCPService(logger log.Logger, label *conf.Label,
	playerRT *player.RouteTable, playerClient playerv1.TunnelServiceClient,
	roomRT *room.RouteTable, roomClient roomv1.TunnelServiceClient,
) *Service {
	return &Service{
		logger:       logger,
		encrypted:    label.Encrypted,
		playerClient: playerClient,
		playerRT:     playerRT,
		roomClient:   roomClient,
		roomRT:       roomRT,
	}
}

func (s *Service) Handle(ctx context.Context, ss xnet.Session, th tunnel.Holder, in []byte) (err error) {
	if err = s.handle(ctx, ss, th, in); err != nil {
		return errors.WithMessagef(err, "uid=%d color=%s status=%d", ss.UID(), ss.Color(), ss.Status())
	}
	return nil
}

func (s *Service) handle(ctx context.Context, ss xnet.Session, th tunnel.Holder, in []byte) (err error) {
	p := pool.GetPacket()
	defer pool.PutPacket(p)

	if err = proto.Unmarshal(in, p); err != nil {
		err = errors.Wrapf(err, "packet unmarshal failed")
		return
	}

	if ss.IsCrypto() {
		index := ss.CSIndex()
		if index != int64(p.Index) {
			err = errors.Errorf("csindex validation failed. want=%d give=%d", index, p.Index)
			return
		}
		ss.IncreaseCSIndex()
	}

	if p.Compress {
		p.Data, err = compress.Decompress(p.Data)
		if err != nil {
			return errors.WithMessagef(err, "mod=%d seq=%d obj=%d", p.Mod, p.Seq, p.Obj)
		}
	}

	if p.Obj == 0 {
		p.Obj = int64(ss.UID())
	}
	ctx = rctx.SetOID(ctx, p.Obj)

	var t tunnel.Tunnel
	if t, err = th.Tunnel(ctx, p.Mod, p.Obj); err != nil {
		return errors.WithMessagef(err, "mod=%d seq=%d obj=%d", p.Mod, p.Seq, p.Obj)
	}
	if err = t.Forward(ctx, p); err != nil {
		return errors.WithMessagef(err, "mod=%d seq=%d obj=%d", p.Mod, p.Seq, p.Obj)
	}
	return nil
}
