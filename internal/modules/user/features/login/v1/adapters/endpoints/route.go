package endpoints

import (
	"github.com/gin-gonic/gin"
	ports2 "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/features/login/v1/adapters/endpoints/http"
	query "github.com/racibaz/go-arch/internal/modules/user/features/login/v1/application/queries"
)

func MapHttpRoute(
	router *gin.Engine,
	queryHandler ports2.QueryHandler[query.LoginQueryV1, *query.LoginQueryResponse],
) {
	loginHandler := http.NewLoginHandler(queryHandler)

	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/auth")
		{
			eg.POST("/login", loginHandler.Login)
		}
	}
}
