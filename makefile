lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2 run ./...

test: lint
	@go test ./..