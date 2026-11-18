.PHONY: all build test lint clean

all: build test

build:
	go build -o bin/qgate ./cmd/qgate
	go build -o bin/webhook ./cmd/webhook

test:
	go test -v -race ./...

lint:
	golangci-lint run

clean:
	rm -rf bin/
