package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func Connection() *amqp.Connection {
	return RabbitMQ
}
