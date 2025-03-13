package room

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-pantheon/vulcan-kit/router/balancer"
	"github.com/go-pantheon/vulcan-kit/router/conn"
	intrav1 "github.com/vulcan-frame/vulcan-gate/gen/api/server/room/intra/v1"
)

const (
	serviceName = "vulcan.room.service"
)

type Conn struct {
	*conn.Conn
}

func NewConn(logger log.Logger, rt *RouteTable, r registry.Discovery) (*Conn, error) {
	conn, err := conn.NewConn(serviceName, balancer.BalancerTypeMaster, logger, rt, r)
	if err != nil {
		return nil, err
	}

	return &Conn{
		Conn: conn,
	}, nil
}

func NewClient(conn *Conn) intrav1.TunnelServiceClient {
	return intrav1.NewTunnelServiceClient(conn.ClientConnInterface)
}
