package providers

import (
	postModule "github.com/racibaz/go-arch/internal/modules/post"
	"github.com/racibaz/go-arch/internal/modules/shared"
	userModule "github.com/racibaz/go-arch/internal/modules/user"
)

// RegisterSeeders registers seeders for different modules
func RegisterSeeders() {
	// Register shared module seeder last
	sharedErr := shared.Seed()
	if sharedErr != nil {
		panic(sharedErr)
	}

	// You can add seeders of your modules here
	postErr := postModule.Seed()
	if postErr != nil {
		panic(postErr)
	}

	// User module seeder
	userErr := userModule.Seed()
	if userErr != nil {
		panic(userErr)
	}

	// Add more module seeders as needed
}
