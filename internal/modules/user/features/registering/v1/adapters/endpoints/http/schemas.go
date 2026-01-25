package http

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/pkg/helper"
)

// Register It gets fields, types, and validation rules of creation requirement
//
//	@Summary	Get fields of creation requirement
//	@Schemes
//	@Description	Get fields of creation requirement
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Router			/schemas/users/register [get]
func Register(c *gin.Context) {
	helper.SchemaResponse(c,
		helper.BuildSchemaFromStruct(RegisterUserRequestDto{}),
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
//	@Router			/schemas/users/update [get]
func Update(c *gin.Context) {
	helper.SchemaResponse(c,
		helper.BuildSchemaFromStruct(RegisterUserRequestDto{}),
		"application/json")
}
