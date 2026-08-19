package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"day51/config"
	"day51/middleware"
)

type HealthResponse struct {
	ServiceName   string    `json:"service_name"`
	Status        string    `json:"status"`
	Environment   string    `json:"environment"`
	CorrelationID string    `json:"correlation_id"`
	Timestamp     time.Time `json:"timestamp"`
}

type HealthHandler struct {
	Cfg *config.Config
}

func NewHealthHandler(cfg *config.Config) *HealthHandler {
	return &HealthHandler{Cfg: cfg}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := HealthResponse{
		ServiceName:   h.Cfg.ServiceName,
		Status:        "UP",
		Environment:   h.Cfg.Environment,
		CorrelationID: middleware.GetCorrelationID(r.Context()),
		Timestamp:     time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res)
}
