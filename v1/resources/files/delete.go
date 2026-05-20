package files

import (
	"encoding/json"
	"time"
)

// FilesDeleteOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type FilesDeleteOutputCreatedByOrganizationActorTeams struct {
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

// FilesDeleteOutputCreatedByOrganizationActor represents the files delete output created by organization actor type.
type FilesDeleteOutputCreatedByOrganizationActor struct {
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
	ImageUrl string                                             `json:"image_url"`
	Teams    []FilesDeleteOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// FilesDeleteOutputCreatedByConsumer represents the files delete output created by consumer type.
type FilesDeleteOutputCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FilesDeleteOutputCreatedBy represents the files delete output created by type.
type FilesDeleteOutputCreatedBy struct {
	Type              string                                       `json:"type"`
	Name              string                                       `json:"name"`
	ImageUrl          *string                                      `json:"image_url,omitempty"`
	Email             *string                                      `json:"email,omitempty"`
	OrganizationActor *FilesDeleteOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *FilesDeleteOutputCreatedByConsumer          `json:"consumer,omitempty"`
}

// FilesDeleteOutput represents the files delete output type.
type FilesDeleteOutput struct {
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
	Purpose   string                      `json:"purpose"`
	CreatedBy *FilesDeleteOutputCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - The files's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The files's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// MapFilesDeleteOutputFromJSON deserializes JSON data into a FilesDeleteOutput.
func MapFilesDeleteOutputFromJSON(data []byte) (*FilesDeleteOutput, error) {
	var v FilesDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFilesDeleteOutputToJSON serializes a FilesDeleteOutput to JSON.
func MapFilesDeleteOutputToJSON(v *FilesDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
