package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	secret []byte
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{secret: []byte(secret)}
}

func (a *AuthMiddleware) ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip authentication for public register/login routes
		path := r.URL.Path
		if strings.HasSuffix(path, "/register") || strings.HasSuffix(path, "/login") || path == "/health" {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "API Gateway: Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		parts := strings.Split(tokenStr, ".")
		if len(parts) != 3 {
			http.Error(w, "API Gateway: Malformed JWT token", http.StatusUnauthorized)
			return
		}

		// Verify signature
		unsignedToken := parts[0] + "." + parts[1]
		h := hmac.New(sha256.New, a.secret)
		h.Write([]byte(unsignedToken))
		expectedSig := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

		if !hmac.Equal([]byte(parts[2]), []byte(expectedSig)) {
			http.Error(w, "API Gateway: Token signature verification failed", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
