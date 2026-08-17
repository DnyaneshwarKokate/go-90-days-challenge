package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

// Response represents a standard API response structure.
type Response struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

// HealthHandler returns system health status for CI/CD probes.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := Response{
		Status:    "UP",
		Message:   "CI/CD Baseline Service is Healthy",
		Version:   "1.0.0",
		Timestamp: time.Now().UTC(),
	}
	_ = json.NewEncoder(w).Encode(res)
}

// InfoHandler returns application build metadata.
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := Response{
		Status:    "OK",
		Message:   "Go CI/CD Automation Day 49",
		Version:   "1.0.0",
		Timestamp: time.Now().UTC(),
	}
	_ = json.NewEncoder(w).Encode(res)
}
