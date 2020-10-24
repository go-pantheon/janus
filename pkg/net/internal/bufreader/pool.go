package bufreader

import (
	"errors"
	"sync"
)

type Pool interface {
	Alloc(int) []byte
	Free([]byte)
}

var _ Pool = (*SyncPool)(nil)

var (
	// ErrInvalidSize when the input size parameters are invalid
	ErrInvalidSize = errors.New("invalid size parameters")
)

// SyncPool is a sync.Pool based slab allocation memory pool
type SyncPool struct {
	classes     []sync.Pool // different size memory pools
	classesSize []int       // size of each memory pool
	minSize     int         // minimum chunk size
	maxSize     int         // maximum chunk size
	sizeLookup  []uint8     // fast lookup table from size to class index
}

// NewSyncPool creates a sync.Pool based slab allocation memory pool
// minSize: minimum chunk size
// maxSize: maximum chunk size
// factor: factor for controlling chunk size growth
func NewSyncPool(minSize, maxSize, factor int) (*SyncPool, error) {
	if minSize <= 0 || maxSize < minSize || factor <= 0 {
		return nil, ErrInvalidSize
	}

	n := 0
	curSize := minSize
	for chunkSize := minSize; chunkSize <= maxSize; chunkSize += minSize * (n/factor*2 + 1) {
		curSize = chunkSize
		n++
	}
	if curSize < maxSize {
		n++
	}

	pool := &SyncPool{
		classes:     make([]sync.Pool, n),
		classesSize: make([]int, n),
		minSize:     minSize,
		maxSize:     maxSize,
		sizeLookup:  make([]uint8, maxSize+1),
	}

	chunkSize := minSize
	for k := range pool.classes {
		chunkSize += minSize * (k/factor*2 + 1)
		size := min(chunkSize, maxSize)
		pool.classesSize[k] = size
		pool.classes[k].New = func() interface{} {
			buf := make([]byte, size)
			return &buf
		}

		// Fill lookup table for size to class index mapping
		start := 0
		if k > 0 {
			start = pool.classesSize[k-1]
		}
		for i := start; i <= size; i++ {
			pool.sizeLookup[i] = uint8(k)
		}
	}

	return pool, nil
}

// Alloc allocates a []byte from the internal slab class
// if there is no free block, it will create a new one
func (pool *SyncPool) Alloc(size int) []byte {
	if size <= 0 {
		return make([]byte, 0)
	}

	if size <= pool.maxSize {
		classIndex := pool.sizeLookup[size]
		mem := pool.classes[classIndex].Get().(*[]byte)
		return (*mem)[:size]
	}

	return make([]byte, size)
}

// Free frees the []byte allocated from Pool.Alloc
func (pool *SyncPool) Free(mem []byte) {
	if mem == nil {
		return
	}

	size := cap(mem)
	if size <= pool.maxSize {
		classIndex := pool.sizeLookup[size]
		// reset the slice to avoid memory leaks
		mem = mem[:cap(mem)]
		pool.classes[classIndex].Put(&mem)
	}
}

// Size returns the actual allocation size for a given size
func (pool *SyncPool) Size(size int) int {
	if size <= 0 || size > pool.maxSize {
		return size
	}
	return pool.classesSize[pool.sizeLookup[size]]
}
