package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"day-36/domain"
	"day-36/handler"
	"day-36/logger"
	"day-36/middleware"
	"day-36/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("==========================================================================")
	fmt.Println("🚀 Day 36: Redis Basics in Go (Strings, Hashes, Lists, Sets, TTL)")
	fmt.Println("==========================================================================")

	logFilePath := "redis_app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	ctx := context.Background()
	zapLog.Info(ctx, "Redis Service initializing", "target_host", "localhost:6379")

	redisSvc := redis.NewRedisService("localhost:6379", "", 0, zapLog)
	redisHandler := handler.NewRedisHandler(redisSvc, zapLog)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	v1 := router.Group("/api/v1/redis")
	{
		v1.POST("/set", redisHandler.SetKeyValue)
		v1.GET("/get/:key", redisHandler.GetKeyValue)
		v1.POST("/hset", redisHandler.SetHash)
		v1.GET("/hgetall/:key", redisHandler.GetHashAll)
	}

	fmt.Println("\n--- 1️⃣ Testing Redis SET Key-Value ---")
	setKV(router, "session:usr_99", "active_jwt_token_payload", 300)

	fmt.Println("\n--- 2️⃣ Testing Redis GET Key-Value ---")
	getKV(router, "session:usr_99")

	fmt.Println("\n--- 3️⃣ Testing Redis HSET Hash ---")
	setHash(router, "user:101", map[string]string{
		"name":  "Dnyaneshwar",
		"email": "dnyanesh@example.com",
		"role":  "ADMIN",
	})

	fmt.Println("\n--- 4️⃣ Testing Redis HGETALL Hash ---")
	getHash(router, "user:101")

	fmt.Println("\n✅ Day 36 Redis Basics executed successfully! Check redis_app.log for logs.")
}

func setKV(router *gin.Engine, key, val string, ttl int) {
	payload := domain.KeyValueInput{Key: key, Value: val, TTLSeconds: ttl}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/redis/set", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST /redis/set -> Status %d | Body: %s\n", w.Code, w.Body.String())
}

func getKV(router *gin.Engine, key string) {
	req, _ := http.NewRequest("GET", "/api/v1/redis/get/"+key, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("GET /redis/get/%s -> Status %d | Body: %s\n", key, w.Code, w.Body.String())
}

func setHash(router *gin.Engine, key string, fields map[string]string) {
	payload := domain.HashInput{Key: key, Fields: fields}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/redis/hset", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST /redis/hset -> Status %d | Body: %s\n", w.Code, w.Body.String())
}

func getHash(router *gin.Engine, key string) {
	req, _ := http.NewRequest("GET", "/api/v1/redis/hgetall/"+key, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("GET /redis/hgetall/%s -> Status %d | Body: %s\n", key, w.Code, w.Body.String())
}
