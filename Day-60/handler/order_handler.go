package handler

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Order struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateOrderRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type OrderHandler struct {
	mu     sync.RWMutex
	orders map[string]*Order
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		orders: make(map[string]*Order),
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.UserID == "" || req.Amount <= 0 {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	idBytes := make([]byte, 8)
	_, _ = rand.Read(idBytes)
	ordID := "ord_" + hex.EncodeToString(idBytes)

	ord := &Order{
		ID:        ordID,
		UserID:    req.UserID,
		Amount:    req.Amount,
		Status:    "PENDING",
		CreatedAt: time.Now().UTC(),
	}

	h.mu.Lock()
	h.orders[ordID] = ord
	h.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(ord)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ordID := r.URL.Query().Get("id")
	if ordID == "" {
		http.Error(w, "Order ID required", http.StatusBadRequest)
		return
	}

	h.mu.RLock()
	ord, ok := h.orders[ordID]
	h.mu.RUnlock()

	if !ok {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ord)
}

func (h *OrderHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"UP","service":"day60-production-api"}`))
}
