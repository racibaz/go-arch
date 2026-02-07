# ADR-0004: Use JWT Token with HMAC

Date: 2026/02/07

## Status

Accepted

## Context

I want to use JWT tokens for authentication and authorization in my application. I need a secure way to sign the tokens, and HMAC (Hash-based Message Authentication Code) is a widely used algorithm for this purpose. It allows me to ensure the integrity and authenticity of the tokens without the need for asymmetric keys, which simplifies the implementation and management of the authentication system.

## Decision

It will implement JWT tokens signed with HMAC for authentication and authorization. This decision is based on the need for a secure, efficient, and easy-to-implement solution for token-based authentication in the application. HMAC provides a strong level of security while being straightforward to implement and manage, making it an ideal choice for our use case.


## Consequences

- Improved security for authentication and authorization processes.
- Simplified implementation and management of the authentication system compared to asymmetric key solutions.
- Reduced overhead in token generation and verification, as HMAC is computationally efficient.
- Potential risks if the secret key used for HMAC is compromised, as it would allow attackers to forge valid tokens. Therefore, it is crucial to implement proper key management practices and ensure the secret key is stored securely.
- Enhanced user experience by enabling stateless authentication, allowing users to authenticate once and access protected resources without needing to re-authenticate for each request.
- Increased flexibility in token design, as JWT allows for custom claims to be included in the token payload, enabling more granular control over user permissions and access levels.