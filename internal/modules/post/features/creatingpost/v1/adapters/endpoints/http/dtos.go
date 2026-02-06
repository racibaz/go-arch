package http

import (
	"errors"
	"strings"
)

var (
	TitleRequiredErr       = errors.New("title is required")
	DescriptionRequiredErr = errors.New("description is required")
	ContentRequiredErr     = errors.New("content is required")
)

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

// Validate validates the CreatePostResponseDto fields
func (r CreatePostRequestDto) Validate() error {
	r.Title = strings.TrimSpace(r.Title)
	r.Description = strings.TrimSpace(r.Description)
	r.Content = strings.TrimSpace(r.Content)

	// Note: These validation just exemplary and you can write your specific validation logic as you want,
	// also you can use validator package for more complex validation

	if r.Title == "" {
		return TitleRequiredErr
	}

	if r.Description == "" {
		return DescriptionRequiredErr
	}

	if r.Content == "" {
		return ContentRequiredErr
	}

	return nil
}

// CreatePostResponseDto
//
//	@Description	CreatePostResponseDto is a data transfer object for reporting the details of a created post
type CreatePostResponseDto struct {
	//	@Description	Title is the title of the post
	Title string `json:"title"`
	//	@Description	Description is the description of the post
	Description string `json:"description"`
	//	@Description	Content is the content of the post
	Content string `json:"content"`
	//	@Description	Status is the status of the post
	Status string `json:"status"`
}
