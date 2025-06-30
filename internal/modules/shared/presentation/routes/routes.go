package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	docs "github.com/racibaz/go-arch/docs"
	"github.com/racibaz/go-arch/pkg/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(router *gin.Engine) {

	configs := config.Get()

	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	url := fmt.Sprintf("http://%s:%s/%s", configs.Swagger.Host, configs.Swagger.Port, configs.Swagger.Path)

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL(url), // The url pointing to API definition
		ginSwagger.DefaultModelsExpandDepth(-1))
}
