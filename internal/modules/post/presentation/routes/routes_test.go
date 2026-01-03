package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func TestRoutes(t *testing.T) {
	// Test that Routes function doesn't panic and registers routes correctly
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Routes function panicked: %v", r)
		}
	}()

	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Call the Routes function
	Routes(router)

	// Check that routes were registered
	routes := router.Routes()

	// Should have at least 2 routes: GET /api/v1/posts/:id and POST /api/v1/posts/
	expectedRoutes := map[string]string{
		"GET":  "/api/v1/posts/:id",
		"POST": "/api/v1/posts/",
	}

	foundRoutes := make(map[string]string)
	for _, route := range routes {
		foundRoutes[route.Method] = route.Path
	}

	for method, path := range expectedRoutes {
		if foundPath, exists := foundRoutes[method]; !exists {
			t.Errorf("Expected route %s %s not found", method, path)
		} else if foundPath != path {
			t.Errorf("Expected route path %s, got %s", path, foundPath)
		}
	}

	if len(foundRoutes) < len(expectedRoutes) {
		t.Errorf("Expected at least %d routes, got %d", len(expectedRoutes), len(foundRoutes))
	}
}

func TestGrpcRoutes(t *testing.T) {
	// Test that GrpcRoutes function doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("GrpcRoutes function panicked: %v", r)
		}
	}()

	// Create a mock gRPC server
	grpcServer := grpc.NewServer()

	// Call the GrpcRoutes function
	GrpcRoutes(grpcServer)

	// The function should complete without panicking
	// We can't easily test the internal gRPC service registration without more complex mocking
	// but we can at least verify the function executes
}

func TestRoutes_Grouping(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	Routes(router)

	routes := router.Routes()

	// Check that all routes start with /api/v1
	for _, route := range routes {
		if len(route.Path) < 7 || route.Path[:7] != "/api/v1" {
			t.Errorf("Route path %s does not start with /api/v1", route.Path)
		}
	}
}

func TestRoutes_PostGroup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	Routes(router)

	routes := router.Routes()

	// Check that routes contain /posts/ in the path
	found := false
	for _, route := range routes {
		if len(route.Path) >= 13 && route.Path[:13] == "/api/v1/posts" {
			found = true
			break
		}
	}

	if !found {
		t.Error("No routes found with /api/v1/posts prefix")
	}
}
