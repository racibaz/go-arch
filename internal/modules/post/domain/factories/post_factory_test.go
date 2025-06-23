package factories

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain/value_objects"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPostFactory(t *testing.T) {

	testCases := []struct {
		name string
		post domain.Post
		err  error
	}{
		{
			name: "valid",
			post: domain.Post{
				ID:          "acb863d4-07b4-4644-b598-7f5cc2494613",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      postValueObject.PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: nil,
		},
		{
			name: "title min length",
			post: domain.Post{
				ID:          "acb863d4-07b4-4644-b598-7f5cc2494613",
				Title:       "ti",
				Description: "Description",
				Content:     "content content content",
				Status:      postValueObject.PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: domain.ErrMinTitleLength,
		},
		{
			name: "description min length",
			post: domain.Post{
				ID:          "acb863d4-07b4-4644-b598-7f5cc2494613",
				Title:       "title with more than 10 characters",
				Description: "Desc",
				Content:     "content content content",
				Status:      postValueObject.PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: domain.ErrMinDescriptionLength,
		},
		{
			name: "content min length",
			post: domain.Post{
				ID:          "acb863d4-07b4-4644-b598-7f5cc2494613",
				Title:       "title with more than 10 characters",
				Description: "Description with more than 10 characters",
				Content:     "cont",
				Status:      postValueObject.PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: domain.ErrMinContentLength,
		},
		{
			name: "invalid status",
			post: domain.Post{
				ID:          "acb863d4-07b4-4644-b598-7f5cc2494613",
				Title:       "title with more than 10 characters",
				Description: "Description with more than 10 characters",
				Content:     "content content content",
				Status:      -1,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: domain.ErrInvalidStatus,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := New(
				tc.post.ID,
				tc.post.Title,
				tc.post.Description,
				tc.post.Content,
				tc.post.Status,
				tc.post.CreatedAt,
				tc.post.UpdatedAt)

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
