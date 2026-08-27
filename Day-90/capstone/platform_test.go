package capstone_test

import (
	"context"
	"testing"

	"day90/capstone"
)

func TestFinalProductionPlatform(t *testing.T) {
	platform := capstone.NewProductionPlatform()

	order, err := platform.ExecuteProductionOrderPipeline(context.Background(), "Dnyaneshwar Kokate", "MacBook Pro M3", 2499.00)
	if err != nil {
		t.Fatalf("Pipeline execution failed: %v", err)
	}

	if order.Status != "FULFILLED_PRODUCTION" || order.ID == 0 || order.TraceID == "" {
		t.Fatalf("Invalid order payload: %+v", order)
	}

	report := platform.SystemReport()
	if report == "" {
		t.Fatalf("Expected non-empty system report")
	}
}
