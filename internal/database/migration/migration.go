package migration

import (
	"fmt"
	postDomain "github.com/racibaz/go-arch/internal/modules/post/domain"
	sharedDomain "github.com/racibaz/go-arch/internal/modules/shared/domain"
	"github.com/racibaz/go-arch/pkg/database"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&postDomain.Post{}, &sharedDomain.Event{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration done ..")
}
