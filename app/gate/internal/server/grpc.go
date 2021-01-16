package server

import (
	"math"

	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/conf"
	pushv1 "github.com/vulcan-frame/vulcan-gate/gen/api/server/gate/service/push/v1"
	"github.com/vulcan-frame/vulcan-pkg-app/metrics"
	"google.golang.org/grpc"
)

func NewGRPCServer(c *conf.Server, logger log.Logger, ps pushv1.PushServiceServer) *kgrpc.Server {
	var opts = []kgrpc.ServerOption{
		kgrpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, kgrpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, kgrpc.Address(c.Grpc.Addr))
	}

	svr := kgrpc.NewServer(opts...)
	pushv1.RegisterPushServiceServer(svr, ps)
	return svr
}
