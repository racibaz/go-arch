package messaging

type MessagePayload struct {
	OwnerID string `json:"ownerId"`
	Data    []byte `json:"data"`
}

type MessagePublisher interface {
	DeclareQueue(queueName string) error
	PublishEvent(queueName string, body interface{}) error
}
