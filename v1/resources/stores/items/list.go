package items

import (
	"encoding/json"
	"time"
)

// StoresItemsListOutputItemsFileCreatedByOrganizationActorTeams - The teams the actor belongs to
type StoresItemsListOutputItemsFileCreatedByOrganizationActorTeams struct {
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

// StoresItemsListOutputItemsFileCreatedByOrganizationActor represents the stores items list output items file created by organization actor type.
type StoresItemsListOutputItemsFileCreatedByOrganizationActor struct {
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
	Teams    []StoresItemsListOutputItemsFileCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsListOutputItemsFileCreatedByConsumer represents the stores items list output items file created by consumer type.
type StoresItemsListOutputItemsFileCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsListOutputItemsFileCreatedBy represents the stores items list output items file created by type.
type StoresItemsListOutputItemsFileCreatedBy struct {
	Type              string                                                    `json:"type"`
	Name              string                                                    `json:"name"`
	ImageUrl          *string                                                   `json:"image_url,omitempty"`
	Email             *string                                                   `json:"email,omitempty"`
	OrganizationActor *StoresItemsListOutputItemsFileCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *StoresItemsListOutputItemsFileCreatedByConsumer          `json:"consumer,omitempty"`
}

// StoresItemsListOutputItemsFile represents the stores items list output items file type.
type StoresItemsListOutputItemsFile struct {
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
	Purpose   string                                   `json:"purpose"`
	CreatedBy *StoresItemsListOutputItemsFileCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - The files's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The files's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsListOutputItemsDocumentCreatedByOrganizationActorTeams - The teams the actor belongs to
type StoresItemsListOutputItemsDocumentCreatedByOrganizationActorTeams struct {
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

// StoresItemsListOutputItemsDocumentCreatedByOrganizationActor represents the stores items list output items document created by organization actor type.
type StoresItemsListOutputItemsDocumentCreatedByOrganizationActor struct {
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
	ImageUrl string                                                              `json:"image_url"`
	Teams    []StoresItemsListOutputItemsDocumentCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsListOutputItemsDocumentCreatedByConsumer represents the stores items list output items document created by consumer type.
type StoresItemsListOutputItemsDocumentCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsListOutputItemsDocumentCreatedBy represents the stores items list output items document created by type.
type StoresItemsListOutputItemsDocumentCreatedBy struct {
	Type              string                                                        `json:"type"`
	Name              string                                                        `json:"name"`
	ImageUrl          *string                                                       `json:"image_url,omitempty"`
	Email             *string                                                       `json:"email,omitempty"`
	OrganizationActor *StoresItemsListOutputItemsDocumentCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *StoresItemsListOutputItemsDocumentCreatedByConsumer          `json:"consumer,omitempty"`
}

// StoresItemsListOutputItemsDocument represents the stores items list output items document type.
type StoresItemsListOutputItemsDocument struct {
	// Object - String representing the object's type
	Object           string                                       `json:"object"`
	Id               string                                       `json:"id"`
	Status           string                                       `json:"status"`
	Title            string                                       `json:"title"`
	Content          string                                       `json:"content"`
	FileId           string                                       `json:"file_id"`
	ParentDocumentId *string                                      `json:"parent_document_id,omitempty"`
	CurrentVersionId *string                                      `json:"current_version_id,omitempty"`
	CreatedBy        *StoresItemsListOutputItemsDocumentCreatedBy `json:"created_by,omitempty"`
	CreatedAt        time.Time                                    `json:"created_at"`
	UpdatedAt        time.Time                                    `json:"updated_at"`
}

// StoresItemsListOutputItems represents the stores items list output items type.
type StoresItemsListOutputItems struct {
	// Object - String representing the object's type
	Object      string                              `json:"object"`
	Id          string                              `json:"id"`
	Kind        string                              `json:"kind"`
	Path        string                              `json:"path"`
	StoreId     string                              `json:"store_id"`
	DirectoryId *string                             `json:"directory_id,omitempty"`
	File        *StoresItemsListOutputItemsFile     `json:"file,omitempty"`
	Document    *StoresItemsListOutputItemsDocument `json:"document,omitempty"`
	CreatedAt   time.Time                           `json:"created_at"`
	UpdatedAt   time.Time                           `json:"updated_at"`
}

// StoresItemsListOutputPagination represents the stores items list output pagination type.
type StoresItemsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// StoresItemsListOutput represents the stores items list output type.
type StoresItemsListOutput struct {
	Items      []StoresItemsListOutputItems    `json:"items"`
	Pagination StoresItemsListOutputPagination `json:"pagination"`
}

// MapStoresItemsListOutputFromJSON deserializes JSON data into a StoresItemsListOutput.
func MapStoresItemsListOutputFromJSON(data []byte) (*StoresItemsListOutput, error) {
	var v StoresItemsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresItemsListOutputToJSON serializes a StoresItemsListOutput to JSON.
func MapStoresItemsListOutputToJSON(v *StoresItemsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// StoresItemsListQueryCreatedAt - Filter Filter by creation time by date range
type StoresItemsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for Filter by creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// StoresItemsListQueryUpdatedAt - Filter Filter by update time by date range
type StoresItemsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for Filter by update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by update time
	Lt *time.Time `json:"lt,omitempty"`
}

// StoresItemsListQuery represents the stores items list query type.
type StoresItemsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by store item ID
	Id *any `json:"id,omitempty"`
	// FileId - Filter by file ID
	FileId *any `json:"file_id,omitempty"`
	// DocumentId - Filter by document ID
	DocumentId *any `json:"document_id,omitempty"`
	// Type - Filter by store item type. Repeat `type` to include multiple values. Defaults to `file` and `document`.
	Type *any `json:"type,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
	CreatedAt *StoresItemsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter Filter by update time by date range
	UpdatedAt *StoresItemsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapStoresItemsListQueryFromJSON deserializes JSON data into a StoresItemsListQuery.
func MapStoresItemsListQueryFromJSON(data []byte) (*StoresItemsListQuery, error) {
	var v StoresItemsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresItemsListQueryToJSON serializes a StoresItemsListQuery to JSON.
func MapStoresItemsListQueryToJSON(v *StoresItemsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
