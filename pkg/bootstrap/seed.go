package bootstrap

import (
	"github.com/racibaz/go-arch/internal/database/seeder"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/database"
)

func Seed() {
	config.Set("./config", "./.env")

	database.Connect()

	seeder.Seed()
}
