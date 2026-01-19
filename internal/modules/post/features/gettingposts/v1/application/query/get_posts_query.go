package query

import (
	"time"

	"github.com/racibaz/go-arch/pkg/helper"
)

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

type GetPostsQuery struct {
	Pagination helper.Pagination
}

type GetPostsQueryResponse struct {
	Posts []*PostView
}
