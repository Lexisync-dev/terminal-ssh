.PHONY: all build run run-skip-boot ssh-serve ssh-connect test test-coverage lint clean

BINARY_NAME=jawad
BUILD_DIR=bin
SSH_PORT?=2222

all: build

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/jawad/main.go

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

run-skip-boot: build
	./$(BUILD_DIR)/$(BINARY_NAME) -skip-boot

ssh-serve: build
	./$(BUILD_DIR)/$(BINARY_NAME) -ssh -port $(SSH_PORT)

ssh-connect:
	ssh localhost -p $(SSH_PORT)

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

lint:
	go vet ./...

clean:
	rm -rf $(BUILD_DIR) coverage.out
