# Constants
BINARY_NAME := aoc15
BIN_DIR := bin
TEST_DIR := ./tests

# All target
.PHONEY: all
all: test build

# Test target
.PHONEY: test
test:
	go test $(TEST_DIR)

# Build target
.PHONEY: build
build: test
	go build -o $(BIN_DIR)/$(BINARY_NAME)

# Clean target
.PHONEY: clean
clean:
	rm $(BIN_DIR)/*
	go clean -testcache

