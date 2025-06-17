package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-pantheon/fabrica-kit/metrics"
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

	var opts = []tcp.Option{
		tcp.ReadFilter(
			middleware.Chain(
				recovery.Recovery(),
				metadata.Server(),
				tracing.Server(),
				metrics.Server(),
				logging.Request(xnet.NetKindTCP),
			),
		),
		tcp.WriteFilter(
			middleware.Chain(
				logging.Reply(xnet.NetKindTCP),
			),
		),
	}

	if logger != nil {
		opts = append(opts, tcp.Logger(logger))
	}

	opts = append(opts, tcp.AfterConnectFunc(afterConnectFunc(rt)))
	opts = append(opts, tcp.AfterDisconnectFunc(afterDisconnectFunc(rt)))

	s, err := tcp.NewServer(c.Tcp.Addr, svc, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "create tcp server failed")
	}

	return s, nil
}

func afterConnectFunc(rt *router.RouteTable) tcp.WrapperFunc {
	return func(ctx context.Context, uid int64, color string) error {
		return router.AddRouteTable(ctx, rt, color, uid)
	}
}

func afterDisconnectFunc(rt *router.RouteTable) tcp.WrapperFunc {
	return func(ctx context.Context, uid int64, color string) error {
		return router.DelRouteTable(ctx, rt, color, uid)
	}
}
