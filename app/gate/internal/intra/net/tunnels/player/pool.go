package player

import (
	"sync"

	"github.com/go-pantheon/fabrica-net/xnet"
	intrav1 "github.com/go-pantheon/janus/gen/api/server/player/intra/v1"
)

var forwardPool = newPool()

func get() *intrav1.TunnelRequest {
	return forwardPool.get()
}

func put(msg xnet.TunnelMessage) {
	forwardPool.put(msg.(*intrav1.TunnelRequest))
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
