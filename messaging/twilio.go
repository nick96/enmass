package messaging

import (
	"github.com/subosito/twilio"
)

// TwilioClient represents a client for interacting with the twilio API.
type TwilioClient struct {
	client *twilio.Twilio
}

// NewTwilioClient creates a new twilio client.
func NewTwilioClient(sid string, token string) *TwilioClient {
	client := twilio.NewTwilio(sid, token)
	return &TwilioClient{
		client: client,
	}
}

// SendMessage sends a message MSG from the FROM number to the TO
// number. If there is an error then it is returned.
func (c *TwilioClient) SendMessage(from string, to string, msg string) error {
	_, err := c.client.SimpleSendSMS(from, to, msg)
	return err
}
