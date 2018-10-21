package peopleapi

import (
	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
	"log"
	"strings"
)

type GoogleContactGroup struct {
	service *people.Service
	Name    string
	groupId string
	members []Person
}

// Create a new contact group representation.
//
// Getting the group ID and members is deferred to when they are first requested.
func NewGoogleContactGroup(service *people.Service, groupName string) ContactGroup {
	return &GoogleContactGroup{
		service: service,
		Name:    groupName,
	}
}

func (g *GoogleContactGroup) GetName() string {
	return g.Name
}

func (g *GoogleContactGroup) GetGroupId() string {
	if len(g.groupId) == 0 {
		g.groupId = g.getGroupId()
	}
	return g.groupId
}

func (g *GoogleContactGroup) GetMembers() []Person {
	if len(g.members) == 0 {
		for _, member := range g.getMembers() {
			g.members = append(g.members, member)
		}
	}
	return g.members
}

func (g *GoogleContactGroup) GetEmails() []*people.EmailAddress {
	var emails []*people.EmailAddress
	for _, member := range g.GetMembers() {
		emails = append(emails, member.GetEmails()...)
	}
	return emails
}

func (g *GoogleContactGroup) GetPhoneNumbers() []*people.PhoneNumber {
	var phoneNumbers []*people.PhoneNumber
	for _, member := range g.GetMembers() {
		phoneNumbers = append(phoneNumbers, member.GetPhoneNumbers()...)
	}
	return phoneNumbers
}

// Get the group's ID.
//
// This is a package private function that does a network call to retrieve the group ID.
func (g *GoogleContactGroup) getGroupId() string {
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
				return strings.Split(group.ResourceName, "/")[1]
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
	resName := "contactGroups/" + g.getGroupId()

	group, err := groupService.Get(resName).Do()
	if err != nil {
		log.Fatalf("Could not get group: %v", err)
	}

	var members []*GooglePerson
	if len(group.MemberResourceNames) > 0 {
		membersGet := peopleService.GetBatchGet().ResourceNames(group.MemberResourceNames...)
		membersGet = membersGet.RequestMaskIncludeField("person.names,person.phoneNumbers,person.emailAddresses")
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
