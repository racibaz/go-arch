package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/racibaz/go-arch/pkg/config"
	"log"
)

func Connect() {
	conf := config.Get()

	conn, err := amqp.Dial(conf.RabbitMQ.Url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	log.Println("Connected to RabbitMQ")
	RabbitMQ = conn
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
