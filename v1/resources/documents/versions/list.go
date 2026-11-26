package versions

import (
	"encoding/json"
	"time"
)

// DocumentsVersionsListOutputItemsEditorsOrganizationActorMember represents the documents versions list output items editors organization actor member type.
type DocumentsVersionsListOutputItemsEditorsOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsVersionsListOutputItemsEditorsOrganizationActorTeams - The teams the actor belongs to
type DocumentsVersionsListOutputItemsEditorsOrganizationActorTeams struct {
	// Id - The team ID
	Id string `json:"id"`
	// Name - The team name
	Name string `json:"name"`
	// Slug - The team slug
	Slug string `json:"slug"`
	// AssignmentId - The team assignment ID
	AssignmentId string `json:"assignment_id"`
	// CreatedAt - The team assignment creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The team assignment last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsVersionsListOutputItemsEditorsOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type DocumentsVersionsListOutputItemsEditorsOrganizationActor struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Type - The organization member's type
	Type string `json:"type"`
	// OrganizationId - The organization member's organization ID
	OrganizationId string `json:"organization_id"`
	// Name - The organization member's name
	Name string `json:"name"`
	// Email - The organization member's email
	Email *string `json:"email,omitempty"`
	// ImageUrl - The organization member's image URL
	ImageUrl string                                                          `json:"image_url"`
	Member   *DocumentsVersionsListOutputItemsEditorsOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsVersionsListOutputItemsEditorsOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsVersionsListOutputItemsEditorsConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsVersionsListOutputItemsEditorsConsumer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address.
	Email string `json:"email"`
	// ImageUrl - Avatar image URL.
	ImageUrl string `json:"image_url"`
	// UserId - User this consumer is linked to. `null` when unlinked.
	UserId *string `json:"user_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsVersionsListOutputItemsEditors - Actor behind a participant.
type DocumentsVersionsListOutputItemsEditors struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsVersionsListOutputItemsEditorsOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsVersionsListOutputItemsEditorsConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsVersionsListOutputItems represents the documents versions list output items type.
type DocumentsVersionsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// DocumentId - Document this version belongs to.
	DocumentId string `json:"document_id"`
	// VersionNumber - Sequential number of this version.
	VersionNumber float64 `json:"version_number"`
	// PreviousVersionId - Previous version, by ID. `null` for the first version.
	PreviousVersionId *string `json:"previous_version_id,omitempty"`
	// ListEditedAt - When the version content was last changed.
	ListEditedAt *time.Time `json:"list_edited_at,omitempty"`
	// Content - Text body of the document at this version.
	Content string `json:"content"`
	// Editors - Actors who edited this version.
	Editors []DocumentsVersionsListOutputItemsEditors `json:"editors"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// DocumentsVersionsListOutputPagination represents the documents versions list output pagination type.
type DocumentsVersionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// DocumentsVersionsListOutput represents the documents versions list output type.
type DocumentsVersionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                `json:"object"`
	Items      []DocumentsVersionsListOutputItems    `json:"items"`
	Pagination DocumentsVersionsListOutputPagination `json:"pagination"`
}

// MapDocumentsVersionsListOutputFromJSON deserializes JSON data into a DocumentsVersionsListOutput.
func MapDocumentsVersionsListOutputFromJSON(data []byte) (*DocumentsVersionsListOutput, error) {
	var v DocumentsVersionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsVersionsListOutputToJSON serializes a DocumentsVersionsListOutput to JSON.
func MapDocumentsVersionsListOutputToJSON(v *DocumentsVersionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsVersionsListQueryCreatedAt - Only return results whose version creation time is in this range. Set `gt`, `lt`, or both.
type DocumentsVersionsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsVersionsListQueryLastEditedAt - Only return results whose version last edit time is in this range. Set `gt`, `lt`, or both.
type DocumentsVersionsListQueryLastEditedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsVersionsListQuery represents the documents versions list query type.
type DocumentsVersionsListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Id - Only return versions with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CreatedAt - Only return results whose version creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *DocumentsVersionsListQueryCreatedAt `json:"created_at,omitempty"`
	// LastEditedAt - Only return results whose version last edit time is in this range. Set `gt`, `lt`, or both.
	LastEditedAt *DocumentsVersionsListQueryLastEditedAt `json:"last_edited_at,omitempty"`
}

// MapDocumentsVersionsListQueryFromJSON deserializes JSON data into a DocumentsVersionsListQuery.
func MapDocumentsVersionsListQueryFromJSON(data []byte) (*DocumentsVersionsListQuery, error) {
	var v DocumentsVersionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsVersionsListQueryToJSON serializes a DocumentsVersionsListQuery to JSON.
func MapDocumentsVersionsListQueryToJSON(v *DocumentsVersionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
