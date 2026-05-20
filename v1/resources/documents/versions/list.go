package versions

import (
	"encoding/json"
	"time"
)

// DocumentsVersionsListOutputItemsEditorsOrganizationActorTeams - The teams the actor belongs to
type DocumentsVersionsListOutputItemsEditorsOrganizationActorTeams struct {
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

// DocumentsVersionsListOutputItemsEditorsOrganizationActor represents the documents versions list output items editors organization actor type.
type DocumentsVersionsListOutputItemsEditorsOrganizationActor struct {
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
	ImageUrl string                                                          `json:"image_url"`
	Teams    []DocumentsVersionsListOutputItemsEditorsOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsVersionsListOutputItemsEditorsConsumer represents the documents versions list output items editors consumer type.
type DocumentsVersionsListOutputItemsEditorsConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsVersionsListOutputItemsEditors represents the documents versions list output items editors type.
type DocumentsVersionsListOutputItemsEditors struct {
	Type              string                                                    `json:"type"`
	Name              string                                                    `json:"name"`
	ImageUrl          *string                                                   `json:"image_url,omitempty"`
	Email             *string                                                   `json:"email,omitempty"`
	OrganizationActor *DocumentsVersionsListOutputItemsEditorsOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsVersionsListOutputItemsEditorsConsumer          `json:"consumer,omitempty"`
}

// DocumentsVersionsListOutputItems represents the documents versions list output items type.
type DocumentsVersionsListOutputItems struct {
	// Object - String representing the object's type
	Object            string                                    `json:"object"`
	Id                string                                    `json:"id"`
	DocumentId        string                                    `json:"document_id"`
	VersionNumber     float64                                   `json:"version_number"`
	PreviousVersionId *string                                   `json:"previous_version_id,omitempty"`
	ListEditedAt      *time.Time                                `json:"list_edited_at,omitempty"`
	Content           string                                    `json:"content"`
	Editors           []DocumentsVersionsListOutputItemsEditors `json:"editors"`
	CreatedAt         time.Time                                 `json:"created_at"`
}

// DocumentsVersionsListOutputPagination represents the documents versions list output pagination type.
type DocumentsVersionsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// DocumentsVersionsListOutput represents the documents versions list output type.
type DocumentsVersionsListOutput struct {
	Items      []DocumentsVersionsListOutputItems    `json:"items"`
	Pagination DocumentsVersionsListOutputPagination `json:"pagination"`
}

// MapDocumentsVersionsListOutputFromJSON deserializes JSON data into a DocumentsVersionsListOutput.
func MapDocumentsVersionsListOutputFromJSON(data []byte) (*DocumentsVersionsListOutput, error) {
	var v DocumentsVersionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsVersionsListOutputToJSON serializes a DocumentsVersionsListOutput to JSON.
func MapDocumentsVersionsListOutputToJSON(v *DocumentsVersionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsVersionsListQueryCreatedAt - Filter Filter by creation time by date range
type DocumentsVersionsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for Filter by creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsVersionsListQueryLastEditedAt - Filter Filter by last edit time by date range
type DocumentsVersionsListQueryLastEditedAt struct {
	// Gt - Only include records after this timestamp for Filter by last edit time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by last edit time
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsVersionsListQuery represents the documents versions list query type.
type DocumentsVersionsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by document version ID
	Id *any `json:"id,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
	CreatedAt *DocumentsVersionsListQueryCreatedAt `json:"created_at,omitempty"`
	// LastEditedAt - Filter Filter by last edit time by date range
	LastEditedAt *DocumentsVersionsListQueryLastEditedAt `json:"last_edited_at,omitempty"`
}

// MapDocumentsVersionsListQueryFromJSON deserializes JSON data into a DocumentsVersionsListQuery.
func MapDocumentsVersionsListQueryFromJSON(data []byte) (*DocumentsVersionsListQuery, error) {
	var v DocumentsVersionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsVersionsListQueryToJSON serializes a DocumentsVersionsListQuery to JSON.
func MapDocumentsVersionsListQueryToJSON(v *DocumentsVersionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
