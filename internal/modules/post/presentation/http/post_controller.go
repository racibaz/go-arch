package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/application/dtos"
	"github.com/racibaz/go-arch/internal/modules/post/application/ports"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain"
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

// CreatePostResponseDto
// @Description CreatePostResponseDto is a data transfer object for reporting the details of a created post
type CreatePostResponseDto struct {
	// @Description Title is the title of the post
	Title string `json:"title"`
	// @Description Description is the description of the post
	Description string `json:"description"`
	// @Description Content is the content of the post
	Content string `json:"content"`
	// @Description Status is the status of the post
	Status string `json:"status"`
}

// GetPostResponseDto
// @Description GetPostResponseDto is a data transfer object for reporting the details of a post
type GetPostResponseDto struct {
	// @Description Title is the title of the post
	Title string `json:"title"`
	// @Description Description is the description of the post
	Description string `json:"description"`
	// @Description Content is the content of the post
	Content string `json:"content"`
	// @Description Status is the status of the post
	Status string `json:"status"`
}

// CreatePostRequestDto
// @Description CreatePostRequestDto is a data transfer object for creating a post
type CreatePostRequestDto struct {
	// @Description UserId is the ID of the user creating the post
	UserId string `json:"user_id" validate:"required,uuid"`
	// @Description Title is the title of the post
	Title string `json:"title" validate:"required,min=10"`
	// @Description Description is the description of the post
	Description string `json:"description" validate:"required,min=10"`
	// @Description Content is the content of the post
	Content string `json:"content" validate:"required,min=10"`
}

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
func (postController PostController) Store(c *gin.Context) {

	tracer := otel.Tracer("go-arch")
	ctx, span := tracer.Start(c.Request.Context(), "PostModule - Restful - PostController - Store")
	defer span.End()

	createPostRequestDto, err := helper.Decode[CreatePostRequestDto](c)

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

	err = postController.Service.CreatePost(ctx, dto.CreatePostInput{
		ID:          newID,
		UserID:      createPostRequestDto.UserId,
		Title:       createPostRequestDto.Title,
		Description: createPostRequestDto.Description,
		Content:     createPostRequestDto.Content,
		Status:      postValueObject.PostStatusDraft,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		span.SetAttributes(attribute.String("error", "Post create failed"))
		span.SetStatus(codes.Error, "Post create failed")
		helper.ErrorResponse(c, "post create failed", err, http.StatusInternalServerError)
		return
	}

	responsePayload := helper.HateoasResponse[CreatePostResponseDto]{
		Data: CreatePostResponseDto{
			Title:       createPostRequestDto.Title,
			Description: createPostRequestDto.Description,
			Content:     createPostRequestDto.Content,
			Status:      postValueObject.PostStatusDraft.String(),
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
func (postController PostController) Show(c *gin.Context) {

	tracer := otel.Tracer("go-arch")
	ctx, span := tracer.Start(c, "PostModule - Restful - PostController - Show")
	defer span.End()

	postID := c.Param("id")

	result, err := postController.Service.GetById(ctx, postID)

	if err != nil {
		helper.ErrorResponse(c, "Invalid request body", err, http.StatusBadRequest)
		return
	}

	responsePayload := helper.HateoasResponse[GetPostResponseDto]{
		Data: GetPostResponseDto{
			Title:       result.Title,
			Description: result.Description,
			Content:     result.Content,
			Status:      result.Status.String(),
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
