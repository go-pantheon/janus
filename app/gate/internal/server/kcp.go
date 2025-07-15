package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-pantheon/fabrica-kit/metrics"
	kcp "github.com/go-pantheon/fabrica-net/kcp/server"
	"github.com/go-pantheon/fabrica-net/server"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/service"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/middleware/logging"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/middleware/metadata"
	"github.com/go-pantheon/janus/app/gate/internal/router"
)

func NewKCPServer(c *conf.Server, logger log.Logger, rt *router.RouteTable, svc *service.Service) (*kcp.Server, error) {
	if rt == nil {
		return nil, errors.New("route table is nil")
	}

	if svc == nil {
		return nil, errors.New("service is nil")
	}

	var opts = []server.Option{
		server.WithReadFilter(
			middleware.Chain(
				recovery.Recovery(),
				metadata.Server(),
				tracing.Server(),
				metrics.Server(),
				logging.Request(xnet.NetKindKCP),
			),
		),
		server.WithWriteFilter(
			middleware.Chain(
				logging.Reply(xnet.NetKindKCP),
			),
		),
	}

	if logger != nil {
		opts = append(opts, server.WithLogger(logger))
	}

	opts = append(opts, server.WithAfterConnect(afterConnect(rt)))

	s, err := kcp.NewServer(c.Kcp.Addr, svc, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "create KCP server failed")
	}

	return s, nil
}
