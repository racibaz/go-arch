package database

import (
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() {
	config := config.Get()

	db, err := gorm.Open(postgres.Open(config.DatabaseConnectionString()), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database", err)
		return
	}

	if db == nil {
		log.Fatal("db object is nil")
		return
	}

	log.Println("Connecting to database...")

	// Integrate OpenTelemetry with GORM
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	DB = db
}
