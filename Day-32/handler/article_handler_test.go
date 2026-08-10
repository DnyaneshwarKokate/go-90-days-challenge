package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"day-32/handler"
	"day-32/logger"
	"day-32/middleware"
	"day-32/repository"
	"day-32/usecase"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	zapLog, _ := logger.NewZapLogger("test", "")

	articleRepo := repository.NewMemoryArticleRepository(zapLog)
	articleUC := usecase.NewArticleUseCase(articleRepo, zapLog)
	articleHandler := handler.NewArticleHandler(articleUC, zapLog)

	router := gin.New()
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	v1 := router.Group("/api/v1/articles")
	{
		v1.POST("/seed", articleHandler.SeedArticles)
		v1.GET("/offset", articleHandler.GetArticlesOffset)
		v1.GET("/cursor", articleHandler.GetArticlesCursor)
	}

	return router
}

func TestPaginationAndFiltering_Offset(t *testing.T) {
	router := setupTestRouter()

	// 1. Seed mock data
	reqSeed, _ := http.NewRequest("POST", "/api/v1/articles/seed?count=20", nil)
	wSeed := httptest.NewRecorder()
	router.ServeHTTP(wSeed, reqSeed)
	if wSeed.Code != http.StatusCreated {
		t.Fatalf("expected seed status 201 Created, got %d", wSeed.Code)
	}

	// 2. Fetch Offset Page 1
	reqP1, _ := http.NewRequest("GET", "/api/v1/articles/offset?page=1&limit=5", nil)
	wP1 := httptest.NewRecorder()
	router.ServeHTTP(wP1, reqP1)

	if wP1.Code != http.StatusOK {
		t.Fatalf("expected 200 OK for page 1, got %d", wP1.Code)
	}

	var resP1 map[string]interface{}
	json.Unmarshal(wP1.Body.Bytes(), &resP1)
	meta1 := resP1["meta"].(map[string]interface{})

	if int(meta1["page"].(float64)) != 1 || int(meta1["total_pages"].(float64)) != 4 {
		t.Errorf("expected page 1 of 4, got page %v of %v", meta1["page"], meta1["total_pages"])
	}

	if meta1["has_next"].(bool) != true || meta1["has_prev"].(bool) != false {
		t.Errorf("expected has_next=true, has_prev=false")
	}
}

func TestPaginationAndFiltering_Cursor(t *testing.T) {
	router := setupTestRouter()

	// 1. Seed mock data
	reqSeed, _ := http.NewRequest("POST", "/api/v1/articles/seed?count=15", nil)
	wSeed := httptest.NewRecorder()
	router.ServeHTTP(wSeed, reqSeed)

	// 2. Fetch Cursor Page 1
	reqC1, _ := http.NewRequest("GET", "/api/v1/articles/cursor?limit=5", nil)
	wC1 := httptest.NewRecorder()
	router.ServeHTTP(wC1, reqC1)

	var resC1 map[string]interface{}
	json.Unmarshal(wC1.Body.Bytes(), &resC1)
	meta1 := resC1["meta"].(map[string]interface{})
	nextCursor := meta1["next_cursor"].(string)

	if nextCursor == "" {
		t.Fatalf("expected non-empty next_cursor for page 1")
	}

	// 3. Fetch Cursor Page 2 using NextCursor
	urlP2 := fmt.Sprintf("/api/v1/articles/cursor?limit=5&cursor=%s", nextCursor)
	reqC2, _ := http.NewRequest("GET", urlP2, nil)
	wC2 := httptest.NewRecorder()
	router.ServeHTTP(wC2, reqC2)

	if wC2.Code != http.StatusOK {
		t.Fatalf("expected 200 OK for cursor page 2, got %d", wC2.Code)
	}
}
