package domain

import (
	"fmt"
	"github.com/racibaz/go-arch/pkg/es"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Post_PostStatus_EqualTo(t *testing.T) {

	testCases := []struct {
		name      string
		post      Post
		otherPost Post
		result    bool
	}{
		{
			name: "valid",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			otherPost: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			result: true,
		},
		{
			name: "it can be false",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			otherPost: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusArchived,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			result: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			postStatus := tc.post.Status
			otherPostStatus := tc.otherPost.Status

			equalTo := postStatus.EqualTo(otherPostStatus)
			if equalTo {
				assert.True(t, equalTo)
			} else {
				assert.False(t, equalTo)
			}
		})
	}
}

func Test_Post_PostStatus_ToInt(t *testing.T) {

	testCases := []struct {
		name         string
		post         Post
		statusNumber int
		result       bool
	}{
		{
			name: "valid",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			statusNumber: 1,
			result:       true,
		},
		{
			name: "it can be true",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusArchived,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			statusNumber: 2,
			result:       true,
		},
		{
			name: "it can be true",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusDraft,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			statusNumber: 0,
			result:       true,
		},
		{
			name: "it can be false",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusDraft,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			statusNumber: 1,
			result:       false,
		},
		{
			name: "it can be false",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusDraft,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			statusNumber: -1,
			result:       false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			toInt, err := tc.post.Status.ToInt()

			if err != nil {
				assert.Equal(t, -1, toInt)
				assert.Contains(t, err, "invalid post status:")
			}

			var boolResult = int(toInt) == tc.statusNumber
			if boolResult {
				assert.True(t, boolResult)
			} else {
				assert.False(t, boolResult)
			}
		})
	}
}

func Test_Post_PostStatus_String(t *testing.T) {

	testCases := []struct {
		name   string
		post   Post
		text   string
		result bool
	}{
		{
			name: "published valid text",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			text:   "published",
			result: true,
		},
		{
			name: "draft valid text",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusDraft,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			text:   "draft",
			result: true,
		},
		{
			name: "archived valid text",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusArchived,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			text:   "archived",
			result: true,
		},
		{
			name: "archived invalid text",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusArchived,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			text:   "draft",
			result: false,
		},
		{
			name: "unknown post status",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      -1,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			text:   "unknown",
			result: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			text := tc.post.Status.String()

			fmt.Printf("Post Status: %s\n", text)

			if tc.result {
				assert.Equal(t, text, tc.text)
			} else {
				assert.NotEqual(t, text, tc.text)
			}
		})
	}
}

func Test_Post_PostStatus(t *testing.T) {

	testCases := []struct {
		name   string
		post   Post
		status PostStatus
		result bool
	}{
		{
			name: "is published",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			status: PostStatusPublished,
			result: true,
		},
		{
			name: "is draft",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusDraft,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			status: PostStatusDraft,
			result: true,
		},
		{
			name: "is archived",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusArchived,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			status: PostStatusArchived,
			result: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			switch tc.status {
			case PostStatusPublished:
				assert.True(t, tc.post.Status.IsPublished())
			case PostStatusArchived:
				assert.True(t, tc.post.Status.IsArchived())
			case PostStatusDraft:
				assert.True(t, tc.post.Status.IsDraft())
			}

			if tc.result {
				assert.Equal(t, tc.post.Status, tc.status)
			} else {
				assert.NotEqual(t, tc.post.Status, tc.status)
			}
		})
	}
}

func Test_Post_NewPostStatus(t *testing.T) {

	testCases := []struct {
		name   string
		post   Post
		status PostStatus
		result bool
	}{
		{
			name: "is published",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			status: PostStatusPublished,
			result: true,
		},
		{
			name: "is draft",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusDraft,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			status: PostStatusDraft,
			result: true,
		},
		{
			name: "is archived",
			post: Post{
				Aggregate:   es.NewAggregate("acb863d4-07b4-4644-b598-7f5cc2494613", PostAggregate),
				UserID:      "ed796bf3-6ff9-4b8f-9fdf-18358c2d9100",
				Title:       "title with more than 10 characters",
				Description: "Description",
				Content:     "content content content",
				Status:      PostStatusArchived,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			status: PostStatusArchived,
			result: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			newStatus := NewPostStatus(tc.status)

			if tc.result {
				assert.Equal(t, newStatus, tc.status)
			} else {
				assert.NotEqual(t, newStatus, tc.status)
			}
		})
	}
}
