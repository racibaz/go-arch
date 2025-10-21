package rabbitmq

import (
	"fmt"
	"github.com/racibaz/go-arch/pkg/config"
	"log"
)

func Connect() *RabbitMQ {
	conf := config.Get()

	conn, err := NewRabbitMQ(conf.RabbitMQ.Url)
	if err != nil {
		log.Panicf(fmt.Sprintf("failed to connect to RabbitMQ : %v", err))
	}

	fmt.Println("Connected to RabbitMQ")

	return conn
}
