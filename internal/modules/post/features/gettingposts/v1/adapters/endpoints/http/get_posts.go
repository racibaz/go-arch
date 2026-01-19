package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingposts/v1/application/query"
	"github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/helper"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	routePath = "/api/v1/posts"

	ValidationErrMessage = "post validation request body does not validate"
	InValidErrMessage    = "Invalid request body"
	NotFoundErrMessage   = "Not found a record"
	ParseErrMessage      = "The Id does not parse correctly"
	ModulePrefix         = "PostModule - Restful"
)

type GetPostsHandler struct {
	Handler ports.QueryHandler[helper.Pagination, query.GetPostsQueryResponse]
}

func NewGetPostsHandler(
	handler ports.QueryHandler[helper.Pagination, query.GetPostsQueryResponse],
) *GetPostsHandler {
	return &GetPostsHandler{
		Handler: handler,
	}
}

// Index
//
//	@Summary	List posts
//	@Schemes
//	@Description	It is a method to retrieve a pagination keys such as page and page_size and list posts
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			page		path		string				true	"page number"
//	@Param			page_size	path		string				true	"page size"
//	@Success		200			{object}	GetPostResponseDto	"Post listed successfully"
//	@Router			/posts [get]
func (h GetPostsHandler) Index(c *gin.Context) {
	tracer := otel.Tracer(config.Get().App.Name)
	path := fmt.Sprintf(
		"%s - %s - %s",
		ModulePrefix,
		helper.StructName(h),
		helper.CurrentFuncName(),
	)
	ctx, span := tracer.Start(c, path)
	defer span.End()

	// todo it should be refactored.
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return
	}
	pageSizeStr := c.DefaultQuery("page_size", "10")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return
	}

	records, handlerErr := h.Handler.Handle(ctx, helper.Pagination{
		Page:     page,
		PageSize: pageSize,
	},
	)
	if handlerErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", NotFoundErrMessage))
			span.SetStatus(codes.Error, NotFoundErrMessage)
			span.RecordError(handlerErr)
		}

		helper.ErrorResponse(c, NotFoundErrMessage, handlerErr, http.StatusInternalServerError)
		return
	}

	responsePayload := helper.Response[GetPostResponseDto]{
		Data: &GetPostResponseDto{
			Posts: FromPostViewToHTTP(&records),
		},
		Links: []helper.Link{
			helper.AddHateoas(
				"self",
				fmt.Sprintf("%s", routePath),
				http.MethodGet,
				"",
			),
		},
	}

	helper.SuccessResponse(c, "List posts", responsePayload, http.StatusOK)
}
