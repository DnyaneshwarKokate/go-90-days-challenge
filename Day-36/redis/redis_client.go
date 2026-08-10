package redis

import (
	"context"
	"sync"
	"time"

	"day-36/domain"

	rdb "github.com/redis/go-redis/v9"
)

type RedisService struct {
	client     *rdb.Client
	isMock     bool
	mockKV     map[string]string
	mockHash   map[string]map[string]string
	mockList   map[string][]string
	mu         sync.RWMutex
	logger     domain.Logger
}

func NewRedisService(addr, password string, db int, logger domain.Logger) *RedisService {
	client := rdb.NewClient(&rdb.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	isMock := false
	if err := client.Ping(ctx).Err(); err != nil {
		logger.Warn(ctx, "Redis server unreachable, starting in Mock Memory Mode", "addr", addr, "error", err)
		isMock = true
	} else {
		logger.Info(ctx, "Connected to Redis server successfully", "addr", addr)
	}

	return &RedisService{
		client:   client,
		isMock:   isMock,
		mockKV:   make(map[string]string),
		mockHash: make(map[string]map[string]string),
		mockList: make(map[string][]string),
		logger:   logger,
	}
}

func (s *RedisService) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	s.logger.Info(ctx, "Redis SET operation", "key", key, "ttl", ttl)
	if s.isMock {
		s.mu.Lock()
		defer s.mu.Unlock()
		s.mockKV[key] = value
		return nil
	}
	return s.client.Set(ctx, key, value, ttl).Err()
}

func (s *RedisService) Get(ctx context.Context, key string) (string, error) {
	s.logger.Info(ctx, "Redis GET operation", "key", key)
	if s.isMock {
		s.mu.RLock()
		defer s.mu.RUnlock()
		val, exists := s.mockKV[key]
		if !exists {
			return "", domain.ErrKeyNotFound
		}
		return val, nil
	}
	val, err := s.client.Get(ctx, key).Result()
	if err == rdb.Nil {
		return "", domain.ErrKeyNotFound
	}
	return val, err
}

func (s *RedisService) HSet(ctx context.Context, key string, fields map[string]string) error {
	s.logger.Info(ctx, "Redis HSET operation", "key", key, "field_count", len(fields))
	if s.isMock {
		s.mu.Lock()
		defer s.mu.Unlock()
		if _, ok := s.mockHash[key]; !ok {
			s.mockHash[key] = make(map[string]string)
		}
		for f, v := range fields {
			s.mockHash[key][f] = v
		}
		return nil
	}
	for f, v := range fields {
		s.client.HSet(ctx, key, f, v)
	}
	return nil
}

func (s *RedisService) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	s.logger.Info(ctx, "Redis HGETALL operation", "key", key)
	if s.isMock {
		s.mu.RLock()
		defer s.mu.RUnlock()
		hash, exists := s.mockHash[key]
		if !exists {
			return nil, domain.ErrKeyNotFound
		}
		return hash, nil
	}
	return s.client.HGetAll(ctx, key).Result()
}

func (s *RedisService) LPush(ctx context.Context, key string, values ...interface{}) error {
	s.logger.Info(ctx, "Redis LPUSH operation", "key", key, "count", len(values))
	if s.isMock {
		s.mu.Lock()
		defer s.mu.Unlock()
		for _, val := range values {
			strVal := val.(string)
			s.mockList[key] = append([]string{strVal}, s.mockList[key]...)
		}
		return nil
	}
	return s.client.LPush(ctx, key, values...).Err()
}

func (s *RedisService) RPop(ctx context.Context, key string) (string, error) {
	s.logger.Info(ctx, "Redis RPOP operation", "key", key)
	if s.isMock {
		s.mu.Lock()
		defer s.mu.Unlock()
		list, exists := s.mockList[key]
		if !exists || len(list) == 0 {
			return "", domain.ErrKeyNotFound
		}
		lastItem := list[len(list)-1]
		s.mockList[key] = list[:len(list)-1]
		return lastItem, nil
	}
	val, err := s.client.RPop(ctx, key).Result()
	if err == rdb.Nil {
		return "", domain.ErrKeyNotFound
	}
	return val, err
}
