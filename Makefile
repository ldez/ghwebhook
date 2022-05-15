.PHONY: clean generate check test-unit build

default: clean generate check test-unit build

clean:
	rm -f cover.out

generate:
	go generate

build:
	go build

test-unit:
	go test -v -cover ./...

check:
	golangci-lint run
