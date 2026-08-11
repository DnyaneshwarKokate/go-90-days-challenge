# ==============================================================================
# Build Stage (Stage 1)
# ==============================================================================
FROM golang:1.24-alpine AS builder

# Install certificates and build dependencies
RUN apk add --no-cache ca-certificates git

WORKDIR /app

# Step 1: Cache Go modules layer
COPY go.mod ./
RUN go mod download

# Step 2: Copy source files & build static binary
COPY . .
ARG VERSION=1.0.0
ARG COMMIT_SHA=unknown

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X main.Version=${VERSION} -X main.CommitSHA=${COMMIT_SHA}" \
    -o /app/api-server main.go

# ==============================================================================
# Minimal Runtime Stage (Stage 2)
# ==============================================================================
FROM alpine:3.20 AS runner

# Install essential runtime utilities (ca-certificates for HTTPS calls, tzdata for timezone)
RUN apk add --no-cache ca-certificates tzdata wget

# Create non-root system user & group for security hardening
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy binary from build stage
COPY --from=builder /app/api-server /app/api-server

# Enforce file permissions
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Set environment variables
ENV PORT=8080 \
    APP_ENV=production \
    APP_VERSION=1.0.0 \
    DB_HOST=postgres-service

# Expose server port
EXPOSE 8080

# Native Container Healthcheck using wget
HEALTHCHECK --interval=15s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/healthz || exit 1

# Execute container entry point
ENTRYPOINT ["/app/api-server"]
