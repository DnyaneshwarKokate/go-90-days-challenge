package discovery

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type HealthStatus string

const (
	StatusPassing HealthStatus = "PASSING"
	StatusFailing HealthStatus = "FAILING"
)

// ServiceInstance represents a registered microservice node.
type ServiceInstance struct {
	ID            string
	Name          string
	Address       string
	Port          int
	Status        HealthStatus
	LastHeartbeat time.Time
}

// ServiceRegistry is an in-memory service discovery registry node.
type ServiceRegistry struct {
	mu        sync.RWMutex
	instances map[string]*ServiceInstance
}

// NewServiceRegistry initializes a registry instance.
func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		instances: make(map[string]*ServiceInstance),
	}
}

// Register adds or updates a service instance in the registry.
func (r *ServiceRegistry) Register(inst ServiceInstance) error {
	if inst.ID == "" || inst.Name == "" {
		return errors.New("instance ID and Name are required")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	inst.Status = StatusPassing
	inst.LastHeartbeat = time.Now()
	r.instances[inst.ID] = &inst
	return nil
}

// Heartbeat refreshes the last heartbeat timestamp for a service instance.
func (r *ServiceRegistry) Heartbeat(instanceID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	inst, exists := r.instances[instanceID]
	if !exists {
		return errors.New("instance not found in registry")
	}

	inst.LastHeartbeat = time.Now()
	inst.Status = StatusPassing
	return nil
}

// GetHealthyInstances returns all active instances registered under a service name.
func (r *ServiceRegistry) GetHealthyInstances(serviceName string) []ServiceInstance {
	r.mu.RLock()
	defer r.mu.RUnlock()

	healthy := make([]ServiceInstance, 0)
	for _, inst := range r.instances {
		if inst.Name == serviceName && inst.Status == StatusPassing {
			healthy = append(healthy, *inst)
		}
	}
	return healthy
}

// EvictUnhealthy removes instances that have not sent a heartbeat within ttl duration.
func (r *ServiceRegistry) EvictUnhealthy(ttl time.Duration) int {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	evictedCount := 0

	for id, inst := range r.instances {
		if now.Sub(inst.LastHeartbeat) > ttl {
			fmt.Printf("[DISCOVERY REGISTRY] Evicting stale instance %s (%s)\n", id, inst.Name)
			delete(r.instances, id)
			evictedCount++
		}
	}
	return evictedCount
}
