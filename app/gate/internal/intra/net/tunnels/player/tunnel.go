package player

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/client/player"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/intra/net/tunnels"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/intra/net/tunnels/base"
	clipkt "github.com/vulcan-frame/vulcan-gate/gen/api/client/packet"
	intrav1 "github.com/vulcan-frame/vulcan-gate/gen/api/server/player/intra/v1"
	"github.com/vulcan-frame/vulcan-gate/pkg/net"
	"github.com/vulcan-frame/vulcan-gate/pkg/net/tunnel"
	"github.com/vulcan-frame/vulcan-pkg-app/router"
	"github.com/vulcan-frame/vulcan-pkg-tool/time"
)

var _ tunnels.AppTunnel = (*Tunnel)(nil)

// Tunnel is the player tunnel which is the main tunnel for the user
// if the player tunnel is closed, the worker will be closed, and the client will be disconnected
type Tunnel struct {
	*base.Tunnel

	worker     tunnel.Worker
	routeTable *player.RouteTable
	stream     intrav1.TunnelService_TunnelClient
}

func NewTunnel(ctx context.Context, cli intrav1.TunnelServiceClient, ss net.Session, log log.Logger, rt *player.RouteTable, worker tunnel.Worker) (*Tunnel, error) {
	stream, err := cli.Tunnel(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "get tunnel stream failed. uid=%d color=%s %+v", ss.UID(), ss.Color(), err)
	}

	// worker reset the player tunnel disconnect expired countdown time
	worker.Reset()

	t := &Tunnel{
		Tunnel:     base.NewTunnel(tunnels.PlayerTunnelType, ss.UID(), ss, log),
		worker:     worker,
		routeTable: rt,
		stream:     stream,
	}
	return t, nil
}

// TransformMessage transform the packet to the forward message
// the forward message is pooled, so it should be put back to the pool after use on [Tunnel.CSHandle]
func (t *Tunnel) TransformMessage(p *clipkt.Packet) (to tunnel.ForwardMessage) {
	msg := getMessage()
	msg.Mod = p.Mod
	msg.Seq = p.Seq
	msg.Obj = p.Obj
	msg.Data = p.Data
	msg.DataVersion = p.DataVersion
	return msg
}

// CSHandle send the message to the service
// the parameter [msg] is pooled, so it will be put back to the pool on the end of the function
func (t *Tunnel) CSHandle(msg tunnel.ForwardMessage) error {
	if err := t.stream.Send(msg.(*intrav1.Message)); err != nil {
		return errors.Wrapf(err, "stream send failed")
	}
	return nil
}

func (t *Tunnel) SCHandle() (tunnel.ForwardMessage, error) {
	out, err := t.stream.Recv()
	if err != nil {
		return nil, errors.Wrapf(err, "stream receive failed")
	}

	return out, nil
}

// OnStop is called when the player tunnel is closed
// it will delete the player route table and close the stream
func (t *Tunnel) OnStop() {
	// the player route table is only for the specific user, so it will be deleted when the player tunnel is closed
	if err := t.routeTable.DelDelay(context.Background(), t.Color(), t.UID(), router.HolderCacheTimeout); err != nil {
		t.Log().Errorf("[player.Tunnel] route table delete failed. uid=%d color=%s oid=%d %+v", t.UID(), t.Color(), t.OID(), err)
	}
	if err := t.stream.CloseSend(); err != nil {
		t.Log().Errorf("[player.Tunnel] stream close failed. uid=%d color=%s oid=%d %+v", t.UID(), t.Color(), t.OID(), err)
	}
}

// OnGroupStop is called when the player tunnel is closed
// it will close the worker when the player tunnel is closed
func (t *Tunnel) OnGroupStop(ctx context.Context, err error) {
	t.worker.SetExpiryTime(time.Now())
	t.Log().Debugf("[player.Tunnel] tunnel group exit. uid=%d color=%s oid=%d %+v", t.UID(), t.Color(), t.OID(), err)
}

func (t *Tunnel) Session() net.Session {
	return t.Session()
}
