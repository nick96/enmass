package main

import (
	"bufio"
	"fmt"
	"google.golang.org/api/people/v1"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 0 {
		Usage()
		os.Exit(1)
	}

	args := os.Args[1:]
	parsed := parseArgs(args)

	config, err := readCredentials(parsed.CredentialsPath)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(config, parsed.TokenPath)

	srv, err := people.New(client)
	if err != nil {
		log.Fatalf("Unable to create people client: %v", err)
	}

	contactGroup := NewContactGroup(srv, parsed.GroupName)
	switch parsed.ActionType {
	case GET:
		doGetAction(contactGroup, parsed.Action)
	case SEND:
		doSendAction(contactGroup, parsed.Action)
	case CHECK:
		checkGroup(contactGroup)
	case STAT:
		getGroupStats(contactGroup)
	default:
		log.Fatalf("%v is an unknown action type", parsed.ActionType)
	}
}
func getGroupStats(group ContactGroup) {
	members := group.GetMembers()
	id := group.GetGroupId()
	memberCount := len(members)
	phoneNumberCount := countPhoneNumbers(members)
	emailCount := countEmails(members)

	fmt.Printf("Group: %s\nID: %s\n%d members\n%d phone numbers\n%d emails\n",
		group.Name, id, memberCount, phoneNumberCount, emailCount)

	for _, member := range  members {
		fmt.Printf("%s (%d phone numbers, %d emails)\n", member.Name, len(member.GetPhoneNumbers()), len(member.GetEmails()))
	}
}

func countEmails(members []*Person) int {
	sum := 0
	for _, member := range members {
		sum += len(member.GetEmails())
	}
	return sum
}

func countPhoneNumbers(members []*Person) int {
	sum := 0
	for _, member := range members {
		sum += len(member.GetPhoneNumbers())
	}
	return sum
}

func checkGroup(group ContactGroup) {
	for _, member := range group.GetMembers() {
		if len(member.GetEmails()) == 0 {
			fmt.Fprintf(os.Stderr, "%s has not email\n", member.Name)
		}

		if len(member.GetPhoneNumbers()) == 0 {
			fmt.Fprintf(os.Stderr, "%s has no phone numbers\n", member.Name)
		}
	}
}


// Do a send action on the given group
func doSendAction(group ContactGroup, action GroupAction) {
	msg, err := readMessage()
	if err != nil {
		log.Fatalf("Could not read message: %v", err)
	}

	switch action {
	case EMAIL:
		sendEmail(group, msg)
	case TEXT:
		twilioSid := os.Getenv("TWILIO_ACCOUNT_SID")
		twilioToken := os.Getenv("TWILIO_AUTH_TOKEN")
		twilioUrl := os.Getenv("TWILIO_URL")
		sender := os.Getenv("SENDER_PHONE_NUMBER")

		reader := bufio.NewReader(os.Stdin)
		fmt.Fprintf(os.Stderr, "Are you sure you want to send a text message to %s (%d members) from %s? [y/N] ",
			group.Name, len(group.GetMembers()), sender)
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("There was a problem reading the answer: %v\n", err)
		}


		if strings.ToLower(strings.TrimSpace(answer)) == "y" {
			sendTextMessage(group, msg, sender, twilioSid, twilioToken, twilioUrl)
		} else {
			log.Println("Text message has been cancelled")
		}
	}
}

// Do a get action on the given group
func doGetAction(group ContactGroup, action GroupAction) {
	var infos []string
	switch action {
	case EMAIL:
		infos = getEmails(group)
		fmt.Fprintf(os.Stderr, "Found %d emails (%d members)\n", len(infos), len(group.GetMembers()))
	case TEXT:
		infos = getPhoneNumbers(group)
		fmt.Fprintf(os.Stderr, "Found %d phone numbers (%d members)\n", len(infos), len(group.GetMembers()))
	default:
		log.Fatalf("%v is an unkown action", action)
	}

	for _, info := range infos {
		fmt.Println(info)
	}
}

// Get all the email addresses for the group
func getEmails(group ContactGroup) []string {
	var emails []string
	for _, member := range group.getMembers() {
		for _, email := range member.GetEmails() {
			emails = append(emails, strings.TrimSpace(email.Value))
		}
	}
	return emails
}

// Get all the phone numbers for the group
func getPhoneNumbers(group ContactGroup) []string {
	var phoneNumbers []string
	for _, member := range group.getMembers() {
		for _, phoneNumber := range member.GetPhoneNumbers() {
			phoneNumbers = append(phoneNumbers, strings.TrimSpace(phoneNumber.CanonicalForm))
		}
	}
	return phoneNumbers
}

// Send a text message to every member of the group
func sendTextMessage(group ContactGroup, msg string, sender string, sid string, token string, twilioUrl string) {
	msgData := url.Values{}
	msgData.Set("Body", msg)
	msgData.Set("From", sender)

	for _, member := range group.GetMembers() {
		for _, phoneNumber := range member.GetPhoneNumbers() {
			msgData.Set("To", phoneNumber.CanonicalForm)
			err := TwilioSendText(msgData, sid, token, twilioUrl)
			if err != nil {
				log.Printf("Could not send text message to %s: %v", member.Name, err)
			}
		}
	}
}

// Send an email to every member of the group
func sendEmail(group ContactGroup, msg string) {
	log.Fatalf("Not implemented")
}
