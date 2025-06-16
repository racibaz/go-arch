package routes

import (
	"github.com/gin-gonic/gin"
	postRoutes "github.com/racibaz/go-arch/internal/modules/post/presentation/routes"
)

func RegisterRoutes(router *gin.Engine) {
	postRoutes.Routes(router)
	// You can add more module routes here in the future
}
