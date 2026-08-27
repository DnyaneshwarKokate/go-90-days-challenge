package loadbalancer_test

import (
	"testing"

	"day72/loadbalancer"
)

func TestRoundRobinSelection(t *testing.T) {
	node1 := &loadbalancer.ServerNode{ID: "node-1", Address: "10.0.0.1"}
	node2 := &loadbalancer.ServerNode{ID: "node-2", Address: "10.0.0.2"}

	lb := loadbalancer.NewLoadBalancer([]*loadbalancer.ServerNode{node1, node2})

	n1, _ := lb.SelectRoundRobin()
	n2, _ := lb.SelectRoundRobin()
	n3, _ := lb.SelectRoundRobin()

	if n1.ID != "node-1" || n2.ID != "node-2" || n3.ID != "node-1" {
		t.Fatalf("Round Robin distribution failed: got %s, %s, %s", n1.ID, n2.ID, n3.ID)
	}
}

func TestLeastConnectionsSelection(t *testing.T) {
	node1 := &loadbalancer.ServerNode{ID: "node-1", Address: "10.0.0.1", ActiveConnections: 10}
	node2 := &loadbalancer.ServerNode{ID: "node-2", Address: "10.0.0.2", ActiveConnections: 2}

	lb := loadbalancer.NewLoadBalancer([]*loadbalancer.ServerNode{node1, node2})

	best, err := lb.SelectLeastConnections()
	if err != nil {
		t.Fatalf("Least connections selection error: %v", err)
	}

	if best.ID != "node-2" {
		t.Fatalf("Expected node-2 with 2 connections, got %s", best.ID)
	}
}
