package presentation

// Post It is a best practice to keep DTOs in stable when you need to change your dto model such as "GetPostResponseDto"
// Post dto model
type Post struct {
	//	@Description	Title is the title of the post
	Title string `json:"title"`
	//	@Description	Description is the description of the post
	Description string `json:"description"`
	//	@Description	Content is the content of the post
	Content string `json:"content"`
	//	@Description	Status is the status of the post
	Status string `json:"status"`
}

// CreatePostResponseDto
//
//	@Description	CreatePostResponseDto is a data transfer object for reporting the details of a created post
type CreatePostResponseDto struct {
	Post *Post `json:"post"`
}

// GetPostResponseDto
//
//	@Description	GetPostResponseDto is a data transfer object for reporting the details of a post
type GetPostResponseDto struct {
	Post *Post `json:"post"`
}

// CreatePostRequestDto
//
//	@Description	CreatePostRequestDto is a data transfer object for creating a post
type CreatePostRequestDto struct {
	//	@Description	UserId is the ID of the user creating the post
	UserId string `json:"user_id"     validate:"required,uuid"`
	//	@Description	Title is the title of the post
	Title string `json:"title"       validate:"required,min=10"`
	//	@Description	Description is the description of the post
	Description string `json:"description" validate:"required,min=10"`
	//	@Description	Content is the content of the post
	Content string `json:"content"     validate:"required,min=10"`
}
