package net

import (
	"context"
)

const (
	PackLenSize = 4
	MaxBodySize = int32(1 << 14)
)

type Service interface {
	Auth(ctx context.Context, in []byte) (out []byte, ss Session, err error)
	OnConnected(ctx context.Context, ss Session) (err error)
	OnDisconnect(ctx context.Context, ss Session) (err error)
	Handle(ctx context.Context, ss Session, in []byte) (err error)
}
