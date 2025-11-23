.PHONY: help build run docker-up docker-down docker-logs clean test web-dev web-build

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

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

dev: ## Run backend in development mode
	@echo "Starting backend development server..."
	go run main.go

web-dev: ## Run frontend in development mode
	@echo "Starting frontend development server..."
	cd web && npm run dev

web-build: ## Build frontend for production
	cd web && npm run build

web-install: ## Install frontend dependencies
	cd web && npm install

full-dev: ## Run both backend and frontend in development
	@echo "This will start both backend and frontend..."
	@echo "Backend: http://localhost:3000"
	@echo "Frontend: http://localhost:3001"
	@echo "Run 'make dev' in one terminal and 'make web-dev' in another"
