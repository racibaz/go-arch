package http

// RefreshTokenRequestDto
type RefreshTokenRequestDto struct {
	//	@Description	The refresh token used to obtain a new access token.
	RefreshToken string `json:"refresh_token" validate:"required,min=10"` // todo min length
}

// RefreshTokenResponseDto
type RefreshTokenResponseDto struct {
	//	@Description	The newly generated access token.
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
}
