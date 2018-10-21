package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nick96/enmass/messaging"
	"github.com/nick96/enmass/peopleapi"
)

// GetGroupStats summarises group info (i.e. id, number of members,
// phone numbers and emails) then summarises the contact info for each
// member.
func GetGroupStats(group peopleapi.ContactGroup) error {
	members := group.GetMembers()
	id := group.GetGroupID()
	memberCount := len(members)
	phoneNumberCount := len(group.GetPhoneNumbers())
	emailCount := len(group.GetEmails())

	fmt.Printf("Group: %s\nID: %s\n%d members\n%d phone numbers\n%d emails\n",
		group.GetName(), id, memberCount, phoneNumberCount, emailCount)

	for _, member := range members {
		fmt.Printf("%s (%d phone numbers, %d emails)\n", member.GetName(),
			len(member.GetPhoneNumbers()),
			len(member.GetEmails()))
	}

	return nil
}

// CheckGroup checks that all members of a group have at least one
// email and phone phone number.
func CheckGroup(group peopleapi.ContactGroup) error {
	for _, member := range group.GetMembers() {
		if len(member.GetEmails()) == 0 {
			fmt.Fprintf(os.Stderr, "%s has not email\n", member.GetName())
		}

		if len(member.GetPhoneNumbers()) == 0 {
			fmt.Fprintf(os.Stderr, "%s has no phone numbers\n", member.GetName())
		}
	}

	return nil
}

// DoSendAction performs a send action, the specifics of which are
// determined by the value of ACTION.
func DoSendAction(group peopleapi.ContactGroup, action GroupAction,
	twilioSid string, twilioToken string, phone string, email string,
	msg string) error {
	switch action {
	case EMAIL:
		sendEmail(group, msg, email)
	case TEXT:
		sendTextMessage(group, msg, phone, twilioSid, twilioToken)
	}

	return nil
}

// DoGetAction performs a get action, the specifics of which are
// determined by the value of ACTION.
func DoGetAction(group peopleapi.ContactGroup, action GroupAction) error {
	switch action {
	case EMAIL:
		var emails []string
		for _, email := range group.GetEmails() {
			emails = append(emails, string(email))
		}

		fmt.Fprintf(os.Stderr, "Found %d emails (%d members)\n",
			len(emails), len(group.GetMembers()))
		fmt.Println(strings.Join(emails, "; "))
	case TEXT:
		phoneNumbers := group.GetPhoneNumbers()
		fmt.Fprintf(os.Stderr, "Found %d phone numbers (%d members)\n",
			len(phoneNumbers), len(group.GetMembers()))

		for _, phoneNumber := range phoneNumbers {
			fmt.Println(phoneNumber)
		}
	default:
		log.Fatalf("%v is an unkown action", action)
	}

	return nil
}

// sendTextMessage sends a text message to every member of the group
func sendTextMessage(group peopleapi.ContactGroup, msg string, sender string, sid string,
	token string)  error {
	msgData := messaging.TwilioMessage{
		Body:   msg,
		Sender: sender,
	}

	auth := messaging.TwilioAuth{
		Sid:   sid,
		Token: token,
	}

	for _, member := range group.GetMembers() {
		for _, phoneNumber := range member.GetPhoneNumbers() {
			msgData.Recipient = string(phoneNumber)
			err := messaging.SendText(&http.Client{}, msgData, auth)
			if err != nil {
				log.Printf("Could not send text message to %s on %s: %v",
					member.GetName(), phoneNumber, err)
			}
		}
	}

	return nil
}

// sendEmail sends an email to every member of the group
func sendEmail(group peopleapi.ContactGroup, msg string, email string) error {
	panic("Not implemented")
}
