package ratelimit

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// ClientWindow tracks request timestamps for a specific client key.
type ClientWindow struct {
	timestamps []time.Time
}

// SlidingWindowLimiter implements per-client sliding window rate limiting.
type SlidingWindowLimiter struct {
	mu          sync.Mutex
	maxRequests int
	windowSize  time.Duration
	clients     map[string]*ClientWindow
}

// NewSlidingWindowLimiter initializes rate limiter parameters.
func NewSlidingWindowLimiter(maxRequests int, windowSize time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		maxRequests: maxRequests,
		windowSize:  windowSize,
		clients:     make(map[string]*ClientWindow),
	}
}

// Allow determines if a request from clientID is permitted under the sliding window limit.
func (l *SlidingWindowLimiter) Allow(clientID string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	threshold := now.Add(-l.windowSize)

	client, exists := l.clients[clientID]
	if !exists {
		client = &ClientWindow{timestamps: make([]time.Time, 0)}
		l.clients[clientID] = client
	}

	// Filter out timestamps outside current sliding window
	validStamps := make([]time.Time, 0, len(client.timestamps))
	for _, ts := range client.timestamps {
		if ts.After(threshold) {
			validStamps = append(validStamps, ts)
		}
	}
	client.timestamps = validStamps

	// Check if limit exceeded
	if len(client.timestamps) >= l.maxRequests {
		return false
	}

	// Record allowed request timestamp
	client.timestamps = append(client.timestamps, now)
	return true
}

// Middleware decorates HTTP handlers with HTTP 429 rate limiting logic.
func (l *SlidingWindowLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := r.RemoteAddr
		if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
			clientID = ip
		}

		if !l.Allow(clientID) {
			w.Header().Set("Retry-After", fmt.Sprintf("%.0f", l.windowSize.Seconds()))
			http.Error(w, "HTTP 429: Too Many Requests - Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
