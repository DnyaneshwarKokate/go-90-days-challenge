package bulkhead

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
)

var ErrBulkheadFull = errors.New("bulkhead capacity reached: request rejected to prevent cascade failure")

// Bulkhead isolates concurrency limits per service dependency.
type Bulkhead struct {
	maxCapacity int32
	activeCount int32
	semaphore   chan struct{}
}

// NewBulkhead initializes a bulkhead instance with a max concurrent capacity.
func NewBulkhead(maxCapacity int) *Bulkhead {
	return &Bulkhead{
		maxCapacity: int32(maxCapacity),
		semaphore:   make(chan struct{}, maxCapacity),
	}
}

// Execute runs the provided task if bulkhead capacity is available.
// Rejects immediately with ErrBulkheadFull if max concurrent executions are active.
func (b *Bulkhead) Execute(ctx context.Context, task func(ctx context.Context) error) error {
	select {
	case b.semaphore <- struct{}{}:
		// Slot acquired
		atomic.AddInt32(&b.activeCount, 1)
		defer func() {
			atomic.AddInt32(&b.activeCount, -1)
			<-b.semaphore
		}()

		return task(ctx)
	default:
		// Bulkhead full, reject fast
		return ErrBulkheadFull
	}
}

// ActiveTasks returns current in-flight concurrent execution count.
func (b *Bulkhead) ActiveTasks() int {
	return int(atomic.LoadInt32(&b.activeCount))
}

// MaxCapacity returns configured bulkhead limit.
func (b *Bulkhead) MaxCapacity() int {
	return int(b.maxCapacity)
}

// String returns formatted bulkhead status metrics.
func (b *Bulkhead) String() string {
	return fmt.Sprintf("Bulkhead Status [Active: %d / Max: %d]", b.ActiveTasks(), b.MaxCapacity())
}
