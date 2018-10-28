package messaging

import (
	"github.com/subosito/twilio"
)

// TwilioClient represents a client for interacting with the twilio API.
type TwilioClient struct {
	client *twilio.Client
}

// NewTwilioClient creates a new twilio client.
func NewTwilioClient(sid string, token string) *TwilioClient {
	client := twilio.NewClient(sid, token, nil)
	client.UserAgent = "enmass"
	return &TwilioClient{
		client: client,
	}
}

// SendMessage sends a message MSG from the FROM number to the TO
// number. If there is an error then it is returned.
func (c *TwilioClient) SendMessage(from string, to string, msg string) error {
	_, _, err := c.client.Messages.SendSMS(from, to, msg)
	return err
}
