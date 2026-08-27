package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"day62/saga"
)

func main() {
	fmt.Println("=== Day 62: Saga Pattern & Distributed Transactions ===")

	orchestrator := saga.NewSagaOrchestrator()

	// Step 1: Order Service
	orchestrator.AddStep(saga.SagaStep{
		Name: "CreateOrderPending",
		Execute: func(ctx context.Context) error {
			fmt.Println("  -> [Order Service] Created order #1001 with status PENDING")
			return nil
		},
		Compensate: func(ctx context.Context) error {
			fmt.Println("  <- [Order Service] Compensated order #1001 status to CANCELLED")
			return nil
		},
	})

	// Step 2: Payment Service
	orchestrator.AddStep(saga.SagaStep{
		Name: "ProcessPayment",
		Execute: func(ctx context.Context) error {
			fmt.Println("  -> [Payment Service] Processing payment of $250.00...")
			return nil
		},
		Compensate: func(ctx context.Context) error {
			fmt.Println("  <- [Payment Service] Issued refund of $250.00 to card")
			return nil
		},
	})

	// Step 3: Shipping Service (Simulate Failure)
	orchestrator.AddStep(saga.SagaStep{
		Name: "ReserveShipmentSlot",
		Execute: func(ctx context.Context) error {
			fmt.Println("  -> [Shipping Service] Contacting logistics provider...")
			return errors.New("no courier available in destination region")
		},
		Compensate: func(ctx context.Context) error {
			fmt.Println("  <- [Shipping Service] Cancelled shipment slot reservation")
			return nil
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("\n--- Initiating Order Placement Saga ---")
	result := orchestrator.Execute(ctx)

	fmt.Printf("\n--- Saga Execution Outcome ---\n")
	fmt.Printf("Success: %t\n", result.Success)
	fmt.Printf("Successful Steps: %v\n", result.SuccessfulSteps)
	fmt.Printf("Rolled Back Steps: %v\n", result.RolledBackSteps)
	if result.ExecutionError != nil {
		fmt.Printf("Failure Cause: %v\n", result.ExecutionError)
	}
}
