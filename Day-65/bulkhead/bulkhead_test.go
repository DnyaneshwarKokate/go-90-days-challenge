package bulkhead_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"day65/bulkhead"
)

func TestBulkheadCapacityLimiting(t *testing.T) {
	// Bulkhead with capacity of 2
	b := bulkhead.NewBulkhead(2)

	var wg sync.WaitGroup
	blockCh := make(chan struct{})

	// Launch 2 background tasks occupying full capacity
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = b.Execute(context.Background(), func(ctx context.Context) error {
				<-blockCh
				return nil
			})
		}()
	}

	// Give goroutines time to acquire slots
	time.Sleep(50 * time.Millisecond)

	if b.ActiveTasks() != 2 {
		t.Fatalf("Expected 2 active tasks, got %d", b.ActiveTasks())
	}

	// 3rd task must be rejected fast
	err := b.Execute(context.Background(), func(ctx context.Context) error {
		return nil
	})

	if !errors.Is(err, bulkhead.ErrBulkheadFull) {
		t.Fatalf("Expected ErrBulkheadFull, got %v", err)
	}

	// Unblock tasks
	close(blockCh)
	wg.Wait()

	if b.ActiveTasks() != 0 {
		t.Fatalf("Expected 0 active tasks after completion, got %d", b.ActiveTasks())
	}
}
