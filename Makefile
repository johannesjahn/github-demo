.PHONY: all build run test clean

all: build

build:
	go build -o demo main.go

run: build
	./demo

test:
	go test ./...

clean:
	go clean
	rm -f demo
	rm -rf controllers/*.test
	rm -rf joke/*.test
