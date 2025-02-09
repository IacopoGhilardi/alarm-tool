.PHONY: build test run migrate docker-up docker-down migrate-create migrate-up migrate-down

# Variabili
APP_NAME=myapp
DOCKER_COMPOSE=docker-compose

build:
	go build -o $(APP_NAME) ./cmd/$(APP_NAME)

test:
	go test -v ./...

run:
	go run ./cmd/$(APP_NAME)

migrate:
	go run ./cmd/migrate

docker-up:
	$(DOCKER_COMPOSE) up -d

docker-down:
	$(DOCKER_COMPOSE) down

lint:
	golangci-lint run

generate-swagger:
	swag init -g cmd/$(APP_NAME)/main.go

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Error: name parameter is required. Usage: make migrate-create name=migration_name"; \
		exit 1; \
	fi
	@next_version=$$(ls internals/db/migrations/*.sql 2>/dev/null | wc -l); \
	next_version=$$((next_version + 1)); \
	version_padded=$$(printf "%06d" $$next_version); \
	migrate create -ext sql -dir internals/db/migrations -seq $$version_padded_$(name)

migrate-up:
	migrate -path internals/db/migrations -database "${DATABASE_URL}" up

migrate-down:
	migrate -path internals/db/migrations -database "${DATABASE_URL}" down

migrate-force:
	migrate -path internals/db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" force $(version)
