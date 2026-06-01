# go-calendar Makefile

PKG_NAME := go-calendar

BIN_NAME := $(PKG_NAME)
BUILD_DIR := build
DIST_NAME := bin

.PHONY: build test clean help

build:
	@echo "Building $(BIN_NAME)..."
	go build -o dist/$(BIN_NAME) ./cmd/$(PKG_NAME)/
	@echo "Built successfully at dist/$(BIN_NAME)"

test:
	@echo "Running unit tests..."
	go test ./internal/... -v -race

clean:
	rm -rf dist/
	rm -f $(BUILD_DIR)/*.out

help:
	@echo "Available targets:"
	@echo "  build           - Compile the application to binary"
	@echo "  test            - Run all tests (unit and integration)"
	@echo "  clean           - Remove generated files"
