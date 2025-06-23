package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	docs "github.com/racibaz/go-arch/docs"
	config "github.com/racibaz/go-arch/pkg/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// Initialize the configuration
	config.Set()

	configs := config.Get()

	r := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	url := fmt.Sprintf("http://%s:%s/%s", configs.Swagger.Host, configs.Swagger.Port, configs.Swagger.Path)

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL(url), // The url pointing to API definition
		ginSwagger.DefaultModelsExpandDepth(-1))

	err := r.Run(fmt.Sprintf(":%s", configs.Swagger.Port))
	if err != nil {
		return
	} // Run on port 8081
}
