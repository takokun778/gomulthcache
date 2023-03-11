include .env
export

.PHONY: help
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: name
name: ## display app name
	@echo ${APP_NAME}

.PHONY: compile
compile: ## go compile
	@go build -v ./... && go clean

.PHONY: fmt
fmt: ## go format
	@go fmt ./...

.PHONY: tidy
tidy: ## go mod tidy
	@go mod tidy

.PHONY: up
up: ## docker compose up with air hot reload
	@docker compose --project-name ${APP_NAME} --file ./.docker/docker-compose.yaml up -d

.PHONY: reup
reup: ## air rebuild
	@touch ./cmd/main.go

.PHONY: down
down: ## docker compose down
	@docker compose --project-name ${APP_NAME} down

.PHONY: e2e
e2e:
	@go test ./e2e/... -tags e2e count=1
