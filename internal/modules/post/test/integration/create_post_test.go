package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/application/commands"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/messaging/rabbitmq"
	inMemoryRepository "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/in_memory"
	postController "github.com/racibaz/go-arch/internal/modules/post/presentation/http"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/logger"
	rabbitmqConn "github.com/racibaz/go-arch/pkg/messaging/rabbitmq"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type TestSuite struct {
	Repo             *inMemoryRepository.Repository
	Logger           logger.Logger
	RabbitMQ         *rabbitmqConn.RabbitMQ
	MessagePublisher *rabbitmq.PostMessagePublisher
	UseCase          *commands.CreatePostService
}

var suite TestSuite

func TestMain(m *testing.M) {
	// ---- SETUP ----

	// Config
	config.Set("../../../../../config/", "../../../../../.env")

	// Repo
	suite.Repo = inMemoryRepository.New()

	// Logger
	suite.Logger, _ = logger.NewZapLogger()

	// RabbitMQ Connect
	rabbitmqConn.Connect()
	conn := rabbitmqConn.Connection()
	if conn == nil {
		fmt.Println("Failed to connect to RabbitMQ")
		os.Exit(1)
	}
	suite.RabbitMQ = rabbitmqConn.Connection()

	// Message Publisher
	suite.MessagePublisher = rabbitmq.NewPostMessagePublisher(
		suite.RabbitMQ,
		suite.Logger,
	)

	// Use Case
	suite.UseCase = commands.NewCreatePostService(
		suite.Repo,
		suite.Logger,
		suite.MessagePublisher,
	)

	// ---- RUN TESTS ----
	code := m.Run()

	// ---- TEARDOWN ----
	suite.RabbitMQ.Close()

	os.Exit(code)
}

func TestCreatePostIntegration(t *testing.T) {
	// Arrange
	uc := commands.NewCreatePostService(suite.Repo, suite.Logger, suite.MessagePublisher)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/posts", postController.NewPostController(uc).Store)

	reqBody := map[string]interface{}{
		"user_id":     "7b3a4d03-bcb9-47ce-b721-a156edd406f0",
		"title":       "test test test 333333",
		"description": "test 11111111112 11111111112",
		"content":     "test 11111111112 11111111112",
	}

	payload, _ := json.Marshal(reqBody)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/posts", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Post created successfully")
}

func TestCreatePostWithoutTitleIntegration(t *testing.T) {

	// Arrange
	uc := commands.NewCreatePostService(suite.Repo, suite.Logger, suite.MessagePublisher)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/posts", postController.NewPostController(uc).Store)
	server := httptest.NewServer(router)
	defer server.Close()

	reqBody := map[string]interface{}{
		"user_id":     "7b3a4d03-bcb9-47ce-b721-a156edd406f0",
		"description": "post description",
		"content":     "post content",
	}
	payload, _ := json.Marshal(reqBody)

	//Act
	w := httptest.NewRecorder()
	// Create a new HTTP request with the payload
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(string(payload)))

	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreatePostWithTitleLessTenLettersIntegration(t *testing.T) {

	// Arrange
	uc := commands.NewCreatePostService(suite.Repo, suite.Logger, suite.MessagePublisher)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/posts", postController.NewPostController(uc).Store)
	server := httptest.NewServer(router)
	defer server.Close()

	reqBody := map[string]interface{}{
		"user_id":     "7b3a4d03-bcb9-47ce-b721-a156edd406f0",
		"title":       "title",
		"description": "post description",
		"content":     "post content",
	}
	payload, _ := json.Marshal(reqBody)

	//Act
	w := httptest.NewRecorder()
	// Create a new HTTP request with the payload
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(string(payload)))

	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreatePostWithTDescriptionLessTenLettersIntegration(t *testing.T) {

	// Arrange
	uc := commands.NewCreatePostService(suite.Repo, suite.Logger, suite.MessagePublisher)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/posts", postController.NewPostController(uc).Store)
	server := httptest.NewServer(router)
	defer server.Close()

	reqBody := map[string]interface{}{
		"title":       "title test test",
		"description": "desc",
		"content":     "post content",
	}
	payload, _ := json.Marshal(reqBody)

	//Act
	w := httptest.NewRecorder()
	// Create a new HTTP request with the payload
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(string(payload)))

	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreatePostWithTContentLessTenLettersIntegration(t *testing.T) {

	// Arrange
	uc := commands.NewCreatePostService(suite.Repo, suite.Logger, suite.MessagePublisher)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/posts", postController.NewPostController(uc).Store)
	server := httptest.NewServer(router)
	defer server.Close()

	reqBody := map[string]interface{}{
		"title":       "title test test",
		"description": "description test",
		"content":     "content",
	}
	payload, _ := json.Marshal(reqBody)

	//Act
	w := httptest.NewRecorder()
	// Create a new HTTP request with the payload
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(string(payload)))

	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
