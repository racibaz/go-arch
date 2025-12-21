package routes

import (
	"github.com/gin-gonic/gin"
	postRoutes "github.com/racibaz/go-arch/internal/modules/post/presentation/routes"
	sharedRoutes "github.com/racibaz/go-arch/internal/modules/shared/presentation/routes"
	googleGrpc "google.golang.org/grpc"
)

// RegisterRoutes registers HTTP routes for different modules
func RegisterRoutes(router *gin.Engine) {
	// Register shared routes first
	// It needs to metrics, swagger, health check, etc.
	sharedRoutes.Routes(router)

	// Register post module routes
	postRoutes.Routes(router)

	// You can add more restful routes of your modules
}

// RegisterGrpcRoutes registers gRPC routes for different modules
func RegisterGrpcRoutes(server *googleGrpc.Server) {
	postRoutes.GrpcRoutes(server)

	// You can add more gRPC routes of your modules
}
