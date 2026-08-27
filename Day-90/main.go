package main

import (
	"context"
	"fmt"

	"day90/capstone"
)

func main() {
	fmt.Println("🎉 ========================================================= 🎉")
	fmt.Println("🚀 DAY 90: FINAL PRODUCTION DISTRIBUTED MICROSERVICES PLATFORM 🚀")
	fmt.Println("🎉 ========================================================= 🎉")

	platform := capstone.NewProductionPlatform()

	fmt.Println("\n--- Executing Final Capstone Multi-Service Order Transaction ---")
	order, err := platform.ExecuteProductionOrderPipeline(context.Background(),
		"Dnyaneshwar Kokate", "Enterprise Go Backend Architecture", 9999.00)

	if err != nil {
		fmt.Printf("Execution Error: %v\n", err)
		return
	}

	fmt.Printf("[FINAL ORDER] ID: %d | User: %s | Item: %s | Amount: $%.2f | Status: %s\n",
		order.ID, order.User, order.Item, order.Amount, order.Status)
	fmt.Printf("[OTEL TRACE]  W3C TraceID: %s\n", order.TraceID)

	fmt.Println("\n--- Final Challenge Completion Status ---")
	fmt.Println(platform.SystemReport())
	fmt.Println("🏆 Congratulations! Go 90 Days Challenge 100% Completed Successfully! 🏆")
}
