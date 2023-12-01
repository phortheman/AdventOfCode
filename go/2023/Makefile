# Constants
BINARY_NAME := aoc23
BIN_DIR := bin
TEST_DIR := ./test/...

# All target
.PHONY: all
all: build test

# Test target
.PHONY: test
test:
	go test $(TEST_DIR)

# Build target
.PHONY: build
build: test
	go build -o $(BIN_DIR)/$(BINARY_NAME) main.go

# Clean target
.PHONY: clean
clean:
	$(RM) $(BIN_DIR)/$(BINARY_NAME)

# Dependency target
.PHONY: dep
dep:
	go mod download