package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"day-37/domain"

	rdb "github.com/redis/go-redis/v9"
)

type RedisProductCache struct {
	client    *rdb.Client
	isMock    bool
	mockStore map[string]string
	mu        sync.RWMutex
	logger    domain.Logger
}

func NewRedisProductCache(addr string, logger domain.Logger) *RedisProductCache {
	client := rdb.NewClient(&rdb.Options{
		Addr: addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	isMock := false
	if err := client.Ping(ctx).Err(); err != nil {
		logger.Warn(ctx, "Redis cache server unreachable, starting in Mock Mode", "addr", addr)
		isMock = true
	} else {
		logger.Info(ctx, "Redis cache server connected successfully", "addr", addr)
	}

	return &RedisProductCache{
		client:    client,
		isMock:    isMock,
		mockStore: make(map[string]string),
		logger:    logger,
	}
}

func (c *RedisProductCache) Get(ctx context.Context, id string) (*domain.Product, error) {
	key := fmt.Sprintf("product:%s", id)
	c.logger.Debug(ctx, "Checking Redis Cache", "cache_key", key)

	var payload string
	if c.isMock {
		c.mu.RLock()
		val, exists := c.mockStore[key]
		c.mu.RUnlock()
		if !exists {
			c.logger.Info(ctx, "CACHE MISS — Product not found in Redis", "id", id)
			return nil, domain.ErrProductNotFound
		}
		payload = val
	} else {
		val, err := c.client.Get(ctx, key).Result()
		if err == rdb.Nil {
			c.logger.Info(ctx, "CACHE MISS — Product not found in Redis", "id", id)
			return nil, domain.ErrProductNotFound
		} else if err != nil {
			return nil, err
		}
		payload = val
	}

	c.logger.Info(ctx, "⚡ CACHE HIT — Product loaded from Redis", "id", id)
	var product domain.Product
	if err := json.Unmarshal([]byte(payload), &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (c *RedisProductCache) Set(ctx context.Context, p *domain.Product, ttl time.Duration) error {
	key := fmt.Sprintf("product:%s", p.ID)
	bytes, err := json.Marshal(p)
	if err != nil {
		return err
	}

	c.logger.Info(ctx, "Populating Redis Cache", "cache_key", key, "ttl", ttl)
	if c.isMock {
		c.mu.Lock()
		c.mockStore[key] = string(bytes)
		c.mu.Unlock()
		return nil
	}

	return c.client.Set(ctx, key, string(bytes), ttl).Err()
}

func (c *RedisProductCache) Delete(ctx context.Context, id string) error {
	key := fmt.Sprintf("product:%s", id)
	c.logger.Info(ctx, "🔥 CACHE INVALIDATION — Removing key from Redis", "cache_key", key)

	if c.isMock {
		c.mu.Lock()
		delete(c.mockStore, key)
		c.mu.Unlock()
		return nil
	}

	return c.client.Del(ctx, key).Err()
}
