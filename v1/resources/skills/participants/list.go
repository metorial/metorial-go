package participants

import (
	"encoding/json"
	"time"
)

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

// SkillsParticipantsListOutputItemsActorOrganizationActor represents the skills participants list output items actor organization actor type.
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
	Teams    []SkillsParticipantsListOutputItemsActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsParticipantsListOutputItemsActorConsumer represents the skills participants list output items actor consumer type.
type SkillsParticipantsListOutputItemsActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsParticipantsListOutputItemsActor represents the skills participants list output items actor type.
type SkillsParticipantsListOutputItemsActor struct {
	Type              string                                                   `json:"type"`
	Name              string                                                   `json:"name"`
	ImageUrl          *string                                                  `json:"image_url,omitempty"`
	Email             *string                                                  `json:"email,omitempty"`
	OrganizationActor *SkillsParticipantsListOutputItemsActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsParticipantsListOutputItemsActorConsumer          `json:"consumer,omitempty"`
}

// SkillsParticipantsListOutputItems represents the skills participants list output items type.
type SkillsParticipantsListOutputItems struct {
	// Object - String representing the object's type
	Object    string                                 `json:"object"`
	Id        string                                 `json:"id"`
	SkillId   string                                 `json:"skill_id"`
	Roles     []string                               `json:"roles"`
	Actor     SkillsParticipantsListOutputItemsActor `json:"actor"`
	CreatedAt time.Time                              `json:"created_at"`
	UpdatedAt time.Time                              `json:"updated_at"`
}

// SkillsParticipantsListOutputPagination represents the skills participants list output pagination type.
type SkillsParticipantsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsParticipantsListOutput represents the skills participants list output type.
type SkillsParticipantsListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
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
