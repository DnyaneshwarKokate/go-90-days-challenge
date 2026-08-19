	package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

type contextKey string

const CorrelationIDKey contextKey = "correlationID"
const HeaderCorrelationID = "X-Correlation-ID"

func generateCorrelationID() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "00000000000000000000000000000000"
	}
	return hex.EncodeToString(bytes)
}

func CorrelationIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get(HeaderCorrelationID)
		if correlationID == "" {
			correlationID = generateCorrelationID()
		}

		ctx := context.WithValue(r.Context(), CorrelationIDKey, correlationID)
		w.Header().Set(HeaderCorrelationID, correlationID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetCorrelationID(ctx context.Context) string {
	if val, ok := ctx.Value(CorrelationIDKey).(string); ok {
		return val
	}
	return ""
}
