# Makefile for Go project

# Variables
APP_NAME := tow
SRC_DIR := .
BUILD_DIR := .
BIN_DIR := bin
PKG_DIR := pkg
MAIN_FILE := main.go

# Commands
GO := go
GO_BUILD := $(GO) build
GO_RUN := $(GO) run
GO_CLEAN := $(GO) clean
GO_TEST := $(GO) test

# Targets
.PHONY: all build run clean test

all: build

build:
	@echo "Building $(APP_NAME)..."
	$(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)/$(MAIN_FILE)

run:
	@echo "Running $(APP_NAME)..."
	$(GO_RUN) $(SRC_DIR)/$(MAIN_FILE)

clean:
	@echo "Cleaning up..."
	$(GO_CLEAN)
	rm -rf $(BUILD_DIR) $(BIN_DIR) $(PKG_DIR)

test:
	@echo "Running tests..."
	$(GO_TEST) ./...
