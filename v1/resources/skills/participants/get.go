package participants

import (
	"encoding/json"
	"time"
)

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

// SkillsParticipantsGetOutputActorOrganizationActor represents the skills participants get output actor organization actor type.
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
	Teams    []SkillsParticipantsGetOutputActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsParticipantsGetOutputActorConsumer represents the skills participants get output actor consumer type.
type SkillsParticipantsGetOutputActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsParticipantsGetOutputActor represents the skills participants get output actor type.
type SkillsParticipantsGetOutputActor struct {
	Type              string                                             `json:"type"`
	Name              string                                             `json:"name"`
	ImageUrl          *string                                            `json:"image_url,omitempty"`
	Email             *string                                            `json:"email,omitempty"`
	OrganizationActor *SkillsParticipantsGetOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsParticipantsGetOutputActorConsumer          `json:"consumer,omitempty"`
}

// SkillsParticipantsGetOutput represents the skills participants get output type.
type SkillsParticipantsGetOutput struct {
	// Object - String representing the object's type
	Object    string                           `json:"object"`
	Id        string                           `json:"id"`
	SkillId   string                           `json:"skill_id"`
	Roles     []string                         `json:"roles"`
	Actor     SkillsParticipantsGetOutputActor `json:"actor"`
	CreatedAt time.Time                        `json:"created_at"`
	UpdatedAt time.Time                        `json:"updated_at"`
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
