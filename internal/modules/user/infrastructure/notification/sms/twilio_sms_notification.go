package sms

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
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

func (t TwilioSmsNotificationAdapter) NotifyUserRegistered(ctx context.Context, UserID string) error {
	//TODO implement me
	panic("implement me")
}

func (t TwilioSmsNotificationAdapter) NotifyUserCanceled(ctx context.Context, UserID string) error {
	//TODO implement me
	panic("implement me")
}

func (t TwilioSmsNotificationAdapter) NotifyUserReady(ctx context.Context, UserID string) error {
	//TODO implement me
	panic("implement me")
}
