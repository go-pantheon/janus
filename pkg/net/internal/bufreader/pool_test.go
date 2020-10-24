package bufreader

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSyncPool_InvalidParams(t *testing.T) {
	tests := []struct {
		name    string
		min     int
		max     int
		factor  int
		wantErr bool
	}{
		{"min <= 0", 0, 1024, 4, true},
		{"max < min", 1024, 512, 4, true},
		{"factor <= 0", 1024, 2048, 0, true},
		{"valid params", 1024, 2048, 2, false},
		{"max over 1MB", 1024, 2 * 1024 * 1024, 4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSyncPool(tt.min, tt.max, tt.factor)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAlloc(t *testing.T) {
	minSize := 128
	maxSize := 512
	factor := 2

	pool, err := NewSyncPool(minSize, maxSize, factor)
	require.NoError(t, err)

	tests := []struct {
		name      string
		size      int
		wantLen   int
		wantCap   int
		wantError bool
	}{
		{"size 0", 0, 0, 0, true},
		{"negative size", -1, 0, 0, true},
		{"smaller than min", minSize / 2, minSize / 2, minSize, false},
		{"exact min size", minSize, minSize, minSize, false},
		{"between classes", minSize + 1, minSize + 1, minSize * 2, false},
		{"max size", maxSize, maxSize, maxSize, false},
		{"exceed max", maxSize * 2, maxSize * 2, maxSize * 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := pool.Alloc(tt.size)
			assert.Equal(t, tt.wantLen, len(buf))
			assert.GreaterOrEqual(t, cap(buf), tt.wantCap)
		})
	}
}

func TestFree(t *testing.T) {
	pool, err := NewSyncPool(128, 512, 2)
	require.NoError(t, err)

	t.Run("free nil", func(t *testing.T) {
		assert.NotPanics(t, func() { pool.Free(nil) })
	})

	t.Run("reuse memory", func(t *testing.T) {
		buf1 := pool.Alloc(128)
		ptr1 := &buf1[:1][0] // get the underlying array pointer
		pool.Free(buf1)

		buf2 := pool.Alloc(128)
		ptr2 := &buf2[:1][0]
		assert.Equal(t, ptr1, ptr2, "should reuse memory")
	})

	t.Run("free over max", func(t *testing.T) {
		buf := make([]byte, 1024)
		assert.NotPanics(t, func() { pool.Free(buf) })
	})
}

func TestSize(t *testing.T) {
	minSize := 128
	maxSize := 512
	pool, err := NewSyncPool(minSize, maxSize, 2)
	require.NoError(t, err)

	tests := []struct {
		size     int
		expected int
	}{
		{0, 0},
		{minSize - 1, minSize},
		{minSize, minSize},
		{minSize + 1, minSize * 2},
		{maxSize, maxSize},
		{maxSize + 1, maxSize + 1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.expected, pool.Size(tt.size))
		})
	}
}

func TestClassSizes(t *testing.T) {
	minSize := 1024
	maxSize := 65536
	factor := 4

	pool, err := NewSyncPool(minSize, maxSize, factor)
	require.NoError(t, err)

	prev := 0
	for i, size := range pool.classesSize {
		t.Logf("Class %d: %d bytes", i, size)
		assert.True(t, size > prev, "sizes should be increasing")
		prev = size

		// check growth factor
		if i > 0 {
			expectedGrowth := minSize * ((i-1)/factor*2 + 1)
			assert.Equal(t, pool.classesSize[i-1]+expectedGrowth, size)
		}
	}
	assert.Equal(t, maxSize, pool.classesSize[len(pool.classesSize)-1])
}

func BenchmarkAllocAndFree(b *testing.B) {
	pool, _ := NewSyncPool(1024, 65536, 4)
	sizes := []int{512, 1024, 4096, 16384, 65536, 131072}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		size := sizes[i%len(sizes)]
		buf := pool.Alloc(size)
		pool.Free(buf)
	}
}

func BenchmarkAllocOnly(b *testing.B) {
	pool, _ := NewSyncPool(1024, 65536, 4)
	sizes := []int{512, 1024, 4096, 16384, 65536, 131072}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		size := sizes[i%len(sizes)]
		buf := pool.Alloc(size)
		_ = buf
	}
}

func BenchmarkParallelAllocFree(b *testing.B) {
	pool, _ := NewSyncPool(1024, 65536, 4)
	sizes := []int{512, 1024, 4096, 16384, 65536, 131072}

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			size := sizes[i%len(sizes)]
			buf := pool.Alloc(size)
			pool.Free(buf)
			i++
		}
	})
}
