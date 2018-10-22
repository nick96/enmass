package messaging

//go:generate mockgen -destination=../mocks/messaging_mocks.go -package mocks github.com/nick96/enmass/messaging Client

// Client represents the contract that must be implemented for a class
// to qualify as a client.
type Client interface {
	SendMessage(from string, to string, msg string) error
}
