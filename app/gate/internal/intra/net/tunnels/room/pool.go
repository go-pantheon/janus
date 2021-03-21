package room

import (
	"sync"

	intrav1 "github.com/vulcan-frame/vulcan-gate/gen/api/server/room/intra/v1"
)

type messagePool struct {
	pool sync.Pool
}

func newMessagePool() *messagePool {
	return &messagePool{
		pool: sync.Pool{
			New: func() interface{} {
				return new(intrav1.Message)
			},
		},
	}
}

func (p *messagePool) get() *intrav1.Message {
	return p.pool.Get().(*intrav1.Message)
}

func (p *messagePool) put(msg *intrav1.Message) {
	if msg == nil {
		return
	}
	msg.Reset()
	p.pool.Put(msg)
}

var globalPool = newMessagePool()

func getMessage() *intrav1.Message {
	return globalPool.get()
}

func putMessage(msg *intrav1.Message) {
	globalPool.put(msg)
}
