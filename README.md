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

A production-ready template for creating Go services following clean architecture principles.

## Features 🚀

- [x] Clean Architecture with detailed examples
- [x] JWT Authentication & Role-Based Access Control
- [x] End-to-End Encryption for requests/responses
- [x] PostgreSQL with Prisma ORM
- [x] Graceful Shutdown
- [x] Structured Logging with Logrus
- [x] Hot Reload Development
- [x] Docker with Docker Compose
- [x] Make commands for development
- [x] Comprehensive test examples
- [x] Swagger documentation
- [ ] Rate Limiting (coming soon)

## Quick Start 🚀

```bash
# Get the template
git clone https://github.com/ibhanu/go-template.git

# Create your project
mkdir -p your-project && cd your-project
cp -r ../go-clean-template/* .
rm -rf .git && git init

# Run the project
make setup    # Install dependencies and setup database
make dev      # Run with hot reload
```

## Project Structure 📂

```bash
.
├── internal/
│   ├── domain/                 # Enterprise business rules
│   │   ├── entity/            # Enterprise entities
│   │   ├── repository/        # Abstract repositories
│   │   └── constants/         # Domain constants
│   ├── application/           # Application business rules
│   │   └── usecase/          # Use cases
│   ├── infrastructure/        # Frameworks, drivers, tools
│   │   ├── config/           # Configuration
│   │   ├── middleware/       # HTTP middleware
│   │   ├── repository/       # Repository implementations
│   │   └── server/          # HTTP server
│   └── interface/            # Interface adapters
│       └── handler/          # HTTP handlers
├── prisma/                   # Database schema and client
└── scripts/                 # Utility scripts
```

## Layer Dependencies 🎯

![Clean Architecture](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

## Available Commands 🛠

```bash
# Development
make setup           # Initialize project
make dev             # Run with hot reload
make test            # Run tests
make lint            # Run linters

# Database
make db-init         # Initialize database
make prisma-generate # Generate Prisma client
make prisma-studio   # Open Prisma Studio

# Docker
make docker-build    # Build image
make deploy          # Deploy with docker-compose
```

## API Documentation 📚

Our API documentation is available in both Swagger/OpenAPI and traditional format.

### Swagger Documentation

You can access the Swagger UI at:
```
http://localhost:8080/swagger/
```

To view the raw OpenAPI specification:
```
http://localhost:8080/swagger/doc.json
```

### API Routes Overview

#### Public Routes
- `POST /api/public/users/register` - Register a new user
- `POST /api/public/users/login` - Login and get JWT token

#### Protected Routes (Requires JWT)
- `GET /api/private/users/:id` - Get user details
- `PUT /api/private/users/:id` - Update user
- `DELETE /api/private/users/:id` - Delete user

#### Admin Routes (Requires Admin Role)
- `GET /api/private/users/admin` - List all users

For detailed request/response schemas and examples, please refer to the Swagger documentation.

## Environment Variables 🔧

```bash
# Copy example environment file
cp .env.example .env

# Required variables
PORT=8080
DATABASE_URL=postgresql://user:pass@localhost:5432/dbname
JWT_SECRET=your-secret-key
```

## Security 🔒

- JWT-based authentication
- Role-based access control
- Password hashing with bcrypt
- Request/Response encryption
- Secure headers middleware

## Tests ✅

```bash
# Run all tests
make test

# Run with coverage
make test-coverage
```

## Docker 🐳

```bash
# Start all services
make deploy

# View logs
make docker-compose-logs
```

## Contributing 🤝

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.

## License 📝

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments 🙏

- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Gin Web Framework](https://gin-gonic.com/)
- [Prisma](https://www.prisma.io/)