package http

type LoginRequestDto struct {
	//	@Description	Email represents the email of the user
	Email string `json:"email"    validate:"required,email"`
	//	@Description	Password represents the password of the user
	Password string `json:"password" validate:"required,min=10"`
}

type LoginResponseDto struct {
	//	@Description	AccessToken represents the access token of the user
	AccessToken string `json:"access_token"`
	//	@Description	RefreshToken represents the refresh token of the user
	RefreshToken string `json:"refresh_token"`
	//	@Description	UserID represents the unique identifier of the user
	UserID string `json:"user_id"`
}
