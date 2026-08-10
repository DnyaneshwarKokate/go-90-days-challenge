package middleware

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"day-41/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/time/rate"
)

type clientLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type IPRateLimiter struct {
	ips    map[string]*clientLimiter
	mu     sync.RWMutex
	r      rate.Limit
	b      int
	logger domain.Logger
}

func NewIPRateLimiter(r rate.Limit, b int, logger domain.Logger) *IPRateLimiter {
	limiter := &IPRateLimiter{
		ips:    make(map[string]*clientLimiter),
		r:      r,
		b:      b,
		logger: logger,
	}

	go limiter.cleanupInactiveIPs()
	return limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	lim, exists := i.ips[ip]
	if !exists {
		l := rate.NewLimiter(i.r, i.b)
		i.ips[ip] = &clientLimiter{limiter: l, lastSeen: time.Now()}
		return l
	}

	lim.lastSeen = time.Now()
	return lim.limiter
}

func (i *IPRateLimiter) cleanupInactiveIPs() {
	for {
		time.Sleep(3 * time.Minute)
		i.mu.Lock()
		for ip, cl := range i.ips {
			if time.Since(cl.lastSeen) > 5*time.Minute {
				delete(i.ips, ip)
			}
		}
		i.mu.Unlock()
	}
}

func RateLimiterMiddleware(limiter *IPRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}
		c.Header("X-Request-ID", reqID)

		ctx := context.WithValue(c.Request.Context(), domain.RequestIDKey, reqID)
		c.Request = c.Request.WithContext(ctx)

		ip := c.ClientIP()
		if ip == "" {
			ip = "127.0.0.1"
		}

		lim := limiter.GetLimiter(ip)
		c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", limiter.b))

		if !lim.Allow() {
			c.Header("X-RateLimit-Remaining", "0")
			c.Header("Retry-After", "1")
			limiter.logger.Warn(ctx, "🚫 RATE LIMIT EXCEEDED", "ip", ip, "path", c.Request.URL.Path)

			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":   domain.ErrRateLimitExceeded.Error(),
				"status":  429,
				"retry_in": "1s",
			})
			c.Abort()
			return
		}

		c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", int(lim.Tokens())))
		c.Next()
	}
}
