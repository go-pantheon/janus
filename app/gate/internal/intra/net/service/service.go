package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/janus/app/gate/internal/client/player"
	"github.com/go-pantheon/janus/app/gate/internal/client/room"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
	playerv1 "github.com/go-pantheon/janus/gen/api/server/player/intra/v1"
	roomv1 "github.com/go-pantheon/janus/gen/api/server/room/intra/v1"
	"github.com/google/wire"
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
