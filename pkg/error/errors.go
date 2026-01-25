package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var (
	ErrValidation   = errors.New("validation error")
	ErrInValid      = errors.New("invalid error")
	ErrInDecode     = errors.New("invalid error")
	ErrNotFound     = errors.New("resource not found")
	ErrExistFound   = errors.New("exist found")
	ErrConflict     = errors.New("resource conflict")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrInternal     = errors.New("internal error")
)

// appError represents a structured application error
type appError struct {
	Status  int                 `json:"status"`
	Type    string              `json:"type"`
	Message string              `json:"message"`
	Cause   map[string][]string `json:"cause"`
}

// ValidationErrors represents validation errors
type ValidationErrors struct {
	Errors map[string][]string `json:"errors"`
}

// Error implements the error interface for appError
func (e *appError) Error() string {
	if len(e.Cause) > 0 {
		return fmt.Sprintf("%s: %s (cause: %v)", e.Type, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// NewValidationError creates a new validation error
func NewValidationError(message string, cause map[string][]string) *appError {
	return &appError{
		Status:  http.StatusBadRequest,
		Type:    ErrValidation.Error(),
		Message: message,
		Cause:   cause,
	}
}

// ShowRegularValidationErrors formats validation errors from the validator package
func ShowRegularValidationErrors(err error) *ValidationErrors {
	validationErrors := make(map[string][]string)
	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()
		validationErrors[fieldName] = append(validationErrors[fieldName], err.Tag())
	}

	return &ValidationErrors{Errors: validationErrors}
}

// NewDecodeError creates a new decode error
func NewDecodeError(message string) *appError {
	return &appError{
		Status:  http.StatusBadRequest,
		Type:    ErrInValid.Error(),
		Message: message,
		Cause:   make(map[string][]string),
	}
}

// NewInValidError creates a new invalid error
func NewInValidError(message string, cause map[string][]string) *appError {
	return &appError{
		Status:  http.StatusBadRequest,
		Type:    ErrInValid.Error(),
		Message: message,
		Cause:   cause,
	}
}

// NewErrExistFoundError creates a new exist found error
func NewErrExistFoundError(message string, cause map[string][]string) *appError {
	return &appError{
		Status:  http.StatusInternalServerError, // StatusInternalServerError
		Type:    ErrExistFound.Error(),
		Message: message,
		Cause:   cause,
	}
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string) *appError {
	return &appError{
		Status:  http.StatusNotFound,
		Type:    ErrNotFound.Error(),
		Message: message,
	}
}

// NewConflictError creates a new conflict error
func NewUnauthorizedError(message string, cause map[string][]string) *appError {
	return &appError{
		Status:  http.StatusUnauthorized,
		Type:    ErrUnauthorized.Error(),
		Message: message,
		Cause:   cause,
	}
}
