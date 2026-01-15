# C4 Architecture Documentation

This directory contains C4 model documentation for the blog platform, providing multiple levels of architectural understanding.

## C4 Model Overview

The C4 model is a hierarchical approach to documenting software architecture:

1. **Level 1: Context Diagram** - System in its environment
2. **Level 2: Container Diagram** - High-level technology choices
3. **Level 3: Component Diagram** - Major components within containers
4. **Level 4: Code Diagram** - Actual code structure and relationships

## Architecture Principles

This project follows **Hexagonal Architecture** (Ports & Adapters) with **CQRS** (Command Query Responsibility Segregation):

### Hexagonal Architecture
```
┌─────────────────────────────────────┐
│          Presentation Layer         │  ← HTTP, gRPC, CLI
│          (Driving Adapters)         │
├─────────────────────────────────────┤
│          Application Layer          │  ← Commands, Queries, DTOs
├─────────────────────────────────────┤
│            Domain Layer             │  ← Business Logic, Entities
├─────────────────────────────────────┤
│        Infrastructure Layer         │  ← Database, Messaging, External
│        (Driven Adapters)            │     APIs (Driven Adapters)
└─────────────────────────────────────┘
```

### CQRS Pattern
- **Commands**: Write operations (create, update, delete)
- **Queries**: Read operations (get, list, search)
- **Separation**: Different models and handlers for reads vs writes

## Documentation Structure

```
docs/architecture/
├── README.md                    # This overview
├── C4-Context-Diagram.md       # Level 1: System context
├── C4-Container-Diagram.md     # Level 2: Technology choices
├── C4-Component-Diagram.md     # Level 3: Component relationships
└── C4-Code-Diagram.md          # Level 4: Code structure
```

## Key Architectural Decisions

### 1. Technology Stack
- **Language**: Go 1.21+
- **Web Framework**: Gin (HTTP), gRPC (RPC)
- **Database**: PostgreSQL with GORM
- **Message Queue**: RabbitMQ
- **Architecture**: Hexagonal + CQRS + DDD

### 2. Module Structure
Each business module follows the same pattern:
```
module/
├── domain/           # Business logic, entities, interfaces
├── application/      # Use cases, commands, queries, DTOs
├── infrastructure/   # External integrations, repositories
├── presentation/     # HTTP/gRPC controllers, routes
└── testing/         # Mocks and test utilities
```

### 3. Dependency Direction
Dependencies always point **inward** (towards domain):
```
Presentation → Application → Domain ← Infrastructure
     ↓           ↓           ↑           ↑
   Adapters    Use Cases   Entities   External
```

## How to Read C4 Diagrams

### Level 1: Context Diagram
- **Purpose**: Understand system boundaries and external interactions
- **Audience**: Stakeholders, product managers, external teams
- **Focus**: What the system does, who uses it, external dependencies

### Level 2: Container Diagram
- **Purpose**: Understand technology choices and deployment units
- **Audience**: Architects, developers, DevOps teams
- **Focus**: How containers communicate, technology stack

### Level 3: Component Diagram
- **Purpose**: Understand major components and their relationships
- **Audience**: Developers, technical leads
- **Focus**: Component responsibilities, interfaces, data flow

### Level 4: Code Diagram
- **Purpose**: Understand actual code structure and implementation
- **Audience**: Developers, code reviewers
- **Focus**: Classes, interfaces, actual implementation details

## Generating Diagrams

The diagrams use PlantUML syntax and can be rendered using:

1. **VS Code Extension**: PlantUML extension
2. **Online Editor**: plantuml.com
3. **CLI Tool**: `plantuml` command
4. **IntelliJ Plugin**: PlantUML Integration

## Architecture Validation

### Checklist for New Features
- [ ] **Context**: Does this change external interfaces?
- [ ] **Container**: Does this require new technology?
- [ ] **Component**: Does this fit hexagonal architecture?
- [ ] **Code**: Does this follow established patterns?

### Testing Strategy
- **Unit Tests**: Domain logic, individual components
- **Integration Tests**: Component interactions
- **Contract Tests**: Interface compliance
- **E2E Tests**: Full user journeys

## Common Patterns

### 1. Adding a New Module
1. Create module structure following the pattern
2. Define domain entities and interfaces
3. Implement application services (CQRS)
4. Add infrastructure adapters
5. Create presentation controllers
6. Add comprehensive tests

### 2. Adding External Integration
1. Define domain interface (port)
2. Create infrastructure adapter (driven)
3. Inject through dependency injection
4. Add configuration and error handling
5. Write integration tests

### 3. Adding Business Logic
1. Identify if it's a command or query
2. Create appropriate handler in application layer
3. Implement domain logic if needed
4. Add validation and error handling
5. Update presentation layer if needed

## Quality Attributes

### Maintainability
- **Hexagonal Architecture**: Clear separation of concerns
- **CQRS**: Independent read/write models
- **Dependency Injection**: Testable, configurable components

### Testability
- **Interface Segregation**: Easy mocking of dependencies
- **Dependency Injection**: Isolated unit testing
- **Test Doubles**: Comprehensive mock implementations

### Scalability
- **CQRS**: Independent scaling of reads/writes
- **Event Sourcing**: Audit trails and replay capabilities
- **Message Queues**: Async processing and decoupling

### Reliability
- **Error Handling**: Comprehensive error propagation
- **Validation**: Input validation at all layers
- **Monitoring**: Metrics and logging throughout

## Tooling and Automation

### Code Generation
- **Mockery**: Automatic mock generation for interfaces
- **Protocol Buffers**: gRPC service definitions
- **OpenAPI**: REST API documentation

### Quality Assurance
- **Testing**: Comprehensive test suite with coverage reporting
- **Linting**: Code quality checks with golangci-lint
- **CI/CD**: Automated testing and deployment pipelines

### Documentation
- **C4 Diagrams**: Architecture documentation
- **API Docs**: OpenAPI/Swagger specifications
- **Code Comments**: Inline documentation and examples

This C4 documentation provides a complete view of the system's architecture, from high-level context down to implementation details, ensuring all stakeholders understand the design and can contribute effectively.