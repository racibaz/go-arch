package rabbitmq

import (
	"github.com/racibaz/go-arch/pkg/config"
	"log"
)

func Connect() {
	conf := config.Get()

	conn, err := NewRabbitMQ(conf.RabbitMQConnectionString())
	if err != nil {
		log.Panicf("failed to connect to RabbitMQ: %v", err)
	}

	log.Println("Connected to RabbitMQ")
	Conn = conn
}
