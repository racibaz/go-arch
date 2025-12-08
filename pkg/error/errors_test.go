package errors_test

import (
	"testing"

	apperrors "github.com/racibaz/go-arch/pkg/error"
)

func TestAppError_Error_WithCause(t *testing.T) {
	err := &apperrors.AppError{
		Type:    "validation error",
		Message: "invalid request",
		Cause:   "email missing",
	}

	expected := "validation error: invalid request (cause: email missing)"
	if err.Error() != expected {
		t.Fatalf("expected '%s', got '%s'", expected, err.Error())
	}
}

func TestAppError_Error_WithoutCause(t *testing.T) {
	err := &apperrors.AppError{
		Type:    "validation error",
		Message: "invalid request",
	}

	expected := "validation error: invalid request"
	if err.Error() != expected {
		t.Fatalf("expected '%s', got '%s'", expected, err.Error())
	}
}

func TestNewInValidError(t *testing.T) {
	err := apperrors.NewInValidError("bad data", "bad field")

	if err.Type != apperrors.ErrInValid.Error() {
		t.Fatalf("wrong type: %s", err.Type)
	}

	if err.Message != "bad data" {
		t.Fatalf("wrong message: %s", err.Message)
	}

	if err.Cause != "bad field" {
		t.Fatalf("wrong cause: %s", err.Cause)
	}
}

func TestNewErrExistFoundError(t *testing.T) {
	err := apperrors.NewErrExistFoundError("already exists", "duplicate email")

	if err.Type != apperrors.ErrExistFound.Error() {
		t.Fatalf("wrong type: %s", err.Type)
	}

	if err.Message != "already exists" {
		t.Fatalf("wrong message: %s", err.Message)
	}

	if err.Cause != "duplicate email" {
		t.Fatalf("wrong cause: %s", err.Cause)
	}
}

func TestNewNotFoundError(t *testing.T) {
	err := apperrors.NewNotFoundError("user not found")

	if err.Type != apperrors.ErrNotFound.Error() {
		t.Fatalf("wrong type: %s", err.Type)
	}

	if err.Message != "user not found" {
		t.Fatalf("wrong message: %s", err.Message)
	}

	if err.Cause != "" {
		t.Fatalf("expected empty cause, got %s", err.Cause)
	}
}

func TestNewUnauthorizedError(t *testing.T) {
	err := apperrors.NewUnauthorizedError("access denied", "invalid token")

	if err.Type != apperrors.ErrUnauthorized.Error() {
		t.Fatalf("wrong type: %s", err.Type)
	}

	if err.Message != "access denied" {
		t.Fatalf("wrong message: %s", err.Message)
	}

	if err.Cause != "invalid token" {
		t.Fatalf("wrong cause: %s", err.Cause)
	}
}
