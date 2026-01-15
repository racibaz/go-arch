# C4 Container Diagram (Level 2)

## Overview
This diagram shows the high-level technology choices and how containers communicate within the system.

```plantuml
@startuml C4-Container-Diagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

Person(user, "User", "End user")
Person(admin, "Administrator", "System admin")

System_Boundary(blog_platform, "Blog Platform") {
    Container(api_gateway, "API Gateway", "Go, Gin", "Routes requests, handles authentication, rate limiting")
    Container(blog_service, "Blog Service", "Go, Hexagonal Architecture", "Core business logic, domain models, CQRS")
    ContainerDb(postgres_db, "PostgreSQL", "PostgreSQL 15", "Primary data store for blog posts, users, metadata")
    ContainerQueue(rabbitmq, "RabbitMQ", "RabbitMQ 3.12", "Message queue for async operations")
}

System_Ext(auth_service, "Auth Service", "External OAuth")
System_Ext(email_service, "Email Service", "SMTP/SES")
System_Ext(cdn, "CDN", "Cloudflare/AWS CloudFront")
System_Ext(monitoring, "Monitoring", "Prometheus/Grafana")

Rel(user, api_gateway, "HTTP/REST", "Creates/reads blog posts")
Rel(admin, api_gateway, "HTTP/REST", "Manages content and users")

Rel(api_gateway, blog_service, "gRPC/HTTP", "Business operations")
Rel(blog_service, postgres_db, "SQL", "CRUD operations")
Rel(blog_service, rabbitmq, "AMQP", "Async messaging")

Rel(blog_service, auth_service, "REST/OAuth", "User authentication")
Rel(blog_service, email_service, "SMTP/API", "Email notifications")
Rel(blog_service, cdn, "API", "Static asset management")
Rel(blog_service, monitoring, "Metrics/Logs", "Observability")

@enduml
```

## Container Descriptions

### 1. API Gateway
- **Technology**: Go, Gin Framework
- **Responsibilities**:
  - HTTP request routing
  - Authentication middleware
  - Rate limiting
  - Request validation
  - Response formatting

### 2. Blog Service
- **Technology**: Go, Hexagonal Architecture
- **Responsibilities**:
  - Domain business logic
  - CQRS command/query handling
  - Domain model validation
  - Event sourcing
  - Business rules enforcement

### 3. PostgreSQL Database
- **Technology**: PostgreSQL 15
- **Responsibilities**:
  - Primary data persistence
  - Blog posts storage
  - User data storage
  - Metadata and configuration

### 4. RabbitMQ Message Queue
- **Technology**: RabbitMQ 3.12
- **Responsibilities**:
  - Async command processing
  - Event publishing
  - Background job processing
  - Email notification queuing


## Communication Patterns

- **Synchronous**: API Gateway ↔ Blog Service (gRPC/HTTP)
- **Database**: Blog Service ↔ PostgreSQL (SQL)
- **Async**: Blog Service ↔ RabbitMQ (AMQP)
- **External**: Blog Service ↔ External Services (REST/API)