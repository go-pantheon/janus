package pool

import (
	"sync"

	clipkt "github.com/vulcan-frame/vulcan-gate/gen/api/client/packet"
)

type packetPool struct {
	pool sync.Pool
}

func newPacketPool() *packetPool {
	return &packetPool{
		pool: sync.Pool{
			New: func() interface{} {
				return new(clipkt.Packet)
			},
		},
	}
}

func (p *packetPool) get() *clipkt.Packet {
	return p.pool.Get().(*clipkt.Packet)
}

func (p *packetPool) put(packet *clipkt.Packet) {
	if packet == nil {
		return
	}
	packet.Reset()
	p.pool.Put(packet)
}

var globalPool = newPacketPool()

func GetPacket() *clipkt.Packet {
	return globalPool.get()
}

func PutPacket(p *clipkt.Packet) {
	globalPool.put(p)
}
