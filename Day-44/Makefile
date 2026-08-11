APP_NAME=go-docker-basics
PORT=8080

.PHONY: help build run test docker-build docker-run docker-logs docker-stop clean

help:
	@echo "Available targets:"
	@echo "  make build        - Compile local Go binary"
	@echo "  make run          - Run Go service locally"
	@echo "  make test         - Run unit tests"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container on port $(PORT)"
	@echo "  make docker-logs  - View container logs"
	@echo "  make docker-stop  - Stop and remove container"
	@echo "  make clean        - Clean build artifacts"

build:
	go build -o server main.go

run:
	go run main.go

test:
	go test -v -race ./...

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run -d -p $(PORT):$(PORT) --name $(APP_NAME)-app $(APP_NAME):latest

docker-logs:
	docker logs -f $(APP_NAME)-app

docker-stop:
	docker stop $(APP_NAME)-app && docker rm $(APP_NAME)-app

clean:
	rm -f server
