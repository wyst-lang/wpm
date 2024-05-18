# Project variables
BINARY_NAME = wpm
SOURCE_DIR = src
OUTPUT_DIR = bin

# Go commands
BUILD_CMD = go build -o $(OUTPUT_DIR)/$(BINARY_NAME) $(SOURCE_DIR)/*.go
RUN_CMD = go run $(SOURCE_DIR)/*.go

# Targets
.PHONY: all build run clean

# Default target
all: build

# Build the project
build:
	@echo "Building the project..."
	@mkdir -p $(OUTPUT_DIR)
	@$(BUILD_CMD)
	@echo "Build complete!"

# Run the project using `go run`
run:
	@echo "Running the project..."
	@$(RUN_CMD)

# Clean the output directory
clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR)
	@echo "Clean complete!"
