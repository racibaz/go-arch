# Go-Arch
Hexagonal Architecture, Domain Driven Design (DDD), Test Driven Design (TDD), RESTful, gRPC, Swagger, Gorm(PostgreSQL), Notification(Twilio), RabbitMQ and Gin in Golang

## Overview
This project is a sample application built using Go that demonstrates the principles of Hexagonal Architecture, Domain Driven Design (DDD), and Test Driven Design (TDD). It provides both RESTful and gRPC APIs for managing posts, along with Swagger documentation for easy API exploration. The application uses Gorm as the ORM for PostgreSQL, Twilio for notifications, and RabbitMQ for messaging.

### Serve the Application
`docker compose up --build`

### Makefile Commands
#### `make run`
#### `make migrate`
#### `make seed`
#### `make mock`

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
- twilio: `github.com/twilio/twilio-go`
- rabbitmq: `github.com/rabbitmq/amqp091-go`


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



## Validation Error Example
When sending a POST request to create a post with invalid data, you might receive a validation error response like this:
```
{
    "type": "validation error",
    "message": "post validation request body does not validate",
    "cause": {
        "Content": [
            "min"
        ],
        "Description": [
            "min"
        ],
        "Title": [
            "min"
        ]
    }
}

```


## Directory Structure
```
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   ├── migrate.go
│   ├── root.go
│   ├── seed.go
│   └── serve.go
├── config
│   ├── config.go
│   └── config.yml
├── docker-compose.yml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── grpc_client
│   ├── main.go
│   └── posts
│       └── create.go
├── internal
│   ├── database
│   │   ├── migration
│   │   │   └── migration.go
│   │   └── seeder
│   │       └── seed.go
│   ├── modules
│   │   ├── post
│   │   │   ├── application
│   │   │   │   ├── ports
│   │   │   │   │   ├── mocks_test.go
│   │   │   │   │   └── post_service.go
│   │   │   │   └── usecases
│   │   │   │       ├── inputs
│   │   │   │       │   └── create_post_input.go
│   │   │   │       └── post_usecase.go
│   │   │   ├── domain
│   │   │   │   ├── factories
│   │   │   │   │   ├── post_factory.go
│   │   │   │   │   └── post_factory_test.go
│   │   │   │   ├── ports
│   │   │   │   │   ├── mocks_test.go
│   │   │   │   │   └── post_repository.go
│   │   │   │   ├── post.go
│   │   │   │   ├── post_status.go
│   │   │   │   └── post_status_test.go
│   │   │   ├── infrastructure
│   │   │   │   └── persistence
│   │   │   │       ├── gorm
│   │   │   │       │   ├── entities
│   │   │   │       │   │   └── post_entity.go
│   │   │   │       │   ├── mappers
│   │   │   │       │   │   ├── post_mapper.go
│   │   │   │       │   │   └── post_mapper_test.go
│   │   │   │       │   ├── repositories
│   │   │   │       │   │   └── post_repository.go
│   │   │   │       │   └── seeders
│   │   │   │       └── in_memory
│   │   │   │           └── post_repository.go
│   │   │   ├── post_module.go
│   │   │   ├── presentation
│   │   │   │   ├── grpc
│   │   │   │   │   ├── post_grpc_controller.go
│   │   │   │   │   └── proto
│   │   │   │   │       ├── post.pb.go
│   │   │   │   │       ├── post.proto
│   │   │   │   │       └── post_grpc.pb.go
│   │   │   │   ├── http
│   │   │   │   │   ├── post_controller.go
│   │   │   │   │   ├── reponse_dtos
│   │   │   │   │   │   ├── create_post_response_dto.go
│   │   │   │   │   │   └── get_post_response_dto.go
│   │   │   │   │   └── request_dtos
│   │   │   │   │       └── create_post_request_dto.go
│   │   │   │   └── routes
│   │   │   │       └── routes.go
│   │   │   └── test
│   │   │       └── integration
│   │   │           └── create_post_test.go
│   │   └── shared
│   │       └── presentation
│   │           └── routes
│   │               └── routes.go
│   └── providers
│       └── routes
│           └── route.go
├── main.go
├── pkg
│   ├── bootstrap
│   │   ├── migrate.go
│   │   ├── seed.go
│   │   └── serve.go
│   ├── config
│   │   ├── common.go
│   │   ├── getter.go
│   │   └── setter.go
│   ├── database
│   │   ├── common.go
│   │   ├── connect.go
│   │   └── connection.go
│   ├── error
│   │   └── errors.go
│   ├── grpc
│   │   └── grpc.go
│   ├── logger
│   │   ├── logger.go
│   │   └── zap_logger.go
│   ├── routing
│   │   ├── common.go
│   │   ├── routing.go
│   │   └── serve.go
│   ├── uuid
│   │   └── uuid.go
│   └── validator
│       └── validator.go
└── tmp
    └── build-errors.log
```


## Test
`go test -v ./...`
