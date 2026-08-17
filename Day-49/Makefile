.PHONY: all fmt lint test build ci

all: ci

fmt:
	go fmt ./...

lint:
	go vet ./...

test:
	go test -v -cover ./...

build:
	go build -o bin/app main.go

ci: fmt lint test build
	@echo "Local CI Pipeline Passed Successfully!"
