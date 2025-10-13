package domain

import (
	"github.com/racibaz/go-arch/pkg/es"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPost_Create(t *testing.T) {

	testCases := []struct {
		name string
		post Post
		err  error
	}{
		{
			name: "valid",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "aa5cbf57-cc15-490f-bf3b-1e1c627097a8",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: nil,
		},
		{
			name: "id can not be empty",
			post: Post{
				UserID:      "aa5cbf57-cc15-490f-bf3b-1e1c627097a8",
				Title:       "title with more than 10 characters",
				Description: "Description with more than 10 characters",
				Content:     "content content content content",
				Status:      PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: ErrEmptyId,
		},
		{
			name: "title min length",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "aa5cbf57-cc15-490f-bf3b-1e1c627097a8",
				Title:       "ti",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: ErrMinTitleLength,
		},
		{
			name: "description min length",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "aa5cbf57-cc15-490f-bf3b-1e1c627097a8",
				Title:       "title with more than 10 characters",
				Description: "Desc",
				Content:     "content content content",
				Status:      PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: ErrMinDescriptionLength,
		},
		{
			name: "content min length",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "aa5cbf57-cc15-490f-bf3b-1e1c627097a8",
				Title:       "title with more than 10 characters",
				Description: "Description with more than 10 characters",
				Content:     "cont",
				Status:      PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: ErrMinContentLength,
		},
		{
			name: "invalid status",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "aa5cbf57-cc15-490f-bf3b-1e1c627097a8",
				Title:       "title with more than 10 characters",
				Description: "Description with more than 10 characters",
				Content:     "content content content",
				Status:      -1,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			err: ErrInvalidStatus,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := Create(
				tc.post.ID(),
				tc.post.UserID,
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
