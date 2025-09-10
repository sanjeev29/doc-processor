.PHONY: build run fmt vet

build:
	go build -o bin/doc-processor ./cmd/server

run:
	go run ./cmd/server

fmt:
	go fmt ./...

vet:
	go vet ./...