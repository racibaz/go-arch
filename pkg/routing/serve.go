package routing

import (
	"fmt"
	"github.com/racibaz/go-arch/pkg/config"
	"log"
)

func Serve() {
	r := GetRouter()

	configs := config.Get()

	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	if err != nil {
		log.Fatal("Error in routing")
		return
	}
}
