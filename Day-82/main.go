package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"day82/proxy"
)

func main() {
	fmt.Println("=== Day 82: Reverse Proxy & Layer 7 Load Balancer Architecture ===")

	l7Proxy := proxy.NewLayer7ReverseProxy()

	// Upstream Order Microservice
	orderApp := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		via := r.Header.Get("X-Proxy-Via")
		fmt.Printf("  [ORDER SERVICE] Received request via: %s\n", via)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"SUCCESS","data":"Order #1001 details"}`))
	})

	l7Proxy.RegisterRoute("/orders", orderApp)

	fmt.Println("\n--- 1. Forwarding HTTP Request via L7 Reverse Proxy ---")
	req := httptest.NewRequest("GET", "/orders/1001", nil)
	rec := httptest.NewRecorder()

	l7Proxy.ServeHTTP(rec, req)
	fmt.Printf("[PROXY RESPONSE] Status: %d | Upstream Header: %s | Body: %s\n",
		rec.Code, rec.Header().Get("X-Upstream-Path"), rec.Body.String())
}
