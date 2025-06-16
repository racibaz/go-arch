package http

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/application/usecases/inputs"
	requestDto "github.com/racibaz/go-arch/internal/modules/post/presentation/http/request_dtos"
	"github.com/racibaz/go-arch/pkg/uuid"

	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain/value_objects"
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

func (postController PostController) Store(c *gin.Context) {

	var createPostRequestDto requestDto.CreatePostRequestDto

	if err := c.ShouldBindJSON(&createPostRequestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUuid := uuid.NewUuid().ToString()

	err := postController.Service.CreatePost(inputs.CreatePostInput{
		ID:          newUuid,
		Title:       createPostRequestDto.Title,
		Description: createPostRequestDto.Description,
		Status:      postValueObject.PostStatusDraft,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"title": "Post did not created..", "message": err.Error()})
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

func (postController PostController) Show(c *gin.Context) {

	postID := c.Param("id")

	result, err := postController.Service.GetById(postID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"title": "Page not found", "message": err.Error()})
		return
	}
	// This method would typically retrieve a post by its ID and return it.
	// For now, we will just return a placeholder response.
	c.JSON(http.StatusOK, gin.H{
		"message": "Show post",
		"data":    result,
	})
}
