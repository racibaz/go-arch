package query

import "time"

type GetPostQuery struct {
	ID string // Unique identifier for the post
}

type PostView struct {
	ID          string
	UserID      string
	Title       string
	Description string
	Content     string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
