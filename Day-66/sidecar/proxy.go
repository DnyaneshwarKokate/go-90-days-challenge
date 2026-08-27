package sidecar

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"time"
)

// SidecarConfig holds policy rules enforced by the sidecar proxy.
type SidecarConfig struct {
	MaxRetries    int
	EnablemTLS    bool
	EnableTracing bool
	RetryBackoff  time.Duration
}

// SidecarProxy decorates upstream HTTP handlers with mesh features.
type SidecarProxy struct {
	targetHandler http.Handler
	config        SidecarConfig
	metrics       struct {
		totalRequests   int64
		retriedRequests int64
		failedRequests  int64
	}
}

// NewSidecarProxy initializes a sidecar interceptor proxy.
func NewSidecarProxy(target http.Handler, config SidecarConfig) *SidecarProxy {
	return &SidecarProxy{
		targetHandler: target,
		config:        config,
	}
}

// ServeHTTP intercepts HTTP requests, applies sidecar policies, and manages retries.
func (p *SidecarProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&p.metrics.totalRequests, 1)

	// 1. mTLS Enforce Check
	if p.config.EnablemTLS {
		if r.Header.Get("X-Client-Cert") == "" {
			http.Error(w, "Sidecar Security: Mutual TLS (mTLS) client certificate required", http.StatusUnauthorized)
			atomic.AddInt64(&p.metrics.failedRequests, 1)
			return
		}
	}

	// 2. Distributed Tracing Header Injection
	if p.config.EnableTracing && r.Header.Get("X-Trace-ID") == "" {
		traceID := fmt.Sprintf("trace-mesh-%d", time.Now().UnixNano())
		r.Header.Set("X-Trace-ID", traceID)
	}

	// Read request body into memory buffer for safe retries
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
		r.Body.Close()
	}

	attempts := 0
	for {
		attempts++
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		rec := httptest.NewRecorder()
		p.targetHandler.ServeHTTP(rec, r)

		// Check if response is successful (< 500) or retries exhausted
		if rec.Code < 500 || attempts > p.config.MaxRetries {
			// Write recorder results to actual response writer
			for k, v := range rec.Header() {
				w.Header()[k] = v
			}
			w.Header().Set("X-Sidecar-Attempts", fmt.Sprintf("%d", attempts))
			w.WriteHeader(rec.Code)
			_, _ = w.Write(rec.Body.Bytes())

			if rec.Code >= 500 {
				atomic.AddInt64(&p.metrics.failedRequests, 1)
			}
			return
		}

		// Retrying failed upstream call
		atomic.AddInt64(&p.metrics.retriedRequests, 1)
		time.Sleep(p.config.RetryBackoff)
	}
}

// Metrics returns total, retried, and failed request counts.
func (p *SidecarProxy) Metrics() (total, retried, failed int64) {
	return atomic.LoadInt64(&p.metrics.totalRequests),
		atomic.LoadInt64(&p.metrics.retriedRequests),
		atomic.LoadInt64(&p.metrics.failedRequests)
}
