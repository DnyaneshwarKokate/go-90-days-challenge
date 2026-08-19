.PHONY: test lint run build

test:
	go test -v -cover ./...

lint:
	go vet ./...

run:
	go run main.go

build:
	go build -o bin/microservice main.go
