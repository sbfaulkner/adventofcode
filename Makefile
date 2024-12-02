# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=adventofcode

.PHONY: all test clean run fmt vet lint tidy

all: test

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

tidy:
	$(GOMOD) tidy
