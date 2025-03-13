package server

import (
	"math"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-pantheon/vulcan-kit/metrics"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/conf"
	pushv1 "github.com/vulcan-frame/vulcan-gate/gen/api/server/gate/service/push/v1"
	"google.golang.org/grpc"
)

func NewGRPCServer(c *conf.Server, logger log.Logger, ps pushv1.PushServiceServer) *kgrpc.Server {
	var opts = []kgrpc.ServerOption{
		kgrpc.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			tracing.Server(),
			metrics.Server(),
			logging.Server(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, kgrpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, kgrpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, kgrpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	opts = append(opts, kgrpc.Options(
		grpc.InitialConnWindowSize(1<<30),
		grpc.InitialWindowSize(1<<30),
		grpc.MaxConcurrentStreams(math.MaxInt32),
	))

	svr := kgrpc.NewServer(opts...)
	pushv1.RegisterPushServiceServer(svr, ps)
	return svr
}
