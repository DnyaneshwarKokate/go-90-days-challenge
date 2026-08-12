package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-47/config"
)

func TestHealthz(t *testing.T) {
	cfg := &config.Config{PodName: "test-pod"}
	h := NewAPIHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()

	h.Healthz(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "alive") {
		t.Fatalf("expected alive in body, got %s", rr.Body.String())
	}
}

func TestReady(t *testing.T) {
	cfg := &config.Config{PodName: "test-pod"}
	h := NewAPIHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/ready", nil)
	rr := httptest.NewRecorder()

	h.Ready(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	h.SetNotReady()

	rrNotReady := httptest.NewRecorder()
	h.Ready(rrNotReady, req)

	if rrNotReady.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status 503 after SetNotReady, got %d", rrNotReady.Code)
	}
}

func TestInfo(t *testing.T) {
	cfg := &config.Config{
		AppEnv:       "production",
		AppVersion:   "1.0.0",
		PodName:      "k8s-pod-1",
		PodNamespace: "production",
		NodeName:     "worker-node-1",
	}
	h := NewAPIHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/info", nil)
	rr := httptest.NewRecorder()

	h.Info(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "k8s-pod-1") {
		t.Fatalf("expected pod name in response, got %s", rr.Body.String())
	}
}

func TestData_Unauthorized(t *testing.T) {
	cfg := &config.Config{APIKey: "secret-key"}
	h := NewAPIHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/data", nil)
	rr := httptest.NewRecorder()

	h.Data(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", rr.Code)
	}
}

func TestData_Success(t *testing.T) {
	cfg := &config.Config{APIKey: "secret-key"}
	h := NewAPIHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/data", nil)
	req.Header.Set("X-API-Key", "secret-key")
	rr := httptest.NewRecorder()

	h.Data(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Access granted") {
		t.Fatalf("expected success message, got %s", rr.Body.String())
	}
}
