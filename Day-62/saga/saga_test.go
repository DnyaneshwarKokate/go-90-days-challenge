package saga_test

import (
	"context"
	"errors"
	"testing"

	"day62/saga"
)

func TestSagaExecutionSuccess(t *testing.T) {
	orchestrator := saga.NewSagaOrchestrator()

	orderCreated := false
	paymentProcessed := false

	orchestrator.AddStep(saga.SagaStep{
		Name: "CreateOrder",
		Execute: func(ctx context.Context) error {
			orderCreated = true
			return nil
		},
		Compensate: func(ctx context.Context) error {
			orderCreated = false
			return nil
		},
	})

	orchestrator.AddStep(saga.SagaStep{
		Name: "ProcessPayment",
		Execute: func(ctx context.Context) error {
			paymentProcessed = true
			return nil
		},
		Compensate: func(ctx context.Context) error {
			paymentProcessed = false
			return nil
		},
	})

	result := orchestrator.Execute(context.Background())

	if !result.Success {
		t.Fatalf("Expected saga to succeed, got error: %v", result.ExecutionError)
	}

	if !orderCreated || !paymentProcessed {
		t.Fatalf("Expected order and payment flags to be true")
	}

	if len(result.SuccessfulSteps) != 2 {
		t.Fatalf("Expected 2 successful steps, got %d", len(result.SuccessfulSteps))
	}
}

func TestSagaRollbackOnFailure(t *testing.T) {
	orchestrator := saga.NewSagaOrchestrator()

	inventoryReserved := false

	orchestrator.AddStep(saga.SagaStep{
		Name: "ReserveInventory",
		Execute: func(ctx context.Context) error {
			inventoryReserved = true
			return nil
		},
		Compensate: func(ctx context.Context) error {
			inventoryReserved = false
			return nil
		},
	})

	orchestrator.AddStep(saga.SagaStep{
		Name: "ChargeCard",
		Execute: func(ctx context.Context) error {
			return errors.New("insufficient funds")
		},
		Compensate: func(ctx context.Context) error {
			return nil
		},
	})

	result := orchestrator.Execute(context.Background())

	if result.Success {
		t.Fatalf("Expected saga to fail due to card charge error")
	}

	if inventoryReserved {
		t.Fatalf("Expected inventory step to be compensated (set back to false)")
	}

	if len(result.RolledBackSteps) != 1 || result.RolledBackSteps[0] != "ReserveInventory" {
		t.Fatalf("Expected 'ReserveInventory' in rolled back steps, got %v", result.RolledBackSteps)
	}
}
