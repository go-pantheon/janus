package internal

import (
	"maps"
	"sync"

	"github.com/vulcan-frame/vulcan-gate/pkg/net/conf"
)

var (
	uidWidMap = &sync.Map{}
)

type Buckets struct {
	buckets    []*Bucket
	bucketSize uint32
}

func NewBuckets(c *conf.Bucket) *Buckets {
	bs := &Buckets{
		buckets:    make([]*Bucket, c.BucketSize),
		bucketSize: uint32(c.BucketSize),
	}

	for i := 0; i < c.BucketSize; i++ {
		bs.buckets[i] = newBucket(c.WorkerSize)
	}

	return bs
}

func (bs *Buckets) Bucket(key uint64) *Bucket {
	idx := key % uint64(bs.bucketSize)
	return bs.buckets[idx]
}

func (bs *Buckets) Worker(key uint64) *Worker {
	return bs.Bucket(key).get(key)
}

func (bs *Buckets) Put(w *Worker) *Worker {
	b := bs.Bucket(w.WID())
	return b.put(w)
}

func (bs *Buckets) Del(w *Worker) {
	if b := bs.Bucket(w.WID()); b != nil {
		b.del(w)
	}
}

func (bs *Buckets) Walk(f func(w *Worker) bool) {
	for _, b := range bs.buckets {
		b.walk(f)
	}
}

func (bs *Buckets) GetByUID(uid int64) *Worker {
	wid, ok := uidWidMap.Load(uid)
	if !ok {
		return nil
	}
	return bs.Worker(wid.(uint64))
}

func (bs *Buckets) GetByUIDs(uids []int64) []*Worker {
	workers := make([]*Worker, 0, len(uids))
	for _, uid := range uids {
		if wid, ok := uidWidMap.Load(uid); ok {
			workers = append(workers, bs.Worker(wid.(uint64)))
		}
	}
	return workers
}

type Bucket struct {
	sync.RWMutex

	workers map[uint64]*Worker
}

func newBucket(workerSize int) (b *Bucket) {
	b = &Bucket{}
	b.workers = make(map[uint64]*Worker, workerSize)
	return
}

func (b *Bucket) put(w *Worker) (old *Worker) {
	b.Lock()
	defer b.Unlock()

	old = b.workers[w.WID()]
	b.workers[w.WID()] = w
	uidWidMap.Store(w.UID(), w.WID())
	return
}

func (b *Bucket) del(dw *Worker) {
	var (
		ok bool
		w  *Worker
	)

	b.Lock()
	defer b.Unlock()

	if w, ok = b.workers[dw.WID()]; ok {
		if w == dw {
			delete(b.workers, w.WID())
			uidWidMap.Delete(w.UID())
		}
	}
}

func (b *Bucket) get(key uint64) (w *Worker) {
	b.RLock()
	defer b.RUnlock()

	w = b.workers[key]
	return
}

func (b *Bucket) walk(f func(w *Worker) (continued bool)) {
	snapshot := b.snapshot()
	for _, w := range snapshot {
		if !f(w) {
			break
		}
	}
}

func (b *Bucket) snapshot() map[uint64]*Worker {
	b.RLock()
	defer b.RUnlock()

	workers := make(map[uint64]*Worker, len(b.workers))
	maps.Copy(workers, b.workers)
	return workers
}
