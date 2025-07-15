package base

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/compress"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/janus/app/gate/internal/intra/net/tunnels"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/pool"
	"google.golang.org/protobuf/proto"
)

var _ xnet.AppBaseTunnel = (*BaseTunnel)(nil)

type BaseTunnel struct {
	log *log.Helper

	tunnelType tunnels.TunnelType
	session    xnet.Session
	oid        int64
}

func New(tp tunnels.TunnelType, oid int64, ss xnet.Session, logger log.Logger) *BaseTunnel {
	t := &BaseTunnel{
		log:        log.NewHelper(log.With(logger, "module", fmt.Sprintf("gate/tunnel/%d", tp))),
		tunnelType: tp,
		session:    ss,
		oid:        oid,
	}

	return t
}

func (t *BaseTunnel) TunnelMsgToPack(ctx context.Context, msg xnet.TunnelMessage) (pack xnet.Pack, err error) {
	p := pool.GetPacket()
	defer pool.PutPacket(p)

	p.Mod = msg.GetMod()
	p.Seq = msg.GetSeq()
	p.Obj = msg.GetObj()
	p.Index = msg.GetIndex()

	if smuxMsg, ok := msg.(xnet.SmuxParam); ok {
		p.ConnId = smuxMsg.GetConnId()
		p.FragId = smuxMsg.GetFragId()
		p.FragCount = smuxMsg.GetFragCount()
		p.FragIndex = smuxMsg.GetFragIndex()
	}

	p.Data, p.Compress, err = compress.Compress(msg.GetData())
	if err != nil {
		return nil, errors.Wrapf(err, "compress data failed")
	}

	pack, err = proto.Marshal(p)
	if err != nil {
		return nil, errors.Wrapf(err, "packet marshal failed")
	}

	return pack, nil
}

func (t *BaseTunnel) Type() int32 {
	return int32(t.tunnelType)
}

func (t *BaseTunnel) Log() *log.Helper {
	return t.log
}

func (t *BaseTunnel) UID() int64 {
	return t.session.UID()
}

func (t *BaseTunnel) OID() int64 {
	return t.oid
}

func (t *BaseTunnel) Color() string {
	return t.session.Color()
}

func (t *BaseTunnel) Session() xnet.Session {
	return t.session
}
