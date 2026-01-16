package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/adapters/endpoints/http"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/application/query"
	"github.com/racibaz/go-arch/internal/modules/shared/application/ports"
)

func MapHttpRoute(
	router *gin.Engine,
	queryHandler ports.QueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse],
) {
	getPostHandler := http.NewGetPostHandler(queryHandler)

	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/posts")
		{
			// If you need to secure this endpoint, you can add middleware here
			eg.GET("/:id", getPostHandler.Show)
		}
	}
}
