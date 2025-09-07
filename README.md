# Go-Arch
Hexagonal Architecture, Domain Driven Design (DDD), Test Driven Design (TDD), RESTful, gRPC, Swagger, Gorm(PostgreSQL), Notification(Twilio), RabbitMQ and Gin in Golang

## 📖 Overview
This project demonstrates clean architectural principles in Go, including:

- **Hexagonal Architecture** for separation of concerns
- **DDD** for domain modeling
- **TDD** for reliable, testable code
- **RESTful** and **gRPC APIs**
- **Swagger UI** for API exploration
- **PostgreSQL with Gorm ORM**
- **RabbitMQ** for messaging
- **Twilio** for notifications

### Run with Docker
```bash
docker compose up --build
```

### Makefile Commands
```bash
make run
```
```bash
make migrate
```
```bash
make seed
```
```bash
make mock
```

### Swagger Documentation UI
`http://127.0.0.1:8080/swagger/index.html#`

#### Generate Swagger Documentation
```bash
  swag init -o api/openapi-spec
```

#### Generate gRPC Code
```bash
make proto
```


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


## 🛠 Roadmap / TODO

- [ ] Add more unit tests
- [ ] Add more integration tests
- [ ] Add more end-to-end tests
- [ ] Extend documentation
- [ ] Add GraphQL API
- [ ] Add more gRPC services
- [ ] MongoDB integration
- [ ] Grafana & Prometheus integration
- [ ] OpenTelemetry & Jaeger integration



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


## Linters
```bash
go vet ./...
```

## Test

```bash
go test -v ./...
```
```bash
go test -v -cover ./...
```
