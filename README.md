# Go Clean Architecture Template

A production-ready Go server template following clean architecture principles with JWT authentication, end-to-end encryption, and Prisma database integration.

## Features

- 🏗️ Clean Architecture
- 🔐 JWT Authentication
- 👮 Role-Based Access Control
- 🔒 End-to-End Encryption
- 📦 Prisma Database Integration
- 🔄 Hot Reload Development
- 📝 Structured Logging
- 🛑 Graceful Shutdown

## Template Usage Sequence

### 1. Prerequisites Installation

```bash
# 1. Install Go (1.24 or higher)
go version

# 2. Install PostgreSQL
# macOS
brew install postgresql@14
brew services start postgresql@14

# Ubuntu/Debian
sudo apt-get update
sudo apt-get install postgresql postgresql-contrib
sudo service postgresql start

# Windows
# Download from: https://www.postgresql.org/download/windows/
```

### 2. Project Setup

```bash
# 1. Clone this template
git clone <repository-url> your-project-name
cd your-project-name

# 2. Initialize your new Git repository
rm -rf .git
git init
git add .
git commit -m "Initial commit from template"

# 3. Update module name in go.mod
# Replace "web-server" with your module name
sed -i '' 's/module web-server/module your-module-name/' go.mod

# 4. Install dependencies
make deps
```

### 3. Database Configuration

```bash
# 1. Set up environment variables
cp .env.example .env

# 2. Update database credentials in .env if different from defaults:
# DATABASE_URL="postgresql://postgres:postgres@localhost:5432/your_db_name?schema=public"

# 3. Initialize database
make db-init
```

### 4. Security Setup

```bash
# 1. Generate JWT secret
openssl rand -base64 32 > jwt_secret.txt

# 2. Generate encryption keys
openssl rand -base64 32 > encryption_key.txt
openssl rand -base64 12 > encryption_nonce.txt

# 3. Update .env with these values
# Copy contents from the generated files to respective .env variables
```

### 5. Development Workflow

```bash
# Start development server with hot reload
make dev

# Run tests
make test

# Format code
make fmt

# Run linter
make lint
```

## Project Structure

```plaintext
.
├── internal/
│   ├── domain/          # Business logic and interfaces
│   │   ├── entity/      # Domain entities
│   │   ├── repository/  # Repository interfaces
│   │   └── constants/   # Domain constants
│   ├── application/     # Use cases
│   ├── infrastructure/  # External implementations
│   │   ├── config/      # Configuration
│   │   ├── middleware/  # HTTP middleware
│   │   ├── repository/  # Repository implementations
│   │   └── server/      # Server setup
│   └── interface/       # Interface adapters
│       └── handler/     # HTTP handlers
├── prisma/             # Database schema and client
└── scripts/            # Utility scripts
```

## API Endpoints

### Public Routes
```http
POST /api/public/users/register
POST /api/public/users/login
```

### Protected Routes (Requires JWT)
```http
GET    /api/private/users/:id
PUT    /api/private/users/:id
DELETE /api/private/users/:id
```

### Admin Routes (Requires Admin Role)
```http
GET /api/private/users/admin/
```

## Available Make Commands

### Development
- `make dev` - Start server with hot reload
- `make setup` - Initialize development environment
- `make fmt` - Format code
- `make lint` - Run linters

### Database
- `make db-init` - Initialize database
- `make prisma-generate` - Generate Prisma client
- `make prisma-db-push` - Push schema changes
- `make prisma-studio` - Open Prisma Studio

### Testing
- `make test` - Run tests
- `make test-coverage` - Run tests with coverage

## Customization Guide

1. Update Domain Entities
   - Modify `/internal/domain/entity/` files
   - Update Prisma schema in `prisma/schema.prisma`
   - Run `make prisma-generate` after schema changes

2. Add New Features
   - Create new entities in `/internal/domain/entity/`
   - Add repository interfaces in `/internal/domain/repository/`
   - Implement use cases in `/internal/application/usecase/`
   - Add handlers in `/internal/interface/handler/`

3. Extend Middleware
   - Add new middleware in `/internal/infrastructure/middleware/`
   - Configure in `/internal/infrastructure/server/server.go`

## Production Deployment

1. Environment Setup
```bash
# Set production values in .env
ENV=production
PORT=8080
```

2. Build for Production
```bash
make build
```

3. Run Server
```bash
./web-server
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

[MIT License](LICENSE)

## Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/)
- [Prisma](https://www.prisma.io/)
- [JWT-Go](https://github.com/golang-jwt/jwt)