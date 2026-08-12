.PHONY: help up down build restart logs ps test clean

help:
	@echo "Available commands:"
	@echo "  make up       - Start all Docker Compose containers in detached mode"
	@echo "  make down     - Stop and remove containers, networks, and volumes"
	@echo "  make build    - Rebuild and start containers"
	@echo "  make restart  - Restart all services"
	@echo "  make logs     - Tail container logs"
	@echo "  make ps       - List running service status"
	@echo "  make test     - Run Go unit tests"
	@echo "  make clean    - Remove local build artifacts"

up:
	docker compose up -d

down:
	docker compose down -v

build:
	docker compose up -d --build

restart:
	docker compose restart

logs:
	docker compose logs -f

ps:
	docker compose ps

test:
	go test -v ./...

clean:
	go clean
