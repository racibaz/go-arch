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
	postModule := postModule.NewPostModule()

	newPostController := postController.NewPostController(postModule.GetService())
	router.GET("/posts/:id", newPostController.Show)
	router.POST("/posts", newPostController.Store)

}

func GrpcRoutes(grpcServer *googleGrpc.Server) {

	//todo it should be singleton
	postModule := postModule.NewPostModule()

	postGrpcController.NewPostGrpcController(grpcServer, postModule.GetService())
}
