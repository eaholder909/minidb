.PHONY: build test vet run clean

BINARY=minidb
BUILD_DIR=bin

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY) ./cmd/minidb

test:
	go test -v ./...

vet:
	go vet ./...

run:
	go run ./cmd/minidb

clean:
	rm -rf $(BUILD_DIR)
