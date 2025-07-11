package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-pantheon/fabrica-kit/metrics"
	"github.com/go-pantheon/fabrica-net/server"
	tcp "github.com/go-pantheon/fabrica-net/tcp/server"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/service"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/middleware/logging"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/middleware/metadata"
	"github.com/go-pantheon/janus/app/gate/internal/router"
)

func NewTCPServer(c *conf.Server, logger log.Logger, rt *router.RouteTable, svc *service.Service) (*tcp.Server, error) {
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
				logging.Request(xnet.NetKindTCP),
			),
		),
		server.WithWriteFilter(
			middleware.Chain(
				logging.Reply(xnet.NetKindTCP),
			),
		),
	}

	if logger != nil {
		opts = append(opts, server.WithLogger(logger))
	}

	opts = append(opts, server.WithAfterConnect(afterConnect(rt)))

	s, err := tcp.NewServer(c.Tcp.Addr, svc, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "create tcp server failed")
	}

	return s, nil
}

func afterConnect(rt *router.RouteTable) server.Inspector {
	return func(f server.InspectorFunc) server.InspectorFunc {
		return func(ctx context.Context, w xnet.Worker) error {
			return router.AddRouteTable(ctx, rt, w.Session().Color(), w.Session().UID())
		}
	}
}
