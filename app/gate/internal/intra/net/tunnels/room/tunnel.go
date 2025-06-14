package room

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/tunnels"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/tunnels/base"
	clipkt "github.com/go-pantheon/janus/gen/api/client/packet"
	intrav1 "github.com/go-pantheon/janus/gen/api/server/room/intra/v1"
)

var _ tunnels.AppTunnel = (*Tunnel)(nil)

type Tunnel struct {
	*base.Tunnel

	stream intrav1.TunnelService_TunnelClient
}

func NewTunnel(ctx context.Context, oid int64, cli intrav1.TunnelServiceClient, ss xnet.Session, logger log.Logger) (*Tunnel, error) {
	stream, err := cli.Tunnel(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get tunnel failed")
	}

	t := &Tunnel{
		Tunnel: base.NewTunnel(tunnels.RoomTunnelType, oid, ss, logger),
		stream: stream,
	}

	return t, nil
}

// TransformMessage transform the packet to the forward message
// the forward message is pooled, so it should be put back to the pool after use on [xnet.CSHandle]
func (t *Tunnel) TransformMessage(p *clipkt.Packet) (to xnet.ForwardMessage) {
	msg := getRequest()

	msg.Mod = p.Mod
	msg.Seq = p.Seq
	msg.Obj = p.Obj
	msg.Data = p.Data
	msg.DataVersion = p.DataVersion
	msg.Index = p.Index

	return msg
}

// CSHandle send the message to the service
// the parameter [msg] is pooled, so it will be put back to the pool on the end of the function
func (t *Tunnel) CSHandle(msg xnet.ForwardMessage) error {
	defer putRequest(msg.(*intrav1.TunnelRequest))

	if err := t.stream.Send(msg.(*intrav1.TunnelRequest)); err != nil {
		return errors.Wrap(err, "stream send failed")
	}

	return nil
}

func (t *Tunnel) SCHandle() (xnet.ForwardMessage, error) {
	out, err := t.stream.Recv()
	if err != nil {
		return nil, errors.Wrap(err, "stream receive failed")
	}

	return out, nil
}

func (t *Tunnel) OnClose(err error) {
	if err := t.stream.CloseSend(); err != nil {
		t.Log().Errorf("[room.Tunnel] stream close failed. uid=%d sid=%d oid=%d color=%s status=%s %+v", t.UID(), t.Session().SID(), t.OID(), t.Color(), t.Session().Status(), err)
	}
}
