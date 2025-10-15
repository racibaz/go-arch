package helper

import (
	"github.com/gin-gonic/gin"
	errors "github.com/racibaz/go-arch/pkg/error"
	"net/http"
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

func ErrorResponse(c *gin.Context, message string, err error, status int) {

	c.JSON(
		status,
		errors.NewInValidError(
			message,
			err.Error(),
		))

	return
}

func SuccessResponse(c *gin.Context, message string, data any, status int) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}
