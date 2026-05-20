package participants

import (
	"encoding/json"
	"time"
)

// DocumentsParticipantsListOutputItemsActorOrganizationActorTeams - The teams the actor belongs to
type DocumentsParticipantsListOutputItemsActorOrganizationActorTeams struct {
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

// DocumentsParticipantsListOutputItemsActorOrganizationActor represents the documents participants list output items actor organization actor type.
type DocumentsParticipantsListOutputItemsActorOrganizationActor struct {
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
	ImageUrl string                                                            `json:"image_url"`
	Teams    []DocumentsParticipantsListOutputItemsActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsParticipantsListOutputItemsActorConsumer represents the documents participants list output items actor consumer type.
type DocumentsParticipantsListOutputItemsActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsParticipantsListOutputItemsActor represents the documents participants list output items actor type.
type DocumentsParticipantsListOutputItemsActor struct {
	Type              string                                                      `json:"type"`
	Name              string                                                      `json:"name"`
	ImageUrl          *string                                                     `json:"image_url,omitempty"`
	Email             *string                                                     `json:"email,omitempty"`
	OrganizationActor *DocumentsParticipantsListOutputItemsActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsParticipantsListOutputItemsActorConsumer          `json:"consumer,omitempty"`
}

// DocumentsParticipantsListOutputItems represents the documents participants list output items type.
type DocumentsParticipantsListOutputItems struct {
	// Object - String representing the object's type
	Object       string                                    `json:"object"`
	Id           string                                    `json:"id"`
	Role         string                                    `json:"role"`
	EditCount    float64                                   `json:"edit_count"`
	LastEditedAt *time.Time                                `json:"last_edited_at,omitempty"`
	LastViewedAt *time.Time                                `json:"last_viewed_at,omitempty"`
	Actor        DocumentsParticipantsListOutputItemsActor `json:"actor"`
	CreatedAt    time.Time                                 `json:"created_at"`
}

// DocumentsParticipantsListOutputPagination represents the documents participants list output pagination type.
type DocumentsParticipantsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// DocumentsParticipantsListOutput represents the documents participants list output type.
type DocumentsParticipantsListOutput struct {
	Items      []DocumentsParticipantsListOutputItems    `json:"items"`
	Pagination DocumentsParticipantsListOutputPagination `json:"pagination"`
}

// MapDocumentsParticipantsListOutputFromJSON deserializes JSON data into a DocumentsParticipantsListOutput.
func MapDocumentsParticipantsListOutputFromJSON(data []byte) (*DocumentsParticipantsListOutput, error) {
	var v DocumentsParticipantsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsParticipantsListOutputToJSON serializes a DocumentsParticipantsListOutput to JSON.
func MapDocumentsParticipantsListOutputToJSON(v *DocumentsParticipantsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsParticipantsListQueryCreatedAt - Filter Filter by creation time by date range
type DocumentsParticipantsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for Filter by creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsParticipantsListQueryUpdatedAt - Filter Filter by update time by date range
type DocumentsParticipantsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for Filter by update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by update time
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsParticipantsListQuery represents the documents participants list query type.
type DocumentsParticipantsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by document participant ID
	Id *any `json:"id,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
	CreatedAt *DocumentsParticipantsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter Filter by update time by date range
	UpdatedAt *DocumentsParticipantsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapDocumentsParticipantsListQueryFromJSON deserializes JSON data into a DocumentsParticipantsListQuery.
func MapDocumentsParticipantsListQueryFromJSON(data []byte) (*DocumentsParticipantsListQuery, error) {
	var v DocumentsParticipantsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsParticipantsListQueryToJSON serializes a DocumentsParticipantsListQuery to JSON.
func MapDocumentsParticipantsListQueryToJSON(v *DocumentsParticipantsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
