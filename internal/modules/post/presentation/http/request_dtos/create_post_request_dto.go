package request_dtos

// CreatePostRequestDto
// @Description CreatePostRequestDto is a data transfer object for creating a post
type CreatePostRequestDto struct {
	// @Description Title is the title of the post
	Title string `json:"title" validate:"required,min=10"`
	// @Description Description is the description of the post
	Description string `json:"description" validate:"required,min=10"`
	// @Description Content is the content of the post
	Content string `json:"content" validate:"required,min=10"`
}
