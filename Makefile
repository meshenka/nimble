run:
	go run cmd/rnd/main.go

.PHONY: api
api:
	go run cmd/api/main.go

lint: ## Run linters
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

fix: ## Fix linter errors automatically
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix

test: unit ## run all tests
	go test -v -cover -race ./...
unit: ## run unit tests
	go test -v -cover -short -race ./...
