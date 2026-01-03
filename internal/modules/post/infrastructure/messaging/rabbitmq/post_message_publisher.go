package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/messaging"
	"github.com/racibaz/go-arch/pkg/messaging/rabbitmq"
)

type PostMessagePublisher struct {
	rabbitmq *rabbitmq.RabbitMQ
	logger   logger.Logger
}

func NewPostMessagePublisher(
	rabbitmq *rabbitmq.RabbitMQ,
	logger logger.Logger,
) *PostMessagePublisher {
	return &PostMessagePublisher{
		rabbitmq: rabbitmq,
		logger:   logger,
	}
}

func (p *PostMessagePublisher) PublishPostCreated(ctx context.Context, payload *domain.Post) error {
	/*
		todo : change to post event data and it can be change as grpc proto model
	*/

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return p.rabbitmq.PublishMessage(ctx, messaging.PostEventCreated, messaging.MessagePayload{
		OwnerID: payload.UserID, // todo it can get auth user id
		Data:    payloadJSON,
	})
}
