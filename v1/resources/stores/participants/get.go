package participants

import (
	"encoding/json"
	"time"
)

// StoresParticipantsGetOutputActorOrganizationActorMember represents the stores participants get output actor organization actor member type.
type StoresParticipantsGetOutputActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// StoresParticipantsGetOutputActorOrganizationActorTeams - The teams the actor belongs to
type StoresParticipantsGetOutputActorOrganizationActorTeams struct {
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

// StoresParticipantsGetOutputActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type StoresParticipantsGetOutputActorOrganizationActor struct {
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
	ImageUrl string                                                   `json:"image_url"`
	Member   *StoresParticipantsGetOutputActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []StoresParticipantsGetOutputActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresParticipantsGetOutputActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type StoresParticipantsGetOutputActorConsumer struct {
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

// StoresParticipantsGetOutputActor - Actor behind a participant.
type StoresParticipantsGetOutputActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *StoresParticipantsGetOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *StoresParticipantsGetOutputActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// StoresParticipantsGetOutputSurfacesActorOrganizationActorMember represents the stores participants get output surfaces actor organization actor member type.
type StoresParticipantsGetOutputSurfacesActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// StoresParticipantsGetOutputSurfacesActorOrganizationActorTeams - The teams the actor belongs to
type StoresParticipantsGetOutputSurfacesActorOrganizationActorTeams struct {
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

// StoresParticipantsGetOutputSurfacesActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type StoresParticipantsGetOutputSurfacesActorOrganizationActor struct {
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
	ImageUrl string                                                           `json:"image_url"`
	Member   *StoresParticipantsGetOutputSurfacesActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []StoresParticipantsGetOutputSurfacesActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresParticipantsGetOutputSurfacesActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type StoresParticipantsGetOutputSurfacesActorConsumer struct {
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

// StoresParticipantsGetOutputSurfacesActor - Actor behind a participant.
type StoresParticipantsGetOutputSurfacesActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *StoresParticipantsGetOutputSurfacesActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *StoresParticipantsGetOutputSurfacesActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// StoresParticipantsGetOutputSurfaces - Identity linked to a store participant.
type StoresParticipantsGetOutputSurfaces struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Roles - Roles the participant holds.
	Roles []string `json:"roles"`
	// Actor - Actor behind a participant.
	Actor StoresParticipantsGetOutputSurfacesActor `json:"actor"`
	// LastActiveAt - When this identity was last active.
	LastActiveAt *time.Time `json:"last_active_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// StoresParticipantsGetOutput represents the stores participants get output type.
type StoresParticipantsGetOutput struct {
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
	Actor StoresParticipantsGetOutputActor `json:"actor"`
	// Surfaces - Other identities linked to this participant.
	Surfaces *[]StoresParticipantsGetOutputSurfaces `json:"surfaces,omitempty"`
	// LastActiveAt - When the participant last acted in the store.
	LastActiveAt *time.Time `json:"last_active_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapStoresParticipantsGetOutputFromJSON deserializes JSON data into a StoresParticipantsGetOutput.
func MapStoresParticipantsGetOutputFromJSON(data []byte) (*StoresParticipantsGetOutput, error) {
	var v StoresParticipantsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresParticipantsGetOutputToJSON serializes a StoresParticipantsGetOutput to JSON.
func MapStoresParticipantsGetOutputToJSON(v *StoresParticipantsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
