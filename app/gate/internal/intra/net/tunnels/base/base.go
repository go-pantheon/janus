package base

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/tunnels"
)

var _ tunnels.AppTunnelBase = (*Tunnel)(nil)

type Tunnel struct {
	log *log.Helper

	tunnelType tunnels.TunnelType
	session    xnet.Session
	oid        int64
}

func NewTunnel(tp tunnels.TunnelType, oid int64, ss xnet.Session, logger log.Logger) *Tunnel {
	t := &Tunnel{
		log:        log.NewHelper(log.With(logger, "module", fmt.Sprintf("gate/tunnel/%d", tp))),
		tunnelType: tp,
		session:    ss,
		oid:        oid,
	}

	return t
}

func (t *Tunnel) Type() int32 {
	return int32(t.tunnelType)
}

func (t *Tunnel) Log() *log.Helper {
	return t.log
}

func (t *Tunnel) UID() int64 {
	return t.session.UID()
}

func (t *Tunnel) OID() int64 {
	return t.oid
}

func (t *Tunnel) Color() string {
	return t.session.Color()
}

func (t *Tunnel) Session() xnet.Session {
	return t.session
}
