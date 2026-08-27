package cache

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// CacheAsideManager handles cache-aside reading and invalidation on database updates.
type CacheAsideManager struct {
	mu         sync.RWMutex
	cacheStore map[string]string
	dbStore    map[string]string
	cacheHits  int64
	cacheMisses int64
}

// NewCacheAsideManager initializes cache and database mock stores.
func NewCacheAsideManager() *CacheAsideManager {
	return &CacheAsideManager{
		cacheStore: make(map[string]string),
		dbStore:    make(map[string]string),
	}
}

// SetDB populates mock database records.
func (m *CacheAsideManager) SetDB(key string, val string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.dbStore[key] = val
}

// Get implements Cache-Aside read logic.
func (m *CacheAsideManager) Get(key string) (string, bool, error) {
	m.mu.RLock()
	val, exists := m.cacheStore[key]
	m.mu.RUnlock()

	if exists {
		atomic.AddInt64(&m.cacheHits, 1)
		return val, true, nil // Cache Hit
	}

	atomic.AddInt64(&m.cacheMisses, 1)

	// Cache Miss: Fetch from primary Database
	m.mu.Lock()
	defer m.mu.Unlock()

	dbVal, dbExists := m.dbStore[key]
	if !dbExists {
		return "", false, fmt.Errorf("key %s not found in database", key)
	}

	// Populate cache for subsequent reads
	m.cacheStore[key] = dbVal
	return dbVal, false, nil
}

// Update writes new state to DB and invalidates current cache record.
func (m *CacheAsideManager) Update(key string, newVal string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Write-Through to Primary DB
	m.dbStore[key] = newVal

	// Invalidate Cache Key to prevent stale reads
	delete(m.cacheStore, key)
}

// Metrics returns hit and miss statistics.
func (m *CacheAsideManager) Metrics() (hits, misses int64) {
	return atomic.LoadInt64(&m.cacheHits), atomic.LoadInt64(&m.cacheMisses)
}
