package system_test

import (
	"context"
	"testing"

	"day70/system"
)

func TestEnterpriseSystemPipeline(t *testing.T) {
	sys := system.NewEnterpriseSystem()

	order := system.SystemOrder{
		ID:       "ORD-9990",
		Customer: "Grace Hopper",
		Amount:   1999.00,
	}

	if err := sys.ProcessFulfillmentPipeline(context.Background(), order); err != nil {
		t.Fatalf("ProcessFulfillmentPipeline failed: %v", err)
	}

	sys.RoutePoisonToDLQ("msg-poison-001", "Payload corrupted")

	orders, outbox, readViews, dlq := sys.Summary()

	if orders != 1 {
		t.Fatalf("Expected 1 order, got %d", orders)
	}
	if outbox != 1 {
		t.Fatalf("Expected 1 outbox event, got %d", outbox)
	}
	if readViews != 1 {
		t.Fatalf("Expected 1 read view projection, got %d", readViews)
	}
	if dlq != 1 {
		t.Fatalf("Expected 1 DLQ entry, got %d", dlq)
	}
}
