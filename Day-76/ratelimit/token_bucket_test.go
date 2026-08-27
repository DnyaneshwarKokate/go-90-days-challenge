package ratelimit_test

import (
	"testing"
	"time"

	"day76/ratelimit"
)

func TestTokenBucketBurstAndRefill(t *testing.T) {
	// Capacity 2 tokens, refills 10 tokens per second
	limiter := ratelimit.NewTokenBucketLimiter(2, 10)

	// Consume 2 tokens (Burst)
	if !limiter.Allow() || !limiter.Allow() {
		t.Fatalf("Burst tokens should be allowed")
	}

	// 3rd token must fail
	if limiter.Allow() {
		t.Fatalf("3rd request should be rejected due to empty bucket")
	}

	// Wait 200ms (re-fills ~2 tokens)
	time.Sleep(200 * time.Millisecond)

	if !limiter.Allow() {
		t.Fatalf("Request after refill should be allowed")
	}
}
