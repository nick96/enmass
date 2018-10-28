// Auto generated code DO NOT EDIT
package peopleiface

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
)

type Source interface {
	Etag() string
	Id() string
	ProfileMetadata() *ProfileMetadata
	Type() string
	UpdateTime() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Nickname interface {
	Metadata() *FieldMetadata
	Type() string
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ContactGroupsMembersModifyCall interface {
	Fields(s ...googleapi.Field) ContactGroupsMembersModifyCall
	Context(ctx context.Context) ContactGroupsMembersModifyCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (ModifyContactGroupMembersResponse, error)
}

type Biography interface {
	ContentType() string
	Metadata() *FieldMetadata
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ProfileMetadata interface {
	ObjectType() string
	UserTypes() []string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Photo interface {
	Default() bool
	Metadata() *FieldMetadata
	Url() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type SipAddress interface {
	FormattedType() string
	Metadata() *FieldMetadata
	Type() string
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type BraggingRights interface {
	Metadata() *FieldMetadata
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type PhoneNumber interface {
	CanonicalForm() string
	FormattedType() string
	Metadata() *FieldMetadata
	Type() string
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ListConnectionsResponse interface {
	Connections() []*Person
	NextPageToken() string
	NextSyncToken() string
	TotalItems() int64
	TotalPeople() int64
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type DomainMembership interface {
	InViewerDomain() bool
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Birthday interface {
	Date() *Date
	Metadata() *FieldMetadata
	Text() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ModifyContactGroupMembersRequest interface {
	ResourceNamesToAdd() []string
	ResourceNamesToRemove() []string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ContactGroupsCreateCall interface {
	Fields(s ...googleapi.Field) *ContactGroupsCreateCall
	Context(ctx context.Context) *ContactGroupsCreateCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (*ContactGroup, error)
}

type ContactGroupMembership interface {
	ContactGroupId() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ContactGroupsListCall interface {
	PageSize(pageSize int64) *ContactGroupsListCall
	PageToken(pageToken string) *ContactGroupsListCall
	SyncToken(syncToken string) *ContactGroupsListCall
	Fields(s ...googleapi.Field) *ContactGroupsListCall
	IfNoneMatch(entityTag string) *ContactGroupsListCall
	Context(ctx context.Context) *ContactGroupsListCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (*ListContactGroupsResponse, error)
	Pages(ctx context.Context, f func(*ListContactGroupsResponse) error) error
}

type Url interface {
	FormattedType() string
	Metadata() *FieldMetadata
	Type() string
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Membership interface {
	ContactGroupMembership() *ContactGroupMembership
	DomainMembership() *DomainMembership
	Metadata() *FieldMetadata
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ContactGroup interface {
	Etag() string
	FormattedName() string
	GroupType() string
	MemberCount() int64
	MemberResourceNames() []string
	Metadata() *ContactGroupMetadata
	Name() string
	ResourceName() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ContactGroupResponse interface {
	ContactGroup() *ContactGroup
	RequestedResourceName() string
	Status() *Status
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type FieldMetadata interface {
	Primary() bool
	Source() Source
	Verified() bool
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ContactGroupsUpdateCall interface {
	Fields(s ...googleapi.Field) *ContactGroupsUpdateCall
	Context(ctx context.Context) *ContactGroupsUpdateCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (*ContactGroup, error)
}

type Skill interface {
	Metadata() *FieldMetadata
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ContactGroupsService interface {
	Members() *ContactGroupsMembersService
	BatchGet() *ContactGroupsBatchGetCall
	Create(createcontactgrouprequest *CreateContactGroupRequest) *ContactGroupsCreateCall
	Delete(resourceName string) *ContactGroupsDeleteCall
	Get(resourceName string) *ContactGroupsGetCall
	List() *ContactGroupsListCall
	Update(resourceName string, updatecontactgrouprequest *UpdateContactGroupRequest) *ContactGroupsUpdateCall
}

type Address interface {
	City() string
	Country() string
	CountryCode() string
	ExtendedAddress() string
	FormattedType() string
	FormattedValue() string
	Metadata() *FieldMetadata
	PoBox() string
	PostalCode() string
	Region() string
	StreetAddress() string
	Type() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ModifyContactGroupMembersResponse interface {
	NotFoundResourceNames() []string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type PersonMetadata interface {
	Deleted() bool
	LinkedPeopleResourceNames() []string
	ObjectType() string
	PreviousResourceNames() []string
	Sources() []*Source
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Organization interface {
	Current() bool
	Department() string
	Domain() string
	EndDate() *Date
	FormattedType() string
	JobDescription() string
	Location() string
	Metadata() *FieldMetadata
	Name() string
	PhoneticName() string
	StartDate() *Date
	Symbol() string
	Title() string
	Type() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type PeopleCreateContactCall interface {
	Parent(parent string) *PeopleCreateContactCall
	Fields(s ...googleapi.Field) *PeopleCreateContactCall
	Context(ctx context.Context) *PeopleCreateContactCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (*Person, error)
}

type ContactGroupsBatchGetCall interface {
	MaxMembers(maxMembers int64) *ContactGroupsBatchGetCall
	ResourceNames(resourceNames ...string) *ContactGroupsBatchGetCall
	Fields(s ...googleapi.Field) *ContactGroupsBatchGetCall
	IfNoneMatch(entityTag string) *ContactGroupsBatchGetCall
	Context(ctx context.Context) *ContactGroupsBatchGetCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (*BatchGetContactGroupsResponse, error)
}

type UpdateContactGroupRequest interface {
	ContactGroup() *ContactGroup
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type EmailAddress interface {
	DisplayName() string
	FormattedType() string
	Metadata() *FieldMetadata
	Type() string
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type AgeRangeType interface {
	AgeRange() string
	Metadata() *FieldMetadata
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Event interface {
	Date() *Date
	FormattedType() string
	Metadata() *FieldMetadata
	Type() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type PeopleConnectionsService interface {
	List(resourceName string) *PeopleConnectionsListCall
}

type Tagline interface {
	Metadata() *FieldMetadata
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type CreateContactGroupRequest interface {
	ContactGroup() *ContactGroup
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Status interface {
	Code() int64
	Details() []googleapi.RawMessage
	Message() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type PeopleConnectionsListCall interface {
	PageSize(pageSize int64) *PeopleConnectionsListCall
	PageToken(pageToken string) *PeopleConnectionsListCall
	PersonFields(personFields string) *PeopleConnectionsListCall
	RequestMaskIncludeField(requestMaskIncludeField string) *PeopleConnectionsListCall
	RequestSyncToken(requestSyncToken bool) *PeopleConnectionsListCall
	SortOrder(sortOrder string) *PeopleConnectionsListCall
	SyncToken(syncToken string) *PeopleConnectionsListCall
	Fields(s ...googleapi.Field) *PeopleConnectionsListCall
	IfNoneMatch(entityTag string) *PeopleConnectionsListCall
	Context(ctx context.Context) *PeopleConnectionsListCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (*ListConnectionsResponse, error)
	Pages(ctx context.Context, f func(*ListConnectionsResponse) error) error
}

type Date interface {
	Day() int64
	Month() int64
	Year() int64
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type UserDefined interface {
	Key() string
	Metadata() *FieldMetadata
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type GetPeopleResponse interface {
	Responses() []*PersonResponse
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type PeopleUpdateContactCall interface {
	UpdatePersonFields(updatePersonFields string) *PeopleUpdateContactCall
	Fields(s ...googleapi.Field) *PeopleUpdateContactCall
	Context(ctx context.Context) *PeopleUpdateContactCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (*Person, error)
}

type Service interface {
	BasePath() string
	UserAgent() string
	ContactGroups() *ContactGroupsService
	People() *PeopleService
}

type ContactGroupMetadata interface {
	Deleted() bool
	UpdateTime() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Interest interface {
	Metadata() *FieldMetadata
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type PeopleDeleteContactCall interface {
	Fields(s ...googleapi.Field) *PeopleDeleteContactCall
	Context(ctx context.Context) *PeopleDeleteContactCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (*Empty, error)
}

type Person interface {
	Addresses() []*Address
	AgeRange() string
	AgeRanges() []*AgeRangeType
	Biographies() []*Biography
	Birthdays() []*Birthday
	BraggingRights() []*BraggingRights
	CoverPhotos() []*CoverPhoto
	EmailAddresses() []*EmailAddress
	Etag() string
	Events() []*Event
	Genders() []*Gender
	ImClients() []*ImClient
	Interests() []*Interest
	Locales() []*Locale
	Memberships() []*Membership
	Metadata() *PersonMetadata
	Names() []*Name
	Nicknames() []*Nickname
	Occupations() []*Occupation
	Organizations() []*Organization
	PhoneNumbers() []*PhoneNumber
	Photos() []*Photo
	Relations() []*Relation
	RelationshipInterests() []*RelationshipInterest
	RelationshipStatuses() []*RelationshipStatus
	Residences() []*Residence
	ResourceName() string
	SipAddresses() []*SipAddress
	Skills() []*Skill
	Taglines() []*Tagline
	Urls() []*Url
	UserDefined() []*UserDefined
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ContactGroupsGetCall interface {
	MaxMembers(maxMembers int64) *ContactGroupsGetCall
	Fields(s ...googleapi.Field) *ContactGroupsGetCall
	IfNoneMatch(entityTag string) *ContactGroupsGetCall
	Context(ctx context.Context) *ContactGroupsGetCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (*ContactGroup, error)
}

type Locale interface {
	Metadata() *FieldMetadata
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ContactGroupsDeleteCall interface {
	DeleteContacts(deleteContacts bool) *ContactGroupsDeleteCall
	Fields(s ...googleapi.Field) *ContactGroupsDeleteCall
	Context(ctx context.Context) *ContactGroupsDeleteCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (*Empty, error)
}

type Name interface {
	DisplayName() string
	DisplayNameLastFirst() string
	FamilyName() string
	GivenName() string
	HonorificPrefix() string
	HonorificSuffix() string
	Metadata() *FieldMetadata
	MiddleName() string
	PhoneticFamilyName() string
	PhoneticFullName() string
	PhoneticGivenName() string
	PhoneticHonorificPrefix() string
	PhoneticHonorificSuffix() string
	PhoneticMiddleName() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type ContactGroupsMembersService interface {
	Modify(resourceName string, modifycontactgroupmembersrequest *ModifyContactGroupMembersRequest) *ContactGroupsMembersModifyCall
}

type RelationshipStatus interface {
	FormattedValue() string
	Metadata() *FieldMetadata
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type PeopleGetBatchGetCall interface {
	PersonFields(personFields string) *PeopleGetBatchGetCall
	RequestMaskIncludeField(requestMaskIncludeField string) *PeopleGetBatchGetCall
	ResourceNames(resourceNames ...string) *PeopleGetBatchGetCall
	Fields(s ...googleapi.Field) *PeopleGetBatchGetCall
	IfNoneMatch(entityTag string) *PeopleGetBatchGetCall
	Context(ctx context.Context) *PeopleGetBatchGetCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (*GetPeopleResponse, error)
}

type BatchGetContactGroupsResponse interface {
	Responses() []*ContactGroupResponse
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Residence interface {
	Current() bool
	Metadata() *FieldMetadata
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Empty interface {
}

type ImClient interface {
	FormattedProtocol() string
	FormattedType() string
	Metadata() *FieldMetadata
	Protocol() string
	Type() string
	Username() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type RelationshipInterest interface {
	FormattedValue() string
	Metadata() *FieldMetadata
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Relation interface {
	FormattedType() string
	Metadata() *FieldMetadata
	Person() string
	Type() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Gender interface {
	FormattedValue() string
	Metadata() *FieldMetadata
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type PeopleService interface {
	Connections() *PeopleConnectionsService
	CreateContact(person *Person) *PeopleCreateContactCall
	DeleteContact(resourceName string) *PeopleDeleteContactCall
	Get(resourceName string) *PeopleGetCall
	GetBatchGet() *PeopleGetBatchGetCall
	UpdateContact(resourceName string, person *Person) *PeopleUpdateContactCall
}

type PeopleGetCall interface {
	PersonFields(personFields string) *PeopleGetCall
	RequestMaskIncludeField(requestMaskIncludeField string) *PeopleGetCall
	Fields(s ...googleapi.Field) *PeopleGetCall
	IfNoneMatch(entityTag string) *PeopleGetCall
	Context(ctx context.Context) *PeopleGetCall
	Header() http.Header
	Do(opts ...googleapi.CallOption) (*Person, error)
}

type ListContactGroupsResponse interface {
	ContactGroups() []*ContactGroup
	NextPageToken() string
	NextSyncToken() string
	TotalItems() int64
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type CoverPhoto interface {
	Default() bool
	Metadata() *FieldMetadata
	Url() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type PersonResponse interface {
	HttpStatusCode() int64
	Person() *Person
	RequestedResourceName() string
	Status() *Status
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}

type Occupation interface {
	Metadata() *FieldMetadata
	Value() string
	ForceSendFields() []string
	NullFields() []string
	MarshalJSON() ([]byte, error)
}
