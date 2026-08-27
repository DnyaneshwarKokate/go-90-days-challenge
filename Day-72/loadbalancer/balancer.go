package loadbalancer

import (
	"errors"
	"sync"
	"sync/atomic"
)

// ServerNode represents a microservice instance node.
type ServerNode struct {
	ID                string
	Address           string
	ActiveConnections int64
	Weight            int
}

// LoadBalancer manages node selection using Round-Robin and Least-Connections strategy.
type LoadBalancer struct {
	mu           sync.RWMutex
	nodes        []*ServerNode
	rrCounter    uint64
}

// NewLoadBalancer initializes load balancer with server nodes.
func NewLoadBalancer(nodes []*ServerNode) *LoadBalancer {
	return &LoadBalancer{
		nodes: nodes,
	}
}

// SelectRoundRobin picks the next node sequentially.
func (lb *LoadBalancer) SelectRoundRobin() (*ServerNode, error) {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	if len(lb.nodes) == 0 {
		return nil, errors.New("no healthy server nodes available")
	}

	idx := atomic.AddUint64(&lb.rrCounter, 1) - 1
	selected := lb.nodes[idx%uint64(len(lb.nodes))]
	return selected, nil
}

// SelectLeastConnections picks the server node with the lowest active connection count.
func (lb *LoadBalancer) SelectLeastConnections() (*ServerNode, error) {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	if len(lb.nodes) == 0 {
		return nil, errors.New("no healthy server nodes available")
	}

	var bestNode *ServerNode
	minConn := int64(1<<63 - 1)

	for _, node := range lb.nodes {
		conns := atomic.LoadInt64(&node.ActiveConnections)
		if conns < minConn {
			minConn = conns
			bestNode = node
		}
	}

	return bestNode, nil
}

// IncrementConn increases connection count on a node.
func (node *ServerNode) IncrementConn() {
	atomic.AddInt64(&node.ActiveConnections, 1)
}

// DecrementConn decreases connection count on a node.
func (node *ServerNode) DecrementConn() {
	atomic.AddInt64(&node.ActiveConnections, -1)
}
