package snowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	epoch         int64  = 1672531200000 // Custom epoch (2023-01-01 00:00:00 UTC in ms)
	workerBits    uint8  = 10            // Up to 1024 unique workers
	sequenceBits  uint8  = 12            // Up to 4096 IDs per millisecond per worker
	maxWorkerID   int64  = -1 ^ (-1 << workerBits)
	maxSequence   int64  = -1 ^ (-1 << sequenceBits)
	workerShift   uint8  = sequenceBits
	timestampShift uint8 = sequenceBits + workerBits
)

// SnowflakeNode generates 64-bit k-ordered unique IDs.
type SnowflakeNode struct {
	mu        sync.Mutex
	workerID  int64
	lastStamp int64
	sequence  int64
}

// NewSnowflakeNode initializes a generator with a specific worker ID (0 to 1023).
func NewSnowflakeNode(workerID int64) (*SnowflakeNode, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, errors.New("worker ID out of valid range (0-1023)")
	}
	return &SnowflakeNode{
		workerID:  workerID,
		lastStamp: -1,
		sequence:  0,
	}, nil
}

// NextID generates the next unique 64-bit integer ID.
func (n *SnowflakeNode) NextID() (uint64, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Now().UnixMilli()

	if now < n.lastStamp {
		return 0, errors.New("clock moved backwards! Refusing to generate ID")
	}

	if now == n.lastStamp {
		n.sequence = (n.sequence + 1) & maxSequence
		if n.sequence == 0 {
			// Sequence exhausted in current ms, wait for next ms
			for now <= n.lastStamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		n.sequence = 0
	}

	n.lastStamp = now

	id := uint64((now-epoch)<<timestampShift | (n.workerID << workerShift) | n.sequence)
	return id, nil
}
