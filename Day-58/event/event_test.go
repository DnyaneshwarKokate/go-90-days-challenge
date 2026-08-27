package event

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestEventBusPublishAndOutbox(t *testing.T) {
	bus := NewEventBus()

	var counter int32

	bus.Subscribe("OrderCreated", func(ctx context.Context, event Event) error {
		atomic.AddInt32(&counter, 1)
		return nil
	})

	bus.Subscribe("OrderCreated", func(ctx context.Context, event Event) error {
		atomic.AddInt32(&counter, 1)
		return nil
	})

	evt := Event{
		ID:          "evt_101",
		Type:        "OrderCreated",
		AggregateID: "ord_555",
		Data:        map[string]interface{}{"amount": 199.99},
		Timestamp:   time.Now(),
	}

	err := bus.Publish(context.Background(), evt)
	if err != nil {
		t.Fatalf("Failed to publish event: %v", err)
	}

	if atomic.LoadInt32(&counter) != 2 {
		t.Errorf("Expected 2 handler invocations, got %d", counter)
	}

	outbox := bus.GetOutboxEvents()
	if len(outbox) != 1 {
		t.Fatalf("Expected 1 event in Transactional Outbox, got %d", len(outbox))
	}

	if outbox[0].ID != "evt_101" {
		t.Errorf("Expected outbox event ID evt_101, got %s", outbox[0].ID)
	}
}
