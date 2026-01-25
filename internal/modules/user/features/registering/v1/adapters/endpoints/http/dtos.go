package http

// RegisterUserRequestDto @Description	RegisterUserRequestDto is a data transfer object for user registration requests
type RegisterUserRequestDto struct {
	//	@Description	UserName represents the username of the user
	UserName string `json:"user_name" validate:"required,min=3"`
	//	@Description	Email represents the email of the user
	Email string `json:"email" validate:"required,email"`
	//	@Description	Password represents the password of the user
	Password string `json:"password"   validate:"required,min=10"`
}

// RegisterUserResponseDto @Description	RegisterUserResponseDto is a data transfer object for user registration responses
type RegisterUserResponseDto struct {
	//	@Description	UserName represents the username of the user
	UserName string `json:"user_name"`
	//	@Description	Email represents the email of the user
	Email string `json:"email"`
	//	@Description	Password represents the password of the user
	Password string `json:"password"`
}
