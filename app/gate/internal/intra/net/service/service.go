package service

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/client/player"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/conf"
	playerv1 "github.com/vulcan-frame/vulcan-gate/gen/api/server/player/intra/v1"
	xnet "github.com/vulcan-frame/vulcan-gate/pkg/net"
	rctx "github.com/vulcan-frame/vulcan-gate/pkg/net/context"
	"github.com/vulcan-frame/vulcan-gate/pkg/net/tunnel"
	"github.com/vulcan-frame/vulcan-pkg-tool/compress"
	"google.golang.org/protobuf/proto"
)

var ProviderSet = wire.NewSet(NewTCPService)

var _ xnet.Service = (*Service)(nil)

type Service struct {
	logger    log.Logger
	encrypted bool

	playerClient playerv1.TunnelServiceClient
	playerRT     *player.RouteTable
}

func NewTCPService(logger log.Logger, label *conf.Label,
	playerRT *player.RouteTable, playerClient playerv1.TunnelServiceClient,
) *Service {
	return &Service{
		logger:       logger,
		encrypted:    label.Encrypted,
		playerClient: playerClient,
		playerRT:     playerRT,
	}
}

func (s *Service) Handle(ctx context.Context, ss xnet.Session, th tunnel.Holder, in []byte) (err error) {
	if err = s.handle(ctx, ss, th, in); err != nil {
		return errors.WithMessagef(err, "uid=%d color=%s status=%d", ss.UID(), ss.Color(), ss.Status())
	}
	return nil
}

func (s *Service) handle(ctx context.Context, ss xnet.Session, th tunnel.Holder, in []byte) (err error) {
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
