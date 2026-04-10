.DEFAULT_GOAL := build
.PHONY: api cli frontend

npm-install: ## install frontend
	pnpm install

frontend: ## Build front end
	pnpm run build

backend: ## Build backend
	go build -o api cmd/api/main.go
	go build -o cli cmd/rnd/main.go

build: docs types frontend backend ## build all

cli:
	go run cmd/rnd/main.go

api:
	go run cmd/api/main.go

run: docs types frontend api  ## Build all and run api
	go run cmd/api/main.go

lint: ## Run linters
	go tool golangci-lint run

fix: ## Fix linter errors automatically
	go tool golangci-lint run --fix
	go fix ./...

test: ## run all tests
	go test -cover -race ./...

unit: ## run unit tests
	go test -cover -short -race ./...

types: docs/swagger.yaml ## generate frontend types
	npx swagger-typescript-api generate -p ./docs/swagger.yaml -o ./frontend/src -n types.ts

docs: ## generate api documentation
	go tool swag init \
		--requiredByDefault \
		-g ./nimble.go \
		--dir ./ \
		-ot yaml

help: ## Makefile help
	@grep -E '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'
