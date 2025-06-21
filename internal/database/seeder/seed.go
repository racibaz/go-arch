package seeder

import (
	postDomain "github.com/racibaz/go-arch/internal/modules/post/domain"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain/value_objects"
	"github.com/racibaz/go-arch/pkg/database"
	"github.com/racibaz/go-arch/pkg/uuid"
	"log"
)

func Seed() {
	db := database.Connection()

	postId := uuid.NewUuid().ToString()
	post, _ := postDomain.NewPost(
		postId,
		"test title",
		"test description",
		"test content",
		postValueObject.PostStatusPublished)

	db.Create(&post) // pass pointer of data to Create

	log.Println("Seeder done ..")
}
