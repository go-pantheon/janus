package room

import (
	"sync"

	intrav1 "github.com/go-pantheon/janus/gen/api/server/room/intra/v1"
)

var forwardPool = newPool()

func get() *intrav1.TunnelRequest {
	return forwardPool.get()
}

func put(msg *intrav1.TunnelRequest) {
	forwardPool.put(msg)
}

type pool struct {
	pool sync.Pool
}

func newPool() *pool {
	return &pool{
		pool: sync.Pool{
			New: func() any {
				return new(intrav1.TunnelRequest)
			},
		},
	}
}

func (p *pool) get() *intrav1.TunnelRequest {
	return p.pool.Get().(*intrav1.TunnelRequest)
}

func (p *pool) put(msg *intrav1.TunnelRequest) {
	if msg == nil {
		return
	}

	msg.Reset()
	p.pool.Put(msg)
}
