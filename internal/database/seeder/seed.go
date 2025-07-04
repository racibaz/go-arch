package seeder

import (
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain"
	postFactory "github.com/racibaz/go-arch/internal/modules/post/domain/factories"
	"github.com/racibaz/go-arch/pkg/database"
	"github.com/racibaz/go-arch/pkg/uuid"
	"log"
	"time"
)

func Seed() {
	db := database.Connection()

	log.Println("Seeding database...")

	post, err := postFactory.New(
		uuid.NewID(),
		"test title 1",
		"test description 1",
		"test content content 3",
		postValueObject.PostStatusDraft,
		time.Now(),
		time.Now())

	if err != nil {
		log.Fatalf("Error creating post: %v", err)
	}

	db.Create(&post) // pass pointer of data to Create

	log.Println("Seeder done ..")
}
