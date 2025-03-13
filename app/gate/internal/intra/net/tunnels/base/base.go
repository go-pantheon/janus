package base

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/intra/net/tunnels"
	"github.com/vulcan-frame/vulcan-gate/pkg/net"
)

var _ tunnels.AppTunnelBase = (*Tunnel)(nil)

type Tunnel struct {
	log        *log.Helper
	tunnelType tunnels.TunnelType
	oid        int64
	session    net.Session
}

func NewTunnel(tp tunnels.TunnelType, oid int64, ss net.Session, logger log.Logger) *Tunnel {
	t := &Tunnel{
		log:        log.NewHelper(log.With(logger, "module", fmt.Sprintf("gate/tunnel/%d", tp))),
		tunnelType: tp,
		oid:        oid,
		session:    ss,
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

func (t *Tunnel) Session() net.Session {
	return t.session
}
