# Go Clean Template

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/go-clean-template)](https://goreportcard.com/report/github.com/yourusername/go-clean-template)
[![Go Reference](https://pkg.go.dev/badge/github.com/yourusername/go-clean-template.svg)](https://pkg.go.dev/github.com/yourusername/go-clean-template)
[![Release](https://img.shields.io/github/v/release/yourusername/go-clean-template.svg)](https://github.com/yourusername/go-clean-template/releases)

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
- [x] Swagger documentation (coming soon)
- [x] Rate Limiting (coming soon)

## Quick Start 🚀

```bash
# Get the template
git clone https://github.com/yourusername/go-clean-template.git

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

### Public Routes
```http
POST /api/public/users/register
POST /api/public/users/login
```

### Protected Routes
```http
GET    /api/private/users/:id     # Requires JWT
PUT    /api/private/users/:id     # Requires JWT
DELETE /api/private/users/:id     # Requires JWT
GET    /api/private/users/admin/  # Requires Admin Role
```

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