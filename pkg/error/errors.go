package errors

import (
	"errors"
	"fmt"
)

var (
	ErrValidation   = errors.New("validation error")
	ErrInValid      = errors.New("invalid error")
	ErrNotFound     = errors.New("resource not found")
	ErrExistFound   = errors.New("exist found")
	ErrConflict     = errors.New("resource conflict")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrInternal     = errors.New("internal error")
)

type AppError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Cause   string `json:"cause"`
}

func (e *AppError) Error() string {
	if e.Cause != "" {
		return fmt.Sprintf("%s: %s (cause: %v)", e.Type, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func NewInValidError(message string, cause string) *AppError {
	return &AppError{
		Type:    ErrInValid.Error(),
		Message: message,
		Cause:   cause,
	}
}

func NewErrExistFoundError(message string, cause string) *AppError {
	return &AppError{
		Type:    ErrExistFound.Error(),
		Message: message,
		Cause:   cause,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:    ErrNotFound.Error(),
		Message: message,
	}
}

func NewUnauthorizedError(message string, cause string) *AppError {
	return &AppError{
		Type:    ErrUnauthorized.Error(),
		Message: message,
		Cause:   cause,
	}
}

//if you want to add more general error types, you can do so here
