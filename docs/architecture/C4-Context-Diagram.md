# C4 Context Diagram (Level 1)

## Overview
This diagram shows the system in its environment, including users and external systems.

```plantuml
@startuml C4-Context-Diagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

Person(user, "User", "End user who interacts with the blog platform")
Person(admin, "Administrator", "System administrator who manages the platform")

System(blog_system, "Blog Platform", "A hexagonal architecture blog system built with Go")

System_Ext(auth_service, "Authentication Service", "External OAuth/Social login service")
System_Ext(email_service, "Email Service", "External email delivery service")
System_Ext(cdn, "CDN", "Content Delivery Network for static assets")
System_Ext(monitoring, "Monitoring System", "External monitoring and alerting system")

Rel(user, blog_system, "Creates, reads, updates blog posts")
Rel(admin, blog_system, "Manages users, content, system configuration")

Rel(blog_system, auth_service, "Authenticates users")
Rel(blog_system, email_service, "Sends notifications")
Rel(blog_system, cdn, "Serves static content")
Rel(blog_system, monitoring, "Reports metrics and alerts")

@enduml
```

## Key Elements

### Primary Actors
- **User**: End users who read and create blog content
- **Administrator**: System operators who manage the platform

### External Systems
- **Authentication Service**: Handles user login/authentication
- **Email Service**: Manages email notifications
- **CDN**: Distributes static assets globally
- **Monitoring System**: Collects metrics and alerts

### System Boundary
The blog platform is the core system that provides blog functionality with hexagonal architecture principles.