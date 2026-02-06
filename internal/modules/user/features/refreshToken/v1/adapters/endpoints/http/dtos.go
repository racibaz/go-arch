package http

import (
	"errors"
	"strings"
)

var ErrInvalidRefreshToken = errors.New("invalid refresh token")

// RefreshTokenRequestDto represents the data required to request a new access token using a refresh token.
type RefreshTokenRequestDto struct {
	//	@Description	The refresh token used to obtain a new access token.
	RefreshToken string `json:"refresh_token" validate:"required,min=10"` // todo min length
}

func (r RefreshTokenRequestDto) Validate() error {
	r.RefreshToken = strings.TrimSpace(r.RefreshToken)

	// Note: These validation just exemplary and you can write your specific validation logic as you want,
	// also you can use validator package for more complex validation

	if r.RefreshToken == "" {
		return ErrInvalidRefreshToken
	}

	return nil
}

// RefreshTokenResponseDto represents the data returned after successfully refreshing an access token.
type RefreshTokenResponseDto struct {
	//	@Description	The newly generated access token.
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
}
