package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-45/config"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-45/repository"
)

type APIHandler struct {
	cfg  *config.Config
	repo repository.UserRepository
}

func NewAPIHandler(cfg *config.Config, repo repository.UserRepository) *APIHandler {
	return &APIHandler{
		cfg:  cfg,
		repo: repo,
	}
}

// Healthz handles Kubernetes / Docker Liveness probes
func (h *APIHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// Ready handles Kubernetes / Docker Readiness probes
func (h *APIHandler) Ready(w http.ResponseWriter, r *http.Request) {
	// In production, check DB connection or cache availability here
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"status":    "ready",
		"db_status": "connected",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// Info returns container metadata & environment details
func (h *APIHandler) Info(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"service":     "go-dockerized-api",
		"environment": h.cfg.AppEnv,
		"version":     h.cfg.AppVersion,
		"port":        h.cfg.Port,
		"db_host":     h.cfg.DBHost,
	})
}

// HandleUsers routes GET and POST requests for /api/v1/users
func (h *APIHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/users")
	path = strings.TrimPrefix(path, "/")

	if path != "" {
		h.handleUserByID(w, r, path)
		return
	}

	switch r.Method {
	case http.MethodGet:
		users := h.repo.GetAll()
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"count": len(users),
			"data":  users,
		})
	case http.MethodPost:
		var newUser repository.User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			writeError(w, http.StatusBadRequest, "Invalid JSON payload")
			return
		}
		if newUser.ID == "" || newUser.Name == "" {
			writeError(w, http.StatusBadRequest, "ID and Name are required fields")
			return
		}

		if err := h.repo.Create(newUser); err != nil {
			if err == repository.ErrUserExists {
				writeError(w, http.StatusConflict, err.Error())
				return
			}
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		writeJSON(w, http.StatusCreated, map[string]interface{}{
			"message": "User created successfully",
			"data":    newUser,
		})
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *APIHandler) handleUserByID(w http.ResponseWriter, r *http.Request, id string) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	user, err := h.repo.GetByID(id)
	if err != nil {
		if err == repository.ErrUserNotFound {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data": user,
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeJSON(w, statusCode, map[string]string{
		"error": message,
	})
}
