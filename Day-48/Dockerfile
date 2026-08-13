# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install git and ca-certificates
RUN apk add --no-cache git ca-certificates

# Copy go.mod and download dependencies
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /app/go-k8s-app .

# Production Runtime Stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

# Create non-root group and user (UID 10001) for Kubernetes Security Context
RUN addgroup -g 10001 -S appuser && \
    adduser -u 10001 -S appuser -G appuser

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/go-k8s-app /app/go-k8s-app

# Create data directory for PersistentVolume mounts
RUN mkdir -p /data && chown -R appuser:appuser /data /app

USER 10001:10001

EXPOSE 8080

ENTRYPOINT ["/app/go-k8s-app"]
