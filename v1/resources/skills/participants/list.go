package participants

import (
	"encoding/json"
	"time"
)

// SkillsParticipantsListOutputItemsActorOrganizationActorMember represents the skills participants list output items actor organization actor member type.
type SkillsParticipantsListOutputItemsActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsParticipantsListOutputItemsActorOrganizationActorTeams - The teams the actor belongs to
type SkillsParticipantsListOutputItemsActorOrganizationActorTeams struct {
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

// SkillsParticipantsListOutputItemsActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsParticipantsListOutputItemsActorOrganizationActor struct {
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
	ImageUrl string                                                         `json:"image_url"`
	Member   *SkillsParticipantsListOutputItemsActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsParticipantsListOutputItemsActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsParticipantsListOutputItemsActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsParticipantsListOutputItemsActorConsumer struct {
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

// SkillsParticipantsListOutputItemsActor - Actor behind a participant.
type SkillsParticipantsListOutputItemsActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsParticipantsListOutputItemsActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsParticipantsListOutputItemsActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsParticipantsListOutputItemsSurfacesActorOrganizationActorMember represents the skills participants list output items surfaces actor organization actor member type.
type SkillsParticipantsListOutputItemsSurfacesActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsParticipantsListOutputItemsSurfacesActorOrganizationActorTeams - The teams the actor belongs to
type SkillsParticipantsListOutputItemsSurfacesActorOrganizationActorTeams struct {
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

// SkillsParticipantsListOutputItemsSurfacesActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsParticipantsListOutputItemsSurfacesActorOrganizationActor struct {
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
	ImageUrl string                                                                 `json:"image_url"`
	Member   *SkillsParticipantsListOutputItemsSurfacesActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsParticipantsListOutputItemsSurfacesActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsParticipantsListOutputItemsSurfacesActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsParticipantsListOutputItemsSurfacesActorConsumer struct {
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

// SkillsParticipantsListOutputItemsSurfacesActor - Actor behind a participant.
type SkillsParticipantsListOutputItemsSurfacesActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsParticipantsListOutputItemsSurfacesActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsParticipantsListOutputItemsSurfacesActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsParticipantsListOutputItemsSurfaces - Identity linked to a store participant.
type SkillsParticipantsListOutputItemsSurfaces struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Roles - Roles the participant holds.
	Roles []string `json:"roles"`
	// Actor - Actor behind a participant.
	Actor SkillsParticipantsListOutputItemsSurfacesActor `json:"actor"`
	// LastActiveAt - When this identity was last active.
	LastActiveAt *time.Time `json:"last_active_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SkillsParticipantsListOutputItems represents the skills participants list output items type.
type SkillsParticipantsListOutputItems struct {
	// Object - String representing the object's type
	Object  string   `json:"object"`
	Id      string   `json:"id"`
	SkillId string   `json:"skill_id"`
	Roles   []string `json:"roles"`
	// Actor - Actor behind a participant.
	Actor        SkillsParticipantsListOutputItemsActor       `json:"actor"`
	Surfaces     *[]SkillsParticipantsListOutputItemsSurfaces `json:"surfaces,omitempty"`
	LastActiveAt *time.Time                                   `json:"last_active_at,omitempty"`
	CreatedAt    time.Time                                    `json:"created_at"`
	UpdatedAt    time.Time                                    `json:"updated_at"`
}

// SkillsParticipantsListOutputPagination represents the skills participants list output pagination type.
type SkillsParticipantsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsParticipantsListOutput represents the skills participants list output type.
type SkillsParticipantsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                 `json:"object"`
	Items      []SkillsParticipantsListOutputItems    `json:"items"`
	Pagination SkillsParticipantsListOutputPagination `json:"pagination"`
}

// MapSkillsParticipantsListOutputFromJSON deserializes JSON data into a SkillsParticipantsListOutput.
func MapSkillsParticipantsListOutputFromJSON(data []byte) (*SkillsParticipantsListOutput, error) {
	var v SkillsParticipantsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsParticipantsListOutputToJSON serializes a SkillsParticipantsListOutput to JSON.
func MapSkillsParticipantsListOutputToJSON(v *SkillsParticipantsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsParticipantsListQuery represents the skills participants list query type.
type SkillsParticipantsListQuery struct {
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
}

// MapSkillsParticipantsListQueryFromJSON deserializes JSON data into a SkillsParticipantsListQuery.
func MapSkillsParticipantsListQueryFromJSON(data []byte) (*SkillsParticipantsListQuery, error) {
	var v SkillsParticipantsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsParticipantsListQueryToJSON serializes a SkillsParticipantsListQuery to JSON.
func MapSkillsParticipantsListQueryToJSON(v *SkillsParticipantsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
