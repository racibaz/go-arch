# Go-Arch
Hexagonal Architecture, DDD, TDD, RESTful, gRPC, Swagger, Gorm(PostgreSQL) and Gin in Golang

## Overview
Go-Arch is a sample application that demonstrates the principles of Hexagonal Architecture, Domain-Driven Design (DDD), and Test-Driven Development (TDD) using Golang. It includes both RESTful and gRPC APIs, utilizing Gorm for ORM and Gin as the web framework.

### Serve the Application
`docker compose up --build`

### Swagger Documentation UI
`http://127.0.0.1:8080/swagger/index.html#`

## Dependencies
- uuid: `github.com/google/uuid`
- cli: `github.com/spf13/cobra`
- config: `github.com/spf13/viper`
- framework: `github.com/gin-gonic/gin`
- protobuf: `github.com/golang/protobuf`
- grpc: `google.golang.org/grpc`
- grpc-gen: `google.golang.org/genproto/googleapis/rpc`
- orm: `gorm.io/gorm`
- live reload: `github.com/air-verse/air`
- open api: `github.com/swaggo/swag`
- open api gin: `github.com/swaggo/gin-swagger`
- testing: `github.com/stretchr/testify`
- mocking: `github.com/vektra/mockery`
- logger: `github.com/uber-go/zap`


## TODO
- Add more unit tests
- Add more integration tests
- Add more end-to-end tests
- Add more documentation
- GraphQL API
- Add more gRPC services
- MongoDB integration
- Grafana, Prometheus integration
- Opentelemetry, Jaeger integration



## Test
`go test -v ./...`
