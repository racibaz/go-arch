package sms

import (
	"context"
	"fmt"

	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/notification/sms"
	"github.com/twilio/twilio-go"
)

type TwilioSmsNotificationAdapter struct {
	client *twilio.RestClient
}

var _ ports.NotificationAdapter = (*TwilioSmsNotificationAdapter)(nil)

func NewTwilioSmsNotificationAdapter() TwilioSmsNotificationAdapter {
	client := twilio.NewRestClient()

	return TwilioSmsNotificationAdapter{client: client}
}

func (t TwilioSmsNotificationAdapter) NotifyPostCreated(
	ctx context.Context,
	postID, UserID string,
) error {
	params := sms.TwilioInit()
	params.SetBody("Hello from Golang!")
	_, err := t.client.Api.CreateMessage(params)
	if err != nil {
		return err
	} else {
		fmt.Println("SMS sent successfully!")
		return nil
	}
}

func (t TwilioSmsNotificationAdapter) NotifyPostCanceled(
	ctx context.Context,
	postID, UserID string,
) error {
	// TODO implement me
	panic("implement me")
}

func (t TwilioSmsNotificationAdapter) NotifyPostReady(
	ctx context.Context,
	postID, UserID string,
) error {
	// TODO implement me
	panic("implement me")
}
