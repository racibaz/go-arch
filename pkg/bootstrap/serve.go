package bootstrap

import (
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/routing"
)

func Serve() {
	config.Set()

	//database.Connect()

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()
}
