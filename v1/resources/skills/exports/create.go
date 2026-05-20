package exports

import (
	"encoding/json"
	"time"
)

// SkillsExportsCreateOutputFileCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsExportsCreateOutputFileCreatedByOrganizationActorTeams struct {
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

// SkillsExportsCreateOutputFileCreatedByOrganizationActor represents the skills exports create output file created by organization actor type.
type SkillsExportsCreateOutputFileCreatedByOrganizationActor struct {
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
	Teams    []SkillsExportsCreateOutputFileCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsCreateOutputFileCreatedByConsumer represents the skills exports create output file created by consumer type.
type SkillsExportsCreateOutputFileCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsCreateOutputFileCreatedBy represents the skills exports create output file created by type.
type SkillsExportsCreateOutputFileCreatedBy struct {
	Type              string                                                   `json:"type"`
	Name              string                                                   `json:"name"`
	ImageUrl          *string                                                  `json:"image_url,omitempty"`
	Email             *string                                                  `json:"email,omitempty"`
	OrganizationActor *SkillsExportsCreateOutputFileCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsExportsCreateOutputFileCreatedByConsumer          `json:"consumer,omitempty"`
}

// SkillsExportsCreateOutputFile represents the skills exports create output file type.
type SkillsExportsCreateOutputFile struct {
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
	Purpose   string                                  `json:"purpose"`
	CreatedBy *SkillsExportsCreateOutputFileCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - The files's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The files's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsCreateOutputFileLink represents the skills exports create output file link type.
type SkillsExportsCreateOutputFileLink struct {
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

// SkillsExportsCreateOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsExportsCreateOutputCreatedByOrganizationActorTeams struct {
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

// SkillsExportsCreateOutputCreatedByOrganizationActor represents the skills exports create output created by organization actor type.
type SkillsExportsCreateOutputCreatedByOrganizationActor struct {
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
	ImageUrl string                                                     `json:"image_url"`
	Teams    []SkillsExportsCreateOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsCreateOutputCreatedByConsumer represents the skills exports create output created by consumer type.
type SkillsExportsCreateOutputCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsCreateOutputCreatedBy represents the skills exports create output created by type.
type SkillsExportsCreateOutputCreatedBy struct {
	Type              string                                               `json:"type"`
	Name              string                                               `json:"name"`
	ImageUrl          *string                                              `json:"image_url,omitempty"`
	Email             *string                                              `json:"email,omitempty"`
	OrganizationActor *SkillsExportsCreateOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsExportsCreateOutputCreatedByConsumer          `json:"consumer,omitempty"`
}

// SkillsExportsCreateOutput represents the skills exports create output type.
type SkillsExportsCreateOutput struct {
	Object      string                              `json:"object"`
	Id          string                              `json:"id"`
	Target      string                              `json:"target"`
	Status      string                              `json:"status"`
	File        *SkillsExportsCreateOutputFile      `json:"file,omitempty"`
	FileLink    *SkillsExportsCreateOutputFileLink  `json:"file_link,omitempty"`
	CreatedBy   *SkillsExportsCreateOutputCreatedBy `json:"created_by,omitempty"`
	CreatedAt   time.Time                           `json:"created_at"`
	StartedAt   *time.Time                          `json:"started_at,omitempty"`
	CompletedAt *time.Time                          `json:"completed_at,omitempty"`
}

// MapSkillsExportsCreateOutputFromJSON deserializes JSON data into a SkillsExportsCreateOutput.
func MapSkillsExportsCreateOutputFromJSON(data []byte) (*SkillsExportsCreateOutput, error) {
	var v SkillsExportsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsExportsCreateOutputToJSON serializes a SkillsExportsCreateOutput to JSON.
func MapSkillsExportsCreateOutputToJSON(v *SkillsExportsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsExportsCreateBody represents the skills exports create body type.
type SkillsExportsCreateBody struct {
	Target             string  `json:"target"`
	SkillId            *string `json:"skill_id,omitempty"`
	SkillPluginId      *string `json:"skill_plugin_id,omitempty"`
	SkillMarketplaceId *string `json:"skill_marketplace_id,omitempty"`
}

// MapSkillsExportsCreateBodyFromJSON deserializes JSON data into a SkillsExportsCreateBody.
func MapSkillsExportsCreateBodyFromJSON(data []byte) (*SkillsExportsCreateBody, error) {
	var v SkillsExportsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsExportsCreateBodyToJSON serializes a SkillsExportsCreateBody to JSON.
func MapSkillsExportsCreateBodyToJSON(v *SkillsExportsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
