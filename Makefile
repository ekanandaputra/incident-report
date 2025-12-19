.PHONY: help build run clean test deps fmt lint

# Variables
APP_NAME=incident-report
MAIN_PATH=cmd/main.go
BINARY_PATH=./$(APP_NAME)

# Help command
help:
	@echo "Available commands:"
	@echo "  make build      - Build the application"
	@echo "  make run        - Run the application"
	@echo "  make clean      - Clean build artifacts"
	@echo "  make deps       - Download and tidy dependencies"
	@echo "  make fmt        - Format Go code"
	@echo "  make test       - Run tests"
	@echo "  make dev        - Run in development mode with hot reload"

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "✓ Build complete: $(BINARY_PATH)"

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	@$(BINARY_PATH)

# Run in development mode
dev:
	@echo "Running $(APP_NAME) in development mode..."
	@go run $(MAIN_PATH)

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f $(BINARY_PATH)
	@go clean
	@echo "✓ Cleanup complete"

# Download and tidy dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "✓ Dependencies updated"

# Format Go code
fmt:
	@echo "Formatting Go code..."
	@go fmt ./...
	@echo "✓ Code formatted"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "✓ Tests complete"

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	@go vet ./...
	@echo "✓ Linting complete"

# Install dependencies and run
install-and-run: deps build run
