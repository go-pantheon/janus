package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-pantheon/vulcan-kit/metrics"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/conf"
	pushv1 "github.com/vulcan-frame/vulcan-gate/gen/api/server/gate/service/push/v1"
)

func NewHTTPServer(c *conf.Server, logger log.Logger, ps pushv1.PushServiceServer) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				metadata.Server(),
				tracing.Server(),
				metrics.Server(),
				logging.Server(logger),
			)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	svr := http.NewServer(opts...)
	pushv1.RegisterPushServiceHTTPServer(svr, ps)
	return svr
}
