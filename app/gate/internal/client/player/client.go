package player

import (
	"github.com/vulcan-frame/vulcan-pkg-app/router/conn"
)

type Conn struct {
	*conn.Conn
}

func NewConn(logger log.Logger, rt *RouteTable) (*Conn, error) {
	conn, err := conn.NewConn(serviceName, logger, rt)
	if err != nil {
		return nil, err
	}

	return &Conn{
		Conn: conn,
	}, nil
}
