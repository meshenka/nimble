.DEFAULT_GOAL := build
.PHONY: api cli

build:
	go build -o api cmd/api/main.go
	go build -o cli cmd/rnd/main.go

cli:
	go run cmd/rnd/main.go

api:
	go run cmd/api/main.go

lint: ## Run linters
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

fix: ## Fix linter errors automatically
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix

test: ## run all tests
	go test -v -cover -race ./...
unit: ## run unit tests
	go test -v -cover -short -race ./...
