include .env
export $(shell sed 's/=.*//' .env)

DB_URL=postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable&search_path=public

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