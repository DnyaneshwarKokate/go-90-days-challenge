package cache_test

import (
	"testing"

	"day79/cache"
)

func TestCacheAsideHitMissAndInvalidation(t *testing.T) {
	manager := cache.NewCacheAsideManager()
	manager.SetDB("product:101", "Laptop Pro $1200")

	// Read 1: Cache Miss -> Fetches DB & Populates Cache
	val1, hit1, err1 := manager.Get("product:101")
	if err1 != nil || hit1 || val1 != "Laptop Pro $1200" {
		t.Fatalf("Expected cache miss with DB value, got hit=%t val=%s", hit1, val1)
	}

	// Read 2: Cache Hit
	val2, hit2, err2 := manager.Get("product:101")
	if err2 != nil || !hit2 || val2 != "Laptop Pro $1200" {
		t.Fatalf("Expected cache hit, got hit=%t val=%s", hit2, val2)
	}

	// Update DB record -> Invalidates Cache
	manager.Update("product:101", "Laptop Pro $1100")

	// Read 3: Cache Miss (Invalidated) -> Fetches Updated DB
	val3, hit3, err3 := manager.Get("product:101")
	if err3 != nil || hit3 || val3 != "Laptop Pro $1100" {
		t.Fatalf("Expected cache miss after invalidation, got hit=%t val=%s", hit3, val3)
	}

	hits, misses := manager.Metrics()
	if hits != 1 || misses != 2 {
		t.Fatalf("Expected 1 hit and 2 misses, got hits=%d misses=%d", hits, misses)
	}
}
