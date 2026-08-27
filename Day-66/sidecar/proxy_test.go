package sidecar_test

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"day66/sidecar"
)

func TestSidecarProxyInterceptionAndRetry(t *testing.T) {
	var attemptCounter int32

	// Flaky Handler: fails on attempt 1, succeeds on attempt 2
	flakyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		current := atomic.AddInt32(&attemptCounter, 1)
		if r.Header.Get("X-Trace-ID") == "" {
			t.Errorf("Expected X-Trace-ID header to be set by sidecar")
		}

		if current == 1 {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Success from upstream"))
	})

	proxy := sidecar.NewSidecarProxy(flakyHandler, sidecar.SidecarConfig{
		MaxRetries:    2,
		EnablemTLS:    true,
		EnableTracing: true,
		RetryBackoff:  10 * time.Millisecond,
	})

	req := httptest.NewRequest("GET", "/api/v1/user", nil)
	req.Header.Set("X-Client-Cert", "valid-cert-thumbprint")
	rec := httptest.NewRecorder()

	proxy.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected HTTP 200 OK after retry, got %d", rec.Code)
	}

	if atomic.LoadInt32(&attemptCounter) != 2 {
		t.Fatalf("Expected 2 attempts, got %d", attemptCounter)
	}

	total, retried, failed := proxy.Metrics()
	if total != 1 || retried != 1 || failed != 0 {
		t.Fatalf("Metrics mismatch: total=%d, retried=%d, failed=%d", total, retried, failed)
	}
}

func TestSidecarProxymTLSRejection(t *testing.T) {
	target := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	proxy := sidecar.NewSidecarProxy(target, sidecar.SidecarConfig{
		EnablemTLS: true,
	})

	// Request without X-Client-Cert
	req := httptest.NewRequest("GET", "/api/v1/secure", nil)
	rec := httptest.NewRecorder()

	proxy.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("Expected HTTP 401 Unauthorized due to missing mTLS header, got %d", rec.Code)
	}
}
