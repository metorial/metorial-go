package exports

import (
	"encoding/json"
	"time"
)

// SkillsExportsListOutputItemsFileCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsExportsListOutputItemsFileCreatedByOrganizationActorTeams struct {
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

// SkillsExportsListOutputItemsFileCreatedByOrganizationActor represents the skills exports list output items file created by organization actor type.
type SkillsExportsListOutputItemsFileCreatedByOrganizationActor struct {
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
	Teams    []SkillsExportsListOutputItemsFileCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsListOutputItemsFileCreatedByConsumer represents the skills exports list output items file created by consumer type.
type SkillsExportsListOutputItemsFileCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsListOutputItemsFileCreatedBy represents the skills exports list output items file created by type.
type SkillsExportsListOutputItemsFileCreatedBy struct {
	Type              string                                                      `json:"type"`
	Name              string                                                      `json:"name"`
	ImageUrl          *string                                                     `json:"image_url,omitempty"`
	Email             *string                                                     `json:"email,omitempty"`
	OrganizationActor *SkillsExportsListOutputItemsFileCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsExportsListOutputItemsFileCreatedByConsumer          `json:"consumer,omitempty"`
}

// SkillsExportsListOutputItemsFile represents the skills exports list output items file type.
type SkillsExportsListOutputItemsFile struct {
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
	Purpose   string                                     `json:"purpose"`
	CreatedBy *SkillsExportsListOutputItemsFileCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - The files's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The files's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsListOutputItemsFileLink represents the skills exports list output items file link type.
type SkillsExportsListOutputItemsFileLink struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The links's unique identifier
	Id string `json:"id"`
	// FileId - The file's unique identifier
	FileId string `json:"file_id"`
	// Url - The file's public URL
	Url string `json:"url"`
	// CreatedAt - The links's creation date
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - The file's expiration date
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// SkillsExportsListOutputItemsCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsExportsListOutputItemsCreatedByOrganizationActorTeams struct {
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

// SkillsExportsListOutputItemsCreatedByOrganizationActor represents the skills exports list output items created by organization actor type.
type SkillsExportsListOutputItemsCreatedByOrganizationActor struct {
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
	ImageUrl string                                                        `json:"image_url"`
	Teams    []SkillsExportsListOutputItemsCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsListOutputItemsCreatedByConsumer represents the skills exports list output items created by consumer type.
type SkillsExportsListOutputItemsCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsListOutputItemsCreatedBy represents the skills exports list output items created by type.
type SkillsExportsListOutputItemsCreatedBy struct {
	Type              string                                                  `json:"type"`
	Name              string                                                  `json:"name"`
	ImageUrl          *string                                                 `json:"image_url,omitempty"`
	Email             *string                                                 `json:"email,omitempty"`
	OrganizationActor *SkillsExportsListOutputItemsCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsExportsListOutputItemsCreatedByConsumer          `json:"consumer,omitempty"`
}

// SkillsExportsListOutputItems represents the skills exports list output items type.
type SkillsExportsListOutputItems struct {
	Object      string                                 `json:"object"`
	Id          string                                 `json:"id"`
	Target      string                                 `json:"target"`
	Status      string                                 `json:"status"`
	File        *SkillsExportsListOutputItemsFile      `json:"file,omitempty"`
	FileLink    *SkillsExportsListOutputItemsFileLink  `json:"file_link,omitempty"`
	CreatedBy   *SkillsExportsListOutputItemsCreatedBy `json:"created_by,omitempty"`
	CreatedAt   time.Time                              `json:"created_at"`
	StartedAt   *time.Time                             `json:"started_at,omitempty"`
	CompletedAt *time.Time                             `json:"completed_at,omitempty"`
}

// SkillsExportsListOutputPagination represents the skills exports list output pagination type.
type SkillsExportsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsExportsListOutput represents the skills exports list output type.
type SkillsExportsListOutput struct {
	Items      []SkillsExportsListOutputItems    `json:"items"`
	Pagination SkillsExportsListOutputPagination `json:"pagination"`
}

// MapSkillsExportsListOutputFromJSON deserializes JSON data into a SkillsExportsListOutput.
func MapSkillsExportsListOutputFromJSON(data []byte) (*SkillsExportsListOutput, error) {
	var v SkillsExportsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsExportsListOutputToJSON serializes a SkillsExportsListOutput to JSON.
func MapSkillsExportsListOutputToJSON(v *SkillsExportsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsExportsListQuery represents the skills exports list query type.
type SkillsExportsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Id     *any     `json:"id,omitempty"`
	Target *any     `json:"target,omitempty"`
	Status *any     `json:"status,omitempty"`
}

// MapSkillsExportsListQueryFromJSON deserializes JSON data into a SkillsExportsListQuery.
func MapSkillsExportsListQueryFromJSON(data []byte) (*SkillsExportsListQuery, error) {
	var v SkillsExportsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsExportsListQueryToJSON serializes a SkillsExportsListQuery to JSON.
func MapSkillsExportsListQueryToJSON(v *SkillsExportsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
