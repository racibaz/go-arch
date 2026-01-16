package query

import "time"

type GetPostByIdQuery struct {
	ID string // Unique identifier for the post
}

type GetPostByIdQueryResponse struct {
	ID          string
	UserID      string
	Title       string
	Description string
	Content     string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
