# Build Stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server main.go

# Production Stage
FROM alpine:3.19
RUN adduser -D -u 10001 appuser
WORKDIR /app
COPY --from=builder /app/server .
USER 10001
EXPOSE 8080
ENTRYPOINT ["./server"]
