.PHONY: test lint build docker-build

test:
	go test -v -cover ./...

lint:
	go vet ./...

build:
	go build -o server main.go

docker-build:
	docker build -t day50-app:latest .
