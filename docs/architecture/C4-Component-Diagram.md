# C4 Component Diagram (Level 3)

## Overview
This diagram shows the major components within the Blog Service container, highlighting the hexagonal architecture pattern.

```plantuml
@startuml C4-Component-Diagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml

Container_Boundary(blog_service, "Blog Service") {
    Component(module_orchestrator, "Module Orchestrator", "Go", "Coordinates modules, dependency injection")
    Component(post_module, "Post Module", "Hexagonal Architecture", "Post domain logic and operations")

    Component(api_controllers, "API Controllers", "Go, Gin", "HTTP request/response handling")
    Component(grpc_handlers, "gRPC Handlers", "Go, gRPC", "gRPC service implementations")

    Component(command_handlers, "Command Handlers", "Go, CQRS", "Write operations, business logic")
    Component(query_handlers, "Query Handlers", "Go, CQRS", "Read operations, data retrieval")

    Component(domain_services, "Domain Services", "Go, Domain Logic", "Core business rules, validation")
    Component(domain_models, "Domain Models", "Go, DDD", "Aggregate roots, entities, value objects")

    Component(infrastructure, "Infrastructure Layer", "Go, Adapters", "External service integrations")
    ComponentDb(repositories, "Repositories", "Go, GORM", "Data access layer")
    ComponentQueue(message_publishers, "Message Publishers", "Go, RabbitMQ", "Event publishing")
}

Rel(api_controllers, command_handlers, "Commands", "Write operations")
Rel(api_controllers, query_handlers, "Queries", "Read operations")
Rel(grpc_handlers, command_handlers, "Commands", "Write operations")
Rel(grpc_handlers, query_handlers, "Queries", "Read operations")

Rel(command_handlers, domain_services, "Domain Logic", "Business rules")
Rel(query_handlers, domain_services, "Domain Logic", "Data retrieval")

Rel(domain_services, domain_models, "Domain Objects", "Business entities")
Rel(domain_models, repositories, "Persistence", "Data storage/retrieval")

Rel(command_handlers, message_publishers, "Events", "Async notifications")
Rel(infrastructure, message_publishers, "External APIs", "Integration calls")

Rel(module_orchestrator, post_module, "Dependency Injection", "Module coordination")
Rel(post_module, domain_models, "Domain Logic", "Business operations")

@enduml
```

## Component Descriptions

### Core Components

#### 1. Module Orchestrator
- **Purpose**: Coordinates between modules, manages dependency injection
- **Responsibilities**:
  - Module initialization
  - Dependency resolution
  - Service registration

#### 2. Post Module
- **Pattern**: Hexagonal Architecture
- **Responsibilities**:
  - Post domain operations
  - Business rule enforcement
  - CQRS implementation

### Presentation Layer

#### 3. API Controllers
- **Framework**: Gin HTTP Framework
- **Responsibilities**:
  - HTTP request handling
  - Request validation
  - Response formatting
  - Middleware integration

#### 4. gRPC Handlers
- **Framework**: Google gRPC
- **Responsibilities**:
  - Protocol buffer handling
  - Service definitions
  - Remote procedure calls

### Application Layer

#### 5. Command Handlers
- **Pattern**: CQRS Commands
- **Responsibilities**:
  - Write operations
  - Business logic execution
  - Event generation
  - Transaction management

#### 6. Query Handlers
- **Pattern**: CQRS Queries
- **Responsibilities**:
  - Read operations
  - Data retrieval
  - DTO mapping
  - Query optimization

### Domain Layer

#### 7. Domain Services
- **Pattern**: Domain Services
- **Responsibilities**:
  - Complex business logic
  - Cross-aggregate operations
  - Domain validation

#### 8. Domain Models
- **Pattern**: Domain-Driven Design
- **Components**:
  - Aggregate Roots
  - Entities
  - Value Objects
  - Domain Events

### Infrastructure Layer

#### 9. Infrastructure Layer
- **Pattern**: Adapter Pattern
- **Responsibilities**:
  - External service integration
  - Protocol translation
  - Error handling

#### 10. Repositories
- **Framework**: GORM
- **Responsibilities**:
  - Data persistence
  - Query execution
  - Entity mapping

#### 11. Message Publishers
- **Framework**: RabbitMQ
- **Responsibilities**:
  - Event publishing
  - Message queuing
  - Async communication

## Architecture Patterns Used

### Hexagonal Architecture (Ports & Adapters)
- **Domain Layer**: Core business logic (innermost hexagon)
- **Application Layer**: Use cases, CQRS commands/queries
- **Infrastructure Layer**: External concerns (adapters)
- **Presentation Layer**: API interfaces (controllers)

### CQRS (Command Query Responsibility Segregation)
- **Commands**: Write operations through Command Handlers
- **Queries**: Read operations through Query Handlers
- **Separation**: Different models for read/write operations

### Domain-Driven Design (DDD)
- **Aggregates**: Post as aggregate root
- **Entities**: Domain objects with identity
- **Value Objects**: Immutable domain concepts
- **Repositories**: Data access abstraction

### Dependency Injection
- **Module Orchestrator**: Manages component lifecycle
- **Clean Architecture**: Dependencies point inward
- **Testability**: Easy mocking and testing