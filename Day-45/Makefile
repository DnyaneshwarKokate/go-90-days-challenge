# Day 45 - Dockerizing Go API Makefile

APP_NAME = go-docker-api
IMAGE_TAG = v1.0.0
PORT = 8080

.PHONY: help run test build docker-build docker-build-scratch docker-run docker-stop docker-logs docker-clean

help:
	@echo "Available commands:"
	@echo "  make run                   - Run Go REST API locally"
	@echo "  make test                  - Run unit tests with verbosity"
	@echo "  make build                 - Build Go static binary locally"
	@echo "  make docker-build          - Build Alpine multi-stage Docker image"
	@echo "  make docker-build-scratch  - Build Scratch ultra-minimal Docker image"
	@echo "  make docker-run            - Run container background on port $(PORT)"
	@echo "  make docker-stop           - Stop and remove running container"
	@echo "  make docker-logs           - Stream container logs"
	@echo "  make docker-clean          - Clean container and local images"

run:
	go run main.go

test:
	go test -v -race ./...

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/api-server main.go

docker-build:
	docker build -t $(APP_NAME):$(IMAGE_TAG) .

docker-build-scratch:
	docker build -f Dockerfile.scratch -t $(APP_NAME):scratch .

docker-run:
	docker run -d -p $(PORT):$(PORT) --name $(APP_NAME) $(APP_NAME):$(IMAGE_TAG)

docker-stop:
	docker stop $(APP_NAME) || true
	docker rm $(APP_NAME) || true

docker-logs:
	docker logs -f $(APP_NAME)

docker-clean: docker-stop
	docker rmi $(APP_NAME):$(IMAGE_TAG) || true
	docker rmi $(APP_NAME):scratch || true
