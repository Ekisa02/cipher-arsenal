BINARY=server

.PHONY: build run test

build:
	go build -o $(BINARY) ./cmd/server

run: build
	./$(BINARY)

test:
	go test ./... -v
