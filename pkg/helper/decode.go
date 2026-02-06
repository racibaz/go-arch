package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Validator interface {
	Validate() error
}

// Decode decodes the JSON body of a request into a struct of type T
func Decode[T Validator](c *gin.Context) (*T, error) {
	var v T

	if c == nil {
		return nil, fmt.Errorf("gin context is nil")
	}

	if err := c.ShouldBindJSON(&v); err != nil {
		return nil, err
	}

	if err := v.Validate(); err != nil {
		return &v, err
	}

	return &v, nil
}
