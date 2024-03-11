PROJECT_NAME=Phoenicia-Digital-Base-API
MAIN_PACKAGE=main.go
BUILD_DIR=dist

.PHONY: all build clean

all: build

build:
	@echo "Building $(PROJECT_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -v -o $(BUILD_DIR) $(MAIN_PACKAGE)

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)