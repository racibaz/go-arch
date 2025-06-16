package routes

import (
	"github.com/gin-gonic/gin"
	postModule "github.com/racibaz/go-arch/internal/modules/post"
	postController "github.com/racibaz/go-arch/internal/modules/post/presentation/http"
)

func Routes(router *gin.Engine) {

	postModule := postModule.NewPostModule()

	newPostController := postController.NewPostController(postModule.GetService())
	router.GET("/posts/:id", newPostController.Show)
	router.POST("/posts", newPostController.Store)

}
