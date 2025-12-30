package module

import (
	postDomain "github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/mappers"
	"github.com/racibaz/go-arch/pkg/database"
	"github.com/racibaz/go-arch/pkg/uuid"
	"log"
	"time"
)

// Seed seeds the database with initial data for the post module.
func Seed() {
	// Get database connection
	db := database.Connection()

	// Implement module-specific seeding logic here
	log.Println("Seeding database...")

	post, err := postDomain.Create(
		uuid.NewID(),
		uuid.NewID(),
		"test title 1",
		"test description 1",
		"test content content 3",
		postDomain.PostStatusDraft,
		time.Now(),
		time.Now())

	postEntity := mappers.ToPersistence(*post)

	if err != nil {
		log.Fatalf("Error creating post: %v", err)
	}

	db.Create(&postEntity) // pass pointer of data to Create

	log.Println("Post Seeder Done ..")
}
