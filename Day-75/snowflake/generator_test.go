package snowflake_test

import (
	"sync"
	"testing"

	"day75/snowflake"
)

func TestSnowflakeUniquenessAndMonotonicity(t *testing.T) {
	node, err := snowflake.NewSnowflakeNode(7)
	if err != nil {
		t.Fatalf("Failed to create Snowflake node: %v", err)
	}

	const count = 1000
	idMap := make(map[uint64]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id, err := node.NextID()
			if err != nil {
				t.Errorf("NextID error: %v", err)
				return
			}

			mu.Lock()
			if idMap[id] {
				t.Errorf("Duplicate ID generated: %d", id)
			}
			idMap[id] = true
			mu.Unlock()
		}()
	}

	wg.Wait()

	if len(idMap) != count {
		t.Fatalf("Expected %d unique IDs, got %d", count, len(idMap))
	}
}
