package room

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/tunnels"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/tunnels/base"
	intrav1 "github.com/go-pantheon/janus/gen/api/server/room/intra/v1"
)

var _ xnet.AppTunnel = (*Tunnel)(nil)

type Tunnel struct {
	*base.BaseTunnel

	stream intrav1.TunnelService_TunnelClient
}

func NewTunnel(ctx context.Context, oid int64, cli intrav1.TunnelServiceClient, ss xnet.Session, logger log.Logger) (*Tunnel, error) {
	stream, err := cli.Tunnel(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get tunnel failed")
	}

	t := &Tunnel{
		BaseTunnel: base.New(tunnels.RoomTunnelType, oid, ss, logger),
		stream:     stream,
	}

	return t, nil
}

// PacketToForwardMsg transform the packet to the forward message
// the forward message is pooled, so it should be put back to the pool after use on [xnet.CSHandle]
func (t *Tunnel) PacketToTunnelMsg(from xnet.PacketMessage) (to xnet.TunnelMessage) {
	msg := get()

	msg.Mod = from.GetMod()
	msg.Seq = from.GetSeq()
	msg.Obj = from.GetObj()
	msg.Data = from.GetData()
	msg.DataVersion = from.GetDataVersion()
	msg.Index = from.GetIndex()
	msg.ConnId = from.GetConnId()
	msg.FragId = from.GetFragId()
	msg.FragCount = from.GetFragCount()
	msg.FragIndex = from.GetFragIndex()

	return msg
}

// CSHandle send the message to the service
// the parameter [msg] is pooled, so it will be put back to the pool on the end of the function
func (t *Tunnel) CSHandle(msg xnet.TunnelMessage) error {
	defer put(msg.(*intrav1.TunnelRequest))

	if err := t.stream.Send(msg.(*intrav1.TunnelRequest)); err != nil {
		return errors.Wrap(err, "stream send failed")
	}

	return nil
}

func (t *Tunnel) SCHandle() (xnet.TunnelMessage, error) {
	out, err := t.stream.Recv()
	if err != nil {
		return nil, errors.Wrap(err, "stream receive failed")
	}

	return out, nil
}

func (t *Tunnel) OnStop(ctx context.Context, erreason error) (err error) {
	if streamErr := t.stream.CloseSend(); streamErr != nil {
		err = errors.Join(err, streamErr)
	}

	return err
}
