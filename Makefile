# Go parameters
BINARY_NAME=web-server
MAIN_FILE=main.go

# Build the project
build:
	go build -o ${BINARY_NAME} ${MAIN_FILE}

# Run the project with normal go run
run:
	go run ${MAIN_FILE}

# Run the project with hot reload
dev:
	@if [ ! -f ./bin/air ]; then \
		echo "Air not found. Running setup..." && \
		$(MAKE) setup; \
	fi
	@echo "Starting server with hot reload..."
	@./bin/air

# Clean build files
clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -rf ./tmp

# Download dependencies
deps:
	@echo "Downloading Go dependencies..."
	@go mod download
	@echo "✓ Dependencies downloaded successfully"

# Run tests
test:
	go test ./...

# Run tests with coverage
test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# Generate mocks (requires mockgen)
mocks:
	go generate ./...

# Format code
fmt:
	go fmt ./...

# Check for linting issues
lint:
	go vet ./...
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint is not installed"; \
	fi

# Prisma commands
prisma-generate:
	@echo "Generating Prisma client..."
	@go run github.com/steebchen/prisma-client-go generate

prisma-db-push:
	@echo "Pushing schema to database..."
	@go run github.com/steebchen/prisma-client-go db push

prisma-studio:
	@echo "Starting Prisma Studio..."
	@go run github.com/steebchen/prisma-client-go studio

# Database initialization
db-init: deps
	@echo "Initializing database..."
	@./scripts/init-db.sh || (echo "⚠️  Database initialization failed. Make sure PostgreSQL is running and .env is configured properly." && exit 1)

# Setup development environment
setup: deps
	@echo "Setting up development environment..."
	@mkdir -p bin
	@if [ ! -f ./bin/air ]; then \
		echo "Installing Air for hot reload..." && \
		curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b ./bin; \
	fi
	@if [ ! -f .env ]; then \
		echo "Creating .env file from example..." && \
		cp .env.example .env && \
		echo "⚠️  Please update your .env file with proper values"; \
	else \
		echo "✓ .env file already exists"; \
	fi
	@mkdir -p tmp
	@$(MAKE) db-init || (echo "⚠️  Setup failed. Please check the error messages above." && exit 1)
	@echo "✓ Development environment setup complete"
	@echo "Run 'make dev' to start the server with hot reload"

# Docker commands
docker-build:
@echo "Building Docker image..."
@docker build -t go-server .

docker-run:
@echo "Running Docker container..."
@docker run -p 8080:8080 --env-file .env go-server

docker-compose-up:
@echo "Starting services with Docker Compose..."
@docker-compose up -d

docker-compose-down:
@echo "Stopping services..."
@docker-compose down

docker-compose-logs:
@docker-compose logs -f

docker-clean:
@echo "Cleaning Docker resources..."
@docker-compose down -v
@docker rmi go-server

# Production deployment
deploy: docker-build docker-compose-up
@echo "✓ Application deployed successfully"
@echo "Run 'make docker-compose-logs' to view logs"

.PHONY: build run dev clean deps test test-coverage mocks fmt lint setup \
prisma-generate prisma-db-push prisma-studio db-init \
docker-build docker-run docker-compose-up docker-compose-down docker-compose-logs docker-clean deploy