package providers

import (
	postModule "github.com/racibaz/go-arch/internal/modules/post"
	"github.com/racibaz/go-arch/internal/modules/shared"
)

// RegisterSeeders registers seeders for different modules
func RegisterSeeders() {
	// Register shared module seeder last
	shared.Seed()

	// You can add seeders of your modules here
	postModule.Seed()

	// Add more module seeders as needed
}
