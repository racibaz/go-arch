package http

import (
	"fmt"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	dto "github.com/racibaz/go-arch/internal/modules/post/application/dtos"
	"github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/presentation"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/uuid"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

const (
	routePath            = "/api/v1/posts"
	ValidationErrMessage = "post validation request body does not validate"
	InValidErrMessage    = "Invalid request body"
	NotFoundErrMessage   = "Not found a record"
	ParseErrMessage      = "The Id does not parse correctly"
	ModulePrefix         = "PostModule - Restful"
)

type PostController struct {
	Service ports.PostService
}

func NewPostController(service ports.PostService) *PostController {
	return &PostController{
		Service: service,
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
//	@Param			post	body		presentation.CreatePostRequestDto	true	"Create Post Request DTO"
//	@Success		201		{object}	presentation.CreatePostResponseDto	"Post created successfully"
//	@Failure		400		{object}	errors.AppError						"Invalid request body"
//	@Router			/posts [post]
func (controller PostController) Store(c *gin.Context) {
	tracer := otel.Tracer(config.Get().App.Name) // go-arch
	path := fmt.Sprintf(
		"%s - %s - %s",
		ModulePrefix,
		helper.StructName(controller),
		helper.CurrentFuncName(),
	)
	//todo context should be passed from gin context
	ctx, span := tracer.Start(c.Request.Context(), path)
	defer span.End()

	// Decode the request body into CreatePostRequestDto
	createPostRequestDto, decodeErr := helper.Decode[presentation.CreatePostRequestDto](c)
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

	serviceErr := controller.Service.CreatePost(ctx, dto.CreatePostInput{
		ID:          newID,
		UserID:      createPostRequestDto.UserId,
		Title:       createPostRequestDto.Title,
		Description: createPostRequestDto.Description,
		Content:     createPostRequestDto.Content,
		Status:      domain.PostStatusDraft,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if serviceErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", "post create failed"))
			span.SetStatus(codes.Error, "post create failed")
			span.RecordError(serviceErr)
		}

		helper.ErrorResponse(c, "post create failed", serviceErr, http.StatusInternalServerError)
	}

	responsePayload := helper.Response[presentation.CreatePostResponseDto]{
		Data: &presentation.CreatePostResponseDto{
			Post: &presentation.Post{
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
//	@Failure		404	{object}	errors.AppError					"Page not found"
//	@Router			/posts/{id} [get]
func (controller PostController) Show(c *gin.Context) {
	tracer := otel.Tracer(config.Get().App.Name)
	path := fmt.Sprintf(
		"%s - %s - %s",
		ModulePrefix,
		helper.StructName(controller),
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
	result, serviceErr := controller.Service.GetById(ctx, postID.ToString())
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
			Post: presentation.FromPostCoreToHTTP(result),
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
