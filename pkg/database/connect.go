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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DB.Host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database", err)
		return
	}

	fmt.Println("Connecting to database...")

	DB = db
}
