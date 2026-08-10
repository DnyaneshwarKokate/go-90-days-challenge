package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"day-30/config"
	"day-30/di"
	"day-30/domain"
)

func setupTestApp() (*di.Container, http.Handler) {
	cfg := config.LoadConfig()
	cfg.Environment = "test"

	container, _ := di.NewContainer(cfg)
	router := container.SetupRouter()
	return container, router
}

func getAuthToken(t *testing.T, handler http.Handler, email, password, role string) string {
	regPayload := domain.RegisterInput{Email: email, Password: password, Role: role}
	regBody, _ := json.Marshal(regPayload)
	req1, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(regBody))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req1)

	loginPayload := domain.LoginInput{Email: email, Password: password}
	loginBody, _ := json.Marshal(loginPayload)
	req2, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(loginBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Fatalf("expected login status 200, got %d. Body: %s", w2.Code, w2.Body.String())
	}

	var res map[string]interface{}
	json.Unmarshal(w2.Body.Bytes(), &res)
	dataMap := res["data"].(map[string]interface{})
	return dataMap["token"].(string)
}

func TestStudentAPI_FullLifecycle(t *testing.T) {
	_, handler := setupTestApp()

	adminToken := getAuthToken(t, handler, "test.admin@school.com", "AdminPassword123", "ADMIN")
	studentToken := getAuthToken(t, handler, "test.student@school.com", "StudentPassword123", "STUDENT")

	createPayload := domain.CreateStudentInput{
		FullName:   "Alice Johnson",
		Email:      "alice.johnson@school.com",
		Department: "PHYSICS",
		GPA:        3.88,
	}
	createBody, _ := json.Marshal(createPayload)

	reqCreate, _ := http.NewRequest("POST", "/api/v1/students", bytes.NewBuffer(createBody))
	reqCreate.Header.Set("Content-Type", "application/json")
	reqCreate.Header.Set("Authorization", "Bearer "+adminToken)
	wCreate := httptest.NewRecorder()
	handler.ServeHTTP(wCreate, reqCreate)

	if wCreate.Code != http.StatusCreated {
		t.Fatalf("expected status 201 Created, got %d. Body: %s", wCreate.Code, wCreate.Body.String())
	}

	var createRes map[string]interface{}
	json.Unmarshal(wCreate.Body.Bytes(), &createRes)
	studentData := createRes["data"].(map[string]interface{})
	studentID := studentData["id"].(string)

	reqGet, _ := http.NewRequest("GET", "/api/v1/students/"+studentID, nil)
	reqGet.Header.Set("Authorization", "Bearer "+studentToken)
	wGet := httptest.NewRecorder()
	handler.ServeHTTP(wGet, reqGet)

	if wGet.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", wGet.Code)
	}

	reqForbiddenDelete, _ := http.NewRequest("DELETE", "/api/v1/students/"+studentID, nil)
	reqForbiddenDelete.Header.Set("Authorization", "Bearer "+studentToken)
	wForbiddenDelete := httptest.NewRecorder()
	handler.ServeHTTP(wForbiddenDelete, reqForbiddenDelete)

	if wForbiddenDelete.Code != http.StatusForbidden {
		t.Errorf("expected status 403 Forbidden for STUDENT role deletion, got %d", wForbiddenDelete.Code)
	}

	reqDelete, _ := http.NewRequest("DELETE", "/api/v1/students/"+studentID, nil)
	reqDelete.Header.Set("Authorization", "Bearer "+adminToken)
	wDelete := httptest.NewRecorder()
	handler.ServeHTTP(wDelete, reqDelete)

	if wDelete.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK for ADMIN role deletion, got %d", wDelete.Code)
	}
}
