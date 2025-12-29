<p align="center">



<a href="https://github.com/racibaz/go-arch/actions">
    <img src="https://img.shields.io/github/actions/workflow/status/racibaz/go-arch/ci.yaml" alt="CI" />
</a>

<a href="https://goreportcard.com/report/github.com/racibaz/go-arch">
    <img src="https://goreportcard.com/badge/github.com/racibaz/go-arch" alt="Go Report Card" />
</a>

<img src="https://img.shields.io/github/go-mod/go-version/racibaz/go-arch" alt="Go Version" />

<img src="https://img.shields.io/docker/pulls/racibaz/go-arch" alt="Docker Image" />

<img src="https://img.shields.io/github/license/racibaz/go-arch" alt="License" />

<img src="https://img.shields.io/github/languages/code-size/racibaz/go-arch" alt="Repo Size" />

<a href="https://codecov.io/github/racibaz/go-arch" > 
    <img src="https://codecov.io/github/racibaz/go-arch/graph/badge.svg?token=2YAP23FY1G" alt="codecov"/> 
</a>


</p>

# Go-Arch
Go-Arch provides a full-featured template for building modern backend services in Go, combining:
- Hexagonal (ports & adapters) architecture + Domain-Driven Design (DDD)
- Modular monolith structure
- Module code generator for rapid development
- RESTful APIs and gRPC support
- Database integration via Gorm + PostgreSQL + migrations
- Message queue support (RabbitMQ) & async notifications
- Swagger UI for API documentation + auto-generated docs / protos
- Built-in config management, logging, graceful shutdown, and Docker / docker-compose setup — ready for production or microservice environments.

Use Go-Arch as a starting point boilerplate to launch Go services rapidly: fork, configure, build — and go.

## 📚 Table of Contents

- [📖 Overview](#-overview)
- [📝 Notes](#notes)
- [🔐 GitHub Secrets](#github-secrets)
- [🐳 Run with Docker (air for live reload)](#run-with-docker)
- [🧩 Build Your First Module](#creating-a-new-module)
    - [Step 1: Generate the Module](#step-1-generate-the-module)
    - [Step 2: Register Routes](#step-2-register-routes)
    - [Step 3: Add Database Migrations](#step-3-add-database-migrations)
    - [Step 4: Implement Module Logic](#step-4-implement-module-logic)
    - [Step 5: Generate Swagger Documentation](#step-5-generate-swagger-documentation)
    - [Module Creation Flow](#module-creation-flow)
- [⚙️ Application Runtime Modes](#--application-runtime-modes)
- [🪲 Debugging Mode](#-debugging-mode)
- [🚀 GitHub Actions CI Workflow](#-github-actions-ci-workflow)
- [🔧 Makefile Commands](#makefile-commands)
- [📂 Project Structure](#project-structure)
- [⚙️ Generate gRPC Code](#-generate-grpc-code)
- [gRPC Client Example](#-grpc-client-example)
- [📑 Swagger Documentation UI](#-swagger-documentation-ui)
    - [Generate Swagger Documentation](#generate-swagger-documentation)
- [📬 RabbitMQ UI](#rabbitmq-ui)
- [📡 Prometheus UI](#prometheus-ui)
- [📊 Grafana UI](#grafana-ui)
- [🔎 Jaeger UI](#jaeger-ui)
- [🗄️ Elasticsearch](#elasticsearch)
- [🌐 Kibana UI](#kibana-ui)
- [📦 Dependencies](#-dependencies)
- [🛠 Roadmap / TODO](#-roadmap--todo)
- [🚪 API Requests](#api-requests)
- [📬 Postman Collection](#-postman-collection)
- [❌ Validation Error Example](#validation-error-example)
- [✔️ Linters](#-linters)
- [🧪 Test](#-tests)
- [🤝 Code of Conduct](#code-of-conduct)
- [👥 Contributing](#contributing)
- [📜 License](#license)


## 📖 Overview
This project demonstrates clean architectural principles in Go, including:

- **Hexagonal Architecture** for separation of concerns
- **DDD** for domain modeling
- **Modular Monolith Structure**
- **TDD** for reliable, testable code
- **RESTful APIs** 
- **gRPC APIs**
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
- **Postman Collection** for API testing
- **EFK Stack** for logging
- CI with GitHub Actions
- **Build Docker Images** and **Push to Docker Hub**
- **Module Generator** for rapid module creation
- **Hateoas** for API responses
- **CodeQL Analysis** for security
- **Codecov** coverage reports
- **Interface Assertions** for better type safety
- And more...


## 📝Notes
- There are two config files that are .env and config.yaml. You can override config.yaml values with environment variables defined in the .env file.
- You can use two ways to run database migrations:
    1. Using golang-migrate package via Makefile commands.
       The Command examples:
        - name=init_schema make db_create_migration
        - make db_migrate_up
        - make db_migrate_down
        - make db_migrate_force
        - make db_migrate_drop
        - make db_migrate_version
    2. Using the `cmd/migrate` CLI application provided in the project.
       The Command example:
        - make migrate
    3. If you want use air (live reload), you can change the `entrypoint.sh` file in the root directory.
        - change the command `make run` to `exec air` or `exec air -d` 


### 🔐GitHub Secrets
To enable automatic Docker image builds and pushes to Docker Hub via GitHub Actions, set the following secrets in your GitHub repository settings:
- `DOCKERHUB_USERNAME`: Your Docker Hub username.
- `DOCKERHUB_PASSWORD`: Your Docker Hub password or access token.
- `DOCKERHUB_REPOSITORY`: The name of your Docker Hub repository (e.g., `racibaz/go-arch`).
- `DOCKERHUB_IMAGE_TAG`: The tag for the Docker image (e.g., `latest` or a specific version).
- `CODECOV_TOKEN`: Your Codecov token for code coverage reporting.


### Run with Docker
```bash
git clone https://github.com/racibaz/go-arch.git

cd go-arch

cp .env.example .env

docker compose up --build
```
To access the running container shell:
```bash
docker exec -it Blog-app sh
```
To run database migrations using Makefile commands inside the container shell:
```bash
make db_migrate_up
make seed
```
Elasticsearch Enrollment Token & Kibana Verification Code:
```bash
docker exec -it elasticsearch bin/elasticsearch-create-enrollment-token --scope kibana
docker exec -it kibana bin/kibana-verification-code
```


## Creating a New Module

Follow the steps below to create and integrate a new module into the application.

---

### Step 1: Generate the Module

Run the following command from the project root:

```bash
make module name=YourModuleName
```

This command generates the standard module skeleton under:

`internal/modules/YourModuleName/`



### Step 2: Register Routes

Modules are not registered automatically.
You must explicitly add their routes to the main router registry.

Location:

`internal/providers/routers/router.go`

- Add HTTP routes to the RegisterRoutes function
- Add gRPC routes (if any) to the RegisterGrpcRoutes function

This keeps routing centralized and predictable.

### Step 3: Add Database Migrations

If your module introduces database changes, add your SQL migration files to:

`migrations/`

Make sure to follow the existing migration naming and versioning conventions.
and run the migration commands to apply them:

```bash
make db_migrate_up
``` 

### Step 4: Implement Module Logic

Implement your module inside the generated directory:

`internal/modules/YourModuleName/`

Follow the structure of existing modules (for example, the post module).

Typical responsibilities include:

- Controllers
- Services (business logic)
- Repositories and queries
- DTOs and validation logic

Step 5: Generate Swagger Documentation

After adding or modifying API endpoints, update the Swagger documentation:

`make generate_swagger`

See [Generate Swagger Documentation](#generate-swagger-documentation) for details.

### Module Creation Flow

    Generate module
        ↓
    Register routes
        ↓
    Add migrations
        ↓
    Implement logic
        ↓
    Generate Swagger



###⚙️ Application Runtime Modes
You can set the application environment by changing the `APP_ENV` variable in the `.env` file.


| **APP_ENV** | **Gin Mode** | **Description**                                                                |
| ----------- | ------------ | ------------------------------------------------------------------------------ |
| `local`     | `debug`      | Local development mode with full debug logs and detailed error output.         |
| `dev`       | `debug`      | Development mode; debugging features and verbose logs are enabled.             |
| `test`      | `test`       | Test mode with minimal logs, optimized for automated tests.                    |
| `prod`      | `release`    | Production mode; highest performance with simplified logs and no debug output. |


### 🪲 Debugging Mode

- Stop the app container if it's running.
- Edit the `.env` file to set `APP_ENV` to `local`.
- In the main.go file, uncomment the following line:
```go
	//cmd.Execute() // if you want  use  cobra cli
    bootstrap.Serve() //uncomment this line, if you want to local debugging 
```
- Debug it with your IDE or command line.




### 🚀 GitHub Actions CI Workflow
The project includes a GitHub Actions workflow for continuous integration (CI). The workflow is defined in the `.github/workflows/ci.yaml` file and includes the following steps:
#### Step 1 - (push and release):
- Checkout code
- Set up Go environment
- Install dependencies
- Run linters
- Run tests

#### Step 2 - If you give tag to your repository, it will trigger the release job:
- Build the application
- Build Docker image
- Push Docker image to Docker Hub
 


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
```bash
make coverage
```
```bash
make test
```
```bash
make lint
```

#### 🛠️ Generate gRPC Code
```bash
make generate_proto
```

#### 🧪 gRPC Client Example
```   

package main

import (
	"context"
	"fmt"
	"github.com/racibaz/go-arch/internal/modules/post/presentation/grpc/proto"
	"github.com/racibaz/go-arch/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	PostAggregate = "posts.Post"
)

func main() {

	config.Set("./../config", "./../.env")
	config := config.Get()

	addr := fmt.Sprintf("%s:%s", config.Grpc.Host, config.Grpc.Port)

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	if err != nil {
		log.Fatalf("Couldn't connect to grpc client: %v\n", err)
	}

	defer conn.Close()
	c := proto.NewPostServiceClient(conn)

	CreatePost(c)

}

// CreatePost creates a new post via gRPC client
func CreatePost(c proto.PostServiceClient) string {

	var payload = &proto.CreatePostInput{
		UserID:      "7b3a4d03-bcb9-47ce-b721-a156edd406f0",
		Title:       "test title title title grpc",
		Description: "test description description grpc",
		Content:     "test content content content grpc",
	}

	res, err := c.CreatePost(context.Background(), payload)

	if err != nil {
		log.Fatalf("Could not create post: %v\n", err)
	}

	log.Printf("Post has been created with ID: %s\n", res.GetId())

	return res.GetId()
}
``` 

### 📘 Swagger Documentation UI
http://127.0.0.1:3001/swagger/index.html#

#### 🧬 Generate Swagger Documentation
```bash
  make generate_swagger
```

![Swagger UI](https://github.com/user-attachments/assets/2fd68ccb-5e1e-406f-a30f-744aa51a836a)


### RabbitMQ UI
http://localhost:15672/#/

#### Username: guest
#### Password: guest

![RabbitMQ UI](https://github.com/user-attachments/assets/76b78666-c44a-487b-91e9-a6d8fc72d980)
![RabbitMQ UI](https://github.com/user-attachments/assets/6642d40c-dfa2-416d-a512-0069c05de376)


### Prometheus UI
#### http://localhost:9090/
#### http://localhost:3001/metrics

![Prometheus UI](https://github.com/user-attachments/assets/ca863e64-cb7f-4d2e-92bb-d64892ae3f37)

### Grafana UI
http://localhost:3002/login

#### Username: admin
#### Password: admin

![Grafana UI](https://github.com/user-attachments/assets/fa8d87e5-2257-4267-aba7-9823ecbc6774)

### Jaeger UI

http://localhost:16686/search

![Jaeger UI](https://github.com/user-attachments/assets/74793b87-2adc-4974-abe2-d894c95e2e39)

### Elasticsearch
http://localhost:9200/


### Kibana UI
http://127.0.0.1:5601/app/home#/


![Jaeger UI](https://github.com/user-attachments/assets/ffe97ba4-e9c1-49d2-98cd-94656bfe8cc9)




## 📦 Dependencies
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
- open telemetry: `go.opentelemetry.io/otel`
- jaeger: `go.opentelemetry.io/otel/exporters/jaeger`

## 📂Project Structure

Minimal Structure

```lua
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── cmd
├── config
├── deployments
├── docker-compose.yml
├── docs
├── entrypoint.sh
├── go.mod
├── go.sum
├── internal
│   ├── database
│   │   ├── migration
│   │   └── seeder
│   ├── modules
│   │   ├── post                          # Your module directory
│   │   │   ├── application         # The application bussiness logics live here such as services, commands, queries, so on.
│   │   │   ├── domain              # The core business logic and entities
│   │   │   ├── infrastructure      # You can implemenet 3. party services integration here such as PostgreSQL and  Mysql implementation
│   │   │   ├── module.go           # The module's main entry point for registration
│   │   │   ├── presentation        # Handles HTTP and gRPC requests and responses
│   │   │   └── testing             # Mocks and Integration tests for the module
│   │   └── shared                        # Shared module for common functionalities across modules
│   │       ├── domain
│   │       ├── infrastructure
│   │       └── presentation
│   │       └── testing
│   └── providers                                
├── main.go
├── migrations
├── pkg
```

Expanded Structure

```lua
├── cmd
├── internal
│   ├── database
│   │   ├── migration
│   │   └── seeder
│   ├── modules
│   │   ├── post
│   │   │   ├── application
│   │   │   │   ├── commands
│   │   │   │   ├── dtos
│   │   │   │   ├── handlers
│   │   │   │   ├── ports
│   │   │   │   └── queries
│   │   │   ├── domain
│   │   │   ├── infrastructure
│   │   │   │   ├── messaging
│   │   │   │   │   └── rabbitmq
│   │   │   │   ├── notification
│   │   │   │   │   └── sms
│   │   │   │   └── persistence
│   │   │   │       ├── gorm
│   │   │   │       │   ├── entities
│   │   │   │       │   ├── mappers
│   │   │   │       │   └── repositories
│   │   │   │       └── in_memory
│   │   │   ├── module.go
│   │   │   ├── presentation
│   │   │   │   ├── grpc
│   │   │   │   │   └── proto
│   │   │   │   ├── http
│   │   │   └── testing
│   │   │       └── mocks
│   │   │       └── integration
│   │   └── shared
│   │       ├── domain
│   │       ├── infrastructure
│   │       └── presentation
│   │       └── testing
│   └── providers
├── main.go
├── migrations
├── pkg
│   ├── bootstrap
│   │   ├── migrate.go
│   │   ├── seed.go
│   │   └── serve.go
│   ├── config
│   ├── database
│   ├── ddd
│   ├── env
│   ├── error
│   ├── es
│   ├── grpc
│   ├── helper
│   ├── logger
│   ├── messaging
│   ├── notification
│   ├── prometheus
│   ├── registry
│   ├── routing
│   ├── trace
│   ├── uuid
│   └── validator


```


## 🛠 Roadmap / TODO

- [x] Module Code Generator
- [x] Push Docker Image to Docker Hub via GitHub Actions
- [x] Grafana & Prometheus integration
- [x] OpenTelemetry & Jaeger integration
- [x] Tracing with Jaeger
- [x] EFK Stack for logging
- [x] Single environment (override config.yaml file with .env file)
- [x] Alternative migration usage with cmd/migrate CLI app and golang-migrate package
- [x] GitHub Actions Workflow for CI
- [ ] Add more unit tests
- [ ] Add more integration tests
- [ ] Add more end-to-end tests
- [ ] Extend documentation
- [ ] Add GraphQL API
- [ ] Add more gRPC services
- [ ] MongoDB integration
- [ ] Add correlationId support
- [ ] Add Auth Module
- [ ] Kubernetes deployment manifests
- [ ] Helm charts for easy deployment
- [ ] Support for more notification channels (e.g., Email, Push Notifications)
- [ ] Implement rate limiting
- [ ] Implement API versioning
- [ ] Implement feature toggles


## 🚪API Requests

| Endpoint                    | HTTP Method |    Description    |
|-----------------------------|:-----------:|:-----------------:|
| `/api/v1/posts`             |   `POST`    |  `Create a post`  |
| `/api/v1/posts/{{post_id}}` |    `GET`    |   `Get a post`    |
| `/api/health`               |    `GET`    | `Health endpoint` |
| `/metrics`                  |    `GET`    |  `List metrics`   |

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


## 🧹 Linters
```bash
make lint
```

## 🧪 Tests
For testing, when you are working in you local, change the APP_ENV variable to "test" in the .env file.
```bash 
APP_ENV="test" 
```
```bash
make test
```
```bash
make coverage
```


## Code of Conduct

Please note that this project is governed by a [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Contributing

Please see the [CONTRIBUTING](CONTRIBUTING.md) file.

## License

This project is licensed under the Apache 2.0 License. For further details, please see the [LICENSE](LICENSE) file.
