package server

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/conf"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/intra/net/service"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/pkg/middleware/logging"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/pkg/middleware/metadata"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/router"
	"github.com/vulcan-frame/vulcan-gate/pkg/net"
	tcp "github.com/vulcan-frame/vulcan-gate/pkg/net/tcp/server"
	"github.com/vulcan-frame/vulcan-pkg-app/metrics"
	"github.com/vulcan-frame/vulcan-pkg-app/router/routetable"
)

func NewTCPServer(c *conf.Server, logger log.Logger, rt *router.RouteTable, svc *service.Service) (*tcp.Server, error) {
	var opts = []tcp.Option{
		tcp.ReadFilter(
			middleware.Chain(
				recovery.Recovery(),
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
