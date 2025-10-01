.PHONY: run build test clean docker-build docker-run

# Run the application
run:
	go run cmd/api/main.go

# Build the application
build:
	go build -o bin/api cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod download
	go mod tidy

# Docker commands
docker-build:
	docker build -t backend .

docker-run:
	docker run -p 3000:3000 backend

# Development with hot reload (requires air)
dev:
	air

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run
