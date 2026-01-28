package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/features/signup/v1/adapters/endpoints/http"
	"github.com/racibaz/go-arch/internal/modules/user/features/signup/v1/application/commands"
	googleGrpc "google.golang.org/grpc"
)

func MapHttpRoute(
	router *gin.Engine,
	commandHandler ports.CommandHandler[commands.RegisterUserCommandV1],
) {
	registerUserHandler := http.NewRegisterUserHandler(commandHandler)

	v1 := router.Group("/api/v1")
	{
		schemas := v1.Group("/schemas")
		{
			schemas.GET("/auth/signup", http.Register)
			schemas.GET("/auth/update", http.Update)
		}

		eg := v1.Group("/auth/signup")
		{
			eg.POST("", registerUserHandler.Store)
		}
	}
}

func MapGrpcRoute(
	grpcServer *googleGrpc.Server,
	postHandler ports.CommandHandler[commands.RegisterUserCommandV1],
) {
	//	todo implement grpc endpoint mapping when needed
}
