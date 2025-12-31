package bootstrap

import (
	"github.com/racibaz/go-arch/internal/providers"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/database"
	"log"
)

func Seed() {
	log.Println("Seeder is starting")

	config.Set("./config", "./.env")

	database.Connect()

	providers.RegisterSeeders()

	log.Println("Seeder has finished")
}
