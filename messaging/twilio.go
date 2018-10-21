package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type TwilioMessage struct {
	Sender    string
	Body      string
	Recipient string
	Url       string
}

type TwilioAuth struct {
	Sid   string
	Token string
}

func (m TwilioAuth) GetSid() string {
	return m.Sid
}

func (m TwilioAuth) GetToken() string {
	return m.Token
}

func (m TwilioMessage) GetSender() string {
	return m.Sender
}

func (m TwilioMessage) GetBody() string {
	return m.Body
}

func (m TwilioMessage) GetRecipient() string {
	return m.Recipient
}

func (m TwilioMessage) GetUrl() string {
	return m.Url
}

func (m TwilioAuth) GetUrl() string {
	return fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json",
		m.Sid)
}

func SendText(client HttpClient, message Message, auth Authentication) error {
	msgData := url.Values{}
	msgData.Set("Body", message.GetBody())
	msgData.Set("From", message.GetSender())
	msgData.Set("To", message.GetRecipient())

	msgDataReader := *strings.NewReader(msgData.Encode())
	req, err := buildTwilioRequest(msgDataReader, auth, auth.GetUrl())
	if err != nil {
		return err
	}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	var data map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("Text message sent (sid %s) to %s\n", data["sid"], msgData.Get("To"))
		return nil
	} else {
		log.Printf("%s: Failed to send message to %s (see %s)", resp.Status, message.GetRecipient(), data["more_info"])
	}

	return err
}

func buildTwilioRequest(reader strings.Reader, authentication Authentication, url string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, &reader)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(authentication.GetSid(), authentication.GetToken())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}
