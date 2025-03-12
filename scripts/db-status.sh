#!/bin/bash

# Source environment variables if .env exists
if [ -f .env ]; then
    source .env
fi

# Default values if not set in .env
DB_HOST=${DB_HOST:-"localhost"}
DB_PORT=${DB_PORT:-"5432"}
DB_USER=${POSTGRES_USER:-"postgres"}
DB_NAME=${POSTGRES_DB:-"go_server_db"}

# Try to connect to PostgreSQL and get its status
psql "host=$DB_HOST port=$DB_PORT user=$DB_USER dbname=$DB_NAME" -c "SELECT version();" > /dev/null 2>&1

if [ $? -eq 0 ]; then
    echo "✓ Database is running and accessible"
    exit 0
else
    echo "⚠️ Database is not running or not accessible"
    echo "Please check that:"
    echo "  1. PostgreSQL service is running"
    echo "  2. Database credentials in .env are correct"
    echo "  3. Database '$DB_NAME' exists"
    exit 1
fi