# Gola Score Keeper - Build Configuration

# Variables
APP_NAME := gola-score-keeper
VERSION := 1.0.0
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT := $(shell git rev-parse --short HEAD)

# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOMOD := $(GOCMD) mod

# Build flags
LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -s -w"
BUILD_FLAGS := -trimpath

# Apple Silicon optimization flags
ARM64_FLAGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -s -w" -trimpath

# Directories
CMD_DIR := cmd
INTERNAL_DIR := internal
WEB_DIR := web
BUILD_DIR := build
DIST_DIR := dist

# Default target
.PHONY: all
all: clean deps build

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  build-arm64  - Build optimized for Apple Silicon"
	@echo "  test         - Run tests"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Download dependencies"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"
	@echo "  run          - Run the application"
	@echo "  setup        - Setup database"
	@echo "  docker       - Build Docker image"
	@echo "  install      - Install the application"

# Dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Build targets
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME) ./$(CMD_DIR)/server

.PHONY: build-arm64
build-arm64:
	@echo "Building $(APP_NAME) for Apple Silicon..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 $(GOBUILD) $(ARM64_FLAGS) -o $(BUILD_DIR)/$(APP_NAME)-arm64 ./$(CMD_DIR)/server

.PHONY: build-all
build-all: build build-arm64

# Test targets
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

.PHONY: test-benchmark
test-benchmark:
	@echo "Running benchmarks..."
	$(GOTEST) -bench=. -benchmem ./...

# Code quality
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

.PHONY: lint
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

.PHONY: vet
vet:
	@echo "Running go vet..."
	$(GOCMD) vet ./...

# Clean targets
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)
	rm -f coverage.out coverage.html

# Run targets
.PHONY: run
run:
	@echo "Running $(APP_NAME)..."
	$(GOCMD) run ./$(CMD_DIR)/server

.PHONY: run-setup
run-setup:
	@echo "Running database setup..."
	$(GOCMD) run ./$(CMD_DIR)/setup

.PHONY: run-camera-test
run-camera-test:
	@echo "Running camera test..."
	$(GOCMD) run ./$(CMD_DIR)/camera-test

# Development targets
.PHONY: dev
dev:
	@echo "Starting development server with hot reload..."
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "air not installed. Install with: go install github.com/cosmtrek/air@latest"; \
		$(GOCMD) run ./$(CMD_DIR)/server; \
	fi

# Database targets
.PHONY: db-migrate
db-migrate:
	@echo "Running database migrations..."
	$(GOCMD) run ./$(CMD_DIR)/setup

.PHONY: db-reset
db-reset:
	@echo "Resetting database..."
	rm -f data/score_keeper.db
	$(GOCMD) run ./$(CMD_DIR)/setup

# Docker targets
.PHONY: docker
docker:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):$(VERSION) .

.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(APP_NAME):$(VERSION)

# Installation
.PHONY: install
install: build
	@echo "Installing $(APP_NAME)..."
	sudo cp $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/

# Release targets
.PHONY: release
release: clean test build-all
	@echo "Creating release package..."
	@mkdir -p $(DIST_DIR)
	cp $(BUILD_DIR)/$(APP_NAME) $(DIST_DIR)/
	cp $(BUILD_DIR)/$(APP_NAME)-arm64 $(DIST_DIR)/
	cp README.md $(DIST_DIR)/
	cp LICENSE $(DIST_DIR)/
	@echo "Release package created in $(DIST_DIR)/"

# Development setup
.PHONY: setup-dev
setup-dev:
	@echo "Setting up development environment..."
	$(GOMOD) download
	$(GOMOD) tidy
	@echo "Installing development tools..."
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	@if ! command -v air >/dev/null 2>&1; then \
		go install github.com/cosmtrek/air@latest; \
	fi
	@echo "Development environment setup complete!"

# Check targets
.PHONY: check
check: fmt vet lint test

# CI targets
.PHONY: ci
ci: deps fmt vet lint test build

# Default target
.DEFAULT_GOAL := help
