package participants

import (
	"encoding/json"
	"time"
)

// DocumentsParticipantsGetOutputActorOrganizationActorMember represents the documents participants get output actor organization actor member type.
type DocumentsParticipantsGetOutputActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsParticipantsGetOutputActorOrganizationActorTeams - The teams the actor belongs to
type DocumentsParticipantsGetOutputActorOrganizationActorTeams struct {
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

// DocumentsParticipantsGetOutputActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type DocumentsParticipantsGetOutputActorOrganizationActor struct {
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
	ImageUrl string                                                      `json:"image_url"`
	Member   *DocumentsParticipantsGetOutputActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsParticipantsGetOutputActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsParticipantsGetOutputActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsParticipantsGetOutputActorConsumer struct {
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

// DocumentsParticipantsGetOutputActor - Actor behind a participant.
type DocumentsParticipantsGetOutputActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsParticipantsGetOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsParticipantsGetOutputActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsParticipantsGetOutput represents the documents participants get output type.
type DocumentsParticipantsGetOutput struct {
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
	Actor DocumentsParticipantsGetOutputActor `json:"actor"`
	// StoreParticipantId - Linked store participant, by ID.
	StoreParticipantId *string `json:"store_participant_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MapDocumentsParticipantsGetOutputFromJSON deserializes JSON data into a DocumentsParticipantsGetOutput.
func MapDocumentsParticipantsGetOutputFromJSON(data []byte) (*DocumentsParticipantsGetOutput, error) {
	var v DocumentsParticipantsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsParticipantsGetOutputToJSON serializes a DocumentsParticipantsGetOutput to JSON.
func MapDocumentsParticipantsGetOutputToJSON(v *DocumentsParticipantsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
