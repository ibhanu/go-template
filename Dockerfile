# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install prisma-client-go
RUN go install github.com/steebchen/prisma-client-go@latest

# Copy source code
COPY . .

# Generate Prisma client and build
RUN make prisma-generate && make build

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/web-server .
COPY --from=builder /app/.env.example .env

# Expose port
EXPOSE 8080

# Command to run
CMD ["./web-server"]