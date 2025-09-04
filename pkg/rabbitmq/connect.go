package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/racibaz/go-arch/pkg/env"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Connect() {

	rabbitmqURI := env.GetString("RABBITMQ_URI", "")
	conn, err := amqp.Dial(rabbitmqURI)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	fmt.Println("Connecting to rabbitmq...")
	Conn = conn
}
