package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/fabrica-util/security/rsa"
	"github.com/go-pantheon/fabrica-util/xtime"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/pool"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/security"
	climsg "github.com/go-pantheon/janus/gen/api/client/message"
	cliseq "github.com/go-pantheon/janus/gen/api/client/sequence"
	intrav1 "github.com/go-pantheon/janus/gen/api/server/gate/intra/v1"
	"google.golang.org/protobuf/proto"
)

func (s *Service) Auth(ctx context.Context, in []byte) (out []byte, session xnet.Session, err error) {
	if len(in) == 0 {
		return nil, nil, errors.New("proto is empty")
	}

	if s.encrypted {
		if in, err = security.DecryptCSHandshake(in); err != nil {
			return nil, nil, err
		}
	}

	var (
		cs = &climsg.CSHandshake{}
		sc = &climsg.SCHandshake{}

		data []byte
		pub  []byte
	)

	inp := pool.GetPacket()
	defer pool.PutPacket(inp)

	if err = proto.Unmarshal(in, inp); err != nil {
		return nil, nil, err
	}

	if inp.Seq != int32(cliseq.SystemSeq_Handshake) {
		return nil, nil, errors.Errorf("not handshake msg. seq=%d", inp.Seq)
	}

	if err = proto.Unmarshal(inp.Data, cs); err != nil {
		return nil, nil, errors.Wrapf(err, "CSHandshake decode failed. len=%d", len(inp.Data))
	}

	if pub, session, err = s.auth(cs.Token, cs.ServerId); err != nil {
		return nil, nil, err
	}

	sc.StartIndex = int32(session.IncreaseCSIndex())
	sc.Pub = pub

	if data, err = proto.Marshal(sc); err != nil {
		return nil, nil, errors.Wrap(err, "SCHandshake encode failed")
	}

	oup := pool.GetPacket()
	defer pool.PutPacket(oup)

	oup.Mod = inp.Mod
	oup.Seq = inp.Seq
	oup.Data = data

	if out, err = proto.Marshal(oup); err != nil {
		return nil, nil, errors.Wrap(err, "Packet encode failed")
	}

	if s.encrypted {
		pub, err := rsa.ParsePublicKey(cs.Pub)
		if err != nil {
			return nil, nil, errors.Wrap(err, "public key decode failed")
		}

		if out, err = rsa.Encrypt(pub, out); err != nil {
			return nil, nil, errors.WithMessage(err, "Packet encrypt failed")
		}
	}

	return out, session, nil
}

func (s *Service) auth(authToken string, sid int64) (key []byte, ss xnet.Session, err error) {
	if len(authToken) == 0 {
		return nil, nil, errors.New("[net.auth] token is empty")
	}

	var (
		now     = time.Now()
		token   *intrav1.AuthToken
		cryptor xnet.Cryptor
	)

	if token, err = decryptAccountToken(authToken); err != nil {
		return nil, nil, err
	}

	if now.After(xtime.Time(token.Timeout)) {
		return nil, nil, errors.Errorf("token timeout. timeout=%d now=%d", token.Timeout, now.Unix())
	}

	if cryptor, err = security.InitApiCrypto(s.encrypted); err != nil {
		return nil, nil, err
	}

	ss = xnet.NewSession(token.AccountId, sid, now.Unix(), cryptor, token.Color, int64(token.Status))

	return cryptor.Key(), ss, nil
}

func decryptAccountToken(token string) (auth *intrav1.AuthToken, err error) {
	if len(token) == 0 {
		return nil, errors.New("token is empty")
	}

	auth = &intrav1.AuthToken{}

	bytes, err := security.DecryptToken(token)
	if err != nil {
		return nil, errors.Wrap(err, "token decrypt failed")
	}

	if err = proto.Unmarshal(bytes, auth); err != nil {
		return nil, errors.Wrapf(err, "AuthToken proto decode failed")
	}

	return auth, nil
}

func (s *Service) OnConnected(ctx context.Context, ss xnet.Session) (err error) {
	log.Debugf("client connected. uid=%d sid=%d color=%s status=%d ip=%s", ss.UID(), ss.SID(), ss.Color(), ss.Status(), ss.ClientIP())
	return nil
}

func (s *Service) OnDisconnect(ctx context.Context, ss xnet.Session) (err error) {
	log.Debugf("client disconnected. uid=%d sid=%d color=%s status=%d ip=%s", ss.UID(), ss.SID(), ss.Color(), ss.Status(), ss.ClientIP())
	return nil
}
