package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"day65/bulkhead"
)

func main() {
	fmt.Println("=== Day 65: Bulkhead Pattern & Service Resilience Isolation ===")

	// Separate bulkheads for independent downstream dependencies
	paymentBulkhead := bulkhead.NewBulkhead(3)
	recommendationBulkhead := bulkhead.NewBulkhead(1)

	var wg sync.WaitGroup

	fmt.Println("\n--- Simulating High Concurrency Requests ---")

	// Simulate Payment Service Calls (Max 3 concurrent)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		reqID := i
		go func() {
			defer wg.Done()
			err := paymentBulkhead.Execute(context.Background(), func(ctx context.Context) error {
				fmt.Printf("[PAYMENT SERVICE] Processing Request #%d (%s)\n", reqID, paymentBulkhead.String())
				time.Sleep(100 * time.Millisecond)
				return nil
			})

			if err != nil {
				fmt.Printf("[PAYMENT SERVICE] Request #%d REJECTED: %v\n", reqID, err)
			}
		}()
	}

	// Simulate Recommendation Service Calls (Max 1 concurrent)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		reqID := i
		go func() {
			defer wg.Done()
			err := recommendationBulkhead.Execute(context.Background(), func(ctx context.Context) error {
				fmt.Printf("  ==> [RECOMMENDATION SERVICE] Executing Request #%d (%s)\n", reqID, recommendationBulkhead.String())
				time.Sleep(150 * time.Millisecond)
				return nil
			})

			if err != nil {
				fmt.Printf("  ==> [RECOMMENDATION SERVICE] Request #%d REJECTED: %v\n", reqID, err)
			}
		}()
	}

	wg.Wait()
	fmt.Println("\n--- All Concurrent Requests Processed ---")
	fmt.Printf("Payment Bulkhead Final State: %s\n", paymentBulkhead.String())
	fmt.Printf("Recommendation Bulkhead Final State: %s\n", recommendationBulkhead.String())
}
