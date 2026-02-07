package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/adapters/endpoints/grpc"
	"github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/adapters/endpoints/http"
	"github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/application/commands"
	"github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/features/_shared/middlewares"
	googleGrpc "google.golang.org/grpc"
)

func MapHttpRoute(
	router *gin.Engine,
	commandHandler ports.CommandHandler[commands.CreatePostCommandV1],
) {
	createPostHandler := http.NewCreatePostHandler(commandHandler)

	v1 := router.Group("/api/v1")
	{
		schemas := v1.Group("/schemas")
		{
			schemas.GET("/posts/create", http.Create)
			schemas.GET("/posts/update", http.Update)
		}

		eg := v1.Group("/posts")
		{
			eg.Use(middlewares.Authenticate()).
				POST("", createPostHandler.Store)
		}
	}
}

func MapGrpcRoute(
	grpcServer *googleGrpc.Server,
	postHandler ports.CommandHandler[commands.CreatePostCommandV1],
) {
	createPostHandler := http.NewCreatePostHandler(postHandler)

	grpc.NewCreatePostHandler(grpcServer, createPostHandler.Handler)
}
