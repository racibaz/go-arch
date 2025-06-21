# Go-Arch
  Hexagonal Architecture, DDD, TDD, (RESTful, gRPC), Gorm and Gin in Golang

## Overview
You can start this project with this command.

### Migrate the Application
`go run main.go migrate` 

### Seed the Application
`go run main.go seed`

### Serve the Application
`go run main.go serve` \
or \
`make run`

 

## Dependencies
- uuid: `github.com/google/uuid`
- cli: `github.com/spf13/cobra`
- config: `github.com/spf13/viper`
- framework: `github.com/gin-gonic/gin`
- protobuf: `github.com/golang/protobuf`
- grpc: `google.golang.org/grpc`
- grpc-gen: `google.golang.org/genproto/googleapis/rpc`
- orm: `gorm.io/gorm`

## Test
`go test`
