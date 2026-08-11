package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-44/config"
)

var startTime = time.Now()

// Item represents a product or resource in the system.
type Item struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

// HealthResponse represents standard health check output.
type HealthResponse struct {
	Status    string    `json:"status"`
	Uptime    string    `json:"uptime"`
	Timestamp time.Time `json:"timestamp"`
}

// InfoResponse represents system and container information.
type InfoResponse struct {
	ServiceName string `json:"service_name"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
	Hostname    string `json:"hostname"`
	GoVersion   string `json:"go_version"`
	NumCPU      int    `json:"num_cpu"`
	Goroutines  int    `json:"goroutines"`
}

// ItemService manages in-memory item operations safely.
type ItemService struct {
	mu    sync.RWMutex
	items map[string]Item
	cfg   config.Config
}

// NewItemService initializes an ItemService with default sample data.
func NewItemService(cfg config.Config) *ItemService {
	s := &ItemService{
		items: make(map[string]Item),
		cfg:   cfg,
	}
	// Initial sample data
	s.items["1"] = Item{ID: "1", Name: "Docker Container", Price: 0.00, CreatedAt: time.Now()}
	s.items["2"] = Item{ID: "2", Name: "Go Binary", Price: 49.99, CreatedAt: time.Now()}
	return s
}

// HandleHealth responds with service status and uptime.
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(HealthResponse{
		Status:    "UP",
		Uptime:    time.Since(startTime).String(),
		Timestamp: time.Now().UTC(),
	})
}

// HandleInfo responds with environment and system metadata.
func (s *ItemService) HandleInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(InfoResponse{
		ServiceName: s.cfg.ServiceName,
		Version:     s.cfg.Version,
		Environment: s.cfg.AppEnv,
		Hostname:    hostname,
		GoVersion:   runtime.Version(),
		NumCPU:      runtime.NumCPU(),
		Goroutines:  runtime.NumGoroutine(),
	})
}

// HandleItems routes GET and POST requests for items.
func (s *ItemService) HandleItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		s.mu.RLock()
		itemList := make([]Item, 0, len(s.items))
		for _, item := range s.items {
			itemList = append(itemList, item)
		}
		s.mu.RUnlock()

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(itemList)

	case http.MethodPost:
		var newItem Item
		if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
			http.Error(w, `{"error":"Invalid JSON payload"}`, http.StatusBadRequest)
			return
		}

		if newItem.ID == "" || newItem.Name == "" {
			http.Error(w, `{"error":"ID and Name are required fields"}`, http.StatusBadRequest)
			return
		}

		newItem.CreatedAt = time.Now().UTC()

		s.mu.Lock()
		s.items[newItem.ID] = newItem
		s.mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(newItem)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
