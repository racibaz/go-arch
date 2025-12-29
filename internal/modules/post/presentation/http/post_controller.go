package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/application/dtos"
	"github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/presentation"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/uuid"
	validator "github.com/racibaz/go-arch/pkg/validator"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"net/http"
	"time"
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
//	@Param			post	body		CreatePostRequestDto	true	"Create Post Request DTO"
//	@Success		201		{object}	CreatePostResponseDto	"Post created successfully"
//	@Failure		400		{object}	errors.AppError					"Invalid request body"
//	@Router			/posts [post]
func (controller PostController) Store(c *gin.Context) {

	tracer := otel.Tracer(config.Get().App.Name) //go-arch
	//"PostModule - Restful - PostController - Store"
	path := fmt.Sprintf("PostModule - Restful - %s - %s", helper.StructName(controller), helper.CurrentFuncName())

	ctx, span := tracer.Start(c.Request.Context(), path)
	defer span.End()

	// Decode the request body into CreatePostRequestDto
	createPostRequestDto, err := helper.Decode[presentation.CreatePostRequestDto](c)

	if err != nil {
		helper.ErrorResponse(c, "Invalid request body", err, http.StatusBadRequest)
		return
	}

	// Validate the request body
	if err := validator.Get().Struct(&createPostRequestDto); err != nil {
		// If validation fails, extract the validation errors
		c.JSON(
			http.StatusBadRequest,
			validator.NewValidationError(
				"post validation request body does not validate",
				validator.ShowRegularValidationErrors(err).Errors,
			),
		)
		return
	}

	newID := uuid.NewID()

	err = controller.Service.CreatePost(ctx, dto.CreatePostInput{
		ID:          newID,
		UserID:      createPostRequestDto.UserId,
		Title:       createPostRequestDto.Title,
		Description: createPostRequestDto.Description,
		Content:     createPostRequestDto.Content,
		Status:      domain.PostStatusDraft,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		span.SetAttributes(attribute.String("error", "Post create failed"))
		span.SetStatus(codes.Error, "Post create failed")

		helper.ErrorResponse(c, "post create failed", err, http.StatusInternalServerError)
		return
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
			{
				Rel:  "self",
				Href: fmt.Sprintf("/api/v1/posts/%s", newID),
				Type: "GET",
			},
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
//	@Param			id	path		string		true	"Post ID"
//	@Success		200	{object}	GetPostResponseDto	"Post retrieved successfully"
//	@Failure		404	{object}	errors.AppError		"Page not found"
//	@Router			/posts/{id} [get]
func (controller PostController) Show(c *gin.Context) {

	tracer := otel.Tracer(config.Get().App.Name)
	path := fmt.Sprintf("PostModule - Restful - %s - %s", helper.StructName(controller), helper.CurrentFuncName())
	ctx, span := tracer.Start(c, path)
	defer span.End()

	id := c.Param("id")
	postID, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(c, "Invalid request body", err, http.StatusBadRequest)
		return
	}
	result, err := controller.Service.GetById(ctx, postID.ToString())

	if err != nil {
		helper.ErrorResponse(c, "Not found a post", err, http.StatusInternalServerError)
		return
	}

	responsePayload := helper.Response[presentation.GetPostResponseDto]{
		Data: &presentation.GetPostResponseDto{
			Post: presentation.FromPostCoreToHTTP(result),
		},
		Links: []helper.Link{
			{
				Rel:  "self",
				Href: fmt.Sprintf("/api/v1/posts/%s", postID),
				Type: "GET",
			},
		},
	}

	helper.SuccessResponse(c, "Show post", responsePayload, http.StatusOK)
}
