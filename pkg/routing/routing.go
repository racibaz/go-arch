package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/providers/routes"
	"github.com/racibaz/go-arch/pkg/config"
)

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())
}

func Init() {
	conf := config.Get()
	// Set gin mode
	gin.SetMode(conf.GinMode())

	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}
