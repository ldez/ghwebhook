.PHONY: all

GOLIST := $(shell go list ./... | grep -v '/vendor/')

default: test-unit build

dependencies:
	dep ensure

build:
	go build

test-unit:
	go test -v  $(GOLIST)
