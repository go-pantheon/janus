package player

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-pantheon/fabrica-kit/router/balancer"
	"github.com/go-pantheon/fabrica-kit/router/conn"
	intrav1 "github.com/go-pantheon/janus/gen/api/server/player/intra/v1"
)

const (
	serviceName    = "roma.player.service"
	routeTableName = "player"
)

func NewClient(conn *Conn) intrav1.TunnelServiceClient {
	return intrav1.NewTunnelServiceClient(conn.ClientConnInterface)
}

type Conn struct {
	*conn.Conn
}

func NewConn(logger log.Logger, rt *RouteTable, r registry.Discovery) (*Conn, error) {
	conn, err := conn.NewConn(serviceName, balancer.TypeMaster, logger, rt, r)
	if err != nil {
		return nil, err
	}

	return &Conn{
		Conn: conn,
	}, nil
}
