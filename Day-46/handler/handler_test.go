package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-46/config"
	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-46/repository"
)

type mockRepo struct {
	users map[int]repository.User
	hits  int64
	miss  int64
}

func newMockRepo() *mockRepo {
	return &mockRepo{
		users: map[int]repository.User{
			1: {ID: 1, Name: "Alice", Email: "alice@example.com", Role: "admin", CreatedAt: time.Now()},
			2: {ID: 2, Name: "Bob", Email: "bob@example.com", Role: "user", CreatedAt: time.Now()},
		},
	}
}

func (m *mockRepo) GetByID(ctx context.Context, id int) (*repository.User, bool, error) {
	u, ok := m.users[id]
	if !ok {
		m.miss++
		return nil, false, nil
	}
	m.hits++
	return &u, true, nil
}

func (m *mockRepo) GetAll(ctx context.Context) ([]repository.User, error) {
	list := make([]repository.User, 0, len(m.users))
	for _, u := range m.users {
		list = append(list, u)
	}
	return list, nil
}

func (m *mockRepo) Create(ctx context.Context, name, email, role string) (*repository.User, error) {
	if name == "error" {
		return nil, errors.New("db error")
	}
	id := len(m.users) + 1
	u := repository.User{ID: id, Name: name, Email: email, Role: role, CreatedAt: time.Now()}
	m.users[id] = u
	return &u, nil
}

func (m *mockRepo) GetStats() repository.CacheStats {
	return repository.CacheStats{Hits: m.hits, Misses: m.miss}
}

func TestHealthzEndpoint(t *testing.T) {
	cfg := &config.Config{AppEnv: "testing", AppVersion: "1.0.0"}
	h := NewAPIHandler(cfg, newMockRepo(), nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()

	h.Healthz(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "healthy") {
		t.Fatalf("expected body to contain healthy, got %s", rr.Body.String())
	}
}

func TestReadyEndpoint_NotReady(t *testing.T) {
	cfg := &config.Config{}
	h := NewAPIHandler(cfg, newMockRepo(), nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/ready", nil)
	rr := httptest.NewRecorder()

	h.Ready(rr, req)

	if rr.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status 503, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "not_ready") {
		t.Fatalf("expected not_ready in body, got %s", rr.Body.String())
	}
}

func TestHandleUsers_Get(t *testing.T) {
	cfg := &config.Config{}
	h := NewAPIHandler(cfg, newMockRepo(), nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	rr := httptest.NewRecorder()

	h.HandleUsers(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Alice") {
		t.Fatalf("expected body to contain Alice, got %s", rr.Body.String())
	}
}

func TestHandleUsers_PostSuccess(t *testing.T) {
	cfg := &config.Config{}
	h := NewAPIHandler(cfg, newMockRepo(), nil, nil)

	body := strings.NewReader(`{"name":"Charlie","email":"charlie@example.com","role":"admin"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", body)
	rr := httptest.NewRecorder()

	h.HandleUsers(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Charlie") {
		t.Fatalf("expected body to contain Charlie, got %s", rr.Body.String())
	}
}

func TestHandleUsers_MethodNotAllowed(t *testing.T) {
	cfg := &config.Config{}
	h := NewAPIHandler(cfg, newMockRepo(), nil, nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/users", nil)
	rr := httptest.NewRecorder()

	h.HandleUsers(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rr.Code)
	}
}

func TestHandleUserByID_Found(t *testing.T) {
	cfg := &config.Config{}
	h := NewAPIHandler(cfg, newMockRepo(), nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/1", nil)
	rr := httptest.NewRecorder()

	h.HandleUserByID(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if rr.Header().Get("X-Cache") != "HIT" {
		t.Fatalf("expected X-Cache header to be HIT, got %s", rr.Header().Get("X-Cache"))
	}
}

func TestHandleUserByID_NotFound(t *testing.T) {
	cfg := &config.Config{}
	h := NewAPIHandler(cfg, newMockRepo(), nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/99", nil)
	rr := httptest.NewRecorder()

	h.HandleUserByID(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rr.Code)
	}
}

func TestStatsEndpoint(t *testing.T) {
	cfg := &config.Config{}
	h := NewAPIHandler(cfg, newMockRepo(), nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats", nil)
	rr := httptest.NewRecorder()

	h.Stats(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "hits") {
		t.Fatalf("expected body to contain hits, got %s", rr.Body.String())
	}
}
