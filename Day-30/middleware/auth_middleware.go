package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"day-30/domain"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(jwtSecret string, logger domain.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Warn(c.Request.Context(), "Authentication failed: missing Authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": domain.ErrUnauthorized.Error()})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			logger.Warn(c.Request.Context(), "Authentication failed: malformed Authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must be Bearer token"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			logger.Warn(c.Request.Context(), "Authentication failed: invalid or expired token", "error", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired authentication token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logger.Warn(c.Request.Context(), "Authentication failed: invalid token claims payload")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userID, _ := claims["sub"].(string)
		role, _ := claims["role"].(string)
		email, _ := claims["email"].(string)

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, domain.UserIDKey, userID)
		ctx = context.WithValue(ctx, domain.UserRoleKey, role)
		ctx = context.WithValue(ctx, domain.UserEmailKey, email)

		c.Request = c.Request.WithContext(ctx)

		c.Set(string(domain.UserIDKey), userID)
		c.Set(string(domain.UserRoleKey), role)

		c.Next()
	}
}

func RequireRoleMiddleware(logger domain.Logger, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get(string(domain.UserRoleKey))
		if !exists {
			logger.Warn(c.Request.Context(), "Authorization rejected: Role missing from context")
			c.JSON(http.StatusForbidden, gin.H{"error": domain.ErrForbidden.Error()})
			c.Abort()
			return
		}

		userRole, ok := roleVal.(string)
		if !ok {
			logger.Warn(c.Request.Context(), "Authorization rejected: Invalid role type in context")
			c.JSON(http.StatusForbidden, gin.H{"error": domain.ErrForbidden.Error()})
			c.Abort()
			return
		}

		roleAllowed := false
		for _, allowed := range allowedRoles {
			if strings.EqualFold(userRole, allowed) {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			logger.Warn(c.Request.Context(), "Authorization rejected: Access denied for role", "user_role", userRole, "required_roles", allowedRoles)
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("Role '%s' is not authorized to perform this operation", userRole)})
			c.Abort()
			return
		}

		c.Next()
	}
}
