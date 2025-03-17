package service

import (
	"context"
	"crypto/cipher"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-net"
	"github.com/go-pantheon/fabrica-util/security/rsa"
	"github.com/go-pantheon/fabrica-util/time"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/pool"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/security"
	climsg "github.com/go-pantheon/janus/gen/api/client/message"
	clipkt "github.com/go-pantheon/janus/gen/api/client/packet"
	cliseq "github.com/go-pantheon/janus/gen/api/client/sequence"
	intrav1 "github.com/go-pantheon/janus/gen/api/server/gate/intra/v1"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func (s *Service) OnConnected(ctx context.Context, ss net.Session) (err error) {
	log.Debugf("[net.Service] connected. uid=%d color=%s status=%d", ss.UID(), ss.Color(), ss.Status())
	return nil
}

func (s *Service) OnDisconnect(ctx context.Context, ss net.Session) (err error) {
	log.Debugf("[net.Service] disconnected. uid=%d color=%s status=%d", ss.UID(), ss.Color(), ss.Status())
	return nil
}

func (s *Service) Auth(ctx context.Context, in []byte) (out []byte, session net.Session, err error) {
	if len(in) <= 0 {
		err = errors.New("proto is empty")
		return
	}

	if s.encrypted {
		if in, err = security.DecryptCSHandshake(in); err != nil {
			return
		}
	}

	var (
		inp = &clipkt.Packet{}
		cs  = &climsg.CSHandshake{}
		sc  = &climsg.SCHandshake{}

		data []byte
		pub  []byte
	)

	if err = proto.Unmarshal(in, inp); err != nil {
		return nil, nil, err
	}

	if inp.Seq != int32(cliseq.SystemSeq_Handshake) {
		return nil, nil, errors.New("not handshake msg")
	}
	if err = proto.Unmarshal(inp.Data, cs); err != nil {
		err = errors.Wrap(err, "CSHandshake decode failed")
		return
	}

	log.Debugf("[net.Service] handshake received. len=%d token=%s", len(inp.Data), cs.Token)

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

	oup.Index = int32(session.IncreaseSCIndex())
	oup.Mod = inp.Mod
	oup.Seq = inp.Seq
	oup.Data = data

	if out, err = proto.Marshal(oup); err != nil {
		return nil, nil, errors.Wrap(err, "Packet encode failed")
	}

	if s.encrypted {
		pub, err := rsa.ParsePublicKey(cs.Pub)
		if err != nil {
			return nil, nil, errors.Wrap(err, "RSA public key decode failed")
		}
		if out, err = rsa.Encrypt(pub, out); err != nil {
			return nil, nil, errors.WithMessage(err, "Packet encrypt failed")
		}
	}

	return out, session, nil
}

func (s *Service) auth(authToken string, sid int64) (key []byte, ss net.Session, err error) {
	if len(authToken) <= 0 {
		err = errors.New("[net.auth] token is empty")
		return
	}

	var (
		token *intrav1.AuthToken
		block cipher.Block
	)

	now := time.Now()
	if token, err = decryptAccountToken(authToken); err != nil {
		return
	}
	if now.After(time.Time(token.Timeout)) {
		err = errors.New("token expired")
		return
	}
	if block, key, err = security.InitApiCrypto(); err != nil {
		return
	}

	ss = net.NewSession(token.AccountId, sid, now.Unix(), block, key, s.encrypted, token.Color, int64(token.Status))
	return
}

func decryptAccountToken(token string) (auth *intrav1.AuthToken, err error) {
	if len(token) <= 0 {
		err = errors.New("token is empty")
		return
	}

	auth = &intrav1.AuthToken{}

	bytes, err := security.DecryptToken(token)
	if err != nil {
		err = errors.Wrap(err, "token decrypt failed")
		return
	}

	if err = proto.Unmarshal(bytes, auth); err != nil {
		err = errors.Wrap(err, "AuthToken proto decode failed")
	}
	return
}
