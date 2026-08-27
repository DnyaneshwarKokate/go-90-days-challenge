package proxy_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"day82/proxy"
)

func TestLayer7ReverseProxyRouting(t *testing.T) {
	l7Proxy := proxy.NewLayer7ReverseProxy()

	userApp := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("User Service Response"))
	})

	l7Proxy.RegisterRoute("/api/v1/users", userApp)

	req := httptest.NewRequest("GET", "/api/v1/users/profile", nil)
	rec := httptest.NewRecorder()

	l7Proxy.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK || rec.Body.String() != "User Service Response" {
		t.Fatalf("Routing failed: code=%d body=%s", rec.Code, rec.Body.String())
	}

	// Mark route unhealthy
	l7Proxy.SetRouteHealth("/api/v1/users", false)

	rec2 := httptest.NewRecorder()
	l7Proxy.ServeHTTP(rec2, req)

	if rec2.Code != http.StatusBadGateway {
		t.Fatalf("Expected HTTP 502 Bad Gateway for unhealthy route, got %d", rec2.Code)
	}
}
