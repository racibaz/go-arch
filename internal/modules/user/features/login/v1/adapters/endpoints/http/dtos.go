package http

import (
	"errors"
	"strings"
)

var (
	EmailRequiredErr    = errors.New("email is required")
	PasswordRequiredErr = errors.New("password is required")
)

type LoginRequestDto struct {
	//	@Description	Email represents the email of the user
	Email string `json:"email"    validate:"required,email"`
	//	@Description	Password represents the password of the user
	Password string `json:"password" validate:"required,min=10"`
}

func (r LoginRequestDto) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	r.Password = strings.TrimSpace(r.Password)

	// Note: These validation just exemplary and you can write your specific validation logic as you want,
	// also you can use validator package for more complex validation

	if r.Email == "" {
		return EmailRequiredErr
	}

	if r.Password == "" {
		return PasswordRequiredErr
	}

	return nil
}

type LoginResponseDto struct {
	//	@Description	AccessToken represents the access token of the user
	AccessToken string `json:"access_token"`
	//	@Description	RefreshToken represents the refresh token of the user
	RefreshToken string `json:"refresh_token"`
	//	@Description	UserID represents the unique identifier of the user
	UserID string `json:"user_id"`
}
