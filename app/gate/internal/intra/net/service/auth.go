package service

import (
	"context"
	"time"

	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/fabrica-util/security/ecdh"
	"github.com/go-pantheon/fabrica-util/xtime"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/pool"
	"github.com/go-pantheon/janus/app/gate/internal/pkg/security"
	climsg "github.com/go-pantheon/janus/gen/api/client/message"
	clipkt "github.com/go-pantheon/janus/gen/api/client/packet"
	cliseq "github.com/go-pantheon/janus/gen/api/client/sequence"
	intrav1 "github.com/go-pantheon/janus/gen/api/server/gate/intra/v1"
	"google.golang.org/protobuf/proto"
)

const HandshakeTimeLimit = 5 * time.Second

func (s *Service) Auth(ctx context.Context, in xnet.Pack) (out xnet.Pack, session xnet.Session, err error) {
	if len(in) == 0 {
		return nil, nil, errors.New("proto is empty")
	}

	var (
		cs = &climsg.CSHandshake{}
		sc = &climsg.SCHandshake{}

		svrSign []byte
		scData  []byte
	)

	inp := pool.GetPacket()
	defer pool.PutPacket(inp)

	if err = proto.Unmarshal(in, inp); err != nil {
		return nil, nil, errors.Wrap(err, "packet unmarshal failed")
	}

	if inp.Seq != int32(cliseq.SystemSeq_Handshake) {
		return nil, nil, errors.Errorf("not handshake msg. seq=%d", inp.Seq)
	}

	if err = proto.Unmarshal(inp.Data, cs); err != nil {
		return nil, nil, errors.Wrapf(err, "CSHandshake decode failed. len=%d", len(inp.Data))
	}

	if session, svrSign, err = s.auth(inp, cs); err != nil {
		return nil, nil, err
	}

	sc.Pub = session.SelfPublicKey()
	sc.Sign = svrSign
	sc.Timestamp = time.Now().Unix()
	sc.StartIndex = session.IncreaseCSIndex()

	if scData, err = proto.Marshal(sc); err != nil {
		return nil, nil, errors.Wrap(err, "SCHandshake encode failed")
	}

	oup := pool.GetPacket()
	defer pool.PutPacket(oup)

	oup.Mod = inp.Mod
	oup.Seq = inp.Seq
	oup.Data = scData

	if out, err = proto.Marshal(oup); err != nil {
		return nil, nil, errors.Wrap(err, "Packet encode failed")
	}

	return out, session, nil
}

func (s *Service) auth(inp *clipkt.Packet, cs *climsg.CSHandshake) (xnet.Session, []byte, error) {
	if len(cs.Token) == 0 {
		return nil, nil, errors.New("[net.auth] token is empty")
	}

	now := time.Now()
	before := now.Add(-HandshakeTimeLimit)
	after := now.Add(HandshakeTimeLimit)

	if xtime.Time(cs.Timestamp).Before(before) || xtime.Time(cs.Timestamp).After(after) {
		return nil, nil, errors.Errorf("handshake timestamp not match. timestamp=%d now=%d", cs.Timestamp, now.Unix())
	}

	token, err := decryptAccountToken(cs.Token)
	if err != nil {
		return nil, nil, err
	}

	if now.After(xtime.Time(token.Timeout)) {
		return nil, nil, errors.Errorf("token timeout. timeout=%d now=%d", token.Timeout, now.Unix())
	}

	var (
		ecdhInfo  xnet.ECDHable
		scPubSign []byte
		cryptor   xnet.Cryptor
	)

	if !s.encrypted {
		ecdhInfo = xnet.NewUnECDH()
		cryptor = xnet.NewUnCryptor()
	} else {
		ecdhInfo, scPubSign, err = newECDHInfo(cs.Pub, cs.Sign)
		if err != nil {
			return nil, nil, err
		}

		cryptor, err = xnet.NewCryptor(ecdhInfo.SharedKey())
		if err != nil {
			return nil, nil, err
		}
	}

	sid := cs.ServerId
	if sid == 0 {
		sid = token.ServerId
	}

	ss := xnet.NewSession(token.AccountId, token.Color, int64(token.Status),
		xnet.WithSID(sid),
		xnet.WithConnID(inp.ConnId),
		xnet.WithEncryptor(cryptor),
		xnet.WithECDH(ecdhInfo),
		xnet.WithStartTime(now.Unix()),
	)

	return ss, scPubSign, nil
}

func newECDHInfo(csPub []byte, csSign []byte) (ecdhInfo xnet.ECDHable, scPubSign []byte, err error) {
	if err := security.VerifyECDHCliPubKey(csPub, csSign); err != nil {
		return nil, nil, err
	}

	csPubBytes, err := ecdh.ParseKey(csPub)
	if err != nil {
		return nil, nil, err
	}

	ecdhInfo, err = xnet.NewECDHInfoWithGenSelfPair(csPubBytes)
	if err != nil {
		return nil, nil, err
	}

	scPubSign, err = security.SignECDHSvrPubKey(ecdhInfo.SelfPublicKey())
	if err != nil {
		return nil, nil, err
	}

	return ecdhInfo, scPubSign, nil
}

func decryptAccountToken(token string) (auth *intrav1.AuthToken, err error) {
	if len(token) == 0 {
		return nil, errors.New("token is empty")
	}

	auth = &intrav1.AuthToken{}

	bytes, err := security.DecryptAccountToken(token)
	if err != nil {
		return nil, errors.Wrap(err, "token decrypt failed")
	}

	if err = proto.Unmarshal(bytes, auth); err != nil {
		return nil, errors.Wrapf(err, "AuthToken proto decode failed")
	}

	return auth, nil
}
