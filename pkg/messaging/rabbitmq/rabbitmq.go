package rabbitmq

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"github.com/racibaz/go-arch/pkg/config"
	"log"
)

type RabbitMQConection struct {
	conn  *amqp091.Connection
	queue string
}

func NewRabbitMQConnection(url string) *RabbitMQConection {
	config := config.Get()
	conn, err := amqp091.Dial(url)

	if err != nil {
		panic(fmt.Sprintf("failed to connect to RabbitMQ : %v", err))
	}

	log.Println("RabbitMq has been connected")

	return &RabbitMQConection{
		conn:  conn,
		queue: config.RabbitMQ.DefaultQueue,
	}
}

func (r *RabbitMQConection) Connect() *amqp091.Connection {
	config := config.Get()

	conn, err := amqp091.Dial(config.RabbitMQ.Url)

	if err != nil {
		panic(fmt.Sprintf("failed to connect to RabbitMQ : %v", err))
	}

	log.Println("RabbitMq has been reconnected")
	return conn
}

func (r *RabbitMQConection) DeclareQueue(queueName string) error {
	var err error
	channel, err := r.conn.Channel()
	defer channel.Close()

	_, err = channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

func (r *RabbitMQConection) Connection() *amqp091.Connection {
	if r.conn == nil {
		r.conn = r.Connect()
	}
	return r.conn
}

func (r *RabbitMQConection) Channel() *amqp091.Channel {
	var channel *amqp091.Channel
	connection := r.conn
	if connection == nil {
		connection = r.Connect()
	}
	channel, err := r.conn.Channel()
	if err != nil {
		channel, _ = r.conn.Channel()
		log.Println("channel is nil")
	}
	if channel != nil && channel.IsClosed() {
		channel, err = r.conn.Channel()
		if err != nil {
			log.Println("channel was closed and error creating channel")
		}
	}
	return channel
}

func (r *RabbitMQConection) Queue() string {
	return r.queue
}

func (r *RabbitMQConection) Close() {
	r.conn.Close()
}
