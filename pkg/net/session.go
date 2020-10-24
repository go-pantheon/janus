package net

import (
	"crypto/cipher"

	"go.uber.org/atomic"
)

type Session interface {
	Encryptor

	UID() int64
	SID() int64
	Color() string

	CSIndex() int64
	SCIndex() int64
}

type Encryptor interface {
	Block() cipher.Block
	Key() []byte
}

var _ Session = (*session)(nil)

type session struct {
	*encryptor

	userId   int64
	serverId int64
	color    string

	csIndex *indexInfo
	scIndex *indexInfo
}

func DefaultSession() Session {
	return &session{
		encryptor: &encryptor{},
		csIndex:   newIndexInfo(0),
		scIndex:   newIndexInfo(1),
	}
}

func NewSession(userId int64, sid int64, block cipher.Block, key []byte, color string) Session {
	s := &session{
		encryptor: &encryptor{
			block:   block,
			key:     key,
		},
		userId:    userId,
		color:     color,
		serverId:  sid,
		csIndex:   newIndexInfo(0),
		scIndex:   newIndexInfo(1),
	}
	return s
}

func (s *session) CSIndex() int64 {
	return s.csIndex.Load()
}

func (s *session) SCIndex() int64 {
	return s.scIndex.Load()
}

func (s *session) UID() int64 {
	return s.userId
}

func (s *session) SID() int64 {
	return s.serverId
}

func (s *session) Color() string {
	if len(s.color) == 0 {
		return ""
	}
	return s.color
}

func newIndexInfo(start int64) *indexInfo {
	return &indexInfo{
		start: start,
		index: atomic.NewInt64(start),
	}
}

func (i *indexInfo) Increase() int64 {
	return i.index.Add(1)
}

func (i *indexInfo) Load() int64 {
	return i.index.Load()
}

var _ Encryptor = (*encryptor)(nil)

type encryptor struct {
	encrypt bool
	block   cipher.Block
	key     []byte
}

func (c *encryptor) IsCrypto() bool {
	return c.encrypt
}

func (c *encryptor) Block() cipher.Block {
	return c.block
}

func (c *encryptor) Key() []byte {
	return c.key
}
