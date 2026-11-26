package participants

import (
	"encoding/json"
	"time"
)

// DocumentsParticipantsListOutputItemsActorOrganizationActorMember represents the documents participants list output items actor organization actor member type.
type DocumentsParticipantsListOutputItemsActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsParticipantsListOutputItemsActorOrganizationActorTeams - The teams the actor belongs to
type DocumentsParticipantsListOutputItemsActorOrganizationActorTeams struct {
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

// DocumentsParticipantsListOutputItemsActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type DocumentsParticipantsListOutputItemsActorOrganizationActor struct {
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
	ImageUrl string                                                            `json:"image_url"`
	Member   *DocumentsParticipantsListOutputItemsActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsParticipantsListOutputItemsActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsParticipantsListOutputItemsActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsParticipantsListOutputItemsActorConsumer struct {
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

// DocumentsParticipantsListOutputItemsActor - Actor behind a participant.
type DocumentsParticipantsListOutputItemsActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsParticipantsListOutputItemsActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsParticipantsListOutputItemsActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsParticipantsListOutputItems represents the documents participants list output items type.
type DocumentsParticipantsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Role - What the participant can do in the document.
	Role string `json:"role"`
	// EditCount - Number of edits the participant has made.
	EditCount float64 `json:"edit_count"`
	// LastEditedAt - When the participant last changed the document.
	LastEditedAt *time.Time `json:"last_edited_at,omitempty"`
	// LastViewedAt - When the participant last opened the document.
	LastViewedAt *time.Time `json:"last_viewed_at,omitempty"`
	// Actor - Actor behind a participant.
	Actor DocumentsParticipantsListOutputItemsActor `json:"actor"`
	// StoreParticipantId - Linked store participant, by ID.
	StoreParticipantId *string `json:"store_participant_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// DocumentsParticipantsListOutputPagination represents the documents participants list output pagination type.
type DocumentsParticipantsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// DocumentsParticipantsListOutput represents the documents participants list output type.
type DocumentsParticipantsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                    `json:"object"`
	Items      []DocumentsParticipantsListOutputItems    `json:"items"`
	Pagination DocumentsParticipantsListOutputPagination `json:"pagination"`
}

// MapDocumentsParticipantsListOutputFromJSON deserializes JSON data into a DocumentsParticipantsListOutput.
func MapDocumentsParticipantsListOutputFromJSON(data []byte) (*DocumentsParticipantsListOutput, error) {
	var v DocumentsParticipantsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsParticipantsListOutputToJSON serializes a DocumentsParticipantsListOutput to JSON.
func MapDocumentsParticipantsListOutputToJSON(v *DocumentsParticipantsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsParticipantsListQueryCreatedAt - Only return results whose participant creation time is in this range. Set `gt`, `lt`, or both.
type DocumentsParticipantsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsParticipantsListQueryUpdatedAt - Only return results whose participant last update time is in this range. Set `gt`, `lt`, or both.
type DocumentsParticipantsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsParticipantsListQuery represents the documents participants list query type.
type DocumentsParticipantsListQuery struct {
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
	// Id - Only return participants with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CreatedAt - Only return results whose participant creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *DocumentsParticipantsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose participant last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *DocumentsParticipantsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapDocumentsParticipantsListQueryFromJSON deserializes JSON data into a DocumentsParticipantsListQuery.
func MapDocumentsParticipantsListQueryFromJSON(data []byte) (*DocumentsParticipantsListQuery, error) {
	var v DocumentsParticipantsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsParticipantsListQueryToJSON serializes a DocumentsParticipantsListQuery to JSON.
func MapDocumentsParticipantsListQueryToJSON(v *DocumentsParticipantsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
