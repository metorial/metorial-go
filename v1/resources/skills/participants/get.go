package participants

import (
	"encoding/json"
	"time"
)

// SkillsParticipantsGetOutputActorOrganizationActorMember represents the skills participants get output actor organization actor member type.
type SkillsParticipantsGetOutputActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsParticipantsGetOutputActorOrganizationActorTeams - The teams the actor belongs to
type SkillsParticipantsGetOutputActorOrganizationActorTeams struct {
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

// SkillsParticipantsGetOutputActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsParticipantsGetOutputActorOrganizationActor struct {
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
	Member   *SkillsParticipantsGetOutputActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsParticipantsGetOutputActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsParticipantsGetOutputActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsParticipantsGetOutputActorConsumer struct {
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

// SkillsParticipantsGetOutputActor - Actor behind a participant.
type SkillsParticipantsGetOutputActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsParticipantsGetOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsParticipantsGetOutputActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsParticipantsGetOutputSurfacesActorOrganizationActorMember represents the skills participants get output surfaces actor organization actor member type.
type SkillsParticipantsGetOutputSurfacesActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsParticipantsGetOutputSurfacesActorOrganizationActorTeams - The teams the actor belongs to
type SkillsParticipantsGetOutputSurfacesActorOrganizationActorTeams struct {
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

// SkillsParticipantsGetOutputSurfacesActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsParticipantsGetOutputSurfacesActorOrganizationActor struct {
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
	Member   *SkillsParticipantsGetOutputSurfacesActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsParticipantsGetOutputSurfacesActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsParticipantsGetOutputSurfacesActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsParticipantsGetOutputSurfacesActorConsumer struct {
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

// SkillsParticipantsGetOutputSurfacesActor - Actor behind a participant.
type SkillsParticipantsGetOutputSurfacesActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsParticipantsGetOutputSurfacesActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsParticipantsGetOutputSurfacesActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsParticipantsGetOutputSurfaces - Identity linked to a store participant.
type SkillsParticipantsGetOutputSurfaces struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Roles - Roles the participant holds.
	Roles []string `json:"roles"`
	// Actor - Actor behind a participant.
	Actor SkillsParticipantsGetOutputSurfacesActor `json:"actor"`
	// LastActiveAt - When this identity was last active.
	LastActiveAt *time.Time `json:"last_active_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SkillsParticipantsGetOutput represents the skills participants get output type.
type SkillsParticipantsGetOutput struct {
	// Object - String representing the object's type
	Object  string   `json:"object"`
	Id      string   `json:"id"`
	SkillId string   `json:"skill_id"`
	Roles   []string `json:"roles"`
	// Actor - Actor behind a participant.
	Actor        SkillsParticipantsGetOutputActor       `json:"actor"`
	Surfaces     *[]SkillsParticipantsGetOutputSurfaces `json:"surfaces,omitempty"`
	LastActiveAt *time.Time                             `json:"last_active_at,omitempty"`
	CreatedAt    time.Time                              `json:"created_at"`
	UpdatedAt    time.Time                              `json:"updated_at"`
}

// MapSkillsParticipantsGetOutputFromJSON deserializes JSON data into a SkillsParticipantsGetOutput.
func MapSkillsParticipantsGetOutputFromJSON(data []byte) (*SkillsParticipantsGetOutput, error) {
	var v SkillsParticipantsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsParticipantsGetOutputToJSON serializes a SkillsParticipantsGetOutput to JSON.
func MapSkillsParticipantsGetOutputToJSON(v *SkillsParticipantsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
