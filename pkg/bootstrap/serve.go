package bootstrap

import (
	"context"
	"log"

	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/database"
	"github.com/racibaz/go-arch/pkg/grpc"
	helper "github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/messaging/rabbitmq"
	"github.com/racibaz/go-arch/pkg/routing"
	"github.com/racibaz/go-arch/pkg/trace"
)

func Serve() {
	config.Set("./config", "./.env")

	// Initialize Tracer
	tracerProvider, err := trace.InitTracer()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			panic(err)
		}
	}()

	log.Println("Tracer initialized")

	// Initialize JWT
	helper.InitJWT("" + config.Get().App.JWTKey)

	database.Connect()

	rabbitmq.Connect()

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()

	grpc.Serve()
}
