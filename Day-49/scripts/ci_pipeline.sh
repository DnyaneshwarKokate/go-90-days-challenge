#!/usr/bin/env bash
set -e

echo "=========================================="
echo "      🚀 Starting Go Local CI Pipeline    "
echo "=========================================="

echo "[1/4] Formatting check..."
go fmt ./...

echo "[2/4] Running Go Vet analysis..."
go vet ./...

echo "[3/4] Running Unit Tests with Coverage..."
go test -v -cover ./...

echo "[4/4] Building Binary Artifact..."
mkdir -p bin
go build -o bin/ci_app main.go

echo "=========================================="
echo "  ✅ All CI Pipeline Stages Succeeded!   "
echo "=========================================="
