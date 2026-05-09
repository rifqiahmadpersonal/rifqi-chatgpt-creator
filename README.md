# ChatGPT Account Registration Bot

Full-stack application with Go backend (Gin), Next.js frontend, PostgreSQL database. Features concurrent workers, TLS fingerprint spoofing, automatic email generation, OTP verification, and retry-until-success logic.

## Features

- **Concurrent Registration** — Configurable worker pool for parallel account creation
- **TLS Fingerprinting** — Randomized Chrome TLS profiles to avoid detection
- **Auto Email Generation** — Generates temporary emails via generator.email or custom domains
- **OTP Verification** — Automatic email OTP retrieval and validation
- **Retry Loop** — Automatically retries failed registrations until target count is reached
- **Proxy Support** — Optional HTTP/SOCKS proxy for all requests
- **Configurable Email Domains** — Add, remove, and prioritize email domains via UI
- **Batch Job Management** — Start, stop, and monitor registration jobs
- **PostgreSQL Database** — Persistent storage with full relational data
- **REST API** — Full CRUD operations for all entities
- **Real-time Dashboard** — Monitor registrations in real-time

## Tech Stack

### Backend
- Go 1.25+
- Gin (HTTP framework)
- PostgreSQL
- sqlx (database access)
- golang-migrate (migrations)
- Swagger/OpenAPI

### Frontend
- Next.js 15
- TypeScript
- Tailwind CSS
- React Query

## Quick Start

### Prerequisites
- Go 1.25+
- Node.js 20+
- PostgreSQL 16+
- Docker & docker-compose (optional)

### Using Docker (Recommended)

```bash
# Clone the repository
git clone https://github.com/rifqiahmadpersonal/rifqi-chatgpt-creator.git
cd rifqi-chatgpt-creator

# Copy environment file
cp .env.example .env

# Start all services
docker compose --env-file .env.docker up -d

# View logs
docker compose logs -f
```

Access:
- Frontend: http://localhost:3000
- API: http://localhost:8080
- Swagger: http://localhost:8080/swagger/index.html

### Manual Setup

```bash
# Backend
go mod download
go run cmd/api/main.go

# Frontend (separate terminal)
cd frontend
npm install
npm run dev
```

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
│   ├── register/      # Registration flow
│   ├── email/         # Email generation
│   ├── config/        # Configuration
│   ├── chrome/        # TLS fingerprinting
│   ├── sentinel/      # Anti-bot tokens
│   └── util/          # Utility functions
├── migrations/        # Database migrations
├── frontend/          # Next.js application
├── docs/              # OpenAPI documentation
├── scripts/           # Setup and utility scripts
└── deploy/            # Deployment configurations
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/health` | Health check |
| GET | `/api/accounts` | List accounts |
| POST | `/api/accounts` | Create account |
| DELETE | `/api/accounts/:id` | Delete account |
| GET | `/api/email-domains` | List email domains |
| POST | `/api/email-domains` | Add email domain |
| PUT | `/api/email-domains/:id` | Update email domain |
| DELETE | `/api/email-domains/:id` | Delete email domain |
| POST | `/api/batch-jobs` | Create batch job |
| GET | `/api/batch-jobs` | List batch jobs |
| POST | `/api/batch-jobs/:id/start` | Start batch job |
| POST | `/api/batch-jobs/:id/stop` | Stop batch job |
| GET | `/api/configurations` | List configurations |
| PUT | `/api/configurations/:key` | Update configuration |
| GET | `/api/stats/dashboard` | Dashboard statistics |

## Configuration

Environment variables (see `.env.example`):

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_PORT` | API server port | 8080 |
| `DB_HOST` | PostgreSQL host | localhost |
| `DB_PORT` | PostgreSQL port | 5432 |
| `DB_USER` | Database user | chatgpt |
| `DB_PASSWORD` | Database password | - |
| `DB_NAME` | Database name | chatgpt_creator |
| `WORKER_POOL_SIZE` | Max concurrent workers | 5 |
| `DEFAULT_PROXY` | Default proxy URL | - |
| `DEFAULT_DOMAIN` | Default email domain | - |

## Database Migrations

```bash
# Run migrations
make migrate-up

# Rollback
make migrate-down

# Create new migration
migrate create -ext sql -dir migrations -seq <name>
```

## Development

```bash
# Run API
make run-api

# Run CLI
make run-cli

# Run tests
make test

# Run linter
make lint

# Generate swagger docs
make swagger
```

## Frontend Development

```bash
cd frontend
npm install
npm run dev     # Development server
npm run build   # Production build
npm run lint    # Run linter
npm test        # Run tests
```

## Docker Commands

```bash
# Start all services
docker compose up -d

# With optional services
docker compose --profile redis up -d
docker compose --profile frontend up -d

# View logs
docker compose logs -f api
docker compose logs -f frontend

# Stop services
docker compose down

# Remove volumes
docker compose down -v
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Disclaimer

This tool is provided for educational and research purposes only. Use of this tool to create accounts in violation of OpenAI's Terms of Service is solely at your own risk.
