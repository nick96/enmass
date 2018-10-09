package main

import (
	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
	"log"
)

type Person struct {
	service *people.PeopleConnectionsService
	Name string
	memberships []*people.Membership
	phoneNumbers []*people.PhoneNumber
	emails []*people.EmailAddress
}

// Create a representation of a person.
func NewPerson(service *people.Service, name string) Person {
	person := Person{}
	person.service = people.NewPeopleConnectionsService(service)
	person.Name = name
	return person
}

// Check if a person is a member of the given contact group.
func (p *Person) IsMember(grp *ContactGroup) bool {
	for _, membership := range p.GetMemberships() {
		if membership.ContactGroupMembership.ContactGroupId == grp.GetGroupId() {
			return true
		}
	}
	return false
}

// Get the memberships a person has.
func (p *Person) GetMemberships() []*people.Membership {
	if len(p.memberships) == 0 {
		p.setFields()
	}
	return p.memberships
}

// Get the phone numbers a person has associated with them.
func (p *Person) GetPhoneNumbers() []*people.PhoneNumber {
	if len(p.phoneNumbers) == 0 {
		p.setFields()
	}
	return p.phoneNumbers
}

// Get the emails a person has associated with them.
func (p *Person) GetEmails() []*people.EmailAddress {
	if len(p.emails) == 0 {
		p.setFields()
	}
	return p.emails
}

// Set the memberships a person has.
//
// This is generally used for creating a person's presentation.
func (p *Person) SetMemberships(memberships []*people.Membership) {
	p.memberships = memberships
}

// Set the phone numbers associated with a person.
//
// This is generally used for creating a person's presentation.

func (p *Person) SetPhoneNumbers(phoneNumbers []*people.PhoneNumber) {
	p.phoneNumbers = phoneNumbers
}

// Set the email addresses associated with a person.
//
// This is generally used for creating a person's presentation.
func (p *Person) SetEmailAddresses(emailAddresses []*people.EmailAddress) {
	p.emails = emailAddresses
}

// Set the person's information fields.
//
// This is a package private function used by the getter functions when a field is empty. As the information for all
// the fields is stored in the same place is makes sense to grab it all when one of them is requested to save on
// network calls.
func (p *Person) setFields() {
	list := p.service.List("people/me")
	list = list.RequestMaskIncludeField("person.names,person.memberships,person.phoneNumbers,person.emailAddresses")

	var memberships []*people.Membership
	var phoneNumbers []*people.PhoneNumber
	var emails []*people.EmailAddress

	var nextPage googleapi.Field
	for {
		if nextPage != "" {
			list = list.Fields("nextPageToken", nextPage)
		}

		resp, err := list.Do()
		if err != nil {
			log.Fatalf("Could not get people: %v", err)
		}

		for _, connection := range resp.Connections {
			if len(connection.Names) > 0 {
				if connection.Names[0].DisplayName == p.Name {
					memberships = connection.Memberships
					phoneNumbers = connection.PhoneNumbers
					emails = connection.EmailAddresses
					break
				}
			}
		}


		if len(resp.NextPageToken) == 0 {
			break
		}
		nextPage = googleapi.Field(resp.NextPageToken)
	}

	p.memberships = memberships
	p.phoneNumbers = phoneNumbers
	p.emails = emails
}
