package handler

import (
	"encoding/json"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-47/config"
)

type APIHandler struct {
	cfg       *config.Config
	isReady   int32
	startTime time.Time
}

func NewAPIHandler(cfg *config.Config) *APIHandler {
	h := &APIHandler{
		cfg:       cfg,
		isReady:   1,
		startTime: time.Now(),
	}
	return h
}

func (h *APIHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "alive",
		"uptime": time.Since(h.startTime).String(),
		"pod":    h.cfg.PodName,
	})
}

func (h *APIHandler) Ready(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if atomic.LoadInt32(&h.isReady) == 1 {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "ready",
			"pod":    h.cfg.PodName,
		})
		return
	}

	w.WriteHeader(http.StatusServiceUnavailable)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "not_ready",
		"reason": "server is shutting down",
	})
}

func (h *APIHandler) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"service":   "k8s-golang-api",
		"version":   h.cfg.AppVersion,
		"env":       h.cfg.AppEnv,
		"pod":       h.cfg.PodName,
		"namespace": h.cfg.PodNamespace,
		"node":      h.cfg.NodeName,
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func (h *APIHandler) Data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authHeader := r.Header.Get("X-API-Key")
	if authHeader != h.cfg.APIKey {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "unauthorized access - invalid API key from K8s Secret",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Access granted using K8s Secret",
		"pod":     h.cfg.PodName,
		"items": []string{
			"Pod-to-Pod Communication",
			"ClusterIP Service Routing",
			"ConfigMap & Secret Injection",
			"Horizontal Pod Autoscaling",
		},
	})
}

func (h *APIHandler) SetNotReady() {
	atomic.StoreInt32(&h.isReady, 0)
}
