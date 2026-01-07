# 1. Use a vertical slice architecture

Date: 2026/01/07

## Status

Accepted

## Context

I want to add features quickly while keeping the codebase maintainable.

## Decision

I will add features using a vertical slice architecture. Each feature will be developed as an independent slice that includes commands, queries, handlers, middlewares and data access logic.



## Consequences

- Features are isolated, making it easier to maintain and extend the codebase.
- Development speed increases as teams can work on different features simultaneously without conflicts.
- Testing becomes more straightforward since each slice can be tested independently.
