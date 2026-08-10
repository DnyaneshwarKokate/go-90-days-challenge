package middleware

import (
	"context"
	"time"

	"day-38/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}
		c.Header("X-Request-ID", reqID)

		ctx := context.WithValue(c.Request.Context(), domain.RequestIDKey, reqID)
		c.Request = c.Request.WithContext(ctx)
		c.Set(string(domain.RequestIDKey), reqID)
		c.Next()
	}
}

func StructuredLoggerMiddleware(log domain.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()
		ctx := c.Request.Context()

		keysAndValues := []interface{}{
			"status", statusCode,
			"method", c.Request.Method,
			"path", path,
			"latency_ms", latency.Milliseconds(),
		}

		if statusCode >= 500 {
			log.Error(ctx, "HTTP Request Failed", keysAndValues...)
		} else if statusCode >= 400 {
			log.Warn(ctx, "HTTP Request Warning", keysAndValues...)
		} else {
			log.Info(ctx, "HTTP Request Completed", keysAndValues...)
		}
	}
}
