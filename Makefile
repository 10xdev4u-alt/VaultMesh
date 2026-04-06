.PHONY: all build test clean run

BINARY_NAME=vaultmesh
DAEMON_NAME=vaultmeshd

all: build

build:
	@echo "Building..."
	go build -o bin/$(BINARY_NAME) ./cmd/vaultmesh
	go build -o bin/$(DAEMON_NAME) ./cmd/vaultmeshd

run: build
	./bin/$(BINARY_NAME)

run-daemon: build
	./bin/$(DAEMON_NAME)

test:
	go test -v ./...

clean:
	@echo "Cleaning..."
	rm -rf bin/
	go clean

deps:
	go mod tidy
	go mod download
