package v1

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	servicev1 "github.com/vulcan-frame/vulcan-gate/gen/api/server/gate/service/push/v1"
	tcp "github.com/vulcan-frame/vulcan-gate/pkg/net/tcp/server"
)

var _ servicev1.PushServiceServer = (*PushService)(nil)

type PushService struct {
	servicev1.UnimplementedPushServiceServer

	log    *log.Helper
	server *tcp.Server
}

func NewPushService(logger log.Logger, ts *tcp.Server) servicev1.PushServiceServer {
	return &PushService{
		UnimplementedPushServiceServer: servicev1.UnimplementedPushServiceServer{},
		log:                            log.NewHelper(log.With(logger, "module", "gate/service/push")),
		server:                         ts,
	}
}

func (s *PushService) Push(ctx context.Context, req *servicev1.PushRequest) (*servicev1.PushResponse, error) {
	return &servicev1.PushResponse{}, nil
}

func (s *PushService) Multicast(ctx context.Context, req *servicev1.MulticastRequest) (*servicev1.MulticastResponse, error) {
	return &servicev1.MulticastResponse{}, nil
}

func (s *PushService) Broadcast(ctx context.Context, req *servicev1.BroadcastRequest) (*servicev1.BroadcastResponse, error) {
	return &servicev1.BroadcastResponse{}, nil
}
