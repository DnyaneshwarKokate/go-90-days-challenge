package concurrency_test

import (
	"sync"
	"testing"

	"day87/concurrency"
)

func TestAtomicCounterRaceFree(t *testing.T) {
	counter := concurrency.NewSafeCounter()
	var wg sync.WaitGroup

	const goroutines = 100
	const incs = 100

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incs; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	if counter.Value() != int64(goroutines*incs) {
		t.Fatalf("Expected %d, got %d", goroutines*incs, counter.Value())
	}
}

func TestSyncPoolBufferReuse(t *testing.T) {
	pool := concurrency.NewBufferPool()

	buf := pool.Get()
	buf.WriteString("hello concurrency")
	if buf.String() != "hello concurrency" {
		t.Fatalf("Buffer string mismatch")
	}

	pool.Put(buf)

	buf2 := pool.Get()
	if buf2.Len() != 0 {
		t.Fatalf("Expected reset buffer, got length %d", buf2.Len())
	}
}
