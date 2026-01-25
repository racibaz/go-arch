package http

// @Description	User is a data transfer object representing a user
type User struct {
	// @Description	UserName represents the username of the user
	UserName string `json:"user_name"`
	// @Description	Email represents the email of the user
	Email string `json:"email"`
	// @Description	Password represents the password of the user
	Password string `json:"password"`
}

// @Description	RegisterUserRequestDto is a data transfer object for user registration requests
type RegisterUserRequestDto struct {
	// @Description	User represents the user information for registration
	User *User `json:"user"`
}

// @Description	RegisterUserResponseDto is a data transfer object for user registration responses
type RegisterUserResponseDto struct {
	//	@Description	User represents the registered user information
	User *User `json:"user"`
}
