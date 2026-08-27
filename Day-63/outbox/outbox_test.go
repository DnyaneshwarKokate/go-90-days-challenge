package outbox_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"day63/outbox"
)

func TestTransactionalOutboxFlow(t *testing.T) {
	db := outbox.NewDatabaseSimulator()

	order := outbox.OrderRecord{
		ID:        "ORD-5501",
		Customer:  "Alice",
		Amount:    149.99,
		Status:    "PLACED",
		CreatedAt: time.Now(),
	}

	err := db.CreateOrderWithOutbox(order, "OrderCreated", `{"order_id":"ORD-5501","amount":149.99}`)
	if err != nil {
		t.Fatalf("Failed to create order with outbox: %v", err)
	}

	pending := db.FetchPendingEvents(10)
	if len(pending) != 1 {
		t.Fatalf("Expected 1 pending event, got %d", len(pending))
	}

	var publishedEvents []string
	var mu sync.Mutex

	publisher := func(evt *outbox.OutboxEvent) error {
		mu.Lock()
		defer mu.Unlock()
		publishedEvents = append(publishedEvents, evt.ID)
		return nil
	}

	processor := outbox.NewOutboxProcessor(db, publisher)
	count, err := processor.ProcessBatch(context.Background())

	if err != nil {
		t.Fatalf("ProcessBatch failed: %v", err)
	}

	if count != 1 {
		t.Fatalf("Expected 1 event processed, got %d", count)
	}

	remainingPending := db.FetchPendingEvents(10)
	if len(remainingPending) != 0 {
		t.Fatalf("Expected 0 pending events after processor run, got %d", len(remainingPending))
	}
}
