package server

import (
	"math"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-pantheon/fabrica-kit/metrics"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
	pushv1 "github.com/go-pantheon/janus/gen/api/server/gate/service/push/v1"
	grpcgo "google.golang.org/grpc"
)

func NewGRPCServer(c *conf.Server, logger log.Logger, ps pushv1.PushServiceServer) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			tracing.Server(),
			metrics.Server(),
			logging.Server(logger),
		),
	}

	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}

	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}

	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}

	opts = append(opts, grpc.Options(
		grpcgo.InitialConnWindowSize(1<<30),
		grpcgo.InitialWindowSize(1<<30),
		grpcgo.MaxConcurrentStreams(math.MaxInt32),
	))

	svr := grpc.NewServer(opts...)

	pushv1.RegisterPushServiceServer(svr, ps)

	return svr
}
