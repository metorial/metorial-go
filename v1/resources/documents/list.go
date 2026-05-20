package documents

import (
	"encoding/json"
	"time"
)

// DocumentsListOutputItemsCreatedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsListOutputItemsCreatedByOrganizationActorTeams struct {
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

// DocumentsListOutputItemsCreatedByOrganizationActor represents the documents list output items created by organization actor type.
type DocumentsListOutputItemsCreatedByOrganizationActor struct {
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
	ImageUrl string                                                    `json:"image_url"`
	Teams    []DocumentsListOutputItemsCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsListOutputItemsCreatedByConsumer represents the documents list output items created by consumer type.
type DocumentsListOutputItemsCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsListOutputItemsCreatedBy represents the documents list output items created by type.
type DocumentsListOutputItemsCreatedBy struct {
	Type              string                                              `json:"type"`
	Name              string                                              `json:"name"`
	ImageUrl          *string                                             `json:"image_url,omitempty"`
	Email             *string                                             `json:"email,omitempty"`
	OrganizationActor *DocumentsListOutputItemsCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsListOutputItemsCreatedByConsumer          `json:"consumer,omitempty"`
}

// DocumentsListOutputItems represents the documents list output items type.
type DocumentsListOutputItems struct {
	// Object - String representing the object's type
	Object           string                             `json:"object"`
	Id               string                             `json:"id"`
	Status           string                             `json:"status"`
	Title            string                             `json:"title"`
	Content          string                             `json:"content"`
	FileId           string                             `json:"file_id"`
	ParentDocumentId *string                            `json:"parent_document_id,omitempty"`
	CurrentVersionId *string                            `json:"current_version_id,omitempty"`
	CreatedBy        *DocumentsListOutputItemsCreatedBy `json:"created_by,omitempty"`
	CreatedAt        time.Time                          `json:"created_at"`
	UpdatedAt        time.Time                          `json:"updated_at"`
}

// DocumentsListOutputPagination represents the documents list output pagination type.
type DocumentsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// DocumentsListOutput represents the documents list output type.
type DocumentsListOutput struct {
	Items      []DocumentsListOutputItems    `json:"items"`
	Pagination DocumentsListOutputPagination `json:"pagination"`
}

// MapDocumentsListOutputFromJSON deserializes JSON data into a DocumentsListOutput.
func MapDocumentsListOutputFromJSON(data []byte) (*DocumentsListOutput, error) {
	var v DocumentsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsListOutputToJSON serializes a DocumentsListOutput to JSON.
func MapDocumentsListOutputToJSON(v *DocumentsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsListQueryCreatedAt - Filter Filter by creation time by date range
type DocumentsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for Filter by creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsListQueryUpdatedAt - Filter Filter by update time by date range
type DocumentsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for Filter by update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by update time
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsListQuery represents the documents list query type.
type DocumentsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by document ID
	Id *any `json:"id,omitempty"`
	// FileId - Filter by file ID
	FileId *any `json:"file_id,omitempty"`
	// StoreId - Filter by store ID
	StoreId *any `json:"store_id,omitempty"`
	// ParentDocumentId - Filter by parent document ID
	ParentDocumentId *any `json:"parent_document_id,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
	CreatedAt *DocumentsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter Filter by update time by date range
	UpdatedAt *DocumentsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapDocumentsListQueryFromJSON deserializes JSON data into a DocumentsListQuery.
func MapDocumentsListQueryFromJSON(data []byte) (*DocumentsListQuery, error) {
	var v DocumentsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsListQueryToJSON serializes a DocumentsListQuery to JSON.
func MapDocumentsListQueryToJSON(v *DocumentsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
