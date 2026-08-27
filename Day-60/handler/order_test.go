package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrderHandlerLifecycle(t *testing.T) {
	h := NewOrderHandler()

	// 1. Create Order
	body := `{"user_id":"usr_444","amount":150.75}`
	req := httptest.NewRequest("POST", "/api/v1/orders", bytes.NewBufferString(body))
	rr := httptest.NewRecorder()

	h.CreateOrder(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d. Body: %s", rr.Code, rr.Body.String())
	}

	var ord Order
	_ = json.Unmarshal(rr.Body.Bytes(), &ord)

	if ord.ID == "" || ord.UserID != "usr_444" {
		t.Errorf("Unexpected order creation output: %+v", ord)
	}

	// 2. Fetch Order
	reqGet := httptest.NewRequest("GET", "/api/v1/orders?id="+ord.ID, nil)
	rrGet := httptest.NewRecorder()

	h.GetOrder(rrGet, reqGet)

	if rrGet.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK for GetOrder, got %d", rrGet.Code)
	}

	var fetchedOrd Order
	_ = json.Unmarshal(rrGet.Body.Bytes(), &fetchedOrd)

	if fetchedOrd.Amount != 150.75 {
		t.Errorf("Expected amount 150.75, got %f", fetchedOrd.Amount)
	}
}
