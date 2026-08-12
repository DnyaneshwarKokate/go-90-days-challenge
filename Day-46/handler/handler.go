package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-46/config"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-46/repository"
	"github.com/redis/go-redis/v9"
)

type APIHandler struct {
	cfg  *config.Config
	repo repository.UserRepository
	db   *sql.DB
	rdb  *redis.Client
}

func NewAPIHandler(cfg *config.Config, repo repository.UserRepository, db *sql.DB, rdb *redis.Client) *APIHandler {
	return &APIHandler{
		cfg:  cfg,
		repo: repo,
		db:   db,
		rdb:  rdb,
	}
}

func (h *APIHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":  "healthy",
		"env":     h.cfg.AppEnv,
		"version": h.cfg.AppVersion,
	})
}

func (h *APIHandler) Ready(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbStatus := "ok"
	if h.db == nil || h.db.Ping() != nil {
		dbStatus = "unreachable"
	}

	redisStatus := "ok"
	if h.rdb == nil || h.rdb.Ping(r.Context()).Err() != nil {
		redisStatus = "unreachable"
	}

	if dbStatus != "ok" {
		w.WriteHeader(http.StatusServiceUnavailable)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":   "not_ready",
			"postgres": dbStatus,
			"redis":    redisStatus,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":   "ready",
		"postgres": dbStatus,
		"redis":    redisStatus,
	})
}

func (h *APIHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		users, err := h.repo.GetAll(r.Context())
		if err != nil {
			http.Error(w, `{"error":"failed to fetch users"}`, http.StatusInternalServerError)
			return
		}
		_ = json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var req struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Role  string `json:"role"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" || req.Email == "" {
			http.Error(w, `{"error":"invalid request payload"}`, http.StatusBadRequest)
			return
		}
		if req.Role == "" {
			req.Role = "user"
		}

		user, err := h.repo.Create(r.Context(), req.Name, req.Email, req.Role)
		if err != nil {
			http.Error(w, `{"error":"failed to create user"}`, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(user)

	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func (h *APIHandler) HandleUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 4 {
		http.Error(w, `{"error":"invalid user id"}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil || id <= 0 {
		http.Error(w, `{"error":"invalid user id"}`, http.StatusBadRequest)
		return
	}

	user, fromCache, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, `{"error":"failed to fetch user"}`, http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, `{"error":"user not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("X-Cache", map[bool]string{true: "HIT", false: "MISS"}[fromCache])
	_ = json.NewEncoder(w).Encode(user)
}

func (h *APIHandler) Stats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(h.repo.GetStats())
}
