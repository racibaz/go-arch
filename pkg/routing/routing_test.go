package routing

import (
	"github.com/racibaz/go-arch/pkg/config"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// ---- SETUP ----
	// Set environment variable for testing
	os.Setenv("APP_ENV", "test")

	config.Set("../../config/", "../../.env")

	// Run tests
	code := m.Run()

	// Clean up environment variable
	os.Unsetenv("APP_ENV")

	os.Exit(code)
}

func TestInit_SetsGinModeAndCreatesRouter(t *testing.T) {

	// Act
	Init()

	// Assert
	router := GetRouter()
	require.NotNil(t, router)
	assert.Equal(t, gin.TestMode, gin.Mode())

	t.Cleanup(func() {
		router = nil
	})
}

func TestRegisterRoutes_CreatesRoutes(t *testing.T) {
	// Arrange
	Init()
	RegisterRoutes()

	// Act
	req, _ := http.NewRequest("GET", "/api/health", nil)
	w := httptest.NewRecorder()

	GetRouter().ServeHTTP(w, req)

	// Assert
	assert.NotEqual(t, http.StatusNotFound, w.Code)
}
