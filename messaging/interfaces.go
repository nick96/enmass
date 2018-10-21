package messaging

import "net/http"

//go:generate mockgen -destination=../mocks/messaging_mocks.go -package mock github.com/nick96/enmass/messaging Message,Authentication,HttpClient

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
