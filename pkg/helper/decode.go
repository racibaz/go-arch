package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Decode decodes the JSON body of a request into a struct of type T
func Decode[T any](c *gin.Context) (*T, error) {
	var v T

	if c == nil {
		return nil, fmt.Errorf("gin context is nil")
	}

	if err := c.ShouldBindJSON(&v); err != nil {
		return nil, err
	}

	return &v, nil
}
