# Backend API

Production-level Go backend using Fiber framework.

## Project Structure

```
backend/
├── cmd/
│   └── api/              # Application entrypoint
│       └── main.go
├── internal/             # Private application code
│   ├── config/          # Configuration management
│   ├── handlers/        # HTTP request handlers
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data models
│   ├── repository/      # Database layer
│   ├── services/        # Business logic
│   └── utils/           # Helper functions
├── pkg/                 # Public libraries
│   ├── logger/         # Logging utilities
│   └── validator/      # Input validation
├── migrations/         # Database migrations
├── docs/              # Documentation
├── go.mod
└── go.sum
```

## Getting Started

### Prerequisites

- Go 1.24+
- Docker (optional)

### Installation

1. Install dependencies:
```bash
go mod download
```

2. Copy environment file:
```bash
cp .env.example .env
```

3. Update environment variables in `.env`

### Running

```bash
# Run locally
go run cmd/api/main.go

# Build binary
go build -o bin/api cmd/api/main.go

# Run binary
./bin/api
```

### Docker

```bash
# Build image
docker build -t backend .

# Run container
docker run -p 3000:3000 backend
```

## API Endpoints

### Health Check
- `GET /api/v1/health` - Check server status

### Authentication
- `POST /api/v1/auth/signup` - Register new user
- `POST /api/v1/auth/login` - Login user

## Development

### Adding New Features

1. **Models**: Add data structures in `internal/models/`
2. **Handlers**: Add HTTP handlers in `internal/handlers/`
3. **Services**: Add business logic in `internal/services/`
4. **Repository**: Add database operations in `internal/repository/`
5. **Routes**: Register routes in `cmd/api/main.go`

### Code Organization

- Use `internal/` for private code that shouldn't be imported by other projects
- Use `pkg/` for reusable libraries
- Keep handlers thin - move business logic to services
- Keep services database-agnostic - use repository pattern
