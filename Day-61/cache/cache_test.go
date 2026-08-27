package cache

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestDistributedCacheSingleflightCoalescing(t *testing.T) {
	c := NewDistributedCache()

	var dbQueryCount int32

	fetcher := func() (string, error) {
		atomic.AddInt32(&dbQueryCount, 1)
		time.Sleep(50 * time.Millisecond) // Simulate slow database query
		return "cached_product_details", nil
	}

	// 10 concurrent requests for the exact same uncached key
	var wg sync.WaitGroup
	results := make([]string, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		idx := i
		go func() {
			defer wg.Done()
			val, err := c.GetOrFetch("prod_999", fetcher, 5*time.Second)
			if err != nil {
				t.Errorf("GetOrFetch failed: %v", err)
			}
			results[idx] = val
		}()
	}

	wg.Wait()

	if atomic.LoadInt32(&dbQueryCount) != 1 {
		t.Errorf("Expected exactly 1 DB fetch call due to singleflight coalescing, got %d", dbQueryCount)
	}

	for _, res := range results {
		if res != "cached_product_details" {
			t.Errorf("Unexpected result: %s", res)
		}
	}
}
