package participants

import (
	"encoding/json"
	"time"
)

// StoresParticipantsListOutputItemsActorOrganizationActorMember represents the stores participants list output items actor organization actor member type.
type StoresParticipantsListOutputItemsActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// StoresParticipantsListOutputItemsActorOrganizationActorTeams - The teams the actor belongs to
type StoresParticipantsListOutputItemsActorOrganizationActorTeams struct {
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

// StoresParticipantsListOutputItemsActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type StoresParticipantsListOutputItemsActorOrganizationActor struct {
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
	Member   *StoresParticipantsListOutputItemsActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []StoresParticipantsListOutputItemsActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresParticipantsListOutputItemsActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type StoresParticipantsListOutputItemsActorConsumer struct {
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

// StoresParticipantsListOutputItemsActor - Actor behind a participant.
type StoresParticipantsListOutputItemsActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *StoresParticipantsListOutputItemsActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *StoresParticipantsListOutputItemsActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// StoresParticipantsListOutputItemsSurfacesActorOrganizationActorMember represents the stores participants list output items surfaces actor organization actor member type.
type StoresParticipantsListOutputItemsSurfacesActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// StoresParticipantsListOutputItemsSurfacesActorOrganizationActorTeams - The teams the actor belongs to
type StoresParticipantsListOutputItemsSurfacesActorOrganizationActorTeams struct {
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

// StoresParticipantsListOutputItemsSurfacesActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type StoresParticipantsListOutputItemsSurfacesActorOrganizationActor struct {
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
	Member   *StoresParticipantsListOutputItemsSurfacesActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []StoresParticipantsListOutputItemsSurfacesActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresParticipantsListOutputItemsSurfacesActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type StoresParticipantsListOutputItemsSurfacesActorConsumer struct {
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

// StoresParticipantsListOutputItemsSurfacesActor - Actor behind a participant.
type StoresParticipantsListOutputItemsSurfacesActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *StoresParticipantsListOutputItemsSurfacesActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *StoresParticipantsListOutputItemsSurfacesActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// StoresParticipantsListOutputItemsSurfaces - Identity linked to a store participant.
type StoresParticipantsListOutputItemsSurfaces struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Roles - Roles the participant holds.
	Roles []string `json:"roles"`
	// Actor - Actor behind a participant.
	Actor StoresParticipantsListOutputItemsSurfacesActor `json:"actor"`
	// LastActiveAt - When this identity was last active.
	LastActiveAt *time.Time `json:"last_active_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// StoresParticipantsListOutputItems represents the stores participants list output items type.
type StoresParticipantsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// StoreId - Store the participant belongs to.
	StoreId string `json:"store_id"`
	// Permissions - What the participant can read or write in the store.
	Permissions []string `json:"permissions"`
	// Roles - Roles the participant holds.
	Roles []string `json:"roles"`
	// Actor - Actor behind a participant.
	Actor StoresParticipantsListOutputItemsActor `json:"actor"`
	// Surfaces - Other identities linked to this participant.
	Surfaces *[]StoresParticipantsListOutputItemsSurfaces `json:"surfaces,omitempty"`
	// LastActiveAt - When the participant last acted in the store.
	LastActiveAt *time.Time `json:"last_active_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresParticipantsListOutputPagination represents the stores participants list output pagination type.
type StoresParticipantsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// StoresParticipantsListOutput represents the stores participants list output type.
type StoresParticipantsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                 `json:"object"`
	Items      []StoresParticipantsListOutputItems    `json:"items"`
	Pagination StoresParticipantsListOutputPagination `json:"pagination"`
}

// MapStoresParticipantsListOutputFromJSON deserializes JSON data into a StoresParticipantsListOutput.
func MapStoresParticipantsListOutputFromJSON(data []byte) (*StoresParticipantsListOutput, error) {
	var v StoresParticipantsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresParticipantsListOutputToJSON serializes a StoresParticipantsListOutput to JSON.
func MapStoresParticipantsListOutputToJSON(v *StoresParticipantsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// StoresParticipantsListQuery represents the stores participants list query type.
type StoresParticipantsListQuery struct {
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

// MapStoresParticipantsListQueryFromJSON deserializes JSON data into a StoresParticipantsListQuery.
func MapStoresParticipantsListQueryFromJSON(data []byte) (*StoresParticipantsListQuery, error) {
	var v StoresParticipantsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresParticipantsListQueryToJSON serializes a StoresParticipantsListQuery to JSON.
func MapStoresParticipantsListQueryToJSON(v *StoresParticipantsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
