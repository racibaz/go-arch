package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"github.com/racibaz/go-arch/pkg/logger"
	"time"
)

type MessagePublisher struct {
	conf   *RabbitMQConection
	logger logger.Logger
}

func NewMessagePublisher(url string, logger logger.Logger) *MessagePublisher {
	rabbitMQConf := NewRabbitMQConnection(url)
	return &MessagePublisher{
		conf:   rabbitMQConf,
		logger: logger,
	}
}

func (mp *MessagePublisher) DeclareQueue(queueName string) error {
	channel := mp.conf.Channel()
	if channel == nil {
		return fmt.Errorf("message channel is nil, please retry")
	}

	_, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

func (mp *MessagePublisher) PublishEvent(queueName string, body any) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if queueName == "" {
		queueName = mp.conf.queue
	}

	channel := mp.conf.Channel()
	if channel == nil {
		panic("Messaging channel is nil, retry !")
	}
	if channel.IsClosed() {
		panic("could not publish event, channel closed")
	}

	mp.logger.Info(fmt.Sprintf("created new channel....%v", &channel))

	err = channel.PublishWithContext(ctx,
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp091.Persistent,
		},
	)

	if err != nil {
		return err
	}

	mp.logger.Info(fmt.Sprintf("Event published: %v", body))
	channel.Close()
	mp.logger.Info(fmt.Sprintf("channel closed: %v", &channel))

	return nil

}
