package http

import (
	"errors"
	"strings"
)

var (
	ErrNameRequired     = errors.New("name is required")
	ErrEmailRequired    = errors.New("email is required")
	ErrPasswordRequired = errors.New("password is required")
)

// RegisterUserRequestDto @Description	RegisterUserRequestDto is a data transfer object for user registration requests
type RegisterUserRequestDto struct {
	//	@Description	Name represents the name of the user
	Name string `json:"name"     validate:"required,min=3"`
	//	@Description	Email represents the email of the user
	Email string `json:"email"    validate:"required,email"`
	//	@Description	Password represents the password of the user
	Password string `json:"password" validate:"required,min=10"`
}

// Validate validates the RegisterUserRequestDto fields
func (r RegisterUserRequestDto) Validate() error {
	r.Name = strings.TrimSpace(r.Name)
	r.Email = strings.TrimSpace(r.Email)
	r.Password = strings.TrimSpace(r.Password)

	// Note: These validation just exemplary and you can write your specific validation logic as you want,
	// also you can use validator package for more complex validation

	if r.Name == "" {
		return ErrNameRequired
	}

	if r.Email == "" {
		return ErrEmailRequired
	}

	if r.Password == "" {
		return ErrPasswordRequired
	}

	return nil
}

// RegisterUserResponseDto @Description	RegisterUserResponseDto is a data transfer object for user registration responses
type RegisterUserResponseDto struct {
	//	@Description	Name represents the name of the user
	Name string `json:"name"`
	//	@Description	Email represents the email of the user
	Email string `json:"email"`
	//	@Description	Password represents the password of the user
	Password string `json:"password"`
}
