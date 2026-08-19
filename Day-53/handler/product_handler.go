package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"day53/domain"
	"day53/service"
)

type ProductHandler struct {
	svc *service.ProductService
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		category := r.URL.Query().Get("category")
		products, err := h.svc.ListProducts(category)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(products)

	case http.MethodPost:
		var req domain.CreateProductRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		prod, err := h.svc.CreateProduct(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(prod)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	parts := strings.Split(path, "/")
	pID := parts[0]

	if pID == "" {
		http.Error(w, "Product ID required", http.StatusBadRequest)
		return
	}

	if len(parts) > 1 && parts[1] == "stock" && r.Method == http.MethodPatch {
		var req domain.UpdateStockRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid stock payload", http.StatusBadRequest)
			return
		}
		prod, err := h.svc.AdjustStock(pID, req.QuantityDelta)
		if err != nil {
			if err == domain.ErrInsufficientStock {
				http.Error(w, err.Error(), http.StatusConflict)
				return
			}
			if err == domain.ErrProductNotFound {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(prod)
		return
	}

	if r.Method == http.MethodGet {
		prod, err := h.svc.GetProduct(pID)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(prod)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
