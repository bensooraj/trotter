.PHONY: build

format:
	@go mod tidy
	@go fmt github.com/bensooraj/trotter/...

build:
	@go build -o build/

make run:
	@./build/trotter

lint:
	golangci-lint run
