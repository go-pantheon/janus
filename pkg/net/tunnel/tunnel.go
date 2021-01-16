package tunnel

import (
	"context"

	"github.com/vulcan-frame/vulcan-pkg-tool/sync"
)

type Holder interface {
	Pusher
	Tunnel(ctx context.Context, key int32, oid int64) (Tunnel, error)
}

type Pusher interface {
	Push(ctx context.Context, pack []byte) error
}

type Worker interface {
	sync.Stoppable
	sync.CountdownStopper
	Holder
	Pusher
}

type Tunnel interface {
	sync.Stoppable
	Pusher

	Type() int32
	Forward(ctx context.Context, msg ForwardMessage) error
}

type ForwardMessage interface {
	GetMod() int32
	GetSeq() int32
	GetObj() int64
	GetData() []byte
}
