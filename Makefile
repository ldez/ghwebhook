.PHONY: all

default: test-unit build

dependencies:
	dep ensure

build:
	go build

test-unit:
	go test -v $(go list ./... | grep -v '/vendor/')
