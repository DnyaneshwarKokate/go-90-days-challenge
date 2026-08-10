package middleware

import (
	"context"
	"time"

	"day-30/domain"

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
		rawQuery := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		fullPath := path
		if rawQuery != "" {
			fullPath = path + "?" + rawQuery
		}

		keysAndValues := []interface{}{
			"status", statusCode,
			"method", method,
			"path", fullPath,
			"ip", clientIP,
			"latency_ms", latency.Milliseconds(),
			"latency_human", latency.String(),
			"user_agent", c.Request.UserAgent(),
		}

		if errorMessage != "" {
			keysAndValues = append(keysAndValues, "error", errorMessage)
		}

		ctx := c.Request.Context()

		if statusCode >= 500 {
			log.Error(ctx, "HTTP Request Failed (Server Error)", keysAndValues...)
		} else if statusCode >= 400 {
			log.Warn(ctx, "HTTP Request Warning (Client Error)", keysAndValues...)
		} else {
			log.Info(ctx, "HTTP Request Completed", keysAndValues...)
		}
	}
}
