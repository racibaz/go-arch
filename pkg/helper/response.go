package helper

import (
	"github.com/gin-gonic/gin"
	errors "github.com/racibaz/go-arch/pkg/error"
	"net/http"
)

// ValidationErrorResponse sends a standardized validation error response
func ValidationErrorResponse(c *gin.Context, message string, err error) {
	c.JSON(
		http.StatusBadRequest,
		errors.NewValidationError(
			message,
			errors.ShowRegularValidationErrors(err).Errors,
		),
	)

	return
}

// ErrorResponse sends a standardized error response
func ErrorResponse(c *gin.Context, message string, err error, status int) {

	errorMap := make(map[string][]string)
	errorMap["error"] = []string{err.Error()}

	c.JSON(
		status,
		errors.NewInValidError(
			message,
			errorMap,
		))

	return
}

// SuccessResponse sends a standardized success response
func SuccessResponse(c *gin.Context, message string, data any, status int) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
