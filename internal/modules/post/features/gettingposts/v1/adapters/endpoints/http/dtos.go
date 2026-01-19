package http

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

// GetPostResponseDto
//
//	@Description	GetPostResponseDto is a data transfer object for reporting the details of a post
type GetPostResponseDto struct {
	Posts []*Post `json:"posts"`
}
