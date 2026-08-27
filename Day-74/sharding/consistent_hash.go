package sharding

import (
	"errors"
	"fmt"
	"hash/fnv"
	"sort"
	"sync"
)

// ConsistentHashRing manages virtual node mapping around a 32-bit hash ring.
type ConsistentHashRing struct {
	mu       sync.RWMutex
	replicas int
	ring     []uint32
	nodesMap map[uint32]string
}

// NewConsistentHashRing initializes a hash ring with a virtual node replica count.
func NewConsistentHashRing(replicas int) *ConsistentHashRing {
	return &ConsistentHashRing{
		replicas: replicas,
		nodesMap: make(map[uint32]string),
	}
}

func hashKey(key string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	return h.Sum32()
}

// AddNode places a storage node and its virtual replicas onto the hash ring.
func (r *ConsistentHashRing) AddNode(node string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := 0; i < r.replicas; i++ {
		vNodeKey := fmt.Sprintf("%s-vnode-%d", node, i)
		hash := hashKey(vNodeKey)
		r.ring = append(r.ring, hash)
		r.nodesMap[hash] = node
	}

	sort.Slice(r.ring, func(i, j int) bool {
		return r.ring[i] < r.ring[j]
	})
}

// RemoveNode removes a node and its virtual replicas from the hash ring.
func (r *ConsistentHashRing) RemoveNode(node string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	newRing := make([]uint32, 0)
	for _, hash := range r.ring {
		if r.nodesMap[hash] == node {
			delete(r.nodesMap, hash)
		} else {
			newRing = append(newRing, hash)
		}
	}
	r.ring = newRing
}

// GetNode routes a key to its designated shard node using binary search on the ring.
func (r *ConsistentHashRing) GetNode(key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.ring) == 0 {
		return "", errors.New("hash ring is empty")
	}

	hash := hashKey(key)

	// Binary search for the first virtual node with hash >= key hash
	idx := sort.Search(len(r.ring), func(i int) bool {
		return r.ring[i] >= hash
	})

	// Wrap around to index 0 if hash > all ring nodes
	if idx == len(r.ring) {
		idx = 0
	}

	return r.nodesMap[r.ring[idx]], nil
}
