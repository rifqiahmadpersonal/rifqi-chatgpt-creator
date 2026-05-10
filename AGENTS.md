# AGENTS.md - AI Assistant Instructions

## Project Overview
ChatGPT Account Registration Bot - Full-stack application with Go backend (Gin), Next.js frontend, PostgreSQL database.

> **Disclaimer**: This tool is for educational/research purposes only. Use at your own risk regarding OpenAI's Terms of Service.

## Build Commands
- Build all: `make build` (API + CLI)
- Build API: `make build-api`
- Build CLI: `make build-cli`
- Build frontend: `cd frontend && npm run build`

## Test Commands
- Run unit tests: `make test` or `go test -short ./... -v -coverprofile=coverage.out`
- Run integration tests: `make test-integration` or `go test -tags=integration ./...`
- Run frontend tests: `cd frontend && npm test`
- Generate coverage: `make test-coverage` (outputs coverage.html)

## Lint Commands
- Go lint: `make lint` or `golangci-lint run ./... --timeout=5m`
- Auto-fix Go: `make lint-fix`
- Frontend lint: `cd frontend && npm run lint`

## Database Commands
- Run migrations: `make migrate-up`
- Rollback migrations: `make migrate-down`
- Create migration: `make migrate-create` (prompts for name)
- Check version: `make migrate-version`

## Docker Commands
- Start all: `make docker-up` (API + PostgreSQL)
- Start with Redis: `docker-compose --profile redis up -d`
- Start with frontend: `docker-compose --profile frontend up -d`
- View logs: `make docker-logs`
- Stop: `make docker-down`

## Development Commands
- Run API: `make run-api` (go run cmd/api/main.go)
- Run CLI: `make run-cli` (go run cmd/register/main.go)
- Run frontend: `cd frontend && npm run dev`
- All checks: `make check` (lint + test)
- Generate swagger: `make swagger`

## Project Structure
```
├── cmd/
│   ├── api/           # API server (Gin)
│   └── register/      # CLI registration tool
├── internal/
│   ├── api/           # HTTP handlers, routes, middleware
│   ├── models/        # Database models (sqlx)
│   ├── repository/    # Data access layer
│   ├── service/       # Business logic
│   ├── register/      # Registration automation
│   ├── email/         # Email/OTP handling
│   ├── chrome/        # TLS fingerprint spoofing
│   ├── sentinel/      # Anti-bot token generation
│   ├── websocket/     # Real-time updates
│   └── config/        # Configuration management
├── migrations/        # SQL migrations
├── frontend/          # Next.js 15 + React + Tailwind
├── docs/              # Swagger/OpenAPI
└── scripts/           # Utility scripts
```

## Key Technologies
- **Backend**: Go 1.25+, Gin, sqlx, golang-migrate, PostgreSQL
- **Frontend**: Next.js 15, TypeScript, Tailwind CSS, React Query
- **Optional**: Redis (job queue), WebSocket (real-time updates)

## Code Style
- Go: Follow Effective Go, gofmt, goimports
- TypeScript: ESLint + Prettier, strict mode enabled
- Commit messages: Conventional Commits (feat/fix/chore/test/docs)

## Environment Variables
All config via `.env` (see `.env.example`):
- Database: `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`
- Registration: `DEFAULT_PROXY`, `DEFAULT_PASSWORD`, `WORKER_POOL_SIZE`, `MAX_RETRIES`
- Redis: `REDIS_ENABLED` (default false)
- Frontend: `NEXT_PUBLIC_API_URL` (default http://localhost:8080)
