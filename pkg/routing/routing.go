package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/providers/routes"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())
}
