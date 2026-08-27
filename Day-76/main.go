package main

import (
	"fmt"
	"time"

	"day76/ratelimit"
)

func main() {
	fmt.Println("=== Day 76: API Rate Limiter & Token Bucket Engine ===")

	// Capacity: 3 tokens, Refill: 5 tokens/sec
	limiter := ratelimit.NewTokenBucketLimiter(3, 5)

	fmt.Println("\n--- Simulating Burst Requests ---")
	for i := 1; i <= 5; i++ {
		allowed := limiter.Allow()
		fmt.Printf("Request #%d: Allowed = %t | Tokens Remaining: %.2f\n",
			i, allowed, limiter.TokensAvailable())
	}

	fmt.Println("\n--- Waiting 400ms for Continuous Token Refill ---")
	time.Sleep(400 * time.Millisecond)

	for i := 6; i <= 7; i++ {
		allowed := limiter.Allow()
		fmt.Printf("Request #%d: Allowed = %t | Tokens Remaining: %.2f\n",
			i, allowed, limiter.TokensAvailable())
	}
}
