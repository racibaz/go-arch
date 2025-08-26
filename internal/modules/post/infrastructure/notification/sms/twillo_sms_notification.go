package sms

import (
	"context"
	"fmt"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilloSmsNotificationRepository struct {
	client *twilio.RestClient
}

var _ ports.NotificationRepository = (*TwilloSmsNotificationRepository)(nil)

func NewTwilloSmsNotificationRepository() TwilloSmsNotificationRepository {

	client := twilio.NewRestClient()

	return TwilloSmsNotificationRepository{client: client}
}

func (t TwilloSmsNotificationRepository) NotifyPostCreated(ctx context.Context, postID, UserID string) error {

	params := &openapi.CreateMessageParams{}
	/*	params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
		params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
		params.SetBody("Hello from Golang!")

		_, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("SMS sent successfully!")
		}*/

	_, err := t.client.Api.CreateMessage(params)
	if err != nil {
		return err.Error()
	} else {
		fmt.Println("SMS sent successfully!")
		return nil
	}
}

func (t TwilloSmsNotificationRepository) NotifyPostCanceled(ctx context.Context, postID, UserID string) error {
	//TODO implement me
	panic("implement me")
}

func (t TwilloSmsNotificationRepository) NotifyPostReady(ctx context.Context, postID, UserID string) error {
	//TODO implement me
	panic("implement me")
}
