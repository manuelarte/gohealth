lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0 run ./...

test: lint
	@go test ./... -coverprofile=coverage.txt

report: test
	@go tool cover -html=coverage.txt -o coverage.html