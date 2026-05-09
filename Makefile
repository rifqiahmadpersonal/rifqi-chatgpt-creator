.PHONY: help build build-api build-cli test test-unit test-integration test-coverage lint run-api run-cli docker-up docker-down docker-logs docker-build migrate-up migrate-down migrate-create migrate-version swagger clean frontend-install frontend-dev frontend-build frontend-test frontend-lint check

help:
	@echo "ChatGPT Creator - Build & Development Commands"
	@echo ""
	@echo "Building:"
	@echo "  make build          Build all binaries"
	@echo "  make build-api      Build API server"
	@echo "  make build-cli      Build CLI tool"
	@echo ""
	@echo "Testing:"
	@echo "  make test           Run all tests"
	@echo "  make test-unit      Run unit tests only"
	@echo "  make test-integration Run integration tests"
	@echo "  make test-coverage  Generate coverage report"
	@echo ""
	@echo "Linting:"
	@echo "  make lint           Run linters"
	@echo ""
	@echo "Running:"
	@echo "  make run-api        Run API server"
	@echo "  make run-cli        Run CLI tool"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-up      Start all services"
	@echo "  make docker-down    Stop all services"
	@echo "  make docker-logs    View container logs"
	@echo ""
	@echo "Database:"
	@echo "  make migrate-up     Run migrations"
	@echo "  make migrate-down   Rollback migrations"
	@echo ""
	@echo "Frontend:"
	@echo "  make frontend-install Install frontend deps"
	@echo "  make frontend-dev   Start frontend dev server"
	@echo "  make frontend-build Build frontend"

build: build-api build-cli

build-api:
	@echo "Building API server..."
	@mkdir -p bin
	go build -ldflags="-s -w" -o bin/api cmd/api/main.go

build-cli:
	@echo "Building CLI tool..."
	@mkdir -p bin
	go build -ldflags="-s -w" -o bin/register cmd/register/main.go

test: test-unit

test-unit:
	@echo "Running unit tests..."
	go test -short ./... -v -coverprofile=coverage.out -covermode=atomic

test-integration:
	@echo "Running integration tests..."
	go test -tags=integration ./... -v -coverprofile=coverage-integration.out -covermode=atomic

test-coverage: test
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

benchmark:
	go test -bench=. -benchmem ./...

lint:
	@echo "Running linters..."
	golangci-lint run ./... --timeout=5m

lint-fix:
	golangci-lint run ./... --fix

run-api:
	go run cmd/api/main.go

run-cli:
	go run cmd/register/main.go

docker-up:
	docker-compose up -d --build
	@echo "Services started. View logs: make docker-logs"

docker-down:
	docker-compose down -v

docker-logs:
	docker-compose logs -f

docker-build:
	docker-compose build

migrate-up:
	@echo "Running migrations..."
	migrate -path migrations -database "postgres://chatgpt:chatgpt_secret@localhost:5432/chatgpt_creator?sslmode=disable" up

migrate-down:
	@echo "Rolling back migrations..."
	migrate -path migrations -database "postgres://chatgpt:chatgpt_secret@localhost:5432/chatgpt_creator?sslmode=disable" down

migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

migrate-version:
	migrate -version

swagger:
	@echo "Generating Swagger docs..."
	swag init -g cmd/api/main.go -o docs

swagger-fmt:
	swag fmt

clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean ./...

frontend-install:
	cd frontend && npm install

frontend-dev:
	cd frontend && npm run dev

frontend-build:
	cd frontend && npm run build

frontend-test:
	cd frontend && npm test

frontend-lint:
	cd frontend && npm run lint

check: lint test
	@echo "All checks passed!"
