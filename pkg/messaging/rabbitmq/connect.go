package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func Connect() {

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
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
