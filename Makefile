include .env
export $(shell sed 's/=.*//' .env)

DB_URL=postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable&search_path=public

default: help

help:
	@echo "Makefile commands:"
	@echo "  make run                   - Run the application"
	@echo "  make migrate               - Run database migrations"
	@echo "  make seed                  - Seed the database with initial data"
	@echo "  make mock                  - Generate mocks using mockery"
	@echo "  make mock-install          - Install mockery tool"
	@echo "  make mock-clean            - Remove all generated mocks"
	@echo "  make generate_proto        - Generate gRPC protobuf code"
	@echo "  make generate_swagger      - Generate Swagger documentation"
	@echo "  make db_create_migration   - Create a new database migration"
	@echo "  make db_migrate_up         - Apply all up database migrations"
	@echo "  make db_migrate_down       - Apply all down database migrations"
	@echo "  make db_migrate_drop       - Drop all database objects"
	@echo "  make db_migrate_version    - Show current migration version"
	@echo "  make db_migrate_force      - Force set migration version"
	@echo "  make test                  - Run unit tests"
	@echo "  make coverage              - Run tests with coverage report"
	@echo "  make lint                  - Run code linting"
	@echo "  make module name=<name>    - Create a new module with the specified name"


run:
	@go run main.go serve

migrate:
	@go run main.go migrate


seed:
	@go run main.go seed

mock:
	@echo "Generating mocks using mockery..."
	@go run github.com/vektra/mockery/v3 --config .mockery.yaml

mock-install:
	@echo "Installing mockery..."
	@go install github.com/vektra/mockery/v3@latest

mock-clean:
	@echo "Cleaning generated mocks..."
	@find internal/modules/post/testing/mocks -name "*_mock.go" -delete

test:
	go test -v ./...

coverage:
	go test -v -cover ./...
	go test -v -coverprofile=coverage.txt ./...
	go tool cover -func=coverage.txt
	go tool cover -html=coverage.txt -o coverage.html

lint:
	go vet ./...

#todo make it generic
generate_proto:
	protoc -Iinternal/modules/post/presentation/grpc/proto \
		--go_out=. \
		--go_opt=module=github.com/racibaz/go-arch \
		--go-grpc_out=.  \
		--go-grpc_opt=module=github.com/racibaz/go-arch  internal/modules/post/presentation/grpc/proto/*.proto

generate_swagger:
	swag init -o api/openapi-spec

db_create_migration:
	migrate create -ext sql -dir migrations -seq $(name)

db_migrate_up:
	migrate -path migrations -database "$(DB_URL)" up

db_migrate_down:
	migrate -path migrations -database "$(DB_URL)" down

db_migrate_drop:
	migrate -path migrations -database "$(DB_URL)" drop

db_migrate_version:
	migrate -path migrations -database "$(DB_URL)" version

db_migrate_force:
	migrate -path migrations -database "$(DB_URL)" force $(version)

module:
	@if [ -z "$(name)" ]; then echo "Usage: make module name=<module_name>"; exit 1; fi
	@bash ./scripts/create_module.sh $(name)