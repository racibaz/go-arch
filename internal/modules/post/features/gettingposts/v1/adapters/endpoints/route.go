package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingposts/v1/adapters/endpoints/http"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingposts/v1/application/query"
	"github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/pkg/helper"
)

func MapHttpRoute(
	router *gin.Engine,
	queryHandler ports.QueryHandler[helper.Pagination, query.GetPostsQueryResponse],
) {
	getPostsHandler := http.NewGetPostsHandler(queryHandler)

	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/posts")
		{
			// If you need to secure this endpoint, you can add middleware here
			eg.GET("", getPostsHandler.Index)
		}
	}
}
