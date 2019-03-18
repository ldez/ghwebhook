.PHONY: clean check test-unit build fmt

GOFILES := $(shell git ls-files '*.go' | grep -v '^vendor/')

default: clean check test-unit build fmt

dependencies:
	dep ensure

clean:
	rm -f cover.out

build:
	go build

test-unit:
	go test -v -cover ./...

check:
	golangci-lint run

fmt:
	gofmt -s -l -w $(GOFILES)
