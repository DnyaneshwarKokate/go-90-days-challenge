package proxy

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
)

type RouteTarget struct {
	PathPrefix string
	Handler    http.Handler
	IsHealthy  bool
}

// Layer7ReverseProxy implements path-based routing and proxying.
type Layer7ReverseProxy struct {
	mu     sync.RWMutex
	routes map[string]*RouteTarget
}

// NewLayer7ReverseProxy initializes proxy instance.
func NewLayer7ReverseProxy() *Layer7ReverseProxy {
	return &Layer7ReverseProxy{
		routes: make(map[string]*RouteTarget),
	}
}

// RegisterRoute registers a URL path prefix to an upstream target handler.
func (p *Layer7ReverseProxy) RegisterRoute(prefix string, target http.Handler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.routes[prefix] = &RouteTarget{
		PathPrefix: prefix,
		Handler:    target,
		IsHealthy:  true,
	}
}

// ServeHTTP intercepts incoming L7 HTTP requests and proxies them to target handlers.
func (p *Layer7ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.mu.RLock()
	var matchedTarget *RouteTarget
	for prefix, target := range p.routes {
		if strings.HasPrefix(r.URL.Path, prefix) && target.IsHealthy {
			matchedTarget = target
			break
		}
	}
	p.mu.RUnlock()

	if matchedTarget == nil {
		http.Error(w, "Layer 7 Proxy: No healthy upstream route target matched", http.StatusBadGateway)
		return
	}

	// Enrich headers
	r.Header.Set("X-Forwarded-For", r.RemoteAddr)
	r.Header.Set("X-Proxy-Via", "Go-L7-ReverseProxy")

	rec := httptest.NewRecorder()
	matchedTarget.Handler.ServeHTTP(rec, r)

	for k, v := range rec.Header() {
		w.Header()[k] = v
	}
	w.Header().Set("X-Upstream-Path", matchedTarget.PathPrefix)
	w.WriteHeader(rec.Code)
	_, _ = w.Write(rec.Body.Bytes())
}

// SetRouteHealth updates the health status of a route.
func (p *Layer7ReverseProxy) SetRouteHealth(prefix string, healthy bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if target, exists := p.routes[prefix]; exists {
		target.IsHealthy = healthy
		fmt.Printf("[L7 REVERSE PROXY] Route '%s' health updated: %t\n", prefix, healthy)
	}
}
