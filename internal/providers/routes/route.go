package routes

import (
	"github.com/gin-gonic/gin"
	postRoutes "github.com/racibaz/go-arch/internal/modules/post/presentation/routes"
	sharedRoutes "github.com/racibaz/go-arch/internal/modules/shared/presentation/routes"
	googleGrpc "google.golang.org/grpc"
)

func RegisterRoutes(router *gin.Engine) {
	postRoutes.Routes(router)
	sharedRoutes.Routes(router)
	// You can add more module routes here in the future
}

func RegisterGrpcRoutes(server *googleGrpc.Server) {
	postRoutes.GrpcRoutes(server)
	// You can add more module routes here in the future
}
