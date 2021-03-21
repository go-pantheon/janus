package room

import (
	intrav1 "github.com/vulcan-frame/vulcan-gate/gen/api/server/room/intra/v1"
	"github.com/vulcan-frame/vulcan-pkg-app/router/balancer"
	"github.com/vulcan-frame/vulcan-pkg-app/router/conn"
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
