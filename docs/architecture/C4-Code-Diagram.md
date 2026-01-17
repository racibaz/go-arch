# C4 Code Diagram (Level 4)

## Overview
This diagram shows the actual code structure using the Post module as an example of the hexagonal architecture implementation.

```plantuml
@startuml C4-Code-Diagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml

Package "internal/modules/post" as post_module {
    Package "domain" as domain {
        Class "Post" as post_entity
        Class "PostStatus" as post_status
        Interface "PostRepository" as post_repo_interface
        Class "PostAggregate" as post_events
    }

    Package "application" as application {
        Package "commands" as commands {
            Class "CreatePostCommand" as create_command
            Class "CreatePostHandler" as create_handler
        }
        Package "queries" as queries {
            Class "GetPostQuery" as get_query
            Class "GetPostHandler" as get_handler
        }
        Package "ports" as app_ports {
            Interface "CommandHandler[C]" as command_handler
            Interface "QueryHandler[Q,R]" as query_handler
        }
        Package "dtos" as dtos {
            Class "CreatePostInput" as create_input
        }
    }

    Package "infrastructure" as infrastructure {
        Package "persistence/gorm" as persistence {
            Package "entities" as entities {
                Class "Post" as post_entity_db
            }
            Package "mappers" as mappers {
                Class "PostMapper" as post_mapper
            }
            Package "repositories" as repositories {
                Class "PostRepository" as post_repo_impl
            }
        }
        Package "messaging/rabbitmq" as messaging {
            Class "PostMessagePublisher" as message_publisher
        }
        Package "notification/sms" as notification {
            Class "TwilioSmsAdapter" as sms_adapter
        }
    }

    Package "presentation" as presentation {
        Package "http" as http {
            Class "CreatePostHandler" as http_create_handler
            Class "GetPostHandler" as http_get_handler
        }
        Package "grpc" as grpc {
            Class "PostGrpcController" as grpc_controller
        }
        Package "routes" as routes {
            Class "Routes" as routes_config
        }
    }

    Package "testing/mocks" as mocks {
        Package "application/ports" as app_mocks {
            Class "MockCommandHandler" as mock_command_handler
            Class "MockQueryHandler" as mock_query_handler
        }
        Package "domain/ports" as domain_mocks {
            Class "MockPostRepository" as mock_post_repo
            Class "MockNotificationAdapter" as mock_notification
        }
    }
}

' Domain Layer Relationships
post_entity --> post_status : uses
post_repo_interface --> post_entity : manages
post_events --> post_entity : aggregates

' Application Layer Relationships
create_command --> post_entity : creates
create_handler --> create_command : handles
create_handler --> post_repo_interface : depends
create_handler --> message_publisher : publishes

get_query --> post_entity : queries
get_handler --> get_query : handles
get_handler --> post_repo_interface : depends

command_handler --> create_handler : interface
query_handler --> get_handler : interface

' Infrastructure Layer Relationships
post_repo_impl --> post_repo_interface : implements
post_repo_impl --> post_entity_db : manages
post_mapper --> post_entity : maps from
post_mapper --> post_entity_db : maps to
message_publisher --> post_entity : publishes events
sms_adapter --> post_entity : sends notifications

' Presentation Layer Relationships
http_create_handler --> command_handler : uses
http_get_handler --> query_handler : uses
grpc_controller --> command_handler : uses
grpc_controller --> query_handler : uses
routes_config --> http_create_handler : routes
routes_config --> http_get_handler : routes
routes_config --> grpc_controller : routes

' Testing Relationships
mock_command_handler --> command_handler : mocks
mock_query_handler --> query_handler : mocks
mock_post_repo --> post_repo_interface : mocks
mock_notification --> sms_adapter : mocks

' Cross-cutting relationships
create_handler --> sms_adapter : notifies
create_handler --> message_publisher : publishes

@enduml
```

## Code Structure Analysis

### Directory Structure

```
internal/modules/post/
├── domain/                          # Domain Layer
│   ├── post.go                     # Aggregate Root
│   ├── post_status.go              # Value Object
│   ├── ports/
│   │   ├── post_repository.go      # Domain Interface
│   │   └── notification_adapter.go # Domain Interface
│   └── post_events.go              # Domain Events
├── features/                        # Application Layer
│   ├── creatingpost/
│   │   ├── v1                        # Feature Version
│   │   ├──── application               # Commands
│   │   │    ├────commmands             # Business Logic
│   │   │    ├────dtos                  # Business Logic DTOs
│   │   ├──── adapters                  # Adapters
│   │   │    ├────endpoints             # Endpoints
│   │   │    │    ├────http                  # Http handler
│   │   │    │    ├────grpc                  # gRCP handler
├── infrastructure/                  # Infrastructure Layer
│   ├── persistence/gorm/
│   │   ├── entities/post_entity.go # Database Entity
│   │   ├── mappers/post_mapper.go  # Data Mapper
│   │   └── repositories/post_repository.go # Repository Impl
│   ├── messaging/rabbitmq/
│   │   └── post_message_publisher.go # Message Publisher
│   └── notification/sms/
│       └── twilio_sms_adapter.go    # Notification Adapter
└── testing/mocks/                   # Test Doubles
    ├── application/ports/
    │   ├── command_handler_mock.go
    │   └── query_handler_mock.go
    └── domain/ports/
        ├── post_repository_mock.go
        └── notification_adapter_mock.go
```

### Key Classes and Interfaces

#### Domain Layer
- **`Post`**: Aggregate root containing business logic
- **`PostStatus`**: Value object for post states
- **`PostRepository`**: Interface for data persistence
- **`NotificationAdapter`**: Interface for notifications

#### Application Layer
- **`CreatePostCommand`**: DTO for create operations
- **`CreatePostHandler`**: Command handler implementation
- **`GetPostQuery`**: DTO for read operations
- **`GetPostHandler`**: Query handler implementation
- **`CommandHandler[T]`**: Generic interface for commands
- **`QueryHandler[Q,R]`**: Generic interface for queries

#### Infrastructure Layer
- **`PostEntity`**: GORM database entity
- **`PostMapper`**: Data mapper between domain and persistence
- **`PostRepository`**: GORM implementation of repository
- **`PostMessagePublisher`**: RabbitMQ event publisher
- **`TwilioSmsAdapter`**: SMS notification implementation

#### Presentation Layer
- **`CreatePostHandler`**: Gin HTTP handler for POST operations
- **`GetPostHandler`**: Gin HTTP handler for GET operations
- **`PostGrpcController`**: gRPC service implementation
- **`Routes`**: Route configuration and middleware setup

### Design Patterns Implemented

#### 1. Hexagonal Architecture
- **Domain**: Core business logic (innermost layer)
- **Application**: Use cases and CQRS (middle layer)
- **Infrastructure**: External concerns (outer layer)
- **Presentation**: API interfaces (outermost layer)

#### 2. CQRS (Command Query Responsibility Segregation)
- **Commands**: Write operations (`CreatePostHandler`)
- **Queries**: Read operations (`GetPostHandler`)
- **Separation**: Different handlers for different concerns

#### 3. Repository Pattern
- **Interface**: `PostRepository` (domain contract)
- **Implementation**: `PostRepository` (GORM implementation)
- **Abstraction**: Domain doesn't know about persistence details

#### 4. Dependency Injection
- **Module**: `PostModule` manages dependencies
- **Constructor**: `NewPostModule()` wires components
- **Testability**: Easy mocking and testing

#### 5. Data Mapper
- **Purpose**: Translates between domain and persistence models
- **Implementation**: `PostMapper.ToDomain()` and `ToPersistence()`
- **Separation**: Keeps domain model clean

### Communication Flow

1. **HTTP Request** → `Routes` → `HTTP Handler`
2. **Handler** → `Command/Query Handler` (CQRS)
3. **Handler** → `Domain Service` (business logic)
4. **Domain** → `Repository Interface` (data access)
5. **Repository** → `GORM Entity` (database)
6. **Events** → `Message Publisher` (async)
7. **Response** ← `Handler` ← `DTO Mapping`

This architecture ensures:
- **Testability**: Each layer can be tested in isolation
- **Maintainability**: Clear separation of concerns
- **Flexibility**: Easy to change implementations
- **Scalability**: CQRS allows independent scaling of reads/writes