package main

import (
	"fmt"

	"day79/cache"
)

func main() {
	fmt.Println("=== Day 79: Cache-Aside Pattern & Multi-Region Sync ===")

	manager := cache.NewCacheAsideManager()
	manager.SetDB("user_profile_55", `{"name":"Carol","tier":"GOLD"}`)

	fmt.Println("\n--- 1. First Read Request (Cache Miss) ---")
	val1, hit1, _ := manager.Get("user_profile_55")
	fmt.Printf("Hit: %t | Payload: %s\n", hit1, val1)

	fmt.Println("\n--- 2. Second Read Request (Cache Hit) ---")
	val2, hit2, _ := manager.Get("user_profile_55")
	fmt.Printf("Hit: %t | Payload: %s\n", hit2, val2)

	fmt.Println("\n--- 3. Updating User Profile (DB Write + Cache Invalidation) ---")
	manager.Update("user_profile_55", `{"name":"Carol","tier":"PLATINUM"}`)

	fmt.Println("\n--- 4. Third Read Request After Invalidation ---")
	val3, hit3, _ := manager.Get("user_profile_55")
	fmt.Printf("Hit: %t | Payload: %s\n", hit3, val3)

	hits, misses := manager.Metrics()
	fmt.Printf("\n--- Cache Performance Telemetry ---\nTotal Cache Hits: %d | Total Cache Misses: %d\n", hits, misses)
}
