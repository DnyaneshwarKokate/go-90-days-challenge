package distlock

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrLockAlreadyAcquired = errors.New("distributed lock is currently held by another node")
	ErrLockNotHeld          = errors.New("cannot release lock: lock not held or owner mismatch")
)

// LockLease represents an acquired lock lease.
type LockLease struct {
	Resource  string
	Owner     string
	ExpiresAt time.Time
}

// DistributedLockManager manages distributed locks with lease expiration.
type DistributedLockManager struct {
	mu    sync.Mutex
	locks map[string]*LockLease
}

// NewDistributedLockManager initializes lock store.
func NewDistributedLockManager() *DistributedLockManager {
	return &DistributedLockManager{
		locks: make(map[string]*LockLease),
	}
}

// Acquire attempts to gain exclusive lock on resource with a lease TTL.
func (m *DistributedLockManager) Acquire(resource string, owner string, ttl time.Duration) (*LockLease, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()

	// Check existing lock and evict if TTL expired
	if lease, exists := m.locks[resource]; exists {
		if now.Before(lease.ExpiresAt) {
			return nil, ErrLockAlreadyAcquired
		}
		// Expired lease auto-evicted
		delete(m.locks, resource)
	}

	lease := &LockLease{
		Resource:  resource,
		Owner:     owner,
		ExpiresAt: now.Add(ttl),
	}

	m.locks[resource] = lease
	return lease, nil
}

// Release frees the lock if requested by the current lease owner.
func (m *DistributedLockManager) Release(resource string, owner string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	lease, exists := m.locks[resource]
	if !exists || lease.Owner != owner {
		return ErrLockNotHeld
	}

	delete(m.locks, resource)
	return nil
}

// String formats lease status.
func (l *LockLease) String() string {
	return fmt.Sprintf("Lease [Resource: %s | Owner: %s | ExpiresIn: %v]",
		l.Resource, l.Owner, time.Until(l.ExpiresAt).Round(time.Millisecond))
}
