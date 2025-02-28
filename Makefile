run:
	go run cmd/rnd/main.go

lint: ## Run linters
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

fix: ## Fix linter errors automatically
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix

