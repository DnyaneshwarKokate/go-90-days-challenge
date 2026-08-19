package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type GatewayProxy struct {
	userTarget    *url.URL
	productTarget *url.URL
	userProxy     *httputil.ReverseProxy
	productProxy  *httputil.ReverseProxy
}

func NewGatewayProxy(userServiceURL, productServiceURL string) (*GatewayProxy, error) {
	uTarget, err := url.Parse(userServiceURL)
	if err != nil {
		return nil, err
	}

	pTarget, err := url.Parse(productServiceURL)
	if err != nil {
		return nil, err
	}

	return &GatewayProxy{
		userTarget:    uTarget,
		productTarget: pTarget,
		userProxy:     httputil.NewSingleHostReverseProxy(uTarget),
		productProxy:  httputil.NewSingleHostReverseProxy(pTarget),
	}, nil
}

func (gp *GatewayProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[API Gateway Router] Incoming Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

	w.Header().Set("X-Gateway-ProxiedBy", "Go-90-Days-API-Gateway")

	if strings.HasPrefix(r.URL.Path, "/api/v1/users") {
		r.Host = gp.userTarget.Host
		gp.userProxy.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/api/v1/products") {
		r.Host = gp.productTarget.Host
		gp.productProxy.ServeHTTP(w, r)
		return
	}

	http.Error(w, "API Gateway: Route not found", http.StatusNotFound)
}
