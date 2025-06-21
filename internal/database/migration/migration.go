package migration

import (
	"fmt"
	postDomain "github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/pkg/database"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&postDomain.Post{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration done ..")
}
