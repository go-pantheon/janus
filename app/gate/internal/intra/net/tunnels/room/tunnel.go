package room

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/intra/net/tunnels"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/intra/net/tunnels/base"
	clipkt "github.com/vulcan-frame/vulcan-gate/gen/api/client/packet"
	intrav1 "github.com/vulcan-frame/vulcan-gate/gen/api/server/room/intra/v1"
	"github.com/vulcan-frame/vulcan-gate/pkg/net"
	"github.com/vulcan-frame/vulcan-gate/pkg/net/tunnel"
)

var _ tunnels.AppTunnel = (*Tunnel)(nil)

type Tunnel struct {
	*base.Tunnel

	stream intrav1.TunnelService_TunnelClient
}

func NewTunnel(ctx context.Context, oid int64, cli intrav1.TunnelServiceClient, ss net.Session, logger log.Logger) (*Tunnel, error) {
	stream, err := cli.Tunnel(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "get tunnel failed. uid=%d color=%s oid=%d %+v", ss.UID(), ss.Color(), oid, err)
	}

	t := &Tunnel{
		Tunnel: base.NewTunnel(tunnels.RoomTunnelType, oid, ss, logger),
		stream: stream,
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
	defer putMessage(msg.(*intrav1.Message))
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

func (t *Tunnel) ResetMessage(msg tunnel.ForwardMessage) {
	if m, ok := msg.(*intrav1.Message); ok {
		putMessage(m)
	}
}

func (t *Tunnel) OnStop() {
	if err := t.stream.CloseSend(); err != nil {
		t.Log().Errorf("[room.Tunnel] stream close failed. uid=%d color=%s oid=%d %+v", t.UID(), t.Color(), t.OID(), err)
	}
}

func (t *Tunnel) OnGroupStop(ctx context.Context, err error) {
	t.Log().Errorf("[room.Tunnel] tunnel group exit. uid=%d color=%s oid=%d %+v", t.UID(), t.Color(), t.OID(), err)
}

func (t *Tunnel) Session() net.Session {
	return t.Session()
}
