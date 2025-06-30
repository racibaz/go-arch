package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var (
	ErrValidation = errors.New("validation error")
)

func Get() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

type AppValidationError struct {
	Type    string              `json:"type"`
	Message string              `json:"message"`
	Cause   map[string][]string `json:"cause"`
}

type ValidationErrors struct {
	Errors map[string][]string `json:"errors"`
}

func NewValidationError(message string, cause map[string][]string) *AppValidationError {
	return &AppValidationError{
		Type:    ErrValidation.Error(),
		Message: message,
		Cause:   cause,
	}
}

func ShowRegularValidationErrors(err error) *ValidationErrors {

	validationErrors := make(map[string][]string)
	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()
		validationErrors[fieldName] = append(validationErrors[fieldName], err.Tag())
	}

	return &ValidationErrors{Errors: validationErrors}
}
