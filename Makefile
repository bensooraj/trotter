format:
	@go mod tidy
	@go fmt github.com/bensooraj/trotter/...

lint:
	golangci-lint run
