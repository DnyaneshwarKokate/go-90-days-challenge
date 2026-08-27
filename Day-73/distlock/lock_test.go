package distlock_test

import (
	"errors"
	"testing"
	"time"

	"day73/distlock"
)

func TestDistributedLockAcquireAndRelease(t *testing.T) {
	manager := distlock.NewDistributedLockManager()

	// Owner A acquires lock
	lease, err := manager.Acquire("order-101-payout", "node-A", 500*time.Millisecond)
	if err != nil {
		t.Fatalf("Failed to acquire lock: %v", err)
	}

	if lease.Owner != "node-A" {
		t.Fatalf("Expected owner node-A, got %s", lease.Owner)
	}

	// Owner B attempts lock acquisition (Must fail)
	_, errB := manager.Acquire("order-101-payout", "node-B", 500*time.Millisecond)
	if !errors.Is(errB, distlock.ErrLockAlreadyAcquired) {
		t.Fatalf("Expected ErrLockAlreadyAcquired, got %v", errB)
	}

	// Owner A releases lock
	if err := manager.Release("order-101-payout", "node-A"); err != nil {
		t.Fatalf("Failed to release lock: %v", err)
	}

	// Owner B can now acquire lock
	leaseB, errB2 := manager.Acquire("order-101-payout", "node-B", 500*time.Millisecond)
	if errB2 != nil || leaseB.Owner != "node-B" {
		t.Fatalf("Owner B failed to acquire freed lock: %v", errB2)
	}
}

func TestDistributedLockTTLExpiration(t *testing.T) {
	manager := distlock.NewDistributedLockManager()

	// Short TTL of 50ms
	_, err := manager.Acquire("inventory-seat-9", "node-X", 50*time.Millisecond)
	if err != nil {
		t.Fatalf("Acquire failed: %v", err)
	}

	// Wait for TTL to expire
	time.Sleep(70 * time.Millisecond)

	// Node-Y acquires expired lock
	leaseY, errY := manager.Acquire("inventory-seat-9", "node-Y", 500*time.Millisecond)
	if errY != nil || leaseY.Owner != "node-Y" {
		t.Fatalf("Node-Y failed to claim expired lease: %v", errY)
	}
}
