package endpoints

import (
	"github.com/gin-gonic/gin"
	ports2 "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/features/_shared/middlewares"
	"github.com/racibaz/go-arch/internal/modules/user/features/me/v1/adapters/endpoints/http"
	query "github.com/racibaz/go-arch/internal/modules/user/features/me/v1/application/query"
)

func MapHttpRoute(
	router *gin.Engine,
	queryHandler ports2.QueryHandler[query.MeQueryHandlerQuery, *query.MeQueryHandlerResponse],
) {
	meHandler := http.NewMeHttpHandler(queryHandler)

	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/users")
		{
			eg.GET("/me", meHandler.Me).
				Use(middlewares.Authenticate())
		}
	}
}
