.DEFAULT_GOAL := build

BIN ?= monkey
REVISION ?= $(shell git describe --always)
LDFLAGS := '-w -s -X main.revision=$(REVISION)'

.PHONY: build
build:
	go build -ldflags $(LDFLAGS) -v -o bin/$(BIN) cmd/$(BIN)/main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test: vet
	richgo test -v -cover -race ./...

.PHONY: coverage
coverage:
	richgo test -v -race -coverprofile=/tmp/profile -covermode=atomic ./...
	go tool cover -html=/tmp/profile

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: ci-test
ci-test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
