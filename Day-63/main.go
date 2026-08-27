package main

import (
	"context"
	"fmt"
	"time"

	"day63/outbox"
)

func main() {
	fmt.Println("=== Day 63: Transactional Outbox Pattern ===")

	db := outbox.NewDatabaseSimulator()

	// Create orders with outbox events in simulated DB transaction
	orders := []outbox.OrderRecord{
		{ID: "ORD-901", Customer: "Bob", Amount: 89.50, Status: "NEW", CreatedAt: time.Now()},
		{ID: "ORD-902", Customer: "Charlie", Amount: 210.00, Status: "NEW", CreatedAt: time.Now()},
	}

	for _, ord := range orders {
		payload := fmt.Sprintf(`{"order_id":"%s","customer":"%s","amount":%.2f}`, ord.ID, ord.Customer, ord.Amount)
		err := db.CreateOrderWithOutbox(ord, "OrderPlaced", payload)
		if err != nil {
			fmt.Printf("Error creating order %s: %v\n", ord.ID, err)
			return
		}
		fmt.Printf("[DB TX] Saved Order %s & Outbox Event atomically.\n", ord.ID)
	}

	// Publisher function simulating message broker (RabbitMQ / Kafka) dispatch
	brokerPublisher := func(evt *outbox.OutboxEvent) error {
		fmt.Printf("  ==> [BROKER PUBLISHED] Queue: '%s.events' | Payload: %s\n", evt.AggregateType, evt.Payload)
		return nil
	}

	processor := outbox.NewOutboxProcessor(db, brokerPublisher)

	fmt.Println("\n--- Starting Outbox Poller Execution ---")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	processedCount, err := processor.ProcessBatch(ctx)
	if err != nil {
		fmt.Printf("Processor encountered error: %v\n", err)
	}

	fmt.Printf("\n--- Batch Processing Finished ---\n")
	fmt.Printf("Successfully Published & Marked Processed: %d events\n", processedCount)
	fmt.Printf("Remaining Pending Outbox Events: %d\n", len(db.FetchPendingEvents(10)))
}
