.DEFAULT_GOAL := build
.PHONY: api cli

##
## LOCAL
## -----
##

npm-install: ## install frontend
	pnpm install


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

##
## BUILD
## -----
##

build: docs types frontend backend ## build all

FE_SRC := $(shell find frontend/src -name "*.tsx" -o -name "*.ts") package.json pnpm-lock.yaml webpack.config.js tsconfig.json

frontend: public/bundle.js ## Build front end

public/bundle.js: $(FE_SRC)
	pnpm run build
	touch public/bundle.js

backend: ## Build backend
	go build -o api cmd/api/main.go
	go build -o cli cmd/rnd/main.go

GO_SRC := $(shell find . -name "*.go")

types: frontend/src/types.ts ## generate frontend types

frontend/src/types.ts: docs/swagger.yaml
	npx swagger-typescript-api generate -p ./docs/swagger.yaml -o ./frontend/src -n types.ts
	touch frontend/src/types.ts

docs/swagger.yaml: $(GO_SRC)
	go tool swag init \
		--requiredByDefault \
		-g ./nimble.go \
		--dir ./ \
		-ot yaml
	touch docs/swagger.yaml

docs: docs/swagger.yaml ## generate api documentation
##
## TESTS
## -----
##

test: ## run all tests
	go test -cover -race ./...

unit: ## run unit tests
	go test -cover -short -race ./...

sqlc: ## generate store code
	go tool sqlc generate

migration: ## create a new migration (usage: make migration NAME=create_users)
	go tool goose -dir migrations create $(NAME) sql

migrate: ## run migrations
	go tool goose -dir migrations sqlite3 nimble.db up

reset-db: ## remove and recreate the database
	rm -f nimble.db
	$(MAKE) migrate

##
## HELP
## ----
##

help: ## Makefile help
	@grep -E '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'
