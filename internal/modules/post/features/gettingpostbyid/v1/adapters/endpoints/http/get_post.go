package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/adapters/endpoints/dtos"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/adapters/endpoints/transformers"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/application/query"
	"github.com/racibaz/go-arch/internal/modules/shared/application/ports"
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

const (
	routePath = "/api/v1/posts"

	ValidationErrMessage = "post validation request body does not validate"
	InValidErrMessage    = "Invalid request body"
	NotFoundErrMessage   = "Not found a record"
	ParseErrMessage      = "The Id does not parse correctly"
	ModulePrefix         = "PostModule - Restful"
)

type GetPostHandler struct {
	Handler ports.QueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse]
}

func NewGetPostHandler(
	handler ports.QueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse],
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
//	@Param			id	path		string					true	"Post ID"
//	@Success		200	{object}	dtos.GetPostResponseDto	"Post retrieved successfully"
//	@Failure		404	{object}	errors.appError			"Page not found"
//	@Failure		400	{object}	errors.appError			"The Id does not parse correctly"
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
	postID, parseErr := uuid.ParseToString(id)
	if parseErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", ParseErrMessage))
			span.SetStatus(codes.Error, ParseErrMessage)
			span.RecordError(parseErr)
		}

		helper.ValidationErrorResponse(c, ParseErrMessage, parseErr)
		return
	}
	postView, handlerErr := h.Handler.Handle(ctx, query.GetPostByIdQuery{ID: postID})
	if handlerErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", NotFoundErrMessage))
			span.SetStatus(codes.Error, NotFoundErrMessage)
			span.RecordError(handlerErr)
		}

		helper.ErrorResponse(c, NotFoundErrMessage, handlerErr, http.StatusInternalServerError)
		return
	}

	responsePayload := helper.Response[dtos.GetPostResponseDto]{
		Data: &dtos.GetPostResponseDto{
			Post: transformers.FromPostViewToHTTP(&postView),
		},
		Links: []helper.Link{
			helper.AddHateoas(
				"self",
				fmt.Sprintf("%s/%s", routePath, postID),
				http.MethodGet,
				"",
			),
			helper.AddHateoas(
				"store",
				fmt.Sprintf("%s/", routePath),
				http.MethodPost,
				"api/v1/schemas/posts/create",
			),
			helper.AddHateoas(
				"update",
				fmt.Sprintf("%s/%s", routePath, postID),
				http.MethodPut,
				"api/v1/schemas/posts/update",
			),
			helper.AddHateoas(
				"delete",
				fmt.Sprintf("%s/%s", routePath, postID),
				http.MethodDelete,
				"",
			),
		},
	}

	helper.SuccessResponse(c, "Show post", responsePayload, http.StatusOK)
}
