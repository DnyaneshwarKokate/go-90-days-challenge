package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"day55/client"
)

type AggregatedOrder struct {
	OrderID   string      `json:"order_id"`
	User      interface{} `json:"user"`
	Product   interface{} `json:"product"`
	Timestamp time.Time   `json:"timestamp"`
}

type OrderService struct {
	userClient    *client.ResilientHTTPClient
	productClient *client.ResilientHTTPClient
	userServiceURL string
	productServiceURL string
}

func NewOrderService(userServiceURL, productServiceURL string) *OrderService {
	return &OrderService{
		userClient:        client.NewResilientHTTPClient(2*time.Second, 2, 3, 5*time.Second),
		productClient:     client.NewResilientHTTPClient(2*time.Second, 2, 3, 5*time.Second),
		userServiceURL:    userServiceURL,
		productServiceURL: productServiceURL,
	}
}

func (s *OrderService) GetAggregatedOrder(ctx context.Context, userID, productID string) (*AggregatedOrder, error) {
	// 1. Fetch User Data
	uURL := fmt.Sprintf("%s/api/v1/users/%s", s.userServiceURL, userID)
	uBytes, err := s.userClient.GetWithRetry(ctx, uURL)
	var userData interface{}
	if err != nil {
		userData = map[string]string{"error": "User service temporarily unavailable (fallback)"}
	} else {
		_ = json.Unmarshal(uBytes, &userData)
	}

	// 2. Fetch Product Data
	pURL := fmt.Sprintf("%s/api/v1/products/%s", s.productServiceURL, productID)
	pBytes, err := s.productClient.GetWithRetry(ctx, pURL)
	var productData interface{}
	if err != nil {
		productData = map[string]string{"error": "Product service temporarily unavailable (fallback)"}
	} else {
		_ = json.Unmarshal(pBytes, &productData)
	}

	return &AggregatedOrder{
		OrderID:   "ord_998877",
		User:      userData,
		Product:   productData,
		Timestamp: time.Now().UTC(),
	}, nil
}
