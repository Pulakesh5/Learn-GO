# Makefile for the Go Restaurant Project

# Define the executable name
EXECUTABLE := restaurant

# Default target
all: build

# Build the project
build:
	@echo "Building the project..."
	@go build -o $(EXECUTABLE) main.go menu.go order.go utils.go bill.go
	@echo "Build complete."

# Run the project
run: build
	@echo "Running the project..."
	@./$(EXECUTABLE)

# Clean the build
clean:
	@echo "Cleaning up..."
	@rm -f $(EXECUTABLE)
	@echo "Clean complete."

# Help message
help:
	@echo "Makefile for the Go Restaurant Project"
	@echo ""
	@echo "Usage:"
	@echo "  make          Build the project"
	@echo "  make build    Build the project"
	@echo "  make run      Build and run the project"
	@echo "  make clean    Clean the build files"
	@echo "  make help     Display this help message"

# Phony targets to avoid conflicts with files named 'all', 'build', 'run', 'clean', or 'help'
.PHONY: all build run clean help
