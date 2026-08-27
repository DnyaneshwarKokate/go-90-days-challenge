package main

import (
	"fmt"

	"day72/loadbalancer"
)

func main() {
	fmt.Println("=== Day 72: Client-Side Load Balancing Algorithms ===")

	nodes := []*loadbalancer.ServerNode{
		{ID: "srv-alpha", Address: "10.0.1.10:8080", ActiveConnections: 12},
		{ID: "srv-beta", Address: "10.0.1.11:8080", ActiveConnections: 3},
		{ID: "srv-gamma", Address: "10.0.1.12:8080", ActiveConnections: 7},
	}

	lb := loadbalancer.NewLoadBalancer(nodes)

	fmt.Println("\n--- 1. Testing Round-Robin Load Balancing ---")
	for i := 1; i <= 6; i++ {
		node, _ := lb.SelectRoundRobin()
		fmt.Printf("Request #%d routed to: %s (%s)\n", i, node.ID, node.Address)
	}

	fmt.Println("\n--- 2. Testing Least-Connections Load Balancing ---")
	bestNode, _ := lb.SelectLeastConnections()
	fmt.Printf("Selected Least-Conns Node: %s (Active Connections: %d)\n",
		bestNode.ID, bestNode.ActiveConnections)
}
