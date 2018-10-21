package peopleapi

import "google.golang.org/api/people/v1"

type ContactGroup interface {
	GetName() string
	GetGroupId() string
	GetMembers() []Person
	GetEmails() []*people.EmailAddress
	GetPhoneNumbers() []*people.PhoneNumber
}

type Person interface {
	GetName() string
	IsMember(group ContactGroup) bool
	GetMemberships() []*people.Membership
	GetPhoneNumbers() []*people.PhoneNumber
	GetEmails() []*people.EmailAddress
	SetMemberships(memberships []*people.Membership)
	SetPhoneNumbers(phoneNumbers []*people.PhoneNumber)
	SetEmailAddresses(emails []*people.EmailAddress)
}
