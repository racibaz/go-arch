package integration

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/application/usecases"
	inMemoryRepository "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/in_memory"
	postController "github.com/racibaz/go-arch/internal/modules/post/presentation/http"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_post_CreatePostIntegration(t *testing.T) {

	//Arrange
	repo := inMemoryRepository.NewInMemoryPostRepository()
	uc := usecases.NewCreatetPostUseCase(repo)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/posts", postController.NewPostController(uc).Store)
	server := httptest.NewServer(router)
	defer server.Close()

	reqBody := map[string]interface{}{
		"title":       "post title",
		"description": "post description",
	}
	payload, _ := json.Marshal(reqBody)

	//Act
	w := httptest.NewRecorder()
	// Create a new HTTP request with the payload
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(string(payload)))
	router.ServeHTTP(w, req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing request body:", err)
		}
	}(req.Body)

	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Post created successfully")
	assert.Contains(t, w.Body.String(), "post title")
	assert.Contains(t, w.Body.String(), "post description")
}
