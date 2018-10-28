package peopleapi

import (
	"log"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

// GooglePerson is an implementation of the Person interface,
// specifically representing a person (contact) from Google Contacts.
type GooglePerson struct {
	service      *people.PeopleConnectionsService
	Name         string
	memberships  []GroupID
	phoneNumbers []PhoneNumber
	emails       []Email
}

// NewGooglePerson creates a new GooglePerson.
func NewGooglePerson(service *people.Service, name string) GooglePerson {
	return GooglePerson{
		service: people.NewPeopleConnectionsService(service),
		Name:    name,
	}
}

// GetName get's a person's name.
func (p *GooglePerson) GetName() string {
	return p.Name
}

// IsMember checks if the person is a member of the given group.
func (p *GooglePerson) IsMember(grp ContactGroup) bool {
	for _, membership := range p.GetMemberships() {
		if membership == grp.GetGroupID() {
			return true
		}
	}
	return false
}

// GetMemberships gets the IDs of the groups the person is a member of
func (p *GooglePerson) GetMemberships() []GroupID {
	if len(p.memberships) == 0 {
		p.setFields()
	}
	return p.memberships
}

// GetPhoneNumbers gets the phone numbers a person has associated with them.
func (p *GooglePerson) GetPhoneNumbers() []PhoneNumber {
	if len(p.phoneNumbers) == 0 {
		p.setFields()
	}
	return p.phoneNumbers
}

// GetEmails gets the emails a person has associated with them.
func (p *GooglePerson) GetEmails() []Email {
	if len(p.emails) == 0 {
		p.setFields()
	}
	return p.emails
}

// SetMemberships sets the memberships a person has.
//
// This is generally used for creating a person's representation.
func (p *GooglePerson) SetMemberships(memberships []GroupID) {
	p.memberships = memberships
}

// SetPhoneNumbers sets the phone numbers associated with a person.
//
// This is generally used for creating a person's representation.
func (p *GooglePerson) SetPhoneNumbers(phoneNumbers []PhoneNumber) {
	p.phoneNumbers = phoneNumbers
}

// SetEmailAddresses sets the email addresses associated with a
// person.
//
// This is generally used for creating a person's presentation.
func (p *GooglePerson) SetEmailAddresses(emailAddresses []Email) {
	p.emails = emailAddresses
}

// Convert a GooglePerson type from the peopleapi API library to the person type of this library.
func fromPerson(peoplePerson *people.Person, service *people.Service) GooglePerson {
	person := NewGooglePerson(service, peoplePerson.Names[0].DisplayName)
	person.SetEmailAddresses(getEmailsFromEmailAddresses(peoplePerson.EmailAddresses))
	person.SetMemberships(getGroupIDsFromMemberships(peoplePerson.Memberships))
	person.SetPhoneNumbers(getNumbersFromPhoneNumbers(peoplePerson.PhoneNumbers))
	return person
}

// Set the person's information fields.
//
// This is a package private function used by the getter functions
// when a field is empty. As the information for all the fields is
// stored in the same place is makes sense to grab it all when one of
// them is requested to save on network calls.
func (p *GooglePerson) setFields() {
	list := p.service.List("people/me")
	list = list.RequestMaskIncludeField("person.names,person.memberships," +
		"person.phoneNumbers,person.emailAddresses")

	var memberships []GroupID
	var phoneNumbers []PhoneNumber
	var emails []Email

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
			if len(connection.Names) > 0 && connection.Names[0].DisplayName == p.Name {
				memberships = getGroupIDsFromMemberships(connection.Memberships)
				phoneNumbers = getNumbersFromPhoneNumbers(connection.PhoneNumbers)
				emails = getEmailsFromEmailAddresses(connection.EmailAddresses)
				break
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

func getGroupIDsFromMemberships(memberships []*people.Membership) []GroupID {
	var groupIDs []GroupID
	for _, membership := range memberships {
		if membership != nil && membership.ContactGroupMembership != nil {
			groupID := membership.ContactGroupMembership.ContactGroupId
			groupIDs = append(groupIDs, GroupID(groupID))
		}
	}

	return groupIDs
}

func getNumbersFromPhoneNumbers(phonenumbers []*people.PhoneNumber) []PhoneNumber {
	var numbers []PhoneNumber

	for _, phonenumber := range phonenumbers {
		if phonenumber != nil {
			numbers = append(numbers, PhoneNumber(phonenumber.CanonicalForm))
		}
	}

	return numbers
}

func getEmailsFromEmailAddresses(emailAddresses []*people.EmailAddress) []Email {
	var emails []Email
	for _, emailAddress := range emailAddresses {
		if emailAddress != nil {
			emails = append(emails, Email(emailAddress.Value))
		}
	}
	return emails
}
