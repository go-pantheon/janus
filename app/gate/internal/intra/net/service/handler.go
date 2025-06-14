package service

import (
	"context"

	"github.com/go-pantheon/fabrica-kit/xcontext"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/compress"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/pool"
	"google.golang.org/protobuf/proto"
)

func (s *Service) Handle(ctx context.Context, ss xnet.Session, th xnet.TunnelManager, in []byte) (err error) {
	p := pool.GetPacket()
	defer pool.PutPacket(p)

	if err = proto.Unmarshal(in, p); err != nil {
		err = errors.Wrapf(err, "packet unmarshal failed")
		return
	}

	defer func() {
		if err != nil {
			err = errors.WithMessagef(err, "mod=%d seq=%d obj=%d", p.Mod, p.Seq, p.Obj)
		}
	}()

	if ss.IsCrypto() {
		index := ss.CSIndex()
		if index != int64(p.Index) {
			err = errors.Errorf("csindex validation failed. want=%d give=%d", index, p.Index)
			return
		}
		ss.IncreaseCSIndex()
	}

	if p.Compress {
		p.Data, err = compress.Decompress(p.Data)
		if err != nil {
			return err
		}
	}

	// the p.Obj is 0 when the mod is belong to player tunnel
	if p.Obj == 0 {
		p.Obj = int64(ss.UID())
	}

	ctx = xcontext.SetOID(ctx, p.Obj)

	t, err := th.Tunnel(ctx, p.Mod, p.Obj)
	if err != nil {
		return err
	}

	msg, err := t.TransformMessage(p)
	if err != nil {
		return err
	}

	return t.Forward(ctx, msg)
}

func (s *Service) Tick(ctx context.Context, ss xnet.Session) (err error) {
	// TODO: check black list
	return nil
}
