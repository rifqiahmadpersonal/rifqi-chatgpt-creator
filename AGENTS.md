# AGENTS.md - AI Assistant Instructions

## Project Overview
ChatGPT Account Registration Bot - Full-stack application with Go backend (Gin), Next.js frontend, PostgreSQL database.

## Build Commands
- Build backend: `make build` or `go build -o bin/api cmd/api/main.go`
- Build CLI: `go build -o bin/register cmd/register/main.go`
- Build frontend: `cd frontend && npm run build`

## Test Commands
- Run all tests: `make test` or `go test ./... -v -coverprofile=coverage.out`
- Run unit tests only: `go test -short ./...`
- Run integration tests: `go test -tags=integration ./...`
- Frontend tests: `cd frontend && npm test`
- Test coverage: `make test-coverage`

## Lint Commands
- Go lint: `make lint` or `golangci-lint run ./...`
- Frontend lint: `cd frontend && npm run lint`

## Database Commands
- Run migrations: `make migrate-up`
- Rollback migrations: `make migrate-down`
- Create migration: `migrate create -ext sql -dir migrations -seq <name>`

## Development Commands
- Start API: `make run-api`
- Start CLI: `make run-cli`
- Start frontend: `make frontend-dev`
- Docker up: `make docker-up`
- Docker down: `make docker-down`

## Project Structure
```
├── cmd/
│   ├── api/           # API server entry point
│   └── register/      # CLI entry point
├── internal/
│   ├── api/           # HTTP handlers, routes, middleware
│   ├── models/        # Database models
│   ├── repository/    # Database access layer
│   ├── service/       # Business logic layer
│   ├── register/      # Registration flow (preserved)
│   ├── email/         # Email generation (preserved)
│   ├── config/        # Configuration
│   ├── chrome/        # TLS fingerprinting (preserved)
│   ├── sentinel/      # Anti-bot tokens (preserved)
│   ├── util/          # Utility functions
│   └── websocket/     # Real-time updates
├── migrations/        # Database migrations
├── frontend/          # Next.js application
├── docs/              # OpenAPI documentation
├── scripts/           # Setup and utility scripts
└── deploy/            # Deployment configurations
```

## Code Style
- Go: Follow Effective Go, use gofmt, goimports
- TypeScript: ESLint + Prettier, strict mode
- Commits: Conventional Commits (feat, fix, chore, test, docs)

## Environment Variables
See .env.example for all required configuration.

## API Endpoints (when implemented)
- `GET /api/health` - Health check
- `GET /api/accounts` - List accounts
- `POST /api/accounts` - Create account
- `GET /api/email-domains` - List email domains
- `POST /api/email-domains` - Add email domain
- `POST /api/batch-jobs` - Start batch registration
- `GET /api/batch-jobs/:id` - Get batch job status
- `GET /api/configurations` - List configurations
- `PUT /api/configurations/:key` - Update configuration
