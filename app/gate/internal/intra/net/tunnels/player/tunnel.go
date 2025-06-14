package player

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/tunnels"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/tunnels/base"
	clipkt "github.com/go-pantheon/janus/gen/api/client/packet"
	intrav1 "github.com/go-pantheon/janus/gen/api/server/player/intra/v1"
)

var _ tunnels.AppTunnel = (*Tunnel)(nil)

// Tunnel is the player tunnel which is the main tunnel for the user
// if the player tunnel is closed, the worker will be closed, and the client will be disconnected
type Tunnel struct {
	*base.Tunnel

	worker xnet.Worker
	stream intrav1.TunnelService_TunnelClient
}

func NewTunnel(ctx context.Context, cli intrav1.TunnelServiceClient, ss xnet.Session, log log.Logger, worker xnet.Worker) (*Tunnel, error) {
	stream, err := cli.Tunnel(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get tunnel stream failed")
	}

	// worker reset the delayClosure countdown time
	worker.Reset()

	t := &Tunnel{
		Tunnel: base.NewTunnel(tunnels.PlayerTunnelType, ss.UID(), ss, log),
		worker: worker,
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

// OnClose is called when the player tunnel is closed
// it will delete the player route table and close the stream
func (t *Tunnel) OnClose(err error) {
	if err := t.stream.CloseSend(); err != nil {
		t.Log().Errorf("[player.Tunnel] stream close failed. uid=%d sid=%d color=%s status=%s %+v", t.UID(), t.Session().SID(), t.Color(), t.Session().Status(), err)
	}

	// TODO do something with error reason
	if err != nil {
		t.worker.SetExpiryTime(time.Now())
	}
}
