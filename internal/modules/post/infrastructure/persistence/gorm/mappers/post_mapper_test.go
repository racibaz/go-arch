package mappers

import (
	domain "github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/pkg/es"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPostMapper_ToDomain(t *testing.T) {

	testCases := []struct {
		name string
		post domain.Post
		err  error
	}{
		{
			name: "valid",
			post: domain.Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", domain.PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      domain.PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			postEntity := ToPersistence(tc.post)

			require.NotEmpty(t, postEntity.ID)
			require.NotEmpty(t, postEntity.Title)
			require.NotEmpty(t, postEntity.Description)
			require.NotEmpty(t, postEntity.Content)
			require.NotEmpty(t, postEntity.Status)

			postModel := ToDomain(postEntity)

			require.NotEmpty(t, postModel.ID())
			require.NotEmpty(t, postModel.Title)
			require.NotEmpty(t, postModel.Description)
			require.NotEmpty(t, postModel.Content)
			require.NotEmpty(t, postModel.Status)

			require.Equal(t, tc.post.ID(), postModel.ID())
			require.Equal(t, tc.post.Title, postModel.Title)
			require.Equal(t, tc.post.Description, postModel.Description)
			require.Equal(t, tc.post.Content, postModel.Content)
			require.Equal(t, tc.post.Status, postModel.Status)
			require.WithinDuration(t, tc.post.CreatedAt, postModel.CreatedAt, time.Second)
			require.WithinDuration(t, tc.post.UpdatedAt, postModel.UpdatedAt, time.Second)
		})
	}
}

func TestPostMapper_ToPersistence(t *testing.T) {

	testCases := []struct {
		name string
		post domain.Post
		err  error
	}{
		{
			name: "valid",
			post: domain.Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", domain.PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      domain.PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			postEntity := ToPersistence(tc.post)

			require.NotEmpty(t, postEntity.ID)
			require.NotEmpty(t, postEntity.Title)
			require.NotEmpty(t, postEntity.Description)
			require.NotEmpty(t, postEntity.Content)
			require.NotEmpty(t, postEntity.Status)

			postModel := ToDomain(postEntity)

			require.NotEmpty(t, postModel.ID())
			require.NotEmpty(t, postModel.Title)
			require.NotEmpty(t, postModel.Description)
			require.NotEmpty(t, postModel.Content)
			require.NotEmpty(t, postModel.Status)

			require.Equal(t, tc.post.ID(), postModel.ID())
			require.Equal(t, tc.post.Title, postModel.Title)
			require.Equal(t, tc.post.Description, postModel.Description)
			require.Equal(t, tc.post.Content, postModel.Content)
			require.Equal(t, tc.post.Status, postModel.Status)
			require.WithinDuration(t, tc.post.CreatedAt, postModel.CreatedAt, time.Second)
			require.WithinDuration(t, tc.post.UpdatedAt, postModel.UpdatedAt, time.Second)
		})
	}
}
