APP_NAME=go-k8s-app
IMAGE_NAME=go-k8s-app:v1.0.0
NAMESPACE=go-app

.PHONY: help build run test test-coverage docker-build k8s-apply k8s-delete k8s-status k8s-logs

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}'

build: ## Build binary locally
	go build -o bin/$(APP_NAME) .

run: build ## Build and run locally
	./bin/$(APP_NAME)

test: ## Run unit tests
	go test -v ./...

test-coverage: ## Run unit tests with coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

docker-build: ## Build Docker container image
	docker build -t $(IMAGE_NAME) .

k8s-apply: ## Deploy all Kubernetes manifests to cluster
	kubectl apply -f k8s/

k8s-delete: ## Delete all deployed Kubernetes resources
	kubectl delete -f k8s/

k8s-status: ## Check status of Pods, Deployments, Services, PVC, and HPA
	kubectl get all,pvc,hpa,ingress -n $(NAMESPACE)

k8s-logs: ## Stream logs from application Pods
	kubectl logs -n $(NAMESPACE) -l app=go-k8s-app -f
