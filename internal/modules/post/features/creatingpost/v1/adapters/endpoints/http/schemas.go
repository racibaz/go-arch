package http

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/pkg/helper"
)

// Create It gets fields, types, and validation rules of creation requirement
//
//	@Summary	Get fields of creation requirement
//	@Schemes
//	@Description	Get fields of creation requirement
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Router			/schemas/posts/create [get]
func Create(c *gin.Context) {
	helper.SchemaResponse(c,
		helper.BuildSchemaFromStruct(CreatePostRequestDto{}),
		"application/json")
}

// todo when add update feature, it will be there.

// Update It gets fields, types, and validation rules of update requirement
//
//	@Summary	Get fields of update requirement
//	@Schemes
//	@Description	Get fields of update requirement
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Router			/schemas/posts/update [get]
func Update(c *gin.Context) {
	helper.SchemaResponse(c,
		helper.BuildSchemaFromStruct(CreatePostRequestDto{}),
		"application/json")
}
