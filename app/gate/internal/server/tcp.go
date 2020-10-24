package server

import (
	"context"

	"github.com/pkg/errors"
	tcp "github.com/vulcan-frame/vulcan-gate/pkg/net/tcp/server"
	"github.com/vulcan-frame/vulcan-pkg-app/router/routetable"
)

func NewTCPServer(addr string, svc *service.Service, rt routetable.RouteTable) (*tcp.Server, error) {
	s, err := tcp.NewServer(addr, svc, rt)
	if err != nil {
		return nil, errors.Wrapf(err, "create tcp server failed")
	}
	return s, nil
}
