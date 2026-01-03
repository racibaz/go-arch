package module

import (
	"errors"
	"fmt"
	"log"
	"time"

	postDomain "github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/mappers"
	"github.com/racibaz/go-arch/pkg/database"
	"github.com/racibaz/go-arch/pkg/es"
	"github.com/racibaz/go-arch/pkg/uuid"
)

// Seed seeds the database with initial data for the post module.
func Seed() error {
	// Implement module-specific seeding logic here
	log.Println("Post Module Seeder Start ..")

	// Get database connection
	db := database.Connection()

	if db == nil {
		return errors.New("database connection is nil")
	}

	posts := []*postDomain.Post{
		{
			Aggregate: es.NewAggregate(
				"2d86263a-eebf-4e7d-867a-0115569d6a3a",
				postDomain.PostAggregate,
			),
			UserID:      uuid.NewID(),
			Title:       "test title title title",
			Description: "test description description",
			Content:     "test content content content",
			Status:      postDomain.PostStatusPublished,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Aggregate:   es.NewAggregate(uuid.NewID(), postDomain.PostAggregate),
			UserID:      uuid.NewID(),
			Title:       "test title title title",
			Description: "test description description",
			Content:     "test content content content",
			Status:      postDomain.PostStatusPublished,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Aggregate:   es.NewAggregate(uuid.NewID(), postDomain.PostAggregate),
			UserID:      uuid.NewID(),
			Title:       "test title title title",
			Description: "test description description",
			Content:     "test content content content",
			Status:      postDomain.PostStatusPublished,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Aggregate:   es.NewAggregate(uuid.NewID(), postDomain.PostAggregate),
			UserID:      uuid.NewID(),
			Title:       "test title title title",
			Description: "test description description",
			Content:     "test content content content",
			Status:      postDomain.PostStatusPublished,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	for _, post := range posts {
		p, err := postDomain.Create(
			post.ID(),
			post.UserID,
			post.Title,
			post.Description,
			post.Content,
			post.Status,
			post.CreatedAt,
			post.UpdatedAt,
		)
		if err != nil {
			return err
		}

		if p == nil {
			return errors.New("created post is nil")
		}

		postEntity, err := mappers.ToPersistence(post)
		if err != nil {
			return errors.New("error mapping post to persistence: " + err.Error())
		}

		log.Println("Post Entity:", postEntity.ID)

		if postEntity.ID == "" {
			return errors.New("post entity ID is empty")
		}
		db.Create(postEntity)

		log.Println(fmt.Sprintf("Seeded Post ID: %s", postEntity.ID))
	}

	log.Println("Post Module Seeder Finish..")

	return nil
}
