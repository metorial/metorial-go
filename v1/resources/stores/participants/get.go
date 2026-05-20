package participants

import (
	"encoding/json"
	"time"
)

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

// StoresParticipantsGetOutputActorOrganizationActor represents the stores participants get output actor organization actor type.
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
	Teams    []StoresParticipantsGetOutputActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresParticipantsGetOutputActorConsumer represents the stores participants get output actor consumer type.
type StoresParticipantsGetOutputActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresParticipantsGetOutputActor represents the stores participants get output actor type.
type StoresParticipantsGetOutputActor struct {
	Type              string                                             `json:"type"`
	Name              string                                             `json:"name"`
	ImageUrl          *string                                            `json:"image_url,omitempty"`
	Email             *string                                            `json:"email,omitempty"`
	OrganizationActor *StoresParticipantsGetOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *StoresParticipantsGetOutputActorConsumer          `json:"consumer,omitempty"`
}

// StoresParticipantsGetOutput represents the stores participants get output type.
type StoresParticipantsGetOutput struct {
	// Object - String representing the object's type
	Object      string                           `json:"object"`
	Id          string                           `json:"id"`
	StoreId     string                           `json:"store_id"`
	Permissions []string                         `json:"permissions"`
	Actor       StoresParticipantsGetOutputActor `json:"actor"`
	CreatedAt   time.Time                        `json:"created_at"`
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
