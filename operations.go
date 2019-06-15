package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/UnwrittenFun/pluralise"
	"github.com/nick96/enmass/messaging"
	"github.com/nick96/enmass/peopleapi"
	"io"
)

// GetGroupStats summarises group info (i.e. id, number of members,
// phone numbers and emails) then summarises the contact info for each
// member.
func GetGroupStats(group peopleapi.ContactGroup) error {
	return getGroupStats(group, os.Stderr)
}

func getGroupStats(group peopleapi.ContactGroup, writer io.Writer) error {
	members := group.GetMembers()
	id := group.GetGroupID()
	memberCount := len(members)
	phoneNumberCount := len(group.GetPhoneNumbers())
	emailCount := len(group.GetEmails())

	fmt.Fprintf(writer, "Group: %s [ID: %s]\n%d members %d phone numbers %d emails\n",
		group.GetName(), id, memberCount, phoneNumberCount, emailCount)

	for _, member := range members {
		numbers := member.GetPhoneNumbers()
		emails := member.GetEmails()
		fmt.Fprintf(writer, "%s (%d %s, %d %s)\n",
			member.GetName(),
			len(numbers),
			pluralise.WithCount("phone numbers", len(numbers)),
			len(emails),
			pluralise.WithCount("email", len(emails)))
	}

	return nil
}

// CheckGroup checks that all members of a group have at least one
// email and phone phone number.
func CheckGroup(group peopleapi.ContactGroup) error {
	return checkGroup(group, os.Stderr)
}

func checkGroup(group peopleapi.ContactGroup, writer io.Writer) error {
	for _, member := range group.GetMembers() {
		emails := member.GetEmails()
		numbers := member.GetPhoneNumbers()
		emailCount := len(emails)
		numberCount := len(numbers)
		if emailCount == 0 {
			fmt.Fprintf(writer, "%s has no email\n", member.GetName())
		}

		if numberCount == 0 {
			fmt.Fprintf(writer, "%s has no phone numbers\n", member.GetName())
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
		panic("Not implemented")
	case TEXT:
		client := messaging.NewTwilioClient(twilioSid, twilioToken)
		sendTextMessage(group, client, phone, msg, os.Stderr)
	}

	return nil
}

// DoGetAction performs a get action, the specifics of which are
// determined by the value of ACTION.
func DoGetAction(group peopleapi.ContactGroup, action GroupAction) error {
	return doGetAction(group, action, os.Stderr, os.Stdout)
}

func doGetAction(group peopleapi.ContactGroup, action GroupAction, writer io.Writer,
	copiableWriter io.Writer) error {
	switch action {
	case EMAIL:
		var emails []string
		for _, email := range group.GetEmails() {
			emails = append(emails, string(email))
		}

		members := group.GetMembers()
		fmt.Fprintf(writer, "Found %d %s (%d %s)\n",
			len(emails), pluralise.WithCount("email", len(emails)),
			len(members),
			pluralise.WithCount("member", len(members)))
		fmt.Fprintln(copiableWriter, strings.Join(emails, "; "))
	case TEXT:
		phoneNumbers := group.GetPhoneNumbers()
		members := group.GetMembers()
		fmt.Fprintf(writer, "Found %d %s (%d %s)\n",
			len(phoneNumbers), pluralise.WithCount("phone number", len(phoneNumbers)),
			len(members),
			pluralise.WithCount("member", len(members)))

		for _, phoneNumber := range phoneNumbers {
			fmt.Fprintln(copiableWriter, phoneNumber)
		}
	default:
		log.Fatalf("%v is an unkown action", action)
	}

	return nil
}

// sendTextMessage sends a text message to every member of the group
func sendTextMessage(group peopleapi.ContactGroup, client messaging.Client, from string,
	msg string, writer io.Writer) error {
	for _, member := range group.GetMembers() {
		for _, to := range member.GetPhoneNumbers() {
			err := client.SendMessage(from, string(to), msg)
			if err != nil {
				fmt.Fprintf(writer, "Failed to send message to %s on %s: %v\n",
					member.GetName(), string(to), err)
			}
		}
	}
	return nil
}
