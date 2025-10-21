package rabbitmq

import (
	"fmt"
	"github.com/racibaz/go-arch/pkg/config"
)

func Connect() *RabbitMQ {
	conf := config.Get()

	conn, err := NewRabbitMQ(conf.RabbitMQ.Url)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to RabbitMQ : %v", err))
	}

	defer conn.Close()

	fmt.Println("Connected to RabbitMQ")

	return conn
}
