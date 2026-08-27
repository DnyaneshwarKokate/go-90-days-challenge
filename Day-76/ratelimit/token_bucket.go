package ratelimit

import (
	"math"
	"sync"
	"time"
)

// TokenBucketLimiter manages continuous token refills and burst requests.
type TokenBucketLimiter struct {
	mu         sync.Mutex
	capacity   float64
	refillRate float64 // Tokens per second
	tokens     float64
	lastRefill time.Time
}

// NewTokenBucketLimiter creates a limiter with max capacity and refill rate per second.
func NewTokenBucketLimiter(capacity float64, refillRate float64) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		capacity:   capacity,
		refillRate: refillRate,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
}

// Allow attempts to consume 1 token. Refills tokens based on elapsed time.
func (b *TokenBucketLimiter) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()
	b.lastRefill = now

	// Refill tokens proportional to elapsed time
	b.tokens = math.Min(b.capacity, b.tokens+(elapsed*b.refillRate))

	if b.tokens >= 1.0 {
		b.tokens -= 1.0
		return true
	}

	return false
}

// TokensAvailable returns current token balance.
func (b *TokenBucketLimiter) TokensAvailable() float64 {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.tokens
}
