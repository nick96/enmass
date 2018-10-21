package peopleapi

import (
	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
	"log"
	"strings"
)

// GoogleContactGroup is an implementation of the ContactGroup
// interface, specifically representing a contact group from Google
// Contacts.
type GoogleContactGroup struct {
	service *people.Service
	Name    string
	groupID GroupID
	members []Person
}

// NewGoogleContactGroup creates a new contact group representation.
//
// Getting the group ID and members is deferred to when they are first requested.
func NewGoogleContactGroup(service *people.Service, groupName string) ContactGroup {
	return &GoogleContactGroup{
		service: service,
		Name:    groupName,
	}
}

// GetName gets the group's name
func (g *GoogleContactGroup) GetName() string {
	return g.Name
}

// GetGroupID gets the groups ID
func (g *GoogleContactGroup) GetGroupID() GroupID {
	if len(g.groupID) == 0 {
		g.groupID = g.getGroupID()
	}
	return g.groupID
}

// GetMembers gets all members of a group
func (g *GoogleContactGroup) GetMembers() []Person {
	if len(g.members) == 0 {
		for _, member := range g.getMembers() {
			g.members = append(g.members, member)
		}
	}
	return g.members
}

// GetEmails gets every email of every member of the group.
func (g *GoogleContactGroup) GetEmails() []Email {
	var emails []Email
	for _, member := range g.GetMembers() {
		emails = append(emails, member.GetEmails()...)
	}
	return emails
}

// GetPhoneNumbers gets all the phone numbers in a contact group
func (g *GoogleContactGroup) GetPhoneNumbers() []PhoneNumber {
	var phoneNumbers []PhoneNumber
	for _, member := range g.GetMembers() {
		phoneNumbers = append(phoneNumbers, member.GetPhoneNumbers()...)
	}
	return phoneNumbers
}

// This is a package private function that does a network call to
// retrieve the group ID.
func (g *GoogleContactGroup) getGroupID() GroupID {
	service := people.NewContactGroupsService(g.service)
	list := service.List()
	var nextPage googleapi.Field
	for {
		if len(nextPage) != 0 {
			list = list.Fields("nextPageToken", nextPage)
		}

		resp, err := list.Do()
		if err != nil {
			log.Fatalf("Could not get group ID: %v", err)
		}

		for _, group := range resp.ContactGroups {
			if group.FormattedName == g.Name {
				return GroupID(strings.Split(group.ResourceName, "/")[1])
			}
		}

		if len(resp.NextPageToken) == 0 {
			break
		}

		nextPage = googleapi.Field(resp.NextPageToken)
	}

	return ""
}

// Get the group's members.
//
// This is a package private function that does a network call to retrieve the group's members.
func (g *GoogleContactGroup) getMembers() []*GooglePerson {
	groupService := people.NewContactGroupsService(g.service)
	peopleService := people.NewPeopleService(g.service)
	resName := string("contactGroups/" + g.getGroupID())

	group, err := groupService.Get(resName).Do()
	if err != nil {
		log.Fatalf("Could not get group: %v", err)
	}

	var members []*GooglePerson
	if len(group.MemberResourceNames) > 0 {
		membersGet := peopleService.GetBatchGet().ResourceNames(group.MemberResourceNames...)
		membersGet = membersGet.RequestMaskIncludeField("person.names,person.phoneNumbers," +
			"person.emailAddresses")
		resp, err := membersGet.Do()

		if err != nil {
			log.Fatalf("Could not get group members: %v", err)
		}

		for _, personResp := range resp.Responses {
			person := fromPerson(personResp.Person, g.service)
			members = append(members, &person)
		}

	}
	return members
}
