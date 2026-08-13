package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-48/config"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-48/domain"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-48/repository"
)

// APIHandler coordinates HTTP handlers with app config and data store.
type APIHandler struct {
	cfg       *config.Config
	store     *repository.Store
	isReady   int32
	startTime time.Time
}

// NewAPIHandler constructs a new APIHandler.
func NewAPIHandler(cfg *config.Config, store *repository.Store) *APIHandler {
	h := &APIHandler{
		cfg:       cfg,
		store:     store,
		startTime: time.Now(),
	}
	atomic.StoreInt32(&h.isReady, 1)
	return h
}

// SetNotReady dynamically sets readiness probe state to 0 for SIGTERM teardown.
func (h *APIHandler) SetNotReady() {
	atomic.StoreInt32(&h.isReady, 0)
}

// Healthz serves Kubernetes Liveness Probe requests.
func (h *APIHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"pod_name":  h.cfg.PodName,
	})
}

// Ready serves Kubernetes Readiness Probe requests.
func (h *APIHandler) Ready(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if atomic.LoadInt32(&h.isReady) == 1 {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    "ready",
			"pod_name":  h.cfg.PodName,
			"timestamp": time.Now().UTC(),
		})
		return
	}

	w.WriteHeader(http.StatusServiceUnavailable)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "unready",
		"reason":    "pod shutting down or initializing",
		"timestamp": time.Now().UTC(),
	})
}

// Info returns pod metadata injected via Downward API and ConfigMaps.
func (h *APIHandler) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"app_env":       h.cfg.AppEnv,
		"pod_name":      h.cfg.PodName,
		"pod_namespace": h.cfg.PodNamespace,
		"node_name":     h.cfg.NodeName,
		"pod_ip":        h.cfg.PodIP,
		"data_path":     h.cfg.DataPath,
		"uptime":        time.Since(h.startTime).String(),
	})
}

// Orders handles GET (list) and POST (create) order endpoints.
func (h *APIHandler) Orders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		orders := h.store.GetAll()
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"count":  len(orders),
			"orders": orders,
		})

	case http.MethodPost:
		var req domain.CreateOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid JSON payload"})
			return
		}

		if err := req.Validate(); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		order := domain.Order{
			ID:        fmt.Sprintf("ORD-%d", time.Now().UnixNano()),
			Customer:  req.Customer,
			Item:      req.Item,
			Quantity:  req.Quantity,
			Price:     req.Price,
			Status:    "PENDING",
			CreatedAt: time.Now().UTC(),
		}

		if err := h.store.Save(order); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "failed to persist order to volume"})
			return
		}

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(order)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
	}
}

// Metrics provides standard metrics for monitoring tools like Prometheus.
func (h *APIHandler) Metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"total_orders": h.store.Count(),
		"uptime_sec":   time.Since(h.startTime).Seconds(),
		"ready_state":  atomic.LoadInt32(&h.isReady),
		"pod_name":     h.cfg.PodName,
	})
}
