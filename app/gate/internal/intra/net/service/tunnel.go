package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/intra/net/tunnels"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/intra/net/tunnels/player"
	xnet "github.com/vulcan-frame/vulcan-gate/pkg/net"
	"github.com/vulcan-frame/vulcan-gate/pkg/net/tunnel"
)

// TunnelType returns the tunnel type based on the module ID
// most modules have the specific tunnel type, except for the player module
func (s *Service) TunnelType(mod int32) (int32, error) {
	return int32(tunnels.PlayerTunnelType), nil
}

func (s *Service) CreateTunnel(ctx context.Context, ss xnet.Session, tp int32, oid int64, worker tunnel.Worker) (tunnel.Tunnel, error) {
	tunnel, err := s.createAppTunnel(ctx, ss, tp, oid, worker)
	if err != nil {
		return nil, err
	}
	return tunnels.NewTunnel(ctx, worker, tunnel), nil
}

func (s *Service) createAppTunnel(ctx context.Context, ss xnet.Session, tp int32, oid int64, worker tunnel.Worker) (tunnels.AppTunnel, error) {
	switch tunnels.TunnelType(tp) {
	case tunnels.PlayerTunnelType:
		return player.NewTunnel(ctx, s.playerClient, ss, s.logger, s.playerRT, worker)
	default:
		return nil, errors.Errorf("TunnelType invalid. TunnelType=%d uid=%d color=%s oid=%d", tp, ss.UID(), ss.Color(), oid)
	}
}
