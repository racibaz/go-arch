package transformers

import (
	"testing"
	"time"

	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/adapters/endpoints/dtos"
	"github.com/stretchr/testify/assert"
)

func TestFromPostCoreToHTTP(t *testing.T) {
	tests := []struct {
		name     string
		input    *domain.Post
		expected *dtos.Post
	}{
		{
			name:     "should return nil when input is nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "should transform draft post correctly",
			input: &domain.Post{
				UserID:      "user-123",
				Title:       "Test Title",
				Description: "Test Description",
				Content:     "Test Content",
				Status:      domain.PostStatusDraft,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			expected: &dtos.Post{
				Title:       "Test Title",
				Description: "Test Description",
				Content:     "Test Content",
				Status:      "draft",
			},
		},
		{
			name: "should transform published post correctly",
			input: &domain.Post{
				UserID:      "user-456",
				Title:       "Published Post Title",
				Description: "Published Post Description",
				Content:     "Published Post Content",
				Status:      domain.PostStatusPublished,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			expected: &dtos.Post{
				Title:       "Published Post Title",
				Description: "Published Post Description",
				Content:     "Published Post Content",
				Status:      "published",
			},
		},
		{
			name: "should transform archived post correctly",
			input: &domain.Post{
				UserID:      "user-789",
				Title:       "Archived Post Title",
				Description: "Archived Post Description",
				Content:     "Archived Post Content",
				Status:      domain.PostStatusArchived,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			expected: &dtos.Post{
				Title:       "Archived Post Title",
				Description: "Archived Post Description",
				Content:     "Archived Post Content",
				Status:      "archived",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromPostCoreToHTTP(tt.input)

			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, tt.expected.Title, result.Title)
				assert.Equal(t, tt.expected.Description, result.Description)
				assert.Equal(t, tt.expected.Content, result.Content)
				assert.Equal(t, tt.expected.Status, result.Status)
			}
		})
	}
}

func TestFromPostCoreToHTTP_FieldMapping(t *testing.T) {
	t.Run("should map all required fields correctly", func(t *testing.T) {
		input := &domain.Post{
			UserID:      "test-user-id",
			Title:       "Sample Title",
			Description: "Sample Description",
			Content:     "Sample Content",
			Status:      domain.PostStatusPublished,
			CreatedAt:   time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 1, 2, 12, 0, 0, 0, time.UTC),
		}

		result := FromPostCoreToHTTP(input)

		assert.NotNil(t, result)
		assert.Equal(t, "Sample Title", result.Title)
		assert.Equal(t, "Sample Description", result.Description)
		assert.Equal(t, "Sample Content", result.Content)
		assert.Equal(t, "published", result.Status)
	})
}

func TestFromPostCoreToHTTP_StatusConversion(t *testing.T) {
	statusTests := []struct {
		name        string
		postStatus  domain.PostStatus
		expectedStr string
	}{
		{"draft status", domain.PostStatusDraft, "draft"},
		{"published status", domain.PostStatusPublished, "published"},
		{"archived status", domain.PostStatusArchived, "archived"},
	}

	for _, st := range statusTests {
		t.Run(st.name, func(t *testing.T) {
			input := &domain.Post{
				UserID:      "user-id",
				Title:       "Title",
				Description: "Description",
				Content:     "Content",
				Status:      st.postStatus,
			}

			result := FromPostCoreToHTTP(input)

			assert.NotNil(t, result)
			assert.Equal(t, st.expectedStr, result.Status)
		})
	}
}
