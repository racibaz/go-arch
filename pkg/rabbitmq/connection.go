package rabbitmq

import "github.com/rabbitmq/amqp091-go"

func Connection() *amqp091.Connection {
	return Conn
}
