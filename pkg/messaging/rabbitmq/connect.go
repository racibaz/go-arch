package rabbitmq

import (
	"fmt"
	"github.com/racibaz/go-arch/pkg/config"
)

func Connect() {
	conf := config.Get()

	conn := NewRabbitMQConnection(conf.RabbitMQ.Url)

	defer conn.Close()

	fmt.Println("Connected to RabbitMQ")
}
