package http

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/pkg/helper"
)

// Login godoc
//
//	@Summary	Login Schema
//	@Schemes
//	@Description	It is the schema for login endpoint
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Router			/schemas/auth/login [post]
func Login(c *gin.Context) {
	helper.SchemaResponse(c,
		helper.BuildSchemaFromStruct(LoginRequestDto{}),
		"application/json")
}
