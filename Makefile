run:
	@go run main.go serve

migrate:
	@go run main.go migrate


seed:
	@go run main.go seed

mock:
	mockery

#todo make it generic
generate-proto:
	protoc -Iinternal/modules/post/presentation/grpc/proto \
		--go_out=. \
		--go_opt=module=github.com/racibaz/go-arch \
		--go-grpc_out=.  \
		--go-grpc_opt=module=github.com/racibaz/go-arch  internal/modules/post/presentation/grpc/proto/*.proto