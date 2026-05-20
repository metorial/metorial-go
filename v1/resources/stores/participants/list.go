package participants

import (
	"encoding/json"
	"time"
)

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

// StoresParticipantsListOutputItemsActorOrganizationActor represents the stores participants list output items actor organization actor type.
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
	Teams    []StoresParticipantsListOutputItemsActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresParticipantsListOutputItemsActorConsumer represents the stores participants list output items actor consumer type.
type StoresParticipantsListOutputItemsActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresParticipantsListOutputItemsActor represents the stores participants list output items actor type.
type StoresParticipantsListOutputItemsActor struct {
	Type              string                                                   `json:"type"`
	Name              string                                                   `json:"name"`
	ImageUrl          *string                                                  `json:"image_url,omitempty"`
	Email             *string                                                  `json:"email,omitempty"`
	OrganizationActor *StoresParticipantsListOutputItemsActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *StoresParticipantsListOutputItemsActorConsumer          `json:"consumer,omitempty"`
}

// StoresParticipantsListOutputItems represents the stores participants list output items type.
type StoresParticipantsListOutputItems struct {
	// Object - String representing the object's type
	Object      string                                 `json:"object"`
	Id          string                                 `json:"id"`
	StoreId     string                                 `json:"store_id"`
	Permissions []string                               `json:"permissions"`
	Actor       StoresParticipantsListOutputItemsActor `json:"actor"`
	CreatedAt   time.Time                              `json:"created_at"`
}

// StoresParticipantsListOutputPagination represents the stores participants list output pagination type.
type StoresParticipantsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// StoresParticipantsListOutput represents the stores participants list output type.
type StoresParticipantsListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
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
