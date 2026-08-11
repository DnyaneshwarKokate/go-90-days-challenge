# ==============================================================================
# Stage 1: Build Stage (Statically Compiles Go Binary)
# ==============================================================================
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache ca-certificates git

# Set working directory inside container
WORKDIR /app

# Copy dependency definition files first (caches layer if go.mod hasn't changed)
COPY go.mod ./
RUN go mod download

# Copy source code into container
COPY . .

# Compile static binary (CGO_ENABLED=0 disables C bindings, -ldflags strip debug info)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /app/server main.go

# ==============================================================================
# Stage 2: Minimal Runtime Stage (Minimal Security Footprint ~15MB)
# ==============================================================================
FROM alpine:3.20 AS runner

# Install SSL root certificates and timezone data
RUN apk add --no-cache ca-certificates tzdata

# Create non-root system user and group for security compliance
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy static binary from build stage
COPY --from=builder /app/server /app/server

# Set permission for non-root execution
RUN chown -R appuser:appgroup /app

# Switch execution user from root (UID 0) to appuser
USER appuser

# Default environment variables
ENV PORT=8080 \
    APP_ENV=production \
    SERVICE_NAME=go-docker-basics \
    VERSION=1.0.0

# Expose server port
EXPOSE 8080

# Define container entrypoint
ENTRYPOINT ["/app/server"]
