package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errors "github.com/racibaz/go-arch/pkg/error"
)

func Decode[T any](c *gin.Context) (T, error) {
	var v T

	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(
			http.StatusBadRequest,
			errors.NewInValidError(
				"Invalid request body",
				err.Error(),
			))
	}

	return v, nil
}
