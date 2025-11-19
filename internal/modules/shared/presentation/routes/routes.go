package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	openapiSpec "github.com/racibaz/go-arch/api/openapi-spec"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/prometheus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

//	@BasePath	/api/v1

// Health Check Endpoint
// @Summary health
// @Schemes
// @Description do health check
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /health [get]
func health(g *gin.Context) {
	g.JSON(http.StatusOK, "ok")
}

func Routes(router *gin.Engine) {

	// Middleware to collect metrics
	router.Use(prometheus.MetricsMiddleware)

	// Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	configs := config.Get()
	openapiSpec.SwaggerInfo.Title = configs.App.Name
	openapiSpec.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	url := fmt.Sprintf("http://%s:%s/%s", configs.Swagger.Host, configs.Swagger.Port, configs.Swagger.Path)

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL(url), // The url pointing to API definition
		ginSwagger.DefaultModelsExpandDepth(-1))

	router.GET("/health", health)
}
