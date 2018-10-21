package messaging

import (
	"net/http"
	"testing"
)

type HttpClientMock struct {}
type MessageMock struct{}
type AuthenticationMock struct {}

func (m AuthenticationMock) GetSid() string {
	panic("implement me")
}

func (m AuthenticationMock) GetToken() string {
	panic("implement me")
}

func (m AuthenticationMock) GetUrl() string {
	panic("implement me")
}

func (m HttpClientMock) Do(r *http.Request) (*http.Response, error) {
	panic("implement me")
}

func (m MessageMock) GetSender() string {
	panic("implement me")
}

func (m MessageMock) GetBody() string {
	panic("implement me")
}

func (m MessageMock) GetRecipients() []string {
	panic("implement me")
}

func (m MessageMock) GetUrl() string {
	panic("implement me")
}

func TestSendTextMessage(t *testing.T) {
	// When there is in one of the components, it should return an error

	//
}
