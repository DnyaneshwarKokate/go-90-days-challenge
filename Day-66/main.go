package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"day66/sidecar"
)

func main() {
	fmt.Println("=== Day 66: Service Mesh & Sidecar Proxy Pattern ===")

	// Mock Upstream Microservice
	upstreamApp := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get("X-Trace-ID")
		cert := r.Header.Get("X-Client-Cert")
		fmt.Printf("  [UPSTREAM MICROSERVICE] Processing request. TraceID: %s | Cert: %s\n", traceID, cert)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"SUCCESS","message":"Upstream processed successfully"}`))
	})

	// Wrap microservice with Service Mesh Sidecar Proxy
	sidecarProxy := sidecar.NewSidecarProxy(upstreamApp, sidecar.SidecarConfig{
		MaxRetries:    3,
		EnablemTLS:    true,
		EnableTracing: true,
		RetryBackoff:  20 * time.Millisecond,
	})

	fmt.Println("\n--- 1. Sending Valid Mesh Request ---")
	req1 := httptest.NewRequest("GET", "/orders/101", nil)
	req1.Header.Set("X-Client-Cert", "SHA256:abc123mtlscertificate")
	rec1 := httptest.NewRecorder()
	sidecarProxy.ServeHTTP(rec1, req1)

	fmt.Printf("[SIDECAR OUTCOME] Status: %d | Retries Used: %s | Body: %s\n",
		rec1.Code, rec1.Header().Get("X-Sidecar-Attempts"), rec1.Body.String())

	fmt.Println("\n--- 2. Sending Request Missing mTLS Header ---")
	req2 := httptest.NewRequest("GET", "/orders/101", nil)
	rec2 := httptest.NewRecorder()
	sidecarProxy.ServeHTTP(rec2, req2)

	fmt.Printf("[SIDECAR OUTCOME] Status: %d | Body: %s\n", rec2.Code, rec2.Body.String())

	total, retried, failed := sidecarProxy.Metrics()
	fmt.Printf("\n--- Sidecar Telemetry Metrics ---\nTotal: %d | Retried: %d | Failed: %d\n", total, retried, failed)
}
