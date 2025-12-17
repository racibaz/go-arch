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
	@echo "  make generate_proto        - Generate gRPC protobuf code"
	@echo "  make generate_swagger      - Generate Swagger documentation"
	@echo "  make db_create_migration   - Create a new database migration"
	@echo "  make db_migrate_up         - Apply all up database migrations"
	@echo "  make db_migrate_down       - Apply all down database migrations"
	@echo "  make db_migrate_drop       - Drop all database objects"
	@echo "  make db_migrate_version    - Show current migration version"
	@echo "  make db_migrate_force      - Force set migration version"


run:
	@go run main.go serve

migrate:
	@go run main.go migrate


seed:
	@go run main.go seed

mock:
	mockery

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