package sms

import (
	"context"
	"fmt"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

type TwilioSmsNotificationAdapter struct {
	client *twilio.RestClient
}

var _ ports.NotificationAdapter = (*TwilioSmsNotificationAdapter)(nil)

func NewTwilioSmsNotificationAdapter() TwilioSmsNotificationAdapter {

	client := twilio.NewRestClient()

	return TwilioSmsNotificationAdapter{client: client}
}

func (t TwilioSmsNotificationAdapter) NotifyPostCreated(ctx context.Context, postID, UserID string) error {

	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody("Hello from Golang!")
	_, err := t.client.Api.CreateMessage(params)
	if err != nil {
		return err.Error()
	} else {
		fmt.Println("SMS sent successfully!")
		return nil
	}
}

func (t TwilioSmsNotificationAdapter) NotifyPostCanceled(ctx context.Context, postID, UserID string) error {
	//TODO implement me
	panic("implement me")
}

func (t TwilioSmsNotificationAdapter) NotifyPostReady(ctx context.Context, postID, UserID string) error {
	//TODO implement me
	panic("implement me")
}
