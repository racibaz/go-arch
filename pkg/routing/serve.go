package routing

import (
	"fmt"
	"github.com/racibaz/go-arch/pkg/config"
)

func Serve() {
	r := GetRouter()

	configs := config.Get()

	go func() {
		err := r.Run(fmt.Sprintf("%s:%s", configs.Server.HttpHost, configs.Server.HttpPort))
		if err != nil {
			panic(fmt.Sprintf("Failed to start HTTP server: %v", err))
		}
	}()
}
