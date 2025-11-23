.PHONY: help build run docker-up docker-down docker-logs clean test

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

build: ## Build the Go application
	go build -o sales-api main.go

run: ## Run the application locally
	go run main.go

docker-up: ## Start all services with Docker Compose
	docker-compose up -d

docker-down: ## Stop all services
	docker-compose down

docker-down-v: ## Stop all services and remove volumes
	docker-compose down -v

docker-logs: ## View logs from all services
	docker-compose logs -f

docker-build: ## Build Docker image
	docker-compose build

clean: ## Clean up build artifacts
	rm -f sales-api
	go clean

test: ## Run tests
	go test -v ./...

deps: ## Download dependencies
	go mod download
	go mod tidy

dev: ## Run in development mode
	@echo "Starting development server..."
	go run main.go
