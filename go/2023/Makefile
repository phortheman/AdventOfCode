# Constants
BINARY_NAME := aoc23
BIN_DIR := bin
TEST_DIR := ./...

MAIN_FILES := $(wildcard days/*/main.go)
SUBDIRS := $(sort $(dir $(MAIN_FILES)))
SUBDIR_NAMES := $(notdir $(patsubst %/,%,$(SUBDIRS)))
BUILD_TARGETS := $(addprefix $(BIN_DIR)/$(BINARY_NAME)_,$(SUBDIR_NAMES))


# All target
.PHONY: all
all: test $(BUILD_TARGETS)

# Test target
.PHONY: test
test:
	go test $(TEST_DIR)

# Build targets
.PHONY: build
build: test $(BUILD_TARGETS)

$(BIN_DIR)/$(BINARY_NAME)_%: days/%/main.go
	go build -o $@ $<

# Clean target
.PHONY: clean
clean:
	$(RM) $(BIN_DIR)/*

# Dependency target
.PHONY: dep
dep:
	go mod download