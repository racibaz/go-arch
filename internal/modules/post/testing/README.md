# Testing Structure

This directory contains mocks and testing utilities for the post module, organized according to hexagonal architecture principles.

## Directory Structure

```
testing/
├── mocks/
│   ├── domain/
│   │   └── ports/              # Domain interface mocks (auto-generated)
│   │       ├── notification_adapter_mock.go
│   │       └── post_repository_mock.go
│   └── infrastructure/         # Infrastructure interface mocks (auto-generated)
│   │   └── ports/
│   │       ├── get_post_handler_mock.go
│   │       ├── post_handler_mock.go
└── README.md
```

## Generating Mocks

This project uses [mockery v3](https://github.com/vektra/mockery) to automatically generate mocks.

### Installation

```bash
make mock-install
```

### Generate Mocks

```bash
make mock
```

This will regenerate all mocks based on the `.mockery.yaml` configuration.

### Clean Mocks

```bash
make mock-clean
```

## Usage

### Domain Interface Mocks

Domain interface mocks should be imported from the testing package:

```go
import (
    postPorts "github.com/racibaz/go-arch/internal/modules/post/testing/mocks/domain/ports"
)

func TestMyUseCase(t *testing.T) {
    mockRepo := postPorts.NewMockPostRepository(t)
    mockAdapter := postPorts.NewMockNotificationAdapter(t)

    // Your test logic here
}
```

### Application Handler Mocks

Application handler mocks remain in their original location:

```go
import (
    appPorts "github.com/racibaz/go-arch/internal/modules/post/features/featureName/v1/application/ports/"
)

func TestMyHandler(t *testing.T) {
    mockHandler := appPorts.NewMockCommandHandler(t)

    // Your test logic here
}
```

## Guidelines

- **Domain Layer**: No mocks should exist in the domain layer itself
- **Application Layer**: Mocks for application handlers can remain in `features/featureName/v1/application/ports/`
- **Infrastructure Layer**: Mocks for external dependencies go in `testing/mocks/infrastructure/`
- **Cross-module**: Shared mocks go in `internal/modules/shared/testing/mocks/`

## Configuration

Mock generation is configured in `.mockery.yaml` in the project root. The configuration specifies:

- **Domain interfaces**: Auto-generated into `testing/mocks/domain/ports/`
- **Application interfaces**: Auto-generated into `application/ports/`

This structure keeps the domain layer clean while providing proper testing infrastructure.