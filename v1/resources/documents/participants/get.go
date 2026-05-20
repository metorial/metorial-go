package participants

import (
	"encoding/json"
	"time"
)

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

// DocumentsParticipantsGetOutputActorOrganizationActor represents the documents participants get output actor organization actor type.
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
	Teams    []DocumentsParticipantsGetOutputActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsParticipantsGetOutputActorConsumer represents the documents participants get output actor consumer type.
type DocumentsParticipantsGetOutputActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsParticipantsGetOutputActor represents the documents participants get output actor type.
type DocumentsParticipantsGetOutputActor struct {
	Type              string                                                `json:"type"`
	Name              string                                                `json:"name"`
	ImageUrl          *string                                               `json:"image_url,omitempty"`
	Email             *string                                               `json:"email,omitempty"`
	OrganizationActor *DocumentsParticipantsGetOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsParticipantsGetOutputActorConsumer          `json:"consumer,omitempty"`
}

// DocumentsParticipantsGetOutput represents the documents participants get output type.
type DocumentsParticipantsGetOutput struct {
	// Object - String representing the object's type
	Object       string                              `json:"object"`
	Id           string                              `json:"id"`
	Role         string                              `json:"role"`
	EditCount    float64                             `json:"edit_count"`
	LastEditedAt *time.Time                          `json:"last_edited_at,omitempty"`
	LastViewedAt *time.Time                          `json:"last_viewed_at,omitempty"`
	Actor        DocumentsParticipantsGetOutputActor `json:"actor"`
	CreatedAt    time.Time                           `json:"created_at"`
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
