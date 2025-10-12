package database

import (
	"fmt"
	"github.com/racibaz/go-arch/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() {
	config := config.Get()

	db, err := gorm.Open(postgres.Open(config.DatabaseUrl()), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database", err)
		return
	}

	if db == nil {
		log.Fatal("db object is nil")
		return
	}

	fmt.Println("Connecting to database...")

	DB = db
}
