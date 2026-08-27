package cache

import (
	"errors"
	"sync"
	"time"
)

var ErrCacheMiss = errors.New("cache miss")

type cacheItem struct {
	value      string
	expiresAt  time.Time
}

type DistributedCache struct {
	mu         sync.RWMutex
	store      map[string]cacheItem
	sfMu       sync.Mutex
	singleCalls map[string]*call
}

type call struct {
	wg  sync.WaitGroup
	val string
	err error
}

func NewDistributedCache() *DistributedCache {
	return &DistributedCache{
		store:       make(map[string]cacheItem),
		singleCalls: make(map[string]*call),
	}
}

func (c *DistributedCache) Get(key string) (string, error) {
	c.mu.RLock()
	item, ok := c.store[key]
	c.mu.RUnlock()

	if !ok {
		return "", ErrCacheMiss
	}

	if time.Now().After(item.expiresAt) {
		c.mu.Lock()
		delete(c.store, key)
		c.mu.Unlock()
		return "", ErrCacheMiss
	}

	return item.value, nil
}

func (c *DistributedCache) Set(key, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[key] = cacheItem{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
}

// GetOrFetch prevents Thundering Herd / Cache Stampede using singleflight request coalescing.
func (c *DistributedCache) GetOrFetch(key string, fetcher func() (string, error), ttl time.Duration) (string, error) {
	if val, err := c.Get(key); err == nil {
		return val, nil
	}

	c.sfMu.Lock()
	if c.singleCalls == nil {
		c.singleCalls = make(map[string]*call)
	}
	if cl, ok := c.singleCalls[key]; ok {
		c.sfMu.Unlock()
		cl.wg.Wait()
		return cl.val, cl.err
	}

	cl := new(call)
	cl.wg.Add(1)
	c.singleCalls[key] = cl
	c.sfMu.Unlock()

	cl.val, cl.err = fetcher()
	if cl.err == nil {
		c.Set(key, cl.val, ttl)
	}

	c.sfMu.Lock()
	delete(c.singleCalls, key)
	c.sfMu.Unlock()

	cl.wg.Done()
	return cl.val, cl.err
}
