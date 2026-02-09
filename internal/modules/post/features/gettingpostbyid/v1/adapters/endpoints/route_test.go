package endpoints

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/application/query"
	sharedPorts "github.com/racibaz/go-arch/internal/modules/shared/testing/mocks/application/ports"
	"github.com/stretchr/testify/assert"
)

func TestMapHttpRoute(t *testing.T) {
	// Test that MapHttpRoute doesn't panic and registers routes correctly
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MapHttpRoute function panicked: %v", r)
		}
	}()

	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Create mock queries handler
	mockQueryHandler := sharedPorts.NewMockQueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse](
		t,
	)

	// Call the MapHttpRoute function
	MapHttpRoute(router, mockQueryHandler)

	// Check that routes were registered
	routes := router.Routes()

	// Should have exactly 1 route: GET /api/v1/posts/:id
	expectedRoute := "/api/v1/posts/:id"
	expectedMethod := "GET"

	found := false
	for _, route := range routes {
		if route.Method == expectedMethod && route.Path == expectedRoute {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected route %s %s not found", expectedMethod, expectedRoute)
	}

	// Verify we have exactly one route
	assert.Len(t, routes, 1, "Should have exactly one route registered")
}

func TestMapHttpRoute_RouteStructure(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockQueryHandler := sharedPorts.NewMockQueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse](
		t,
	)

	MapHttpRoute(router, mockQueryHandler)

	routes := router.Routes()

	// Check that the route has the correct structure
	for _, route := range routes {
		// Should start with /api/v1
		assert.True(t, len(route.Path) >= 7 && route.Path[:7] == "/api/v1",
			"Route path %s should start with /api/v1", route.Path)

		// Should contain /posts/
		assert.Contains(
			t,
			route.Path,
			"/posts/",
			"Route path %s should contain /posts/",
			route.Path,
		)

		// Should be a GET method
		assert.Equal(t, "GET", route.Method, "Route should be a GET method")

		// Should have parameter :id
		assert.Contains(
			t,
			route.Path,
			":id",
			"Route path %s should contain parameter :id",
			route.Path,
		)
	}
}

func TestMapHttpRoute_HandlerCreation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockQueryHandler := sharedPorts.NewMockQueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse](
		t,
	)

	// This should not panic and should create the handler internally
	assert.NotPanics(t, func() {
		MapHttpRoute(router, mockQueryHandler)
	})

	// Verify routes were created
	routes := router.Routes()
	assert.NotEmpty(t, routes, "Routes should be registered")
}

func TestMapHttpRoute_WithNilHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// This should panic or handle nil gracefully
	// Depending on implementation, adjust the test accordingly
	defer func() {
		if r := recover(); r != nil {
			// If it panics, that's acceptable behavior for nil input
			t.Logf("Function panicked with nil handler as expected: %v", r)
		}
	}()

	// This might panic depending on how the HTTP handler handles nil
	MapHttpRoute(router, nil)
}

func TestMapHttpRoute_RoutePathFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockQueryHandler := sharedPorts.NewMockQueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse](
		t,
	)

	MapHttpRoute(router, mockQueryHandler)

	routes := router.Routes()

	for _, route := range routes {
		// Path should follow REST conventions
		assert.Regexp(t, `^/api/v\d+/posts/:\w+$`, route.Path,
			"Route path %s should match REST pattern /api/v{version}/posts/{param}", route.Path)
	}
}

func TestMapHttpRoute_Grouping(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockQueryHandler := sharedPorts.NewMockQueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse](
		t,
	)

	MapHttpRoute(router, mockQueryHandler)

	routes := router.Routes()

	// Verify that routes are properly grouped under /api/v1/posts
	for _, route := range routes {
		assert.Contains(t, route.Path, "/api/v1/posts",
			"Route %s should be grouped under /api/v1/posts", route.Path)
	}
}

func TestMapHttpRoute_MultipleCalls(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockQueryHandler1 := sharedPorts.NewMockQueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse](
		t,
	)
	mockQueryHandler2 := sharedPorts.NewMockQueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse](
		t,
	)

	// Call MapHttpRoute first time - should succeed
	MapHttpRoute(router, mockQueryHandler1)

	routes := router.Routes()
	assert.Len(t, routes, 1, "Should have one route after first call")

	// Call MapHttpRoute second time - should panic because route is already registered
	assert.Panics(t, func() {
		MapHttpRoute(router, mockQueryHandler2)
	}, "Should panic when signup duplicate route")

	// Routes should still be registered (panic happens after registration attempt)
	routesAfter := router.Routes()
	assert.Len(t, routesAfter, 1, "Should still have one route after panic")
}
