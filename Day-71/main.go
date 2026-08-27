package main

import (
	"fmt"
	"time"

	"day71/discovery"
)

func main() {
	fmt.Println("=== Day 71: Service Discovery & Dynamic Health Registry ===")

	registry := discovery.NewServiceRegistry()

	// Register 3 Order Microservice instances
	inst1 := discovery.ServiceInstance{ID: "order-node-1", Name: "order-service", Address: "192.168.1.10", Port: 8081}
	inst2 := discovery.ServiceInstance{ID: "order-node-2", Name: "order-service", Address: "192.168.1.11", Port: 8081}
	inst3 := discovery.ServiceInstance{ID: "order-node-3", Name: "order-service", Address: "192.168.1.12", Port: 8081}

	_ = registry.Register(inst1)
	_ = registry.Register(inst2)
	_ = registry.Register(inst3)

	fmt.Println("\n--- Querying Active Registered Nodes ---")
	nodes := registry.GetHealthyInstances("order-service")
	for _, n := range nodes {
		fmt.Printf("  [ACTIVE NODE] ID: %s | URL: http://%s:%d | Status: %s\n", n.ID, n.Address, n.Port, n.Status)
	}

	fmt.Println("\n--- Simulating Heartbeat Refresh & Stale Eviction ---")
	// Send heartbeat for node-1 and node-2, leave node-3 stale
	time.Sleep(100 * time.Millisecond)
	_ = registry.Heartbeat("order-node-1")
	_ = registry.Heartbeat("order-node-2")

	// Evict nodes silent for > 50ms
	evicted := registry.EvictUnhealthy(50 * time.Millisecond)
	fmt.Printf("Stale Instances Evicted: %d\n", evicted)

	activeAfterEviction := registry.GetHealthyInstances("order-service")
	fmt.Printf("Remaining Healthy Instances for 'order-service': %d\n", len(activeAfterEviction))
}
