# Go-Arch
Hexagonal Architecture, Domain Driven Design (DDD), Test Driven Design (TDD), RESTful, gRPC, Swagger, Gorm(PostgreSQL), Notification(Twilio), RabbitMQ, Prometheus, Grafana, Jeager and Gin in Golang

## 📖 Overview
This project demonstrates clean architectural principles in Go, including:

- **Hexagonal Architecture** for separation of concerns
- **DDD** for domain modeling
- **TDD** for reliable, testable code
- **RESTful** and **gRPC APIs**
- **Swagger UI** for API exploration
- **PostgreSQL with Gorm ORM**
- **RabbitMQ** for messaging
- **Prometheus** for metrics
- **Grafana** for visualization
- **OpenTelemetry** for tracing
- **Jaeger** for distributed tracing
- **Twilio** for notifications
- Graceful Shutdown
- Configuration Management
- **Logging** with Zap
- Docker and Docker Compose
- **Live Reload** with Air
- **Database Migrations** with Golang-Migrate
- **Mocking** with Mockery
- Comprehensive Documentation
- **Makefile** for common tasks

### Run with Docker (air for live reload)
```bash
docker compose up --build

docker exec -it Blog-app sh

make db_migrate_up
```

### Makefile Commands
```bash
make run
```
```bash
make migrate
```
```bash
name=init_schema make db_create_migration
```
```bash
make db_migrate_up
```
```bash
make db_migrate_down
```
```bash
make db_migrate_force
```
```bash
make db_migrate_drop
```
```bash
make db_migrate_version
```
```bash
make seed
```
```bash
make mock
```


#### Generate gRPC Code
```bash
make generate_proto
```


### Swagger Documentation UI
`http://127.0.0.1:3001/swagger/index.html#`

#### Generate Swagger Documentation
```bash
  make generate_swagger
```

![Swagger UI](docs/images/swagger_ui.png)



### RabbitMQ UI
`http://localhost:15672/#/`

#### Username: guest
#### Password: guest

![RabbitMQ UI](docs/images/rabbitmq1.png)
![RabbitMQ UI](docs/images/rabbitmq2.png)



### Prometheus UI
#### `http://localhost:9090/`
#### `http://localhost:3001/metrics`

![Prometheus UI](docs/images/prometheus.png)

### Grafana UI
`http://localhost:3002/login`

#### Username: admin
#### Password: admin

![Grafana UI](docs/images/grafana.png)




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
- migrations: `github.com/golang-migrate/migrate/v4`
- prometheus: `github.com/prometheus/client_golang`
- opentelemetry: `go.opentelemetry.io/otel`
- jaeger: `go.opentelemetry.io/otel/exporters/jaeger`


## 🛠 Roadmap / TODO

- [ ] Add more unit tests
- [ ] Add more integration tests
- [ ] Add more end-to-end tests
- [ ] Extend documentation
- [ ] Add GraphQL API
- [ ] Add more gRPC services
- [ ] MongoDB integration
- [x] Grafana & Prometheus integration
- [x] OpenTelemetry & Jaeger integration
- [ ] Add custom metrics


## 📬 Postman Collection
[Download](docs/postman/baz-arch.postman_collection.json)

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
