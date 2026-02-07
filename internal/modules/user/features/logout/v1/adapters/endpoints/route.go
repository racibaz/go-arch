package endpoints

import (
	"github.com/gin-gonic/gin"
	ports2 "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/features/_shared/middlewares"
	"github.com/racibaz/go-arch/internal/modules/user/features/logout/v1/adapters/endpoints/http"
	"github.com/racibaz/go-arch/internal/modules/user/features/logout/v1/application/commands"
)

func MapHttpRoute(
	router *gin.Engine,
	commandHandler ports2.CommandHandler[commands.LogoutCommandV1],
) {
	logoutHandler := http.NewLogoutHandler(commandHandler)

	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/auth")
		{
			eg.POST("/logout", logoutHandler.Logout).
				Use(middlewares.Authenticate())
		}
	}
}
