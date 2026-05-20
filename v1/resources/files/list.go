package files

import (
	"encoding/json"
	"time"
)

// FilesListOutputItemsCreatedByOrganizationActorTeams - The teams the actor belongs to
type FilesListOutputItemsCreatedByOrganizationActorTeams struct {
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

// FilesListOutputItemsCreatedByOrganizationActor represents the files list output items created by organization actor type.
type FilesListOutputItemsCreatedByOrganizationActor struct {
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
	ImageUrl string                                                `json:"image_url"`
	Teams    []FilesListOutputItemsCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// FilesListOutputItemsCreatedByConsumer represents the files list output items created by consumer type.
type FilesListOutputItemsCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FilesListOutputItemsCreatedBy represents the files list output items created by type.
type FilesListOutputItemsCreatedBy struct {
	Type              string                                          `json:"type"`
	Name              string                                          `json:"name"`
	ImageUrl          *string                                         `json:"image_url,omitempty"`
	Email             *string                                         `json:"email,omitempty"`
	OrganizationActor *FilesListOutputItemsCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *FilesListOutputItemsCreatedByConsumer          `json:"consumer,omitempty"`
}

// FilesListOutputItems represents the files list output items type.
type FilesListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The files's unique identifier
	Id string `json:"id"`
	// Status - The files's status
	Status string `json:"status"`
	// FileName - The file's name
	FileName string `json:"file_name"`
	// FileSize - The file's size in bytes
	FileSize float64 `json:"file_size"`
	// FileType - The file's MIME type
	FileType string `json:"file_type"`
	// Title - The file's title
	Title string `json:"title"`
	// Purpose - The file's purpose identifier
	Purpose   string                         `json:"purpose"`
	CreatedBy *FilesListOutputItemsCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - The files's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The files's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// FilesListOutputPagination represents the files list output pagination type.
type FilesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// FilesListOutput represents the files list output type.
type FilesListOutput struct {
	Items      []FilesListOutputItems    `json:"items"`
	Pagination FilesListOutputPagination `json:"pagination"`
}

// MapFilesListOutputFromJSON deserializes JSON data into a FilesListOutput.
func MapFilesListOutputFromJSON(data []byte) (*FilesListOutput, error) {
	var v FilesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFilesListOutputToJSON serializes a FilesListOutput to JSON.
func MapFilesListOutputToJSON(v *FilesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// FilesListQueryCreatedAt - Filter Filter by creation time by date range
type FilesListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for Filter by creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// FilesListQueryUpdatedAt - Filter Filter by update time by date range
type FilesListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for Filter by update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by update time
	Lt *time.Time `json:"lt,omitempty"`
}

// FilesListQuery represents the files list query type.
type FilesListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by file ID
	Id *any `json:"id,omitempty"`
	// Purpose - Filter by file purpose
	Purpose *any `json:"purpose,omitempty"`
	// StoreId - Filter by store ID
	StoreId *any `json:"store_id,omitempty"`
	// DocumentId - Filter by document ID
	DocumentId *any `json:"document_id,omitempty"`
	// FileLinkId - Filter by file link ID
	FileLinkId *any `json:"file_link_id,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
	CreatedAt *FilesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter Filter by update time by date range
	UpdatedAt *FilesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapFilesListQueryFromJSON deserializes JSON data into a FilesListQuery.
func MapFilesListQueryFromJSON(data []byte) (*FilesListQuery, error) {
	var v FilesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFilesListQueryToJSON serializes a FilesListQuery to JSON.
func MapFilesListQueryToJSON(v *FilesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
