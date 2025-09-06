package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	openapiSpec "github.com/racibaz/go-arch/api/openapi-spec"
	"github.com/racibaz/go-arch/pkg/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(router *gin.Engine) {

	configs := config.Get()
	openapiSpec.SwaggerInfo.Title = configs.App.Name
	openapiSpec.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	url := fmt.Sprintf("http://%s:%s/%s", configs.Swagger.Host, configs.Swagger.Port, configs.Swagger.Path)

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL(url), // The url pointing to API definition
		ginSwagger.DefaultModelsExpandDepth(-1))
}
