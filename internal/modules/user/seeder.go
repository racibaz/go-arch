package module

import (
	"errors"
	"log"
	"time"

	"github.com/racibaz/go-arch/internal/modules/user/domain"
	"github.com/racibaz/go-arch/internal/modules/user/infrastructure/persistence/gorm/mappers"
	"github.com/racibaz/go-arch/pkg/database"
	"github.com/racibaz/go-arch/pkg/es"
)

// Seed seeds the database with initial data for the post module.
func Seed() error {
	// Implement module-specific seeding logic here
	log.Println("User Module Seeder Start ..")

	// Get database connection
	db := database.Connection()

	if db == nil {
		return errors.New("database connection is nil")
	}

	users := []*domain.User{
		{
			Aggregate: es.NewAggregate(
				"2d86263a-eebf-4e7d-867a-0115569d6a3a",
				domain.UserAggregate,
			),
			Email:     "guest@xyz.com",
			Name:      "jackynickname",
			Password:  "jackypassword",
			Status:    domain.StatusPublished,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Aggregate: es.NewAggregate(
				"13dd0ee4-67ed-4dfc-81ef-9cb6684446d0",
				domain.UserAggregate,
			),
			Email:     "raci@xyz.com",
			Name:      "racinickname",
			Password:  "racipassword",
			Status:    domain.StatusPublished,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, user := range users {
		p, err := domain.Create(
			user.ID(),
			user.Name,
			user.Email,
			user.Password,
			user.Status,
			user.CreatedAt,
			user.UpdatedAt,
		)
		if err != nil {
			return err
		}

		if p == nil {
			return errors.New("created user is nil")
		}

		userEntity, err := mappers.ToPersistence(user)
		if err != nil {
			return errors.New("error mapping post to persistence: " + err.Error())
		}

		log.Println("User Entity:", userEntity.ID)

		if userEntity.ID == "" {
			return errors.New("user entity ID is empty")
		}
		db.Create(userEntity)

		log.Printf("Seeded User ID: %s\n", userEntity.ID)
	}

	log.Println("User Module Seeder Finish..")

	return nil
}
