PROJECT_NAME=Phoenicia-Digital-Base-API
MAIN_PACKAGE=main.go
BUILD_DIR=dist # IF THIS IS CHANGED THE DOCKER FILE HAS TO BE EDITED AS WELL!

.PHONY: all build clean

all: build

build:
	@echo "Building $(PROJECT_NAME)..."
	@echo "Creating Build Directory: $(BUILD_DIR)..."
	@mkdir -p $(BUILD_DIR)
	@echo "Build Directory: $(BUILD_DIR) Created!"
	@echo "Compiling Source Code..."
	@go build -v -o $(BUILD_DIR) $(MAIN_PACKAGE)
	@echo "Source Code Compiled!"
	@echo "Creating Config Folder..."
	@mkdir -p $(BUILD_DIR)/config
	@echo "Config Folder Created!"
	@echo "Copying Project Configuration Into Config Folder..."
	@cp config/.env $(BUILD_DIR)/config/.env
	@echo "Copied Project Configuration Into $(BUILD_DIR)/config!"
	@echo "Creating SQL Query Folder..."
	@mkdir -p $(BUILD_DIR)/sql
	@echo "SQL Query Folder Created!"
	@echo "Copying SQL Queries Into SQL Query Folder..."
	@cp sql/* $(BUILD_DIR)/sql
	@echo "Copied SQL Queries Into SQL Query Folder!"
	@echo "$(PROJECT_NAME) Built!"

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

run:
	@./$(BUILD_DIR)/main