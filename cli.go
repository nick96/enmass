package main

import (
	"log"
	"os"

	"github.com/nick96/enmass/peopleapi"
	"github.com/nick96/gapiauth"
	"google.golang.org/api/people/v1"
)

func main() {
	if len(os.Args) == 0 {
		Usage()
		os.Exit(1)
	}

	args := os.Args[1:]
	parsed := parseArgs(args)

	auth := gapiauth.NewAuth(parsed.CredentialsPath, parsed.TokenPath, people.ContactsScope)
	client, err := auth.GetClient()
	if err != nil {
		log.Fatalf("Unable to get people api client: %v", err)
	}

	srv, err := people.New(client)
	if err != nil {
		log.Fatalf("Unable to create peopleapi client: %v", err)
	}

	contactGroup := peopleapi.NewGoogleContactGroup(srv, parsed.GroupName)
	switch parsed.ActionType {
	case GET:
		DoGetAction(contactGroup, parsed.Action)
	case SEND:
		DoSendAction(contactGroup, parsed.Action, RestAuth{
			Sid:    os.Getenv("TWILIO_ACCOUNT_SID"),
			Token:  os.Getenv("TWILIO_AUTH_TOKEN"),
			Url:    os.Getenv("TWILIO_URL"),
			Sender: os.Getenv("SENDER_PHONE_NUMBER"),
		})
	case CHECK:
		CheckGroup(contactGroup)
	case STAT:
		GetGroupStats(contactGroup)
	default:
		log.Fatalf("%v is an unknown action type", parsed.ActionType)
	}
}
