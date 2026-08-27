package concurrency

import (
	"bytes"
	"sync"
	"sync/atomic"
)

// SafeCounter uses atomic operations to guarantee data-race free increments.
type SafeCounter struct {
	value int64
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

func (c *SafeCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *SafeCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

// BufferPool uses sync.Pool to recycle memory allocations across concurrent goroutines.
type BufferPool struct {
	pool sync.Pool
}

func NewBufferPool() *BufferPool {
	return &BufferPool{
		pool: sync.Pool{
			New: func() any {
				return new(bytes.Buffer)
			},
		},
	}
}

func (p *BufferPool) Get() *bytes.Buffer {
	buf := p.pool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func (p *BufferPool) Put(buf *bytes.Buffer) {
	p.pool.Put(buf)
}
