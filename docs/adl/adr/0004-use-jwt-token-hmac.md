# ADR-0003: Use a cqrs

Date: 2026/01/08

## Status

Proposed

## Context

I want to add CQRS to better separate read and write operations in the application.

## Decision

I will implement the CQRS pattern by defining separate models for commands (write operations) and queries (read operations). Each model will have its own handlers and data access logic to ensure clear separation of concerns.



## Consequences

- Improved scalability as read and write operations can be optimized independently.
- Enhanced maintainability due to the clear separation of responsibilities.
- Potential increase in complexity as the architecture becomes more layered.
- Easier to implement event sourcing in the future if needed.
- Testing becomes more focused, allowing for targeted tests for read and write operations.
- Development speed may initially slow down as the team adapts to the new pattern, but will improve over time.
- Facilitates better alignment with domain-driven design principles.
- May require additional infrastructure to handle separate read and write databases if needed.
- Improved performance for read-heavy applications by optimizing query models.
- Potential challenges in keeping read and write models in sync, requiring careful design and implementation.
