// Auto generated code DO NOT EDIT
package people

import (
	"net/http"

	"github.com/nick96/enmass/pkg/peopleapi/peopleiface"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

type ModifyContactGroupMembersRequest struct {
	parent *people.ModifyContactGroupMembersRequest
}

func (x *ModifyContactGroupMembersRequest) ResourceNamesToAdd() []string {
	return x.parent.ResourceNamesToAdd
}
func (x *ModifyContactGroupMembersRequest) ResourceNamesToRemove() []string {
	return x.parent.ResourceNamesToRemove
}
func (x *ModifyContactGroupMembersRequest) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *ModifyContactGroupMembersRequest) NullFields() []string {
	return x.parent.NullFields
}
func (x *ModifyContactGroupMembersRequest) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ContactGroupsService struct {
	parent *people.ContactGroupsService
}

func (x *ContactGroupsService) Members() peopleiface.ContactGroupsMembersService {
	return x.parent.Members
}
func (x *ContactGroupsService) BatchGet() peopleiface.ContactGroupsBatchGetCall {
	return x.parent.BatchGet()
}
func (x *ContactGroupsService) Create(createcontactgrouprequest peopleiface.CreateContactGroupRequest) peopleiface.ContactGroupsCreateCall {
	return x.parent.Create(createcontactgrouprequest)
}
func (x *ContactGroupsService) Delete(resourceName string) peopleiface.ContactGroupsDeleteCall {
	return x.parent.Delete(resourceName)
}
func (x *ContactGroupsService) Get(resourceName string) peopleiface.ContactGroupsGetCall {
	return x.parent.Get(resourceName)
}
func (x *ContactGroupsService) List() peopleiface.ContactGroupsListCall {
	return x.parent.List()
}
func (x *ContactGroupsService) Update(resourceName string, updatecontactgrouprequest peopleiface.UpdateContactGroupRequest) peopleiface.ContactGroupsUpdateCall {
	return x.parent.Update(resourceName, updatecontactgrouprequest)
}

type PeopleGetBatchGetCall struct {
	parent *people.PeopleGetBatchGetCall
}

func (x *PeopleGetBatchGetCall) PersonFields(personFields string) peopleiface.PeopleGetBatchGetCall {
	return x.parent.PersonFields(personFields)
}
func (x *PeopleGetBatchGetCall) RequestMaskIncludeField(requestMaskIncludeField string) peopleiface.PeopleGetBatchGetCall {
	return x.parent.RequestMaskIncludeField(requestMaskIncludeField)
}
func (x *PeopleGetBatchGetCall) ResourceNames(resourceNames ...string) peopleiface.PeopleGetBatchGetCall {
	return x.parent.ResourceNames(resourceNames)
}
func (x *PeopleGetBatchGetCall) Fields(s ...googleapi.Field) peopleiface.PeopleGetBatchGetCall {
	return x.parent.Fields(s)
}
func (x *PeopleGetBatchGetCall) IfNoneMatch(entityTag string) peopleiface.PeopleGetBatchGetCall {
	return x.parent.IfNoneMatch(entityTag)
}
func (x *PeopleGetBatchGetCall) Context(ctx context.Context) peopleiface.PeopleGetBatchGetCall {
	return x.parent.Context(ctx)
}
func (x *PeopleGetBatchGetCall) Header() http.Header {
	return x.parent.Header()
}
func (x *PeopleGetBatchGetCall) Do(opts ...googleapi.CallOption) (peopleiface.GetPeopleResponse, error) {
	return x.parent.Do(opts)
}

type Empty struct {
	parent *people.Empty
}

type Name struct {
	parent *people.Name
}

func (x *Name) DisplayName() string {
	return x.parent.DisplayName
}
func (x *Name) DisplayNameLastFirst() string {
	return x.parent.DisplayNameLastFirst
}
func (x *Name) FamilyName() string {
	return x.parent.FamilyName
}
func (x *Name) GivenName() string {
	return x.parent.GivenName
}
func (x *Name) HonorificPrefix() string {
	return x.parent.HonorificPrefix
}
func (x *Name) HonorificSuffix() string {
	return x.parent.HonorificSuffix
}
func (x *Name) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Name) MiddleName() string {
	return x.parent.MiddleName
}
func (x *Name) PhoneticFamilyName() string {
	return x.parent.PhoneticFamilyName
}
func (x *Name) PhoneticFullName() string {
	return x.parent.PhoneticFullName
}
func (x *Name) PhoneticGivenName() string {
	return x.parent.PhoneticGivenName
}
func (x *Name) PhoneticHonorificPrefix() string {
	return x.parent.PhoneticHonorificPrefix
}
func (x *Name) PhoneticHonorificSuffix() string {
	return x.parent.PhoneticHonorificSuffix
}
func (x *Name) PhoneticMiddleName() string {
	return x.parent.PhoneticMiddleName
}
func (x *Name) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Name) NullFields() []string {
	return x.parent.NullFields
}
func (x *Name) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ContactGroupsBatchGetCall struct {
	parent *people.ContactGroupsBatchGetCall
}

func (x *ContactGroupsBatchGetCall) MaxMembers(maxMembers int64) peopleiface.ContactGroupsBatchGetCall {
	return x.parent.MaxMembers(maxMembers)
}
func (x *ContactGroupsBatchGetCall) ResourceNames(resourceNames ...string) peopleiface.ContactGroupsBatchGetCall {
	return x.parent.ResourceNames(resourceNames)
}
func (x *ContactGroupsBatchGetCall) Fields(s ...googleapi.Field) peopleiface.ContactGroupsBatchGetCall {
	return x.parent.Fields(s)
}
func (x *ContactGroupsBatchGetCall) IfNoneMatch(entityTag string) peopleiface.ContactGroupsBatchGetCall {
	return x.parent.IfNoneMatch(entityTag)
}
func (x *ContactGroupsBatchGetCall) Context(ctx context.Context) peopleiface.ContactGroupsBatchGetCall {
	return x.parent.Context(ctx)
}
func (x *ContactGroupsBatchGetCall) Header() http.Header {
	return x.parent.Header()
}
func (x *ContactGroupsBatchGetCall) Do(opts ...googleapi.CallOption) (peopleiface.BatchGetContactGroupsResponse, error) {
	return x.parent.Do(opts)
}

type RelationshipStatus struct {
	parent *people.RelationshipStatus
}

func (x *RelationshipStatus) FormattedValue() string {
	return x.parent.FormattedValue
}
func (x *RelationshipStatus) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *RelationshipStatus) Value() string {
	return x.parent.Value
}
func (x *RelationshipStatus) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *RelationshipStatus) NullFields() []string {
	return x.parent.NullFields
}
func (x *RelationshipStatus) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Event struct {
	parent *people.Event
}

func (x *Event) Date() peopleiface.Date {
	return x.parent.Date
}
func (x *Event) FormattedType() string {
	return x.parent.FormattedType
}
func (x *Event) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Event) Type() string {
	return x.parent.Type
}
func (x *Event) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Event) NullFields() []string {
	return x.parent.NullFields
}
func (x *Event) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ContactGroupsMembersService struct {
	parent *people.ContactGroupsMembersService
}

func (x *ContactGroupsMembersService) Modify(resourceName string, modifycontactgroupmembersrequest peopleiface.ModifyContactGroupMembersRequest) peopleiface.ContactGroupsMembersModifyCall {
	return x.parent.Modify(resourceName, modifycontactgroupmembersrequest)
}

type BatchGetContactGroupsResponse struct {
	parent *people.BatchGetContactGroupsResponse
}

func (x *BatchGetContactGroupsResponse) Responses() []*ContactGroupResponse {
	return x.parent.Responses
}
func (x *BatchGetContactGroupsResponse) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *BatchGetContactGroupsResponse) NullFields() []string {
	return x.parent.NullFields
}
func (x *BatchGetContactGroupsResponse) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ModifyContactGroupMembersResponse struct {
	parent *people.ModifyContactGroupMembersResponse
}

func (x *ModifyContactGroupMembersResponse) NotFoundResourceNames() []string {
	return x.parent.NotFoundResourceNames
}
func (x *ModifyContactGroupMembersResponse) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *ModifyContactGroupMembersResponse) NullFields() []string {
	return x.parent.NullFields
}
func (x *ModifyContactGroupMembersResponse) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Status struct {
	parent *people.Status
}

func (x *Status) Code() int64 {
	return x.parent.Code
}
func (x *Status) Details() []googleapi.RawMessage {
	return x.parent.Details
}
func (x *Status) Message() string {
	return x.parent.Message
}
func (x *Status) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Status) NullFields() []string {
	return x.parent.NullFields
}
func (x *Status) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Gender struct {
	parent *people.Gender
}

func (x *Gender) FormattedValue() string {
	return x.parent.FormattedValue
}
func (x *Gender) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Gender) Value() string {
	return x.parent.Value
}
func (x *Gender) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Gender) NullFields() []string {
	return x.parent.NullFields
}
func (x *Gender) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Organization struct {
	parent *people.Organization
}

func (x *Organization) Current() bool {
	return x.parent.Current
}
func (x *Organization) Department() string {
	return x.parent.Department
}
func (x *Organization) Domain() string {
	return x.parent.Domain
}
func (x *Organization) EndDate() peopleiface.Date {
	return x.parent.EndDate
}
func (x *Organization) FormattedType() string {
	return x.parent.FormattedType
}
func (x *Organization) JobDescription() string {
	return x.parent.JobDescription
}
func (x *Organization) Location() string {
	return x.parent.Location
}
func (x *Organization) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Organization) Name() string {
	return x.parent.Name
}
func (x *Organization) PhoneticName() string {
	return x.parent.PhoneticName
}
func (x *Organization) StartDate() peopleiface.Date {
	return x.parent.StartDate
}
func (x *Organization) Symbol() string {
	return x.parent.Symbol
}
func (x *Organization) Title() string {
	return x.parent.Title
}
func (x *Organization) Type() string {
	return x.parent.Type
}
func (x *Organization) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Organization) NullFields() []string {
	return x.parent.NullFields
}
func (x *Organization) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Date struct {
	parent *people.Date
}

func (x *Date) Day() int64 {
	return x.parent.Day
}
func (x *Date) Month() int64 {
	return x.parent.Month
}
func (x *Date) Year() int64 {
	return x.parent.Year
}
func (x *Date) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Date) NullFields() []string {
	return x.parent.NullFields
}
func (x *Date) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type PeopleCreateContactCall struct {
	parent *people.PeopleCreateContactCall
}

func (x *PeopleCreateContactCall) Parent(parent string) peopleiface.PeopleCreateContactCall {
	return x.parent.Parent(parent)
}
func (x *PeopleCreateContactCall) Fields(s ...googleapi.Field) peopleiface.PeopleCreateContactCall {
	return x.parent.Fields(s)
}
func (x *PeopleCreateContactCall) Context(ctx context.Context) peopleiface.PeopleCreateContactCall {
	return x.parent.Context(ctx)
}
func (x *PeopleCreateContactCall) Header() http.Header {
	return x.parent.Header()
}
func (x *PeopleCreateContactCall) Do(opts ...googleapi.CallOption) (peopleiface.Person, error) {
	return x.parent.Do(opts)
}

type Address struct {
	parent *people.Address
}

func (x *Address) City() string {
	return x.parent.City
}
func (x *Address) Country() string {
	return x.parent.Country
}
func (x *Address) CountryCode() string {
	return x.parent.CountryCode
}
func (x *Address) ExtendedAddress() string {
	return x.parent.ExtendedAddress
}
func (x *Address) FormattedType() string {
	return x.parent.FormattedType
}
func (x *Address) FormattedValue() string {
	return x.parent.FormattedValue
}
func (x *Address) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Address) PoBox() string {
	return x.parent.PoBox
}
func (x *Address) PostalCode() string {
	return x.parent.PostalCode
}
func (x *Address) Region() string {
	return x.parent.Region
}
func (x *Address) StreetAddress() string {
	return x.parent.StreetAddress
}
func (x *Address) Type() string {
	return x.parent.Type
}
func (x *Address) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Address) NullFields() []string {
	return x.parent.NullFields
}
func (x *Address) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type FieldMetadata struct {
	parent *people.FieldMetadata
}

func (x *FieldMetadata) Primary() bool {
	return x.parent.Primary
}
func (x *FieldMetadata) Source() peopleiface.Source {
	return x.parent.Source
}
func (x *FieldMetadata) Verified() bool {
	return x.parent.Verified
}
func (x *FieldMetadata) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *FieldMetadata) NullFields() []string {
	return x.parent.NullFields
}
func (x *FieldMetadata) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Locale struct {
	parent *people.Locale
}

func (x *Locale) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Locale) Value() string {
	return x.parent.Value
}
func (x *Locale) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Locale) NullFields() []string {
	return x.parent.NullFields
}
func (x *Locale) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type BraggingRights struct {
	parent *people.BraggingRights
}

func (x *BraggingRights) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *BraggingRights) Value() string {
	return x.parent.Value
}
func (x *BraggingRights) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *BraggingRights) NullFields() []string {
	return x.parent.NullFields
}
func (x *BraggingRights) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ContactGroupsMembersModifyCall struct {
	parent *people.ContactGroupsMembersModifyCall
}

func (x *ContactGroupsMembersModifyCall) Fields(s ...googleapi.Field) peopleiface.ContactGroupsMembersModifyCall {
	return x.parent.Fields(s)
}
func (x *ContactGroupsMembersModifyCall) Context(ctx context.Context) peopleiface.ContactGroupsMembersModifyCall {
	return x.parent.Context(ctx)
}
func (x *ContactGroupsMembersModifyCall) Header() http.Header {
	return x.parent.Header()
}
func (x *ContactGroupsMembersModifyCall) Do(opts ...googleapi.CallOption) (peopleiface.ModifyContactGroupMembersResponse, error) {
	return x.parent.Do(opts)
}

type PeopleConnectionsListCall struct {
	parent *people.PeopleConnectionsListCall
}

func (x *PeopleConnectionsListCall) PageSize(pageSize int64) peopleiface.PeopleConnectionsListCall {
	return x.parent.PageSize(pageSize)
}
func (x *PeopleConnectionsListCall) PageToken(pageToken string) peopleiface.PeopleConnectionsListCall {
	return x.parent.PageToken(pageToken)
}
func (x *PeopleConnectionsListCall) PersonFields(personFields string) peopleiface.PeopleConnectionsListCall {
	return x.parent.PersonFields(personFields)
}
func (x *PeopleConnectionsListCall) RequestMaskIncludeField(requestMaskIncludeField string) peopleiface.PeopleConnectionsListCall {
	return x.parent.RequestMaskIncludeField(requestMaskIncludeField)
}
func (x *PeopleConnectionsListCall) RequestSyncToken(requestSyncToken bool) peopleiface.PeopleConnectionsListCall {
	return x.parent.RequestSyncToken(requestSyncToken)
}
func (x *PeopleConnectionsListCall) SortOrder(sortOrder string) peopleiface.PeopleConnectionsListCall {
	return x.parent.SortOrder(sortOrder)
}
func (x *PeopleConnectionsListCall) SyncToken(syncToken string) peopleiface.PeopleConnectionsListCall {
	return x.parent.SyncToken(syncToken)
}
func (x *PeopleConnectionsListCall) Fields(s ...googleapi.Field) peopleiface.PeopleConnectionsListCall {
	return x.parent.Fields(s)
}
func (x *PeopleConnectionsListCall) IfNoneMatch(entityTag string) peopleiface.PeopleConnectionsListCall {
	return x.parent.IfNoneMatch(entityTag)
}
func (x *PeopleConnectionsListCall) Context(ctx context.Context) peopleiface.PeopleConnectionsListCall {
	return x.parent.Context(ctx)
}
func (x *PeopleConnectionsListCall) Header() http.Header {
	return x.parent.Header()
}
func (x *PeopleConnectionsListCall) Do(opts ...googleapi.CallOption) (peopleiface.ListConnectionsResponse, error) {
	return x.parent.Do(opts)
}
func (x *PeopleConnectionsListCall) Pages(ctx context.Context, f func(*ListConnectionsResponse) error) error {
	return x.parent.Pages(ctx, f)
}

type Photo struct {
	parent *people.Photo
}

func (x *Photo) Default() bool {
	return x.parent.Default
}
func (x *Photo) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Photo) Url() string {
	return x.parent.Url
}
func (x *Photo) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Photo) NullFields() []string {
	return x.parent.NullFields
}
func (x *Photo) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ContactGroupResponse struct {
	parent *people.ContactGroupResponse
}

func (x *ContactGroupResponse) ContactGroup() peopleiface.ContactGroup {
	return x.parent.ContactGroup
}
func (x *ContactGroupResponse) RequestedResourceName() string {
	return x.parent.RequestedResourceName
}
func (x *ContactGroupResponse) Status() peopleiface.Status {
	return x.parent.Status
}
func (x *ContactGroupResponse) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *ContactGroupResponse) NullFields() []string {
	return x.parent.NullFields
}
func (x *ContactGroupResponse) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type CoverPhoto struct {
	parent *people.CoverPhoto
}

func (x *CoverPhoto) Default() bool {
	return x.parent.Default
}
func (x *CoverPhoto) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *CoverPhoto) Url() string {
	return x.parent.Url
}
func (x *CoverPhoto) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *CoverPhoto) NullFields() []string {
	return x.parent.NullFields
}
func (x *CoverPhoto) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ContactGroupMembership struct {
	parent *people.ContactGroupMembership
}

func (x *ContactGroupMembership) ContactGroupId() string {
	return x.parent.ContactGroupId
}
func (x *ContactGroupMembership) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *ContactGroupMembership) NullFields() []string {
	return x.parent.NullFields
}
func (x *ContactGroupMembership) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ProfileMetadata struct {
	parent *people.ProfileMetadata
}

func (x *ProfileMetadata) ObjectType() string {
	return x.parent.ObjectType
}
func (x *ProfileMetadata) UserTypes() []string {
	return x.parent.UserTypes
}
func (x *ProfileMetadata) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *ProfileMetadata) NullFields() []string {
	return x.parent.NullFields
}
func (x *ProfileMetadata) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Membership struct {
	parent *people.Membership
}

func (x *Membership) ContactGroupMembership() peopleiface.ContactGroupMembership {
	return x.parent.ContactGroupMembership
}
func (x *Membership) DomainMembership() peopleiface.DomainMembership {
	return x.parent.DomainMembership
}
func (x *Membership) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Membership) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Membership) NullFields() []string {
	return x.parent.NullFields
}
func (x *Membership) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type UserDefined struct {
	parent *people.UserDefined
}

func (x *UserDefined) Key() string {
	return x.parent.Key
}
func (x *UserDefined) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *UserDefined) Value() string {
	return x.parent.Value
}
func (x *UserDefined) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *UserDefined) NullFields() []string {
	return x.parent.NullFields
}
func (x *UserDefined) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ContactGroupsDeleteCall struct {
	parent *people.ContactGroupsDeleteCall
}

func (x *ContactGroupsDeleteCall) DeleteContacts(deleteContacts bool) peopleiface.ContactGroupsDeleteCall {
	return x.parent.DeleteContacts(deleteContacts)
}
func (x *ContactGroupsDeleteCall) Fields(s ...googleapi.Field) peopleiface.ContactGroupsDeleteCall {
	return x.parent.Fields(s)
}
func (x *ContactGroupsDeleteCall) Context(ctx context.Context) peopleiface.ContactGroupsDeleteCall {
	return x.parent.Context(ctx)
}
func (x *ContactGroupsDeleteCall) Header() http.Header {
	return x.parent.Header()
}
func (x *ContactGroupsDeleteCall) Do(opts ...googleapi.CallOption) (peopleiface.Empty, error) {
	return x.parent.Do(opts)
}

type Source struct {
	parent *people.Source
}

func (x *Source) Etag() string {
	return x.parent.Etag
}
func (x *Source) Id() string {
	return x.parent.Id
}
func (x *Source) ProfileMetadata() peopleiface.ProfileMetadata {
	return x.parent.ProfileMetadata
}
func (x *Source) Type() string {
	return x.parent.Type
}
func (x *Source) UpdateTime() string {
	return x.parent.UpdateTime
}
func (x *Source) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Source) NullFields() []string {
	return x.parent.NullFields
}
func (x *Source) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Occupation struct {
	parent *people.Occupation
}

func (x *Occupation) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Occupation) Value() string {
	return x.parent.Value
}
func (x *Occupation) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Occupation) NullFields() []string {
	return x.parent.NullFields
}
func (x *Occupation) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type RelationshipInterest struct {
	parent *people.RelationshipInterest
}

func (x *RelationshipInterest) FormattedValue() string {
	return x.parent.FormattedValue
}
func (x *RelationshipInterest) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *RelationshipInterest) Value() string {
	return x.parent.Value
}
func (x *RelationshipInterest) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *RelationshipInterest) NullFields() []string {
	return x.parent.NullFields
}
func (x *RelationshipInterest) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type DomainMembership struct {
	parent *people.DomainMembership
}

func (x *DomainMembership) InViewerDomain() bool {
	return x.parent.InViewerDomain
}
func (x *DomainMembership) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *DomainMembership) NullFields() []string {
	return x.parent.NullFields
}
func (x *DomainMembership) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type PeopleGetCall struct {
	parent *people.PeopleGetCall
}

func (x *PeopleGetCall) PersonFields(personFields string) peopleiface.PeopleGetCall {
	return x.parent.PersonFields(personFields)
}
func (x *PeopleGetCall) RequestMaskIncludeField(requestMaskIncludeField string) peopleiface.PeopleGetCall {
	return x.parent.RequestMaskIncludeField(requestMaskIncludeField)
}
func (x *PeopleGetCall) Fields(s ...googleapi.Field) peopleiface.PeopleGetCall {
	return x.parent.Fields(s)
}
func (x *PeopleGetCall) IfNoneMatch(entityTag string) peopleiface.PeopleGetCall {
	return x.parent.IfNoneMatch(entityTag)
}
func (x *PeopleGetCall) Context(ctx context.Context) peopleiface.PeopleGetCall {
	return x.parent.Context(ctx)
}
func (x *PeopleGetCall) Header() http.Header {
	return x.parent.Header()
}
func (x *PeopleGetCall) Do(opts ...googleapi.CallOption) (peopleiface.Person, error) {
	return x.parent.Do(opts)
}

type Relation struct {
	parent *people.Relation
}

func (x *Relation) FormattedType() string {
	return x.parent.FormattedType
}
func (x *Relation) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Relation) Person() string {
	return x.parent.Person
}
func (x *Relation) Type() string {
	return x.parent.Type
}
func (x *Relation) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Relation) NullFields() []string {
	return x.parent.NullFields
}
func (x *Relation) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type PeopleConnectionsService struct {
	parent *people.PeopleConnectionsService
}

func (x *PeopleConnectionsService) List(resourceName string) peopleiface.PeopleConnectionsListCall {
	return x.parent.List(resourceName)
}

type PersonResponse struct {
	parent *people.PersonResponse
}

func (x *PersonResponse) HttpStatusCode() int64 {
	return x.parent.HttpStatusCode
}
func (x *PersonResponse) Person() peopleiface.Person {
	return x.parent.Person
}
func (x *PersonResponse) RequestedResourceName() string {
	return x.parent.RequestedResourceName
}
func (x *PersonResponse) Status() peopleiface.Status {
	return x.parent.Status
}
func (x *PersonResponse) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *PersonResponse) NullFields() []string {
	return x.parent.NullFields
}
func (x *PersonResponse) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ContactGroupsCreateCall struct {
	parent *people.ContactGroupsCreateCall
}

func (x *ContactGroupsCreateCall) Fields(s ...googleapi.Field) peopleiface.ContactGroupsCreateCall {
	return x.parent.Fields(s)
}
func (x *ContactGroupsCreateCall) Context(ctx context.Context) peopleiface.ContactGroupsCreateCall {
	return x.parent.Context(ctx)
}
func (x *ContactGroupsCreateCall) Header() http.Header {
	return x.parent.Header()
}
func (x *ContactGroupsCreateCall) Do(opts ...googleapi.CallOption) (peopleiface.ContactGroup, error) {
	return x.parent.Do(opts)
}

type ContactGroup struct {
	parent *people.ContactGroup
}

func (x *ContactGroup) Etag() string {
	return x.parent.Etag
}
func (x *ContactGroup) FormattedName() string {
	return x.parent.FormattedName
}
func (x *ContactGroup) GroupType() string {
	return x.parent.GroupType
}
func (x *ContactGroup) MemberCount() int64 {
	return x.parent.MemberCount
}
func (x *ContactGroup) MemberResourceNames() []string {
	return x.parent.MemberResourceNames
}
func (x *ContactGroup) Metadata() peopleiface.ContactGroupMetadata {
	return x.parent.Metadata
}
func (x *ContactGroup) Name() string {
	return x.parent.Name
}
func (x *ContactGroup) ResourceName() string {
	return x.parent.ResourceName
}
func (x *ContactGroup) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *ContactGroup) NullFields() []string {
	return x.parent.NullFields
}
func (x *ContactGroup) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ContactGroupMetadata struct {
	parent *people.ContactGroupMetadata
}

func (x *ContactGroupMetadata) Deleted() bool {
	return x.parent.Deleted
}
func (x *ContactGroupMetadata) UpdateTime() string {
	return x.parent.UpdateTime
}
func (x *ContactGroupMetadata) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *ContactGroupMetadata) NullFields() []string {
	return x.parent.NullFields
}
func (x *ContactGroupMetadata) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type PeopleService struct {
	parent *people.PeopleService
}

func (x *PeopleService) Connections() peopleiface.PeopleConnectionsService {
	return x.parent.Connections
}
func (x *PeopleService) CreateContact(person peopleiface.Person) peopleiface.PeopleCreateContactCall {
	return x.parent.CreateContact(person)
}
func (x *PeopleService) DeleteContact(resourceName string) peopleiface.PeopleDeleteContactCall {
	return x.parent.DeleteContact(resourceName)
}
func (x *PeopleService) Get(resourceName string) peopleiface.PeopleGetCall {
	return x.parent.Get(resourceName)
}
func (x *PeopleService) GetBatchGet() peopleiface.PeopleGetBatchGetCall {
	return x.parent.GetBatchGet()
}
func (x *PeopleService) UpdateContact(resourceName string, person peopleiface.Person) peopleiface.PeopleUpdateContactCall {
	return x.parent.UpdateContact(resourceName, person)
}

type Url struct {
	parent *people.Url
}

func (x *Url) FormattedType() string {
	return x.parent.FormattedType
}
func (x *Url) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Url) Type() string {
	return x.parent.Type
}
func (x *Url) Value() string {
	return x.parent.Value
}
func (x *Url) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Url) NullFields() []string {
	return x.parent.NullFields
}
func (x *Url) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Residence struct {
	parent *people.Residence
}

func (x *Residence) Current() bool {
	return x.parent.Current
}
func (x *Residence) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Residence) Value() string {
	return x.parent.Value
}
func (x *Residence) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Residence) NullFields() []string {
	return x.parent.NullFields
}
func (x *Residence) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type EmailAddress struct {
	parent *people.EmailAddress
}

func (x *EmailAddress) DisplayName() string {
	return x.parent.DisplayName
}
func (x *EmailAddress) FormattedType() string {
	return x.parent.FormattedType
}
func (x *EmailAddress) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *EmailAddress) Type() string {
	return x.parent.Type
}
func (x *EmailAddress) Value() string {
	return x.parent.Value
}
func (x *EmailAddress) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *EmailAddress) NullFields() []string {
	return x.parent.NullFields
}
func (x *EmailAddress) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Tagline struct {
	parent *people.Tagline
}

func (x *Tagline) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Tagline) Value() string {
	return x.parent.Value
}
func (x *Tagline) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Tagline) NullFields() []string {
	return x.parent.NullFields
}
func (x *Tagline) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ImClient struct {
	parent *people.ImClient
}

func (x *ImClient) FormattedProtocol() string {
	return x.parent.FormattedProtocol
}
func (x *ImClient) FormattedType() string {
	return x.parent.FormattedType
}
func (x *ImClient) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *ImClient) Protocol() string {
	return x.parent.Protocol
}
func (x *ImClient) Type() string {
	return x.parent.Type
}
func (x *ImClient) Username() string {
	return x.parent.Username
}
func (x *ImClient) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *ImClient) NullFields() []string {
	return x.parent.NullFields
}
func (x *ImClient) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ContactGroupsListCall struct {
	parent *people.ContactGroupsListCall
}

func (x *ContactGroupsListCall) PageSize(pageSize int64) peopleiface.ContactGroupsListCall {
	return x.parent.PageSize(pageSize)
}
func (x *ContactGroupsListCall) PageToken(pageToken string) peopleiface.ContactGroupsListCall {
	return x.parent.PageToken(pageToken)
}
func (x *ContactGroupsListCall) SyncToken(syncToken string) peopleiface.ContactGroupsListCall {
	return x.parent.SyncToken(syncToken)
}
func (x *ContactGroupsListCall) Fields(s ...googleapi.Field) peopleiface.ContactGroupsListCall {
	return x.parent.Fields(s)
}
func (x *ContactGroupsListCall) IfNoneMatch(entityTag string) peopleiface.ContactGroupsListCall {
	return x.parent.IfNoneMatch(entityTag)
}
func (x *ContactGroupsListCall) Context(ctx context.Context) peopleiface.ContactGroupsListCall {
	return x.parent.Context(ctx)
}
func (x *ContactGroupsListCall) Header() http.Header {
	return x.parent.Header()
}
func (x *ContactGroupsListCall) Do(opts ...googleapi.CallOption) (peopleiface.ListContactGroupsResponse, error) {
	return x.parent.Do(opts)
}
func (x *ContactGroupsListCall) Pages(ctx context.Context, f func(*ListContactGroupsResponse) error) error {
	return x.parent.Pages(ctx, f)
}

type AgeRangeType struct {
	parent *people.AgeRangeType
}

func (x *AgeRangeType) AgeRange() string {
	return x.parent.AgeRange
}
func (x *AgeRangeType) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *AgeRangeType) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *AgeRangeType) NullFields() []string {
	return x.parent.NullFields
}
func (x *AgeRangeType) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type CreateContactGroupRequest struct {
	parent *people.CreateContactGroupRequest
}

func (x *CreateContactGroupRequest) ContactGroup() peopleiface.ContactGroup {
	return x.parent.ContactGroup
}
func (x *CreateContactGroupRequest) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *CreateContactGroupRequest) NullFields() []string {
	return x.parent.NullFields
}
func (x *CreateContactGroupRequest) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type GetPeopleResponse struct {
	parent *people.GetPeopleResponse
}

func (x *GetPeopleResponse) Responses() []*PersonResponse {
	return x.parent.Responses
}
func (x *GetPeopleResponse) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *GetPeopleResponse) NullFields() []string {
	return x.parent.NullFields
}
func (x *GetPeopleResponse) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ListContactGroupsResponse struct {
	parent *people.ListContactGroupsResponse
}

func (x *ListContactGroupsResponse) ContactGroups() []*ContactGroup {
	return x.parent.ContactGroups
}
func (x *ListContactGroupsResponse) NextPageToken() string {
	return x.parent.NextPageToken
}
func (x *ListContactGroupsResponse) NextSyncToken() string {
	return x.parent.NextSyncToken
}
func (x *ListContactGroupsResponse) TotalItems() int64 {
	return x.parent.TotalItems
}
func (x *ListContactGroupsResponse) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *ListContactGroupsResponse) NullFields() []string {
	return x.parent.NullFields
}
func (x *ListContactGroupsResponse) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ContactGroupsUpdateCall struct {
	parent *people.ContactGroupsUpdateCall
}

func (x *ContactGroupsUpdateCall) Fields(s ...googleapi.Field) peopleiface.ContactGroupsUpdateCall {
	return x.parent.Fields(s)
}
func (x *ContactGroupsUpdateCall) Context(ctx context.Context) peopleiface.ContactGroupsUpdateCall {
	return x.parent.Context(ctx)
}
func (x *ContactGroupsUpdateCall) Header() http.Header {
	return x.parent.Header()
}
func (x *ContactGroupsUpdateCall) Do(opts ...googleapi.CallOption) (peopleiface.ContactGroup, error) {
	return x.parent.Do(opts)
}

type PersonMetadata struct {
	parent *people.PersonMetadata
}

func (x *PersonMetadata) Deleted() bool {
	return x.parent.Deleted
}
func (x *PersonMetadata) LinkedPeopleResourceNames() []string {
	return x.parent.LinkedPeopleResourceNames
}
func (x *PersonMetadata) ObjectType() string {
	return x.parent.ObjectType
}
func (x *PersonMetadata) PreviousResourceNames() []string {
	return x.parent.PreviousResourceNames
}
func (x *PersonMetadata) Sources() []*Source {
	return x.parent.Sources
}
func (x *PersonMetadata) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *PersonMetadata) NullFields() []string {
	return x.parent.NullFields
}
func (x *PersonMetadata) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ListConnectionsResponse struct {
	parent *people.ListConnectionsResponse
}

func (x *ListConnectionsResponse) Connections() []*Person {
	return x.parent.Connections
}
func (x *ListConnectionsResponse) NextPageToken() string {
	return x.parent.NextPageToken
}
func (x *ListConnectionsResponse) NextSyncToken() string {
	return x.parent.NextSyncToken
}
func (x *ListConnectionsResponse) TotalItems() int64 {
	return x.parent.TotalItems
}
func (x *ListConnectionsResponse) TotalPeople() int64 {
	return x.parent.TotalPeople
}
func (x *ListConnectionsResponse) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *ListConnectionsResponse) NullFields() []string {
	return x.parent.NullFields
}
func (x *ListConnectionsResponse) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type PhoneNumber struct {
	parent *people.PhoneNumber
}

func (x *PhoneNumber) CanonicalForm() string {
	return x.parent.CanonicalForm
}
func (x *PhoneNumber) FormattedType() string {
	return x.parent.FormattedType
}
func (x *PhoneNumber) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *PhoneNumber) Type() string {
	return x.parent.Type
}
func (x *PhoneNumber) Value() string {
	return x.parent.Value
}
func (x *PhoneNumber) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *PhoneNumber) NullFields() []string {
	return x.parent.NullFields
}
func (x *PhoneNumber) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type ContactGroupsGetCall struct {
	parent *people.ContactGroupsGetCall
}

func (x *ContactGroupsGetCall) MaxMembers(maxMembers int64) peopleiface.ContactGroupsGetCall {
	return x.parent.MaxMembers(maxMembers)
}
func (x *ContactGroupsGetCall) Fields(s ...googleapi.Field) peopleiface.ContactGroupsGetCall {
	return x.parent.Fields(s)
}
func (x *ContactGroupsGetCall) IfNoneMatch(entityTag string) peopleiface.ContactGroupsGetCall {
	return x.parent.IfNoneMatch(entityTag)
}
func (x *ContactGroupsGetCall) Context(ctx context.Context) peopleiface.ContactGroupsGetCall {
	return x.parent.Context(ctx)
}
func (x *ContactGroupsGetCall) Header() http.Header {
	return x.parent.Header()
}
func (x *ContactGroupsGetCall) Do(opts ...googleapi.CallOption) (peopleiface.ContactGroup, error) {
	return x.parent.Do(opts)
}

type Service struct {
	parent *people.Service
}

func (x *Service) BasePath() string {
	return x.parent.BasePath
}
func (x *Service) UserAgent() string {
	return x.parent.UserAgent
}
func (x *Service) ContactGroups() peopleiface.ContactGroupsService {
	return x.parent.ContactGroups
}
func (x *Service) People() peopleiface.PeopleService {
	return x.parent.People
}

type Interest struct {
	parent *people.Interest
}

func (x *Interest) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Interest) Value() string {
	return x.parent.Value
}
func (x *Interest) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Interest) NullFields() []string {
	return x.parent.NullFields
}
func (x *Interest) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Skill struct {
	parent *people.Skill
}

func (x *Skill) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Skill) Value() string {
	return x.parent.Value
}
func (x *Skill) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Skill) NullFields() []string {
	return x.parent.NullFields
}
func (x *Skill) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type PeopleUpdateContactCall struct {
	parent *people.PeopleUpdateContactCall
}

func (x *PeopleUpdateContactCall) UpdatePersonFields(updatePersonFields string) peopleiface.PeopleUpdateContactCall {
	return x.parent.UpdatePersonFields(updatePersonFields)
}
func (x *PeopleUpdateContactCall) Fields(s ...googleapi.Field) peopleiface.PeopleUpdateContactCall {
	return x.parent.Fields(s)
}
func (x *PeopleUpdateContactCall) Context(ctx context.Context) peopleiface.PeopleUpdateContactCall {
	return x.parent.Context(ctx)
}
func (x *PeopleUpdateContactCall) Header() http.Header {
	return x.parent.Header()
}
func (x *PeopleUpdateContactCall) Do(opts ...googleapi.CallOption) (peopleiface.Person, error) {
	return x.parent.Do(opts)
}

type Biography struct {
	parent *people.Biography
}

func (x *Biography) ContentType() string {
	return x.parent.ContentType
}
func (x *Biography) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Biography) Value() string {
	return x.parent.Value
}
func (x *Biography) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Biography) NullFields() []string {
	return x.parent.NullFields
}
func (x *Biography) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Nickname struct {
	parent *people.Nickname
}

func (x *Nickname) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Nickname) Type() string {
	return x.parent.Type
}
func (x *Nickname) Value() string {
	return x.parent.Value
}
func (x *Nickname) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Nickname) NullFields() []string {
	return x.parent.NullFields
}
func (x *Nickname) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type UpdateContactGroupRequest struct {
	parent *people.UpdateContactGroupRequest
}

func (x *UpdateContactGroupRequest) ContactGroup() peopleiface.ContactGroup {
	return x.parent.ContactGroup
}
func (x *UpdateContactGroupRequest) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *UpdateContactGroupRequest) NullFields() []string {
	return x.parent.NullFields
}
func (x *UpdateContactGroupRequest) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type PeopleDeleteContactCall struct {
	parent *people.PeopleDeleteContactCall
}

func (x *PeopleDeleteContactCall) Fields(s ...googleapi.Field) peopleiface.PeopleDeleteContactCall {
	return x.parent.Fields(s)
}
func (x *PeopleDeleteContactCall) Context(ctx context.Context) peopleiface.PeopleDeleteContactCall {
	return x.parent.Context(ctx)
}
func (x *PeopleDeleteContactCall) Header() http.Header {
	return x.parent.Header()
}
func (x *PeopleDeleteContactCall) Do(opts ...googleapi.CallOption) (peopleiface.Empty, error) {
	return x.parent.Do(opts)
}

type Birthday struct {
	parent *people.Birthday
}

func (x *Birthday) Date() peopleiface.Date {
	return x.parent.Date
}
func (x *Birthday) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *Birthday) Text() string {
	return x.parent.Text
}
func (x *Birthday) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Birthday) NullFields() []string {
	return x.parent.NullFields
}
func (x *Birthday) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type Person struct {
	parent *people.Person
}

func (x *Person) Addresses() []*Address {
	return x.parent.Addresses
}
func (x *Person) AgeRange() string {
	return x.parent.AgeRange
}
func (x *Person) AgeRanges() []*AgeRangeType {
	return x.parent.AgeRanges
}
func (x *Person) Biographies() []*Biography {
	return x.parent.Biographies
}
func (x *Person) Birthdays() []*Birthday {
	return x.parent.Birthdays
}
func (x *Person) BraggingRights() []*BraggingRights {
	return x.parent.BraggingRights
}
func (x *Person) CoverPhotos() []*CoverPhoto {
	return x.parent.CoverPhotos
}
func (x *Person) EmailAddresses() []*EmailAddress {
	return x.parent.EmailAddresses
}
func (x *Person) Etag() string {
	return x.parent.Etag
}
func (x *Person) Events() []*Event {
	return x.parent.Events
}
func (x *Person) Genders() []*Gender {
	return x.parent.Genders
}
func (x *Person) ImClients() []*ImClient {
	return x.parent.ImClients
}
func (x *Person) Interests() []*Interest {
	return x.parent.Interests
}
func (x *Person) Locales() []*Locale {
	return x.parent.Locales
}
func (x *Person) Memberships() []*Membership {
	return x.parent.Memberships
}
func (x *Person) Metadata() peopleiface.PersonMetadata {
	return x.parent.Metadata
}
func (x *Person) Names() []*Name {
	return x.parent.Names
}
func (x *Person) Nicknames() []*Nickname {
	return x.parent.Nicknames
}
func (x *Person) Occupations() []*Occupation {
	return x.parent.Occupations
}
func (x *Person) Organizations() []*Organization {
	return x.parent.Organizations
}
func (x *Person) PhoneNumbers() []*PhoneNumber {
	return x.parent.PhoneNumbers
}
func (x *Person) Photos() []*Photo {
	return x.parent.Photos
}
func (x *Person) Relations() []*Relation {
	return x.parent.Relations
}
func (x *Person) RelationshipInterests() []*RelationshipInterest {
	return x.parent.RelationshipInterests
}
func (x *Person) RelationshipStatuses() []*RelationshipStatus {
	return x.parent.RelationshipStatuses
}
func (x *Person) Residences() []*Residence {
	return x.parent.Residences
}
func (x *Person) ResourceName() string {
	return x.parent.ResourceName
}
func (x *Person) SipAddresses() []*SipAddress {
	return x.parent.SipAddresses
}
func (x *Person) Skills() []*Skill {
	return x.parent.Skills
}
func (x *Person) Taglines() []*Tagline {
	return x.parent.Taglines
}
func (x *Person) Urls() []*Url {
	return x.parent.Urls
}
func (x *Person) UserDefined() []*UserDefined {
	return x.parent.UserDefined
}
func (x *Person) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *Person) NullFields() []string {
	return x.parent.NullFields
}
func (x *Person) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

type SipAddress struct {
	parent *people.SipAddress
}

func (x *SipAddress) FormattedType() string {
	return x.parent.FormattedType
}
func (x *SipAddress) Metadata() peopleiface.FieldMetadata {
	return x.parent.Metadata
}
func (x *SipAddress) Type() string {
	return x.parent.Type
}
func (x *SipAddress) Value() string {
	return x.parent.Value
}
func (x *SipAddress) ForceSendFields() []string {
	return x.parent.ForceSendFields
}
func (x *SipAddress) NullFields() []string {
	return x.parent.NullFields
}
func (x *SipAddress) MarshalJSON() ([]byte, error) {
	return x.parent.MarshalJSON()
}

func New(client http.Client) (peopleiface.Service, error) {
	return people.New(client)
}

func NewContactGroupsService(s peopleiface.Service) peopleiface.ContactGroupsService {
	return people.NewContactGroupsService(s)
}

func NewContactGroupsMembersService(s peopleiface.Service) peopleiface.ContactGroupsMembersService {
	return people.NewContactGroupsMembersService(s)
}

func NewPeopleService(s peopleiface.Service) peopleiface.PeopleService {
	return people.NewPeopleService(s)
}

func NewPeopleConnectionsService(s peopleiface.Service) peopleiface.PeopleConnectionsService {
	return people.NewPeopleConnectionsService(s)
}
