package exports

import (
	"encoding/json"
	"time"
)

// SkillsExportsGetOutputFileCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsExportsGetOutputFileCreatedByOrganizationActorTeams struct {
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

// SkillsExportsGetOutputFileCreatedByOrganizationActor represents the skills exports get output file created by organization actor type.
type SkillsExportsGetOutputFileCreatedByOrganizationActor struct {
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
	ImageUrl string                                                      `json:"image_url"`
	Teams    []SkillsExportsGetOutputFileCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsGetOutputFileCreatedByConsumer represents the skills exports get output file created by consumer type.
type SkillsExportsGetOutputFileCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsGetOutputFileCreatedBy represents the skills exports get output file created by type.
type SkillsExportsGetOutputFileCreatedBy struct {
	Type              string                                                `json:"type"`
	Name              string                                                `json:"name"`
	ImageUrl          *string                                               `json:"image_url,omitempty"`
	Email             *string                                               `json:"email,omitempty"`
	OrganizationActor *SkillsExportsGetOutputFileCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsExportsGetOutputFileCreatedByConsumer          `json:"consumer,omitempty"`
}

// SkillsExportsGetOutputFile represents the skills exports get output file type.
type SkillsExportsGetOutputFile struct {
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
	Purpose   string                               `json:"purpose"`
	CreatedBy *SkillsExportsGetOutputFileCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - The files's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The files's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsGetOutputFileLink represents the skills exports get output file link type.
type SkillsExportsGetOutputFileLink struct {
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

// SkillsExportsGetOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsExportsGetOutputCreatedByOrganizationActorTeams struct {
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

// SkillsExportsGetOutputCreatedByOrganizationActor represents the skills exports get output created by organization actor type.
type SkillsExportsGetOutputCreatedByOrganizationActor struct {
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
	ImageUrl string                                                  `json:"image_url"`
	Teams    []SkillsExportsGetOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsGetOutputCreatedByConsumer represents the skills exports get output created by consumer type.
type SkillsExportsGetOutputCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsGetOutputCreatedBy represents the skills exports get output created by type.
type SkillsExportsGetOutputCreatedBy struct {
	Type              string                                            `json:"type"`
	Name              string                                            `json:"name"`
	ImageUrl          *string                                           `json:"image_url,omitempty"`
	Email             *string                                           `json:"email,omitempty"`
	OrganizationActor *SkillsExportsGetOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsExportsGetOutputCreatedByConsumer          `json:"consumer,omitempty"`
}

// SkillsExportsGetOutput represents the skills exports get output type.
type SkillsExportsGetOutput struct {
	Object      string                           `json:"object"`
	Id          string                           `json:"id"`
	Target      string                           `json:"target"`
	Status      string                           `json:"status"`
	File        *SkillsExportsGetOutputFile      `json:"file,omitempty"`
	FileLink    *SkillsExportsGetOutputFileLink  `json:"file_link,omitempty"`
	CreatedBy   *SkillsExportsGetOutputCreatedBy `json:"created_by,omitempty"`
	CreatedAt   time.Time                        `json:"created_at"`
	StartedAt   *time.Time                       `json:"started_at,omitempty"`
	CompletedAt *time.Time                       `json:"completed_at,omitempty"`
}

// MapSkillsExportsGetOutputFromJSON deserializes JSON data into a SkillsExportsGetOutput.
func MapSkillsExportsGetOutputFromJSON(data []byte) (*SkillsExportsGetOutput, error) {
	var v SkillsExportsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsExportsGetOutputToJSON serializes a SkillsExportsGetOutput to JSON.
func MapSkillsExportsGetOutputToJSON(v *SkillsExportsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
