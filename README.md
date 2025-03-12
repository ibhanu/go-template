# Go Clean Template

[![Build Status](https://github.com/ibhanu/go-template/workflows/CI/badge.svg)](https://github.com/ibhanu/go-template/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/ibhanu/go-template)](https://goreportcard.com/report/github.com/ibhanu/go-template)
[![Go Reference](https://pkg.go.dev/badge/github.com/ibhanu/go-template.svg)](https://pkg.go.dev/github.com/ibhanu/go-template)
[![Release](https://img.shields.io/github/v/release/ibhanu/go-template.svg)](https://github.com/ibhanu/go-template/releases)
[![codecov](https://codecov.io/gh/ibhanu/go-template/branch/main/graph/badge.svg)](https://codecov.io/gh/ibhanu/go-template)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ibhanu/go-template)](https://go.dev/)
[![License](https://img.shields.io/github/license/ibhanu/go-template)](LICENSE)
[![Issues](https://img.shields.io/github/issues/ibhanu/go-template)](https://github.com/ibhanu/go-template/issues)
[![Pull Requests](https://img.shields.io/github/issues-pr/ibhanu/go-template)](https://github.com/ibhanu/go-template/pulls)
[![Last Commit](https://img.shields.io/github/last-commit/ibhanu/go-template)](https://github.com/ibhanu/go-template/commits/main)

A production-ready template for creating Go services following clean architecture principles. This template provides a robust foundation for building scalable, maintainable, and well-tested microservices in Go.

## Why This Template? 🤔

Building production-grade microservices requires more than just writing code. You need:
- A solid architectural foundation
- Security best practices
- Comprehensive testing
- Performance optimization
- Developer-friendly tooling

This template provides all these essentials out of the box, saving you weeks of setup time and helping you follow established best practices.

## Features 🚀

### Architecture & Design
- [x] **Clean Architecture** implementation with clear layer separation
- [x] Domain-Driven Design (DDD) principles
- [x] SOLID principles and best practices
- [x] Dependency injection pattern
- [x] Interface-driven design

### Security
- [x] **JWT Authentication** with role-based access control
- [x] End-to-End request/response encryption
- [x] Secure password hashing with bcrypt
- [x] HTTP security headers
- [x] Rate limiting (coming soon)
- [x] Input validation and sanitization

### Database & ORM
- [x] **PostgreSQL** integration
- [x] Prisma ORM with type-safe queries
- [x] Migration management
- [x] Connection pooling
- [x] Transaction support

### Development Experience
- [x] **Hot Reload** development mode
- [x] Comprehensive Makefile commands
- [x] Docker & Docker Compose setup
- [x] Linting and formatting tools
- [x] Git hooks for code quality
- [x] Debugging configurations

### Testing & Quality
- [x] **100% test coverage** requirement
- [x] Unit test examples
- [x] Integration test examples
- [x] Benchmark tests
- [x] Mocking examples
- [x] CI/CD pipeline with GitHub Actions

### Documentation
- [x] **Swagger/OpenAPI** documentation
- [x] Godoc comments
- [x] Architecture decision records
- [x] API usage examples
- [x] Contributing guidelines

### Operations
- [x] **Graceful shutdown** handling
- [x] Structured logging with Logrus
- [x] Metrics and monitoring setup
- [x] Health check endpoints
- [x] Multi-platform builds
- [x] Container orchestration

## Prerequisites 📋

- Go 1.24 or higher
- Docker and Docker Compose
- PostgreSQL 14 or higher
- Make

## Quick Start 🚀

### 1. Get the Template

```bash
# Clone the repository
git clone https://github.com/ibhanu/go-template.git

# Create your project
mkdir -p your-project && cd your-project
cp -r ../go-clean-template/* .
rm -rf .git && git init
```

### 2. Configure Environment

```bash
# Copy and edit environment variables
cp .env.example .env

# Edit .env with your settings:
# - Database credentials
# - JWT secret
# - API configurations
```

### 3. Setup and Run

```bash
# Install dependencies and setup database
make setup

# Generate Prisma client
make prisma-generate

# Start development server with hot reload
make dev
```

Visit [http://localhost:8080/swagger/](http://localhost:8080/swagger/) to explore the API.

## Project Structure 📂

```bash
.
├── internal/                 # Application code
│   ├── domain/              # Enterprise business rules
│   │   ├── entity/          # Business entities
│   │   ├── repository/      # Repository interfaces
│   │   └── constants/       # Domain constants
│   ├── application/         # Application business rules
│   │   └── usecase/        # Use case implementations
│   ├── infrastructure/      # External tools and frameworks
│   │   ├── config/         # Configuration
│   │   ├── middleware/     # HTTP middleware
│   │   ├── repository/     # Repository implementations
│   │   └── server/        # HTTP server setup
│   └── interface/          # Interface adapters
│       └── handler/        # HTTP handlers
├── prisma/                 # Database schema and client
├── scripts/               # Utility scripts
└── docs/                  # Documentation
```

## Architecture Overview 🏗

### Clean Architecture Layers

1. **Domain Layer** (innermost)
   - Contains enterprise business rules
   - Pure Go without external dependencies
   - Defines interfaces and entities

2. **Application Layer**
   - Implements use cases
   - Orchestrates domain objects
   - Contains business logic

3. **Interface Layer**
   - Adapts data between layers
   - Handles HTTP requests
   - Manages serialization

4. **Infrastructure Layer** (outermost)
   - Implements interfaces
   - Integrates external services
   - Manages technical concerns

![Clean Architecture](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

## Development Commands 🛠

### Core Commands

```bash
# Development
make setup           # Initialize project
make dev             # Run with hot reload
make build           # Build binary
make run             # Run binary
make clean           # Clean build files

# Testing
make test            # Run all tests
make test-coverage   # Run tests with coverage
make test-race       # Run tests with race detection
make benchmark       # Run benchmark tests

# Code Quality
make lint           # Run linters
make fmt            # Format code
make vet            # Run go vet
make cyclo          # Check cyclomatic complexity

# Database
make db-init        # Initialize database
make db-migrate     # Run migrations
make prisma-generate # Generate Prisma client
make prisma-studio  # Open Prisma Studio

# Docker
make docker-build   # Build image
make docker-push    # Push to registry
make deploy         # Deploy with docker-compose
make docker-logs    # View container logs
```

### Advanced Commands

```bash
# Testing
make test-integration  # Run integration tests
make test-e2e         # Run end-to-end tests
make test-stress      # Run stress tests

# Documentation
make docs             # Generate documentation
make swagger          # Generate Swagger specs

# Maintenance
make clean-docker     # Clean Docker resources
make reset-db         # Reset database
make generate-keys    # Generate JWT keys
```

## API Documentation 📚

### Authentication

All protected routes require a JWT token in the Authorization header:

```
Authorization: Bearer <your-token>
```

### Public Routes

#### Register User
```http
POST /api/public/users/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepass",
  "name": "John Doe"
}
```

#### Login
```http
POST /api/public/users/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepass"
}
```

### Protected Routes

#### Get User Details
```http
GET /api/private/users/:id
Authorization: Bearer <token>
```

#### Update User
```http
PUT /api/private/users/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Updated Name",
  "email": "new@example.com"
}
```

#### Delete User
```http
DELETE /api/private/users/:id
Authorization: Bearer <token>
```

### Admin Routes

#### List All Users
```http
GET /api/private/users/admin
Authorization: Bearer <token>
```

## Environment Configuration 🔧

```bash
# Server Configuration
PORT=8080
ENV=development
LOG_LEVEL=debug
CORS_ALLOWED_ORIGINS=*

# Database Configuration
DATABASE_URL=postgresql://user:pass@localhost:5432/dbname
DB_MAX_CONNECTIONS=100
DB_IDLE_TIMEOUT=300

# Security
JWT_SECRET=your-secret-key
JWT_EXPIRY=24h
ENCRYPTION_KEY=32-byte-encryption-key
RATE_LIMIT=100
```

## Security Considerations 🔒

### Authentication & Authorization
- JWT-based authentication
- Role-based access control (RBAC)
- Token refresh mechanism
- Session management

### Data Protection
- Password hashing with bcrypt
- Request/Response encryption
- HTTPS enforcement
- XSS protection
- CSRF protection

### Infrastructure
- Rate limiting
- Secure headers
- Input validation
- SQL injection prevention
- Error handling security

## Performance Optimization 🚄

### Database
- Connection pooling
- Query optimization
- Indexed lookups
- Efficient pagination

### Application
- Response caching
- Compressed responses
- Optimized routing
- Memory management

### Monitoring
- Performance metrics
- Resource utilization
- Response times
- Error rates

## Error Handling & Logging 🔍

### Error Types
- Domain errors
- Application errors
- Infrastructure errors
- HTTP errors

### Logging
- Structured logging
- Log levels
- Request ID tracking
- Error context

## Testing Strategy ✅

### Unit Tests
- Domain logic
- Use cases
- Utilities
- Middleware

### Integration Tests
- API endpoints
- Database operations
- External services

### Performance Tests
- Load testing
- Stress testing
- Benchmarks

## Deployment 🚀

### Docker Deployment

```bash
# Build and run with Docker Compose
make deploy

# Scale services
docker-compose up -d --scale app=3

# View logs
make docker-compose-logs
```

### Kubernetes Deployment

```bash
# Apply Kubernetes manifests
kubectl apply -f k8s/

# Scale deployment
kubectl scale deployment app --replicas=3

# View pods
kubectl get pods
```

## Troubleshooting 🔧

### Common Issues

1. **Database Connection**
   ```bash
   # Check database status
   make db-status
   
   # Reset database
   make reset-db
   ```

2. **Permission Issues**
   ```bash
   # Fix file permissions
   chmod +x scripts/*
   ```

3. **Build Errors**
   ```bash
   # Clean and rebuild
   make clean && make build
   ```

## Contributing 🤝

We welcome contributions! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on:
- Code of Conduct
- Pull Request Process
- Development Guidelines
- Testing Requirements
- Documentation Standards

## License 📝

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments 🙏

- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Gin Web Framework](https://gin-gonic.com/)
- [Prisma](https://www.prisma.io/)
- Open source community

## Support 💬

- [Report Issues](https://github.com/ibhanu/go-template/issues)
- [Discussions](https://github.com/ibhanu/go-template/discussions)
- [Security](SECURITY.md)