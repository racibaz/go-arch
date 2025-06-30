package http

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/application/usecases/inputs"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain"
	requestDto "github.com/racibaz/go-arch/internal/modules/post/presentation/http/request_dtos"
	errors "github.com/racibaz/go-arch/pkg/error"
	"github.com/racibaz/go-arch/pkg/uuid"
	validator "github.com/racibaz/go-arch/pkg/validator"
	"time"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"net/http"
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
//	@Param			post	body		requestDto.CreatePostRequestDto	true	"Create Post Request DTO"
//	@Success		201		{object}	domain.Post						"Post created successfully"
//	@Failure		400		{object}	errors.AppError					"Invalid request body"
//	@Router			/posts [post]
func (postController PostController) Store(c *gin.Context) {
	var createPostRequestDto requestDto.CreatePostRequestDto

	// Bind the JSON request body to the CreatePostRequestDto struct
	if err := c.ShouldBindJSON(&createPostRequestDto); err != nil {
		c.JSON(
			http.StatusBadRequest,
			errors.NewInValidError(
				"Invalid request body",
				err.Error(),
			))

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

	newUuid := uuid.NewID()

	err := postController.Service.CreatePost(inputs.CreatePostInput{
		ID:          newUuid,
		Title:       createPostRequestDto.Title,
		Description: createPostRequestDto.Description,
		Content:     createPostRequestDto.Content,
		Status:      postValueObject.PostStatusDraft,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			errors.NewInValidError(
				"post create failed",
				err.Error(),
			))

		return
	}

	// This method would typically create a new post and return it.
	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"data": gin.H{
			"id":          newUuid,
			"title":       createPostRequestDto.Title,
			"description": createPostRequestDto.Description,
			"status":      postValueObject.PostStatusDraft,
		},
	})
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
//	@Success		200	{object}	domain.Post	"Post retrieved successfully"
//	@Failure		404	{object}	errors.AppError		"Page not found"
//	@Router			/posts/{id} [get]
func (postController PostController) Show(c *gin.Context) {

	postID := c.Param("id")

	result, err := postController.Service.GetById(postID)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			errors.NewNotFoundError("Page not found"),
		)
		return
	}
	// This method would typically retrieve a post by its ID and return it.
	// For now, we will just return a placeholder response.
	c.JSON(http.StatusOK, gin.H{
		"message": "Show post",
		"data":    result,
	})
}
