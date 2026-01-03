package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/messaging"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to create channel: %v", err)
	}
	log.Println("RabbitMq has been connected")

	rmq := &RabbitMQ{
		conn:    conn,
		Channel: ch,
	}

	if err := rmq.setupExchangesAndQueues(); err != nil {
		// Clean up if setup fails
		rmq.Close()
		return nil, fmt.Errorf("failed to setup exchanges and queues: %v", err)
	}

	return rmq, nil
}

func (r *RabbitMQ) Connect() *amqp.Connection {
	config := config.Get()

	conn, err := amqp.Dial(config.RabbitMQConnectionString())
	if err != nil {
		panic(fmt.Sprintf("failed to connect to RabbitMQ : %v", err))
	}

	log.Println("RabbitMq has been reconnected")
	return conn
}

func (r *RabbitMQ) DeclareQueue(queueName string) error {
	var err error
	channel, err := r.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to create channel: %v", err)
	}
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

func (r *RabbitMQ) Connection() *amqp.Connection {
	if r.conn == nil {
		r.conn = r.Connect()
	}
	return r.conn
}

func (r *RabbitMQ) setupExchangesAndQueues() error {
	// todo we need to check DLQ exchange and queue

	err := r.Channel.ExchangeDeclare(
		messaging.DefaultExchange, // name
		"topic",                   // type
		true,                      // durable
		false,                     // auto-deleted
		false,                     // internal
		false,                     // no-wait
		nil,                       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %s: %v", messaging.DefaultExchange, err)
	}

	if err := r.declareAndBindQueue(
		messaging.PostProcessingQueue,
		[]string{
			messaging.PostEventCreated,
			messaging.PostEventUpdated,
			messaging.PostEventDeleted,
			messaging.PostEventPublished,
		},
		messaging.DefaultExchange,
	); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) declareAndBindQueue(
	queueName string,
	messageTypes []string,
	exchange string,
) error {
	// Add dead letter configuration
	args := amqp.Table{
		"x-dead-letter-exchange": messaging.DeadLetterExchange,
	}

	q, err := r.Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		args,      // arguments with DLX config
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range messageTypes {
		if err := r.Channel.QueueBind(
			q.Name,   // queue name
			msg,      // routing key
			exchange, // exchange
			false,
			nil,
		); err != nil {
			return fmt.Errorf("failed to bind queue to %s: %v", queueName, err)
		}
	}

	return nil
}

func (r *RabbitMQ) GetChannel() *amqp.Channel {
	var channel *amqp.Channel
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

func (r *RabbitMQ) PublishMessage(
	ctx context.Context,
	routingKey string,
	message messaging.MessagePayload,
) error {
	log.Printf("Publishing message with routing key: %s", routingKey)

	jsonMsg, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         jsonMsg,
	}
	// todo add tracing here if needed
	return r.publish(ctx, messaging.DefaultExchange, routingKey, msg)
}

func (r *RabbitMQ) publish(
	ctx context.Context,
	exchange, routingKey string,
	msg amqp.Publishing,
) error {
	return r.Channel.PublishWithContext(ctx,
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		msg,
	)
}

func (r *RabbitMQ) Close() {
	err := r.conn.Close()
	if err != nil {
		return
	}
}
