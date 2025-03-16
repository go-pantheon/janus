package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/service"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/middleware/logging"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/middleware/metadata"
	"github.com/go-pantheon/janus/app/gate/internal/router"
	"github.com/go-pantheon/vulcan-kit/metrics"
	"github.com/go-pantheon/vulcan-kit/router/routetable"
	"github.com/go-pantheon/vulcan-net"
	tcp "github.com/go-pantheon/vulcan-net/tcp/server"
	"github.com/pkg/errors"
)

func NewTCPServer(c *conf.Server, logger log.Logger, rt *router.RouteTable, svc *service.Service) (*tcp.Server, error) {
	var opts = []tcp.Option{
		tcp.ReadFilter(
			middleware.Chain(
				recovery.Recovery(),
				metadata.Server(),
				tracing.Server(),
				metrics.Server(),
				logging.Request(net.NetKindTCP),
			),
		),
		tcp.WriteFilter(
			middleware.Chain(
				logging.Reply(net.NetKindTCP),
			),
		),
	}

	if c.Tcp.Addr != "" {
		opts = append(opts, tcp.Bind(c.Tcp.Addr))
	}
	if logger != nil {
		opts = append(opts, tcp.Logger(logger))
	}
	if rt != nil {
		opts = append(opts, tcp.AfterConnectFunc(afterConnectFunc(rt)))
		opts = append(opts, tcp.AfterDisconnectFunc(afterDisconnectFunc(rt)))
	}

	s, err := tcp.NewServer(svc, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "create tcp server failed. %+v", c)
	}
	return s, nil
}

func afterConnectFunc(rt routetable.RouteTable) func(ctx context.Context, color string, uid int64) error {
	grt := rt.(*router.RouteTable)
	return func(ctx context.Context, color string, uid int64) error {
		return router.AddRouteTable(ctx, grt, color, uid)
	}
}

func afterDisconnectFunc(rt routetable.RouteTable) func(ctx context.Context, color string, uid int64) error {
	grt := rt.(*router.RouteTable)
	return func(ctx context.Context, color string, uid int64) error {
		_ = router.DelRouteTable(ctx, grt, color, uid)
		return nil
	}
}
