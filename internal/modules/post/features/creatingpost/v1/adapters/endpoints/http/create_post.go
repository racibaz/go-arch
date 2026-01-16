package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/application/commands"
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

type CreatePostHandler struct {
	Handler ports.CommandHandler[commands.CreatePostCommandV1]
}

func NewCreatePostHandler(
	handler ports.CommandHandler[commands.CreatePostCommandV1],
) *CreatePostHandler {
	return &CreatePostHandler{
		Handler: handler,
	}
}

//	@BasePath	/api/v1

// Store PostStore Store is a method to create a new post
//
//	@Summary	post store
//	@Schemes
//	@Description	It is a method to create a new post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			post	body		CreatePostRequestDto	true	"Create Post Request DTO"
//	@Success		201		{object}	CreatePostResponseDto	"Post created successfully"
//	@Failure		400		{object}	errors.appError				"Invalid request body"
//	@Failure		422		{object}	errors.appError				"Post validation request body does not validate"
//	@Failure		500		{object}	errors.appError				"Post create failed"
//	@Router			/posts [post]
func (h CreatePostHandler) Store(c *gin.Context) {
	tracer := otel.Tracer(config.Get().App.Name) // go-arch
	path := fmt.Sprintf(
		"%s - %s - %s",
		ModulePrefix,
		helper.StructName(h),
		helper.CurrentFuncName(),
	)
	// todo context should be passed from gin context
	ctx, span := tracer.Start(c.Request.Context(), path)
	defer span.End()

	// Decode the request body into CreatePostRequestDto
	createPostRequestDto, decodeErr := helper.Decode[CreatePostRequestDto](c)
	if decodeErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", InValidErrMessage))
			span.SetStatus(codes.Error, InValidErrMessage)
			span.RecordError(decodeErr)
		}

		helper.ErrorResponse(c, InValidErrMessage, decodeErr, http.StatusBadRequest)
		return
	}

	// Validate the request body
	if validationErr := helper.Get().Struct(createPostRequestDto); validationErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", ValidationErrMessage))
			span.SetStatus(codes.Error, ValidationErrMessage)
			span.RecordError(validationErr)
		}

		// If validation fails, extract the validation errors and return a validation error response
		helper.ValidationErrorResponse(c, ValidationErrMessage, validationErr)
	}

	newID := uuid.NewID()

	handlerErr := h.Handler.Handle(ctx, commands.CreatePostCommandV1{
		ID:          newID,
		UserID:      createPostRequestDto.UserId,
		Title:       createPostRequestDto.Title,
		Description: createPostRequestDto.Description,
		Content:     createPostRequestDto.Content,
		Status:      domain.PostStatusDraft,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if handlerErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", "post create failed"))
			span.SetStatus(codes.Error, "post create failed")
			span.RecordError(handlerErr)
		}

		helper.ErrorResponse(c, "post create failed", handlerErr, http.StatusInternalServerError)
	}

	responsePayload := helper.Response[CreatePostResponseDto]{
		Data: &CreatePostResponseDto{
			Post: &Post{
				Title:       createPostRequestDto.Title,
				Description: createPostRequestDto.Description,
				Content:     createPostRequestDto.Content,
				Status:      domain.PostStatusDraft.String(),
			},
		},
		Links: []helper.Link{
			helper.AddHateoas("self", fmt.Sprintf("%s/%s", routePath, newID), http.MethodGet),
			helper.AddHateoas("store", fmt.Sprintf("%s/", routePath), http.MethodPost),
			helper.AddHateoas("update", fmt.Sprintf("%s/%s", routePath, newID), http.MethodPut),
			helper.AddHateoas("delete", fmt.Sprintf("%s/%s", routePath, newID), http.MethodDelete),
		},
	}

	span.SetAttributes(attribute.String("post.id", newID))

	helper.SuccessResponse(c, "Post created successfully", responsePayload, http.StatusCreated)
}
