package main

import (
	"bufio"
	"fmt"
	"github.com/nick96/enmass/messaging"
	"github.com/nick96/enmass/peopleapi"
	"log"
	"net/http"
	"os"
	"strings"
)

type RestAuth struct {
	Sid    string
	Token  string
	Url    string
	Sender string
}

func GetGroupStats(group peopleapi.ContactGroup) {
	members := group.GetMembers()
	id := group.GetGroupId()
	memberCount := len(members)
	phoneNumberCount := len(group.GetPhoneNumbers())
	emailCount := len(group.GetEmails())

	fmt.Printf("Group: %s\nID: %s\n%d members\n%d phone numbers\n%d emails\n",
		group.GetName(), id, memberCount, phoneNumberCount, emailCount)

	for _, member := range members {
		fmt.Printf("%s (%d phone numbers, %d emails)\n", member.GetName(), len(member.GetPhoneNumbers()),
			len(member.GetEmails()))
	}
}

func CheckGroup(group peopleapi.ContactGroup) {
	for _, member := range group.GetMembers() {
		if len(member.GetEmails()) == 0 {
			fmt.Fprintf(os.Stderr, "%s has not email\n", member.GetName())
		}

		if len(member.GetPhoneNumbers()) == 0 {
			fmt.Fprintf(os.Stderr, "%s has no phone numbers\n", member.GetName())
		}
	}
}

// Do a send action on the given group
func DoSendAction(group peopleapi.ContactGroup, action GroupAction, restAuth RestAuth) {
	msg, err := readMessage()
	if err != nil {
		log.Fatalf("Could not read message: %v", err)
	}

	switch action {
	case EMAIL:
		sendEmail(group, msg)
	case TEXT:
		reader := bufio.NewReader(os.Stdin)
		fmt.Fprintf(os.Stderr, "Are you sure you want to send a text message to %s (%d members) from %s? [y/N] ",
			group.GetName(), len(group.GetMembers()), restAuth.Sender)
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("There was a problem reading the answer: %v", err)
		}

		if strings.ToLower(strings.TrimSpace(answer)) == "y" {
			fmt.Fprintf(os.Stderr, "Do you wish to add a signature to the end of the message? [y/N] ")
			answer, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("There was a problem reading the answer: %v", err)
			}

			if strings.ToLower(strings.TrimSpace(answer)) == "y" {
				signature := os.Getenv("TEXT_MSG_SIGNATURE")
				if len(signature) == 0 {
					fmt.Fprintln(os.Stderr, "No signature was found")
					os.Exit(1)
				} else {
					msg = strings.TrimRight(msg, "\n")
					msg = fmt.Sprintf("%s\n%s", msg, signature)

					fmt.Fprintln(os.Stderr, "Message to be sent:")
					fmt.Fprintln(os.Stderr, msg)
					fmt.Fprint(os.Stderr, "Is this okay? [y/N] ")
					answer, err := reader.ReadString('\n')
					if err != nil {
						log.Fatalf("There was a problem reading the answer: %v", err)
					}
					if strings.TrimSpace(strings.ToLower(answer)) != "y" {
						os.Exit(1)
					}
				}
			}

			sendTextMessage(group, msg, restAuth.Sender, restAuth.Sid, restAuth.Token, restAuth.Url)
		} else {
			log.Println("Text message has been cancelled")
		}
	}
}

// Do a get action on the given group
func DoGetAction(group peopleapi.ContactGroup, action GroupAction) {
	switch action {
	case EMAIL:
		var emails []string
		for _, email:= range group.GetEmails() {
			emails = append(emails, email.Value)
		}
		fmt.Fprintf(os.Stderr, "Found %d emails (%d members)\n", len(emails), len(group.GetMembers()))
		fmt.Println(strings.Join(emails, "; "))
	case TEXT:
		phoneNumbers := group.GetPhoneNumbers()
		fmt.Fprintf(os.Stderr, "Found %d phone numbers (%d members)\n", len(phoneNumbers), len(group.GetMembers()))

		for _, phoneNumber := range phoneNumbers {
			fmt.Println(phoneNumber.CanonicalForm)
		}
	default:
		log.Fatalf("%v is an unkown action", action)
	}

}

// Send a text message to every member of the group
func sendTextMessage(group peopleapi.ContactGroup, msg string, sender string, sid string, token string, twilioUrl string) {
	msgData := messaging.TwilioMessage{
		Body: msg,
		Sender: sender,
	}

	auth := messaging.TwilioAuth{
		Sid: sid,
		Token: token,
	}


	for _, member := range group.GetMembers() {
		for _, phoneNumber := range member.GetPhoneNumbers() {
			msgData.Recipient = phoneNumber.CanonicalForm
			err := messaging.SendText(&http.Client{}, msgData, auth)
			if err != nil {
				log.Printf("Could not send text message to %s on %s: %v", member.GetName(), phoneNumber, err)
			}
		}
	}
}

// Send an email to every member of the group
func sendEmail(group peopleapi.ContactGroup, msg string) {
	panic("Not implemented")
}
