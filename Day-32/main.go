package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"day-32/handler"
	"day-32/logger"
	"day-32/middleware"
	"day-32/repository"
	"day-32/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("==========================================================================")
	fmt.Println("🚀 Day 32: Pagination & Filtering in Go (Offset vs Cursor Keyset)")
	fmt.Println("==========================================================================")

	logFilePath := "article_app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	ctx := context.Background()
	zapLog.Info(ctx, "Pagination & Filtering Server starting", "log_file", logFilePath)

	articleRepo := repository.NewMemoryArticleRepository(zapLog)
	articleUC := usecase.NewArticleUseCase(articleRepo, zapLog)
	articleHandler := handler.NewArticleHandler(articleUC, zapLog)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	v1 := router.Group("/api/v1/articles")
	{
		v1.POST("/seed", articleHandler.SeedArticles)
		v1.GET("/offset", articleHandler.GetArticlesOffset)
		v1.GET("/cursor", articleHandler.GetArticlesCursor)
	}

	fmt.Println("\n--- 1️⃣ Seeding 30 Mock Articles ---")
	seedArticles(router, 30)

	fmt.Println("\n--- 2️⃣ Testing Offset Pagination: Page 1 (Limit=5, SortBy=views DESC) ---")
	fetchOffset(router, "/api/v1/articles/offset?page=1&limit=5&sort_by=views&order=desc")

	fmt.Println("\n--- 3️⃣ Testing Offset Pagination: Page 2 (Limit=5, SortBy=views DESC) ---")
	fetchOffset(router, "/api/v1/articles/offset?page=2&limit=5&sort_by=views&order=desc")

	fmt.Println("\n--- 4️⃣ Testing Multi-Attribute Filtering (min_views=500, search=Backend) ---")
	fetchOffset(router, "/api/v1/articles/offset?page=1&limit=5&min_views=500&search=Backend")

	fmt.Println("\n--- 5️⃣ Testing Cursor Keyset Pagination: Page 1 (Limit=3) ---")
	nextCursor := fetchCursorPage1(router, "/api/v1/articles/cursor?limit=3")

	if nextCursor != "" {
		fmt.Println("\n--- 6️⃣ Testing Cursor Keyset Pagination: Page 2 using NextCursor ---")
		fetchOffset(router, fmt.Sprintf("/api/v1/articles/cursor?limit=3&cursor=%s", nextCursor))
	}

	fmt.Println("\n✅ Day 32 Pagination & Filtering executed successfully! Check article_app.log for logs.")
}

func seedArticles(router *gin.Engine, count int) {
	url := fmt.Sprintf("/api/v1/articles/seed?count=%d", count)
	req, _ := http.NewRequest("POST", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST %s -> Status %d | Body: %s\n", url, w.Code, w.Body.String())
}

func fetchOffset(router *gin.Engine, url string) {
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("GET %s -> Status %d | Body: %s\n", url, w.Code, w.Body.String())
}

func fetchCursorPage1(router *gin.Engine, url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("GET %s -> Status %d | Body: %s\n", url, w.Code, w.Body.String())

	var res map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &res)
	if meta, ok := res["meta"].(map[string]interface{}); ok {
		if cur, ok := meta["next_cursor"].(string); ok {
			return cur
		}
	}
	return ""
}
