package http

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingposts/v1/application/query"
)

func FromPostViewToHTTP(inputs *query.GetPostsQueryResponse) []*Post {
	if inputs.Posts == nil {
		return nil
	}

	posts := make([]*Post, 0, len(inputs.Posts))

	for _, post := range inputs.Posts {

		status := domain.PostStatus(post.Status)

		p := &Post{
			Title:       post.Title,
			Description: post.Description,
			Content:     post.Content,
			Status:      status.String(),
		}

		posts = append(posts, p)
	}

	return posts
}
