# Stage 1: Build
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Stage 2: Final image
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .
# Copy migrations for database setup
COPY --from=builder /app/migrations ./migrations

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
