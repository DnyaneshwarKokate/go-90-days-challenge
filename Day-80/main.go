package main

import (
	"context"
	"fmt"

	"day80/system"
)

func main() {
	fmt.Println("=== Day 80: High-Scale Distributed Microservices E-Commerce Platform ===")

	platform := system.NewMicroservicesPlatform(1)

	fmt.Println("\n--- Processing Enterprise Orders across Distributed Architecture ---")
	orders := []system.PlatformOrderRequest{
		{Customer: "Alice Johnson", Amount: 299.99},
		{Customer: "Bob Miller", Amount: 750.50},
		{Customer: "Charlie Davis", Amount: 1400.00},
	}

	for i, req := range orders {
		resp, _ := platform.SubmitOrder(context.Background(), req)
		fmt.Printf("Order #%d: Customer: %-15s | ID: %d | Target Shard: %s | Status: %s\n",
			i+1, req.Customer, resp.ID, resp.ShardNode, resp.Status)
	}

	fmt.Printf("\n--- Capstone Platform Architecture Summary ---\nTotal Distributed Transactions Processed: %d\n", platform.TotalProcessed())
}
