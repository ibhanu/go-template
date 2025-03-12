#!/bin/bash

check_postgres() {
    if ! command -v psql &> /dev/null; then
        echo "PostgreSQL client (psql) is not installed. Please install PostgreSQL:"
        echo ""
        echo "For macOS:"
        echo "  brew install postgresql@14"
        echo "  brew services start postgresql@14"
        echo ""
        echo "For Ubuntu/Debian:"
        echo "  sudo apt-get update"
        echo "  sudo apt-get install postgresql postgresql-contrib"
        echo "  sudo service postgresql start"
        echo ""
        echo "For Windows:"
        echo "  Download and install from: https://www.postgresql.org/download/windows/"
        echo ""
        echo "After installation, please run 'make db-init' again."
        exit 1
    fi
}

# Check if .env file exists
if [ ! -f .env ]; then
    echo "Creating .env file from example..."
    cp .env.example .env
fi

# Source the .env file
source .env

# Check for PostgreSQL installation
check_postgres

# Extract database name from DATABASE_URL
DB_NAME=$(echo $DATABASE_URL | sed -n 's/.*\/\([^?]*\).*/\1/p')
if [ -z "$DB_NAME" ]; then
    echo "Could not extract database name from DATABASE_URL"
    exit 1
fi

# Extract host from DATABASE_URL
DB_HOST=$(echo $DATABASE_URL | sed -n 's/.*@\([^:]*\).*/\1/p')
if [ -z "$DB_HOST" ]; then
    DB_HOST="localhost"
fi

# Extract user from DATABASE_URL
DB_USER=$(echo $DATABASE_URL | sed -n 's/.*:\/\/\([^:]*\).*/\1/p')
if [ -z "$DB_USER" ]; then
    DB_USER="postgres"
fi

# Extract password from DATABASE_URL
DB_PASS=$(echo $DATABASE_URL | sed -n 's/.*:\/\/[^:]*:\([^@]*\).*/\1/p')

echo "Setting up database..."
echo "Database: $DB_NAME"
echo "Host: $DB_HOST"
echo "User: $DB_USER"

# Set PGPASSWORD environment variable
export PGPASSWORD=$DB_PASS

# Check if PostgreSQL server is running
if ! psql -h $DB_HOST -U $DB_USER -d postgres -c '\q' 2>/dev/null; then
    echo "⚠️  Could not connect to PostgreSQL server. Please make sure it's running."
    echo ""
    echo "For macOS:"
    echo "  brew services start postgresql@14"
    echo ""
    echo "For Ubuntu/Debian:"
    echo "  sudo service postgresql start"
    echo ""
    echo "For Windows:"
    echo "  Start PostgreSQL service from Services"
    exit 1
fi

# Create database if it doesn't exist
echo "Creating database if it doesn't exist..."
psql -h $DB_HOST -U $DB_USER -d postgres -tc "SELECT 1 FROM pg_database WHERE datname = '$DB_NAME'" | grep -q 1 || psql -h $DB_HOST -U $DB_USER -d postgres -c "CREATE DATABASE $DB_NAME"

# Generate Prisma client
echo "Generating Prisma client..."
make prisma-generate

# Push schema changes to database
echo "Pushing schema to database..."
make prisma-db-push

echo "✓ Database initialization complete!"
echo "You can now run 'make dev' to start the server"