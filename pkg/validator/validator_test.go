package validator_test

import (
	"errors"
	"testing"

	myvalidator "github.com/racibaz/go-arch/pkg/validator"
)

type TestUser struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=18"`
}

func TestGet_ShouldReturnValidator(t *testing.T) {
	v := myvalidator.Get()
	if v == nil {
		t.Fatal("expected validator instance, got nil")
	}
}

func TestNewValidationError_ShouldCreateCorrectErrorObject(t *testing.T) {
	cause := map[string][]string{
		"Name": {"required"},
	}

	err := myvalidator.NewValidationError("invalid input", cause)

	if err.Type != "validation error" {
		t.Fatalf("expected type 'validation error', got '%s'", err.Type)
	}

	if err.Message != "invalid input" {
		t.Fatalf("unexpected message: %s", err.Message)
	}

	if err.Cause["Name"][0] != "required" {
		t.Fatal("expected cause to contain required tag")
	}
}

func TestShowRegularValidationErrors_ShouldConvertErrors(t *testing.T) {
	v := myvalidator.Get()

	user := TestUser{
		Name:  "",
		Email: "invalid-email",
		Age:   10,
	}

	err := v.Struct(user)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	result := myvalidator.ShowRegularValidationErrors(err)

	if result == nil {
		t.Fatal("expected ValidationErrors, got nil")
	}

	if len(result.Errors) == 0 {
		t.Fatal("expected some validation errors")
	}

	if result.Errors["Name"][0] != "required" {
		t.Fatalf("unexpected error for Name: %v", result.Errors["Name"])
	}

	if result.Errors["Email"][0] != "email" {
		t.Fatalf("unexpected error for Email: %v", result.Errors["Email"])
	}

	if result.Errors["Age"][0] != "gte" {
		t.Fatalf("unexpected error for Age: %v", result.Errors["Age"])
	}
}

func TestShowRegularValidationErrors_ShouldPanic_OnWrongErrorType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on invalid error type")
		}
	}()

	myvalidator.ShowRegularValidationErrors(
		errors.New("some random error"),
	)
}
