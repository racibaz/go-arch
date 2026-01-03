package sms

import (
	"os"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func TwilioInit() *openapi.CreateMessageParams {
	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	return params
}
