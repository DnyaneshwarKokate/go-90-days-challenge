package main

import (
	"context"
	"fmt"

	"day70/system"
)

func main() {
	fmt.Println("=== Day 70: Production Microservices Integration Capstone Project ===")

	sys := system.NewEnterpriseSystem()

	fmt.Println("\n--- Executing Multi-Service Order Pipeline ---")
	order1 := system.SystemOrder{ID: "ORD-8001", Customer: "Alice Smith", Amount: 450.00}
	order2 := system.SystemOrder{ID: "ORD-8002", Customer: "Bob Vance", Amount: 1200.00}

	_ = sys.ProcessFulfillmentPipeline(context.Background(), order1)
	_ = sys.ProcessFulfillmentPipeline(context.Background(), order2)

	fmt.Println("  [SYSTEM] Processed Orders #ORD-8001 and #ORD-8002 through Saga + Outbox + CQRS + Tracing pipeline.")

	// Simulate Poison Message DLQ Routing
	sys.RoutePoisonToDLQ("MSG-MALFORMED-77", "Schema validation failure: missing mandatory user token")

	orders, outbox, readViews, dlq := sys.Summary()

	fmt.Println("\n--- Capstone Enterprise Architecture Summary ---")
	fmt.Printf("Active Domain Orders:   %d\n", orders)
	fmt.Printf("Outbox Events Queued:  %d\n", outbox)
	fmt.Printf("CQRS Read Projections: %d\n", readViews)
	fmt.Printf("Dead Letter Queue:     %d\n", dlq)
}
