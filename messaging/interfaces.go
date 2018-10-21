package messaging

import "net/http"

type Message interface {
	GetSender() string
	GetBody() string
	GetRecipient() string
	GetUrl() string
}

type Authentication interface {
	GetSid() string
	GetToken() string
	GetUrl() string
}

type HttpClient interface {
	Do(r *http.Request) (*http.Response, error)
}
