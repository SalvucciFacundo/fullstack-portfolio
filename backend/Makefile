.PHONY: run test migrate-up migrate-down build docker-up docker-down

# Variables
BINARY_NAME=api
CMD_PATH=./cmd/api/main.go

run:
	go run $(CMD_PATH)

build:
	go build -o bin/$(BINARY_NAME) $(CMD_PATH)

test:
	go test ./...

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Placeholder for migrations (will be updated when tools are installed)
migrate-up:
	@echo "Corriendo migraciones (Up)..."

migrate-down:
	@echo "Corriendo migraciones (Down)..."
