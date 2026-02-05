package endpoints

import (
	"github.com/gin-gonic/gin"
	ports2 "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/features/_shared/middlewares"
	"github.com/racibaz/go-arch/internal/modules/user/features/refreshToken/v1/adapters/endpoints/http"
	"github.com/racibaz/go-arch/internal/modules/user/features/refreshToken/v1/application/query"
)

func MapHttpRoute(
	router *gin.Engine,
	queryHandler ports2.QueryHandler[query.RefreshTokenQueryV1, *query.RefreshTokenQueryResponseV1],
) {
	refreshTokenHandler := http.NewRefreshTokenHandler(queryHandler)

	router.Use(middlewares.Authenticate())

	v1 := router.Group("/api/v1")
	{

		eg := v1.Group("/auth")
		{
			eg.POST("/refresh-token", refreshTokenHandler.RefreshToken)
		}
	}
}
