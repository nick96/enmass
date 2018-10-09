package main

import (
	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
	"log"
	"strings"
)

type ContactGroup struct {
	service *people.Service
	Name    string
	groupId      string
	members []*Person
}

// Create a new contact group representation.
//
// Getting the group ID and members is deferred to when they are first requested.
func NewContactGroup(service *people.Service, groupName string) ContactGroup {
	contactGroup := ContactGroup{}
	contactGroup.service = service
	contactGroup.Name = groupName
	return contactGroup
}

// Get the groups ID.
//
// If the group's ID is already stored in the struct then that is returned, otherwise the people API is used to retrieve
// it.
func (g *ContactGroup) GetGroupId() string {
	if len(g.groupId) == 0 {
		g.groupId = g.getGroupId()
	}
	return g.groupId
}


// Get all the members of the group.
//
// If the group's members are already stored in the struct then they are just returned, otherwise the people API is used
// to retrieve them.
func (g *ContactGroup) GetMembers() []*Person {
	if len(g.members) == 0 {
		g.members = g.getMembers()
	}
	return g.members
}

// Get the group's ID.
//
// This is a package private function that does a network call to retrieve the group ID.
func (g *ContactGroup) getGroupId() string {
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
func (g *ContactGroup) getMembers() []*Person {
	service := people.NewPeopleConnectionsService(g.service)
	list := service.List("people/me")
	list = list.RequestMaskIncludeField("person.names,person.memberships")

	if len(g.groupId) == 0 {
		g.groupId = g.getGroupId()
	}

	var members []*Person

	var nextPage googleapi.Field
	for {
		if len(nextPage) != 0 {
			list = list.Fields("nextPageToken", nextPage)
		}

		resp, err := list.Do()
		if err != nil {
			log.Fatalf("Could not get group members: %v", err)
		}

		for _, connection := range resp.Connections {
			for _, membership := range connection.Memberships {
				if membership.ContactGroupMembership.ContactGroupId == g.groupId {
					person := NewPerson(g.service, connection.Names[0].DisplayName)
					person.SetMemberships(connection.Memberships)
					person.SetPhoneNumbers(connection.PhoneNumbers)
					person.SetEmailAddresses(connection.EmailAddresses)
					members = append(members, &person)
				}
			}
		}

		if len(resp.NextPageToken) == 0 {
			break
		}

		nextPage = googleapi.Field(resp.NextPageToken)
	}

	return members
}
