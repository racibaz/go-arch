package helper

import (
	"github.com/gin-gonic/gin"
	errors "github.com/racibaz/go-arch/pkg/error"
)

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
