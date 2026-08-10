package redis_test

import (
	"context"
	"testing"
	"time"

	"day-36/logger"
	"day-36/redis"
)

func TestRedisService_Operations(t *testing.T) {
	zapLog, _ := logger.NewZapLogger("test", "")
	svc := redis.NewRedisService("localhost:6379", "", 0, zapLog)

	ctx := context.Background()

	// 1. Test SET / GET
	err := svc.Set(ctx, "test_key", "hello_world", 10*time.Second)
	if err != nil {
		t.Fatalf("failed to set redis key: %v", err)
	}

	val, err := svc.Get(ctx, "test_key")
	if err != nil || val != "hello_world" {
		t.Errorf("expected 'hello_world', got '%s', err: %v", val, err)
	}

	// 2. Test HSET / HGETALL
	fields := map[string]string{"foo": "bar", "baz": "qux"}
	err = svc.HSet(ctx, "test_hash", fields)
	if err != nil {
		t.Fatalf("failed to hset: %v", err)
	}

	hashResult, err := svc.HGetAll(ctx, "test_hash")
	if err != nil || hashResult["foo"] != "bar" {
		t.Errorf("expected foo=bar in hash, got %v", hashResult)
	}
}
