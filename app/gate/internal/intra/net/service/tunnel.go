package service

import (
	"context"

	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/tunnels"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/tunnels/player"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/tunnels/room"
	climod "github.com/go-pantheon/janus/gen/api/client/module"
	"github.com/pkg/errors"
)

const (
	playerTunnelCapacity = 1
	roomTunnelCapacity   = 8
)

// TunnelType returns the tunnel type based on the module ID
// most modules have the specific tunnel type, except for the player module
func (s *Service) TunnelType(mod int32) (t int32, initCapacity int, err error) {
	switch climod.ModuleID(mod) {
	case climod.ModuleID_Room:
		return int32(tunnels.RoomTunnelType), roomTunnelCapacity, nil
	default:
		return int32(tunnels.PlayerTunnelType), playerTunnelCapacity, nil
	}
}

func (s *Service) CreateTunnel(ctx context.Context, ss xnet.Session, tp int32, oid int64, worker xnet.Worker) (xnet.Tunnel, error) {
	tunnel, err := s.createAppTunnel(ctx, ss, tp, oid, worker)
	if err != nil {
		return nil, err
	}

	return tunnels.NewTunnel(ctx, worker, tunnel), nil
}

func (s *Service) createAppTunnel(ctx context.Context, ss xnet.Session, tp int32, oid int64, worker xnet.Worker) (tunnels.AppTunnel, error) {
	switch tunnels.TunnelType(tp) {
	case tunnels.PlayerTunnelType:
		return player.NewTunnel(ctx, s.playerClient, ss, s.logger, worker)
	case tunnels.RoomTunnelType:
		return room.NewTunnel(ctx, oid, s.roomClient, ss, s.logger)
	default:
		return nil, errors.Errorf("unknown tunnel type. tunnel-type=%d", tp)
	}
}
