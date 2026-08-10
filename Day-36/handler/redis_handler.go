package handler

import (
	"net/http"
	"time"

	"day-36/domain"
	"day-36/redis"

	"github.com/gin-gonic/gin"
)

type RedisHandler struct {
	redisService *redis.RedisService
	logger       domain.Logger
}

func NewRedisHandler(redisService *redis.RedisService, logger domain.Logger) *RedisHandler {
	return &RedisHandler{
		redisService: redisService,
		logger:       logger,
	}
}

func (h *RedisHandler) SetKeyValue(c *gin.Context) {
	var input domain.KeyValueInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	ttl := time.Duration(input.TTLSeconds) * time.Second
	if err := h.redisService.Set(c.Request.Context(), input.Key, input.Value, ttl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key stored successfully", "key": input.Key, "value": input.Value})
}

func (h *RedisHandler) GetKeyValue(c *gin.Context) {
	key := c.Param("key")
	val, err := h.redisService.Get(c.Request.Context(), key)
	if err != nil {
		if err == domain.ErrKeyNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found in Redis"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": key, "value": val})
}

func (h *RedisHandler) SetHash(c *gin.Context) {
	var input domain.HashInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.redisService.HSet(c.Request.Context(), input.Key, input.Fields); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hash stored successfully", "key": input.Key})
}

func (h *RedisHandler) GetHashAll(c *gin.Context) {
	key := c.Param("key")
	fields, err := h.redisService.HGetAll(c.Request.Context(), key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": key, "data": fields})
}
