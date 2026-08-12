.PHONY: help test k8s-apply k8s-delete k8s-status build

help:
	@echo "Available commands:"
	@echo "  make test        - Run Go unit tests"
	@echo "  make k8s-apply   - Apply all Kubernetes manifests to cluster"
	@echo "  make k8s-delete  - Delete all Kubernetes manifests from cluster"
	@echo "  make k8s-status  - View status of pods, services, and HPA"
	@echo "  make build       - Build local Docker image for K8s deployment"

test:
	go test -v ./...

k8s-apply:
	kubectl apply -f k8s/

k8s-delete:
	kubectl delete -f k8s/

k8s-status:
	kubectl get pods,services,configmaps,secrets,hpa -o wide

build:
	docker build -t go-api-service:v1 .
