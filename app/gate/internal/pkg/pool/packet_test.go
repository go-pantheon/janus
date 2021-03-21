package pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	clipkt "github.com/vulcan-frame/vulcan-gate/gen/api/client/packet"
	"github.com/vulcan-frame/vulcan-pkg-tool/bitmap"
)

func TestProfile(t *testing.T) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	startAlloc := memStats.Alloc

	count := 1_000_002
	pktChecker := bitmap.NewBitmap(count + 1)

	wg := &sync.WaitGroup{}
	pktQueue := make(chan *clipkt.Packet, count)

	for i := 1; i <= count; i++ {
		wg.Add(1)
		go func(i int32) {
			pkt := GetPacket()
			pkt.Index = i
			pktQueue <- pkt
		}(int32(i))
	}

	for i := 1; i <= 10; i++ {
		go func(i int) {
			for j := i; j <= count; j += 10 {
				pkt := <-pktQueue
				pktChecker.Set(int(pkt.Index))
				PutPacket(pkt)
				wg.Done()
			}
		}(i)
	}

	wg.Wait()

	runtime.ReadMemStats(&memStats)
	endAlloc := memStats.Alloc

	assert.Equal(t, count, pktChecker.Count())

	cpu := runtime.NumGoroutine()
	t.Logf("Memory start: %s end: %s allocated: %s", formatBytes(startAlloc), formatBytes(endAlloc), formatBytes(endAlloc-startAlloc))
	t.Logf("Number of goroutines: %d", cpu)
}

func formatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
