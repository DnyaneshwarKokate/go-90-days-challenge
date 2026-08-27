package discovery_test

import (
	"testing"
	"time"

	"day71/discovery"
)

func TestServiceRegistrationAndHeartbeat(t *testing.T) {
	registry := discovery.NewServiceRegistry()

	inst := discovery.ServiceInstance{
		ID:      "user-svc-node-1",
		Name:    "user-service",
		Address: "10.0.1.20",
		Port:    8080,
	}

	if err := registry.Register(inst); err != nil {
		t.Fatalf("Failed to register instance: %v", err)
	}

	healthy := registry.GetHealthyInstances("user-service")
	if len(healthy) != 1 || healthy[0].Address != "10.0.1.20" {
		t.Fatalf("Expected 1 healthy user-service instance, got %v", healthy)
	}

	// Evict instances with 50ms TTL threshold
	time.Sleep(100 * time.Millisecond)
	evicted := registry.EvictUnhealthy(50 * time.Millisecond)

	if evicted != 1 {
		t.Fatalf("Expected 1 evicted stale instance, got %d", evicted)
	}

	remaining := registry.GetHealthyInstances("user-service")
	if len(remaining) != 0 {
		t.Fatalf("Expected 0 healthy instances after eviction, got %d", len(remaining))
	}
}
