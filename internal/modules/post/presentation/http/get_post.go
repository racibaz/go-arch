package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/application/query"
	"github.com/racibaz/go-arch/internal/modules/post/presentation"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/uuid"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type GetPostHandler struct {
	Handler ports.QueryHandler[query.GetPostQuery, query.PostView]
}

func NewGetPostHandler(
	handler ports.QueryHandler[query.GetPostQuery, query.PostView],
) *GetPostHandler {
	return &GetPostHandler{
		Handler: handler,
	}
}

// Show PostGetById Show is a method to retrieve a post by its ID
//
//	@Summary	Get post by id
//	@Schemes
//	@Description	It is a method to retrieve a post by its ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string							true	"Post ID"
//	@Success		200	{object}	presentation.GetPostResponseDto	"Post retrieved successfully"
//	@Failure		404	{object}	errors.appError					"Page not found"
//	@Failure		400	{object}	errors.appError					"The Id does not parse correctly"
//	@Router			/posts/{id} [get]
func (h GetPostHandler) Show(c *gin.Context) {
	tracer := otel.Tracer(config.Get().App.Name)
	path := fmt.Sprintf(
		"%s - %s - %s",
		ModulePrefix,
		helper.StructName(h),
		helper.CurrentFuncName(),
	)
	ctx, span := tracer.Start(c, path)
	defer span.End()

	id := c.Param("id")
	postID, parseErr := uuid.Parse(id)
	if parseErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", ParseErrMessage))
			span.SetStatus(codes.Error, ParseErrMessage)
			span.RecordError(parseErr)
		}

		helper.ValidationErrorResponse(c, ParseErrMessage, parseErr)
		return
	}
	postView, serviceErr := h.Handler.Handle(ctx, query.GetPostQuery{ID: postID.ToString()})
	if serviceErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", NotFoundErrMessage))
			span.SetStatus(codes.Error, NotFoundErrMessage)
			span.RecordError(serviceErr)
		}

		helper.ErrorResponse(c, NotFoundErrMessage, serviceErr, http.StatusInternalServerError)
		return
	}

	responsePayload := helper.Response[presentation.GetPostResponseDto]{
		Data: &presentation.GetPostResponseDto{
			Post: presentation.FromPostViewToHTTP(&postView),
		},
		Links: []helper.Link{
			helper.AddHateoas("self", fmt.Sprintf("%s/%s", routePath, postID), http.MethodGet),
			helper.AddHateoas("store", fmt.Sprintf("%s/", routePath), http.MethodPost),
			helper.AddHateoas("update", fmt.Sprintf("%s/%s", routePath, postID), http.MethodPut),
			helper.AddHateoas("delete", fmt.Sprintf("%s/%s", routePath, postID), http.MethodDelete),
		},
	}

	helper.SuccessResponse(c, "Show post", responsePayload, http.StatusOK)
}
