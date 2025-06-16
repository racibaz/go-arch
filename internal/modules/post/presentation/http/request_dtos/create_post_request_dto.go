package request_dtos

type CreatePostRequestDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
