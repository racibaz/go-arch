run:
	@go run main.go serve

migrate:
	@go run main.go migrate


seed:
	@go run main.go seed

mock:
	mockery