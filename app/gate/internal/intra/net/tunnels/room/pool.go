package room

import (
	"sync"

	intrav1 "github.com/go-pantheon/janus/gen/api/server/room/intra/v1"
)

type requestPool struct {
	pool sync.Pool
}

func newRequestPool() *requestPool {
	return &requestPool{
		pool: sync.Pool{
			New: func() interface{} {
				return new(intrav1.TunnelRequest)
			},
		},
	}
}

func (p *requestPool) get() *intrav1.TunnelRequest {
	return p.pool.Get().(*intrav1.TunnelRequest)
}

func (p *requestPool) put(msg *intrav1.TunnelRequest) {
	if msg == nil {
		return
	}
	msg.Reset()
	p.pool.Put(msg)
}

var globalPool = newRequestPool()

func getRequest() *intrav1.TunnelRequest {
	return globalPool.get()
}

func putRequest(msg *intrav1.TunnelRequest) {
	globalPool.put(msg)
}
