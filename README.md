# Go Web Server

A production-ready web server implementation in Go following clean architecture principles, featuring JWT authentication, end-to-end encryption, and Prisma database integration.

## Features

- Clean Architecture
- JWT Authentication
- Role-Based Access Control
- End-to-End Encryption
- Prisma Database Integration
- Hot Reload Development
- Structured Logging
- Graceful Shutdown

## Prerequisites

1. Go 1.24 or higher
2. Make
3. PostgreSQL
   ```bash
   # macOS (using Homebrew)
   brew install postgresql@14
   brew services start postgresql@14

   # Ubuntu/Debian
   sudo apt-get update
   sudo apt-get install postgresql postgresql-contrib
   sudo service postgresql start

   # Windows
   # Download and install from: https://www.postgresql.org/download/windows/
   ```

4. Default PostgreSQL Configuration:
   ```
   Host: localhost
   Port: 5432
   User: postgres
   Password: postgres
   Database: go_server_db (will be created automatically)
   ```

   To use different credentials, update the DATABASE_URL in your .env file after setup.

## Quick Start

1. Clone the repository:
```bash
git clone <repository-url>
cd go-server-1
```

2. Run the setup script (this will):
   - Copy .env.example to .env
   - Install required tools
   - Set up the database
   - Generate Prisma client
```bash
make setup
```

3. Start the development server:
```bash
make dev
```

## Project Structure

```
.
├── internal/
│   ├── domain/         # Business logic and interfaces
│   ├── application/    # Use cases
│   ├── infrastructure/ # External implementations
│   └── interface/      # HTTP handlers
├── prisma/
│   └── schema.prisma   # Database schema
└── scripts/
    └── init-db.sh      # Database initialization
```

[Rest of the README content remains unchanged...]