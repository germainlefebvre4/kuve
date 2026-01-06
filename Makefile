.PHONY: build install test clean run help

# Variables
BINARY_NAME=kuve
CMD_DIR := cmd
CMD_MAIN_FILE := $(CMD_DIR)/root.go
BUILD_DIR := bin
GO=go
GOFLAGS := -v
LDFLAGS := -s -w
INSTALL_PATH=/usr/local/bin

# Go commands
GOCMD := $(GO)
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOFMT := $(GOCMD) fmt
GOMOD := $(GOCMD) mod

# Build variables
VERSION ?= $(shell git describe --tags 2>/dev/null || echo "dev")
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME := $(shell date -u '+%Y-%m-%d %H:%M:%S')
BUILD_LDFLAGS := $(LDFLAGS) -X 'main.appVersion=$(VERSION)' -X 'main.buildCommit=$(COMMIT)' -X 'main.buildTime=$(BUILD_TIME)'

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -ldflags="$(BUILD_LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_MAIN_FILE)
	@echo "Build complete: $(BINARY_NAME)"

# Install the binary to system path
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	@sudo mv $(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "Installation complete. Run 'kuve --help' to get started."

# Run tests
test:
	@echo "Running tests..."
	$(GO) test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) test -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -f $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

# Run the application
run: build
	./$(BINARY_NAME)

# Format code
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" && exit 1)
	golangci-lint run

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GO) mod download
	$(GO) mod tidy

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BINARY_NAME)-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 $(GO) build -o $(BINARY_NAME)-linux-arm64 main.go
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(BINARY_NAME)-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 $(GO) build -o $(BINARY_NAME)-darwin-arm64 main.go
	@echo "Multi-platform build complete"

# Docker build and push
docker-build:
	@echo "Building Docker image..."
	docker build -t germainlefebvre4/kuve:dev --no-cache .
	@echo "Docker image built: kuve:dev"

#
docker-build-push: docker-build
	docker push germainlefebvre4/kuve:dev
	@echo "Docker image pushed: germainlefebvre4/kuve:dev"

# Goreleaser targets
goreleser-check:
	goreleaser check

# Goreleaser release (snapshot)
goreleser-release:
	goreleaser release --snapshot --clean

# Display help
help:
	@echo "Available targets:"
	@echo "  build          - Build the application"
	@echo "  install        - Build and install to $(INSTALL_PATH)"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  clean          - Remove build artifacts"
	@echo "  run            - Build and run the application"
	@echo "  fmt            - Format code"
	@echo "  lint           - Lint code (requires golangci-lint)"
	@echo "  deps           - Download and tidy dependencies"
	@echo "  build-all      - Build for multiple platforms"
	@echo "  help           - Show this help message"
