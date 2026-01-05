package helper

import (
	"github.com/go-playground/validator/v10"
)

func Get() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}
