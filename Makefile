# Makefile

# Variables
BINARY_NAME=myvpn
SOURCE_DIR=.
MAIN_FILE=main.go

# Default target
all: build

# Build the binary
build:
		go build

# Run the binary
run: build
		./$(BINARY_NAME)

# Clean the build
clean:
		rm -f $(BINARY_NAME)

# Phony targets
.PHONY: all build run clean