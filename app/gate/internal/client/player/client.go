package player

import (
	intrav1 "github.com/vulcan-frame/vulcan-gate/gen/api/server/player/intra/v1"
	"github.com/vulcan-frame/vulcan-pkg-app/router/conn"
)

const (
	serviceName = "vulcan.player.service"
)

type Conn struct {
	*conn.Conn
}

func NewConn(logger log.Logger, rt *RouteTable, r registry.Discovery) (*Conn, error) {
	conn, err := conn.NewConn(serviceName, logger, rt, r)
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
