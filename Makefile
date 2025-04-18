# Binary name
BINARY_NAME=marketmosaic-product-service
BUILD_DIR=build

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags "-w -s"

# Main entry point
MAIN_PATH=cmd/marketmosaic-product-service/main.go

.PHONY: all build clean test deps tidy

all: clean deps build test

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete!"

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	$(GOCLEAN)
	@echo "Clean complete!"

test:
	@echo "Running tests..."
	$(GOTEST) -v ./...
	@echo "Tests complete!"

deps:
	@echo "Downloading dependencies..."
	$(GOGET) ./...
	@echo "Dependencies downloaded!"

tidy:
	@echo "Tidying up modules..."
	$(GOMOD) tidy
	@echo "Tidy complete!"

# Build for multiple platforms
build-all: clean
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	
	# Linux
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	
	# MacOS
	@echo "Building for MacOS..."
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	
	# Windows
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	
	@echo "Multi-platform build complete!"

# Run the application
run:
	@echo "Running $(BINARY_NAME)..."
	$(GOCMD) run $(MAIN_PATH)

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME) .
	@echo "Docker build complete!"

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(BINARY_NAME) 