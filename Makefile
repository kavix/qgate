.PHONY: build test
build:
	go build -o bin/pqcgate ./cmd/pqcgate
	go build -o bin/webhook ./cmd/webhook
test:
	go test ./...
