package main

import (
	"fmt"
	"time"

	"day73/distlock"
)

func main() {
	fmt.Println("=== Day 73: Distributed Locks & Lease Management ===")

	manager := distlock.NewDistributedLockManager()
	resource := "order_checkout_lock_1001"

	fmt.Println("\n--- 1. Node Alpha Acquiring Lock ---")
	leaseAlpha, err := manager.Acquire(resource, "node-alpha", 200*time.Millisecond)
	if err != nil {
		fmt.Printf("Node Alpha Error: %v\n", err)
		return
	}
	fmt.Printf("[LOCK ACQUIRED] %s\n", leaseAlpha.String())

	fmt.Println("\n--- 2. Node Beta Attempting Lock Acquisition ---")
	_, errBeta := manager.Acquire(resource, "node-beta", 200*time.Millisecond)
	if errBeta != nil {
		fmt.Printf("[LOCK REJECTED] Node Beta: %v\n", errBeta)
	}

	fmt.Println("\n--- 3. Waiting for TTL Expiration (250ms) ---")
	time.Sleep(250 * time.Millisecond)

	fmt.Println("\n--- 4. Node Beta Retrying Acquisition After Expiration ---")
	leaseBeta, errBetaRetry := manager.Acquire(resource, "node-beta", 500*time.Millisecond)
	if errBetaRetry == nil {
		fmt.Printf("[LOCK CLAIMED] %s\n", leaseBeta.String())
	}
}
