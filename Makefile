.DEFAULT_GOAL := build
.PHONY: api cli docs frontend

npm-install:
	pnpm install

frontend: 
	pnpm run build

backend:
	go build -o api cmd/api/main.go
	go build -o cli cmd/rnd/main.go

build: frontend backend

cli:
	go run cmd/rnd/main.go

api:
	go run cmd/api/main.go

lint: ## Run linters
	go tool golangci-lint run

fix: ## Fix linter errors automatically
	go tool golangci-lint run --fix

test: ## run all tests
	go test -cover -race ./...

unit: ## run unit tests
	go test -cover -short -race ./...

types:
	npx swagger-typescript-api generate -p ./docs/swagger.yaml -o ./frontend/src -n types.ts

docs: ## generate api documentation
	go tool swag init \
		--requiredByDefault \
		-g ./nimble.go \
		--dir ./ \
		-ot yaml
