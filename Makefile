.DEFAULT_GOAL := build
.PHONY: api cli docs

npm-install:
	pnpm install

frontend: 
	pnpm run build

backend:
	go build -o api cmd/api/main.go
	go build -o cli cmd/rnd/main.go

build: front backend

cli:
	go run cmd/rnd/main.go

api:
	go run cmd/api/main.go

lint: ## Run linters
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

fix: ## Fix linter errors automatically
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix

test: ## run all tests
	go test -cover -race ./...

unit: ## run unit tests
	go test -cover -short -race ./...

docs: ## generate api documentation
	go run github.com/swaggo/swag/cmd/swag@latest init \
		-g ./nimble.go \
		--dir ./ \
		-ot yaml
