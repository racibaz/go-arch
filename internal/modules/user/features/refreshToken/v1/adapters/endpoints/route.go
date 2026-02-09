package endpoints

import (
	"github.com/gin-gonic/gin"
	ports2 "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/features/_shared/middlewares"
	"github.com/racibaz/go-arch/internal/modules/user/features/refreshToken/v1/adapters/endpoints/http"
	"github.com/racibaz/go-arch/internal/modules/user/features/refreshToken/v1/application/queries"
)

func MapHttpRoute(
	router *gin.Engine,
	queryHandler ports2.QueryHandler[queries.RefreshTokenQueryV1, *queries.RefreshTokenQueryResponseV1],
) {
	refreshTokenHandler := http.NewRefreshTokenHandler(queryHandler)

	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/auth")
		{
			eg.Use(middlewares.Authenticate()).
				POST("/refresh-token", refreshTokenHandler.RefreshToken)
		}
	}
}
