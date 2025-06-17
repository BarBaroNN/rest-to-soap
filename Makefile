# Variables
GEN_CMD = go run ./builder/main.go   # Adjust this to your actual generator command
BUILD_CMD = go build -o app ./proxy/main.go
RUN_CMD = ./app

.PHONY: all generate build run clean

all: generate build run

# Generate all generated files (templates, types, routes, etc)
generate:
	@echo "Generating files..."
	$(GEN_CMD)

# Build the app
build:
	@echo "Building application..."
	$(BUILD_CMD)

# Run the built app
run:
	@echo "Running application..."
	$(RUN_CMD)

# Clean generated files and binary
clean:
	@echo "Cleaning..."
	rm -rf generated
	rm -f app
