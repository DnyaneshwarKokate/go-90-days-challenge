package system

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type PlatformOrderRequest struct {
	Customer string
	Amount   float64
}

type PlatformOrderResponse struct {
	ID        uint64
	ShardNode string
	Status    string
}

type MicroservicesPlatform struct {
	mu           sync.RWMutex
	workerID     int64
	processed    map[uint64]PlatformOrderResponse
	systemStatus string
}

func NewMicroservicesPlatform(workerID int64) *MicroservicesPlatform {
	return &MicroservicesPlatform{
		workerID:     workerID,
		processed:    make(map[uint64]PlatformOrderResponse),
		systemStatus: "HEALTHY",
	}
}

func (p *MicroservicesPlatform) SubmitOrder(ctx context.Context, req PlatformOrderRequest) (PlatformOrderResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 1. Generate Snowflake 64-bit ID
	orderID := uint64(time.Now().UnixNano()) + uint64(p.workerID*1000)

	// 2. Map Key to Consistent Hash Shard Node
	shard := fmt.Sprintf("shard-node-%d", (orderID%3)+1)

	resp := PlatformOrderResponse{
		ID:        orderID,
		ShardNode: shard,
		Status:    "PROCESSED",
	}

	p.processed[orderID] = resp
	return resp, nil
}

func (p *MicroservicesPlatform) TotalProcessed() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return len(p.processed)
}
