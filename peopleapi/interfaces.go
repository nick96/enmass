package peopleapi

// GroupID represents a group ID.
type GroupID string

// Email represents a persons email
type Email string

// PhoneNumber represents a phone number in its cannonical form
type PhoneNumber string

// ContactGroup is the contact of methods that must be implemented for
// something to be considered a contact group
type ContactGroup interface {
	GetName() string
	GetGroupID() GroupID
	GetMembers() []Person
	GetEmails() []Email
	GetPhoneNumbers() []PhoneNumber
}

// Person is the contact of methods that must be implemented for
// something to be considered a person
type Person interface {
	GetName() string
	IsMember(group ContactGroup) bool
	GetMemberships() []GroupID
	GetPhoneNumbers() []PhoneNumber
	GetEmails() []Email
	SetMemberships(memberships []GroupID)
	SetPhoneNumbers(phoneNumbers []PhoneNumber)
	SetEmailAddresses(emails []Email)
}
