package http

import (
	"errors"
	"strings"
)

var ErrInvalidRefreshToken = errors.New("invalid refresh token")

// MeRequestDto represents the data required to request a new access token using a refresh token.
type MeRequestDto struct {
	//	@Description	The refresh token used to obtain a new access token.
	RefreshToken string `json:"refresh_token" validate:"required,min=10"`
}

func (r MeRequestDto) Validate() error {
	r.RefreshToken = strings.TrimSpace(r.RefreshToken)

	// Note: These validation just exemplary and you can write your specific validation logic as you want,
	// also you can use validator package for more complex validation

	if r.RefreshToken == "" {
		return ErrInvalidRefreshToken
	}

	return nil
}

// MeResponseDto @Description	MeResponseDto is a data transfer object for user information response
type MeResponseDto struct {
	//	@Description	Name represents the name of the user
	Name string `json:"name"`
	//	@Description	Email represents the email of the user
	Email string `json:"email"`
	//	@Description	Status represents the status of the user (e.g., active, inactive)
	Status string `json:"status"`
	//	@Description	CreatedAt represents the date and time when the user was created
	CreatedAt string `json:"created_at"`
}
