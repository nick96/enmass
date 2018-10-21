package main

import (
	"log"

	"fmt"
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
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}

	srv, err := people.New(client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}

	contactGroup := peopleapi.NewGoogleContactGroup(srv, parsed.GroupName)
	switch parsed.ActionType {
	case GET:
		err := DoGetAction(contactGroup, parsed.Action)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	case SEND:
		msg, err := readMessage()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		twilioSid := os.Getenv("TWILIO_ACCOUNT_SID")
		twilioToken := os.Getenv("TWILIO_AUTH_TOKEN")
		phone := os.Getenv("ENMASS_PHONE")
		email := os.Getenv("ENMASS_EMAIL")
		err = DoSendAction(contactGroup, parsed.Action,
			twilioSid, twilioToken, phone, email, msg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	case CHECK:
		err := CheckGroup(contactGroup)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	case STAT:
		err := GetGroupStats(contactGroup)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	default:
		log.Fatalf("%v is an unknown action type", parsed.ActionType)
		os.Exit(1)
	}
}
