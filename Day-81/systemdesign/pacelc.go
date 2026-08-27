package systemdesign

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type SystemMode string

const (
	ModeAP_EL SystemMode = "AP/EL" // Availability over Consistency during Partition; Latency over Consistency Else (e.g., Cassandra, DynamoDB)
	ModeCP_EC SystemMode = "CP/EC" // Consistency over Availability during Partition; Consistency over Latency Else (e.g., HBase, Spanner)
)

type NodeState struct {
	ID        string
	Value     string
	Version   int64
	Reachable bool
}

// DistributedCluster simulates PACELC theorem trade-offs under network partition.
type DistributedCluster struct {
	mu            sync.RWMutex
	mode          SystemMode
	nodes         map[string]*NodeState
	partitioned   bool
	latencyBuffer time.Duration
}

// NewDistributedCluster initializes a cluster with a PACELC mode setting.
func NewDistributedCluster(mode SystemMode) *DistributedCluster {
	return &DistributedCluster{
		mode:  mode,
		nodes: make(map[string]*NodeState),
	}
}

// AddNode adds a server node to the cluster.
func (c *DistributedCluster) AddNode(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.nodes[id] = &NodeState{ID: id, Reachable: true}
}

// SetPartition Toggles network partition state across the cluster.
func (c *DistributedCluster) SetPartition(partitioned bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.partitioned = partitioned
}

// Write simulates write operations enforcing PACELC rules.
func (c *DistributedCluster) Write(key string, val string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.partitioned {
		if c.mode == ModeCP_EC {
			// CP Mode: Reject writes during partition to guarantee strong consistency
			return errors.New("CP Mode: Network partition active. Write rejected to preserve consistency")
		}
		// AP Mode: Accept write on reachable node (Eventual Consistency)
		fmt.Println("[PACELC] AP Mode: Partition active. Accepting local node write for high availability.")
	} else {
		if c.mode == ModeCP_EC {
			// EC Mode: Synchronously replicate to all nodes (higher latency, strict consistency)
			c.latencyBuffer = 50 * time.Millisecond
		} else {
			// EL Mode: Asynchronous background replication (low latency)
			c.latencyBuffer = 5 * time.Millisecond
		}
	}

	for _, n := range c.nodes {
		if n.Reachable {
			n.Value = val
			n.Version++
		}
	}

	time.Sleep(c.latencyBuffer)
	return nil
}

// Read returns key value enforcing PACELC consistency rules.
func (c *DistributedCluster) Read(nodeID string) (string, int64, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	n, exists := c.nodes[nodeID]
	if !exists {
		return "", 0, errors.New("node not found")
	}

	return n.Value, n.Version, nil
}
