.PHONY: clean check test-unit build

default: clean check test-unit build

clean:
	rm -f cover.out

build:
	go build

test-unit:
	go test -v -cover ./...

check:
	golangci-lint run
