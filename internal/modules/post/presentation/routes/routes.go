package routes

import (
	"github.com/gin-gonic/gin"
	postModule "github.com/racibaz/go-arch/internal/modules/post"
	postGrpcController "github.com/racibaz/go-arch/internal/modules/post/presentation/grpc"
	postController "github.com/racibaz/go-arch/internal/modules/post/presentation/http"
	googleGrpc "google.golang.org/grpc"
)

func Routes(router *gin.Engine) {

	//todo it should be singleton
	module := postModule.NewPostModule()
	newPostController := postController.NewPostController(module.Service())

	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/posts")
		{
			eg.GET("/:id", newPostController.Show)
			eg.POST("/", newPostController.Store)
		}
	}
}

func GrpcRoutes(grpcServer *googleGrpc.Server) {

	//todo it should be singleton
	module := postModule.NewPostModule()

	postGrpcController.NewPostGrpcController(grpcServer, module.Service())
}
