package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/racibaz/go-arch/internal/modules/user/domain"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/messaging"
	"github.com/racibaz/go-arch/pkg/messaging/rabbitmq"
)

type UserMessagePublisher struct {
	rabbitmq *rabbitmq.RabbitMQ
	logger   logger.Logger
}

func NewUserMessagePublisher(
	rabbitmq *rabbitmq.RabbitMQ,
	logger logger.Logger,
) *UserMessagePublisher {
	return &UserMessagePublisher{
		rabbitmq: rabbitmq,
		logger:   logger,
	}
}

func (p *UserMessagePublisher) PublishUserRegistered(
	ctx context.Context,
	payload *domain.User,
) error {
	/*
		todo : change to user event data and it can be change as grpc proto model
	*/

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return p.rabbitmq.PublishMessage(ctx, messaging.UserEventCreated, messaging.MessagePayload{
		OwnerID: payload.ID(), // todo it can get auth user id
		Data:    payloadJSON,
	})
}

func (p *UserMessagePublisher) PublishUserLoggedIn(
	ctx context.Context,
	payload *domain.User,
) error {
	/*
		todo : change to user event data and it can be change as grpc proto model
	*/

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return p.rabbitmq.PublishMessage(ctx, messaging.UserEventLoggedIn, messaging.MessagePayload{
		OwnerID: payload.ID(), // todo it can get auth user id
		Data:    payloadJSON,
	})
}
