package service

import (
	"context"
	"fmt"

	"github.com/go-pantheon/fabrica-kit/profile"
	"github.com/go-pantheon/fabrica-kit/xcontext"
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

func (s *Service) CreateAppTunnel(ctx context.Context, ss xnet.Session, tp int32, oid int64, worker xnet.Worker) (xnet.AppTunnel, error) {
	ctx = xcontext.AppendToOutgoingContext(ctx,
		xcontext.CtxUID, fmt.Sprintf("%d", ss.UID()),
		xcontext.CtxSID, fmt.Sprintf("%d", ss.SID()),
		xcontext.CtxStatus, fmt.Sprintf("%d", ss.Status()),
		xcontext.CtxColor, ss.Color(),
		xcontext.CtxGateReferer, fmt.Sprintf("%s#%d", profile.GRPCEndpoint(), worker.WID()),
	)

	switch tunnels.TunnelType(tp) {
	case tunnels.PlayerTunnelType:
		return player.NewTunnel(ctx, s.playerClient, ss, s.logger, worker)
	case tunnels.RoomTunnelType:
		return room.NewTunnel(ctx, oid, s.roomClient, ss, s.logger)
	default:
		return nil, errors.Errorf("unknown tunnel type. tunnel-type=%d", tp)
	}
}
