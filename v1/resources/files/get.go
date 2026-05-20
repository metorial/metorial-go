package files

import (
	"encoding/json"
	"time"
)

// FilesGetOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type FilesGetOutputCreatedByOrganizationActorTeams struct {
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

// FilesGetOutputCreatedByOrganizationActor represents the files get output created by organization actor type.
type FilesGetOutputCreatedByOrganizationActor struct {
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
	ImageUrl string                                          `json:"image_url"`
	Teams    []FilesGetOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// FilesGetOutputCreatedByConsumer represents the files get output created by consumer type.
type FilesGetOutputCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FilesGetOutputCreatedBy represents the files get output created by type.
type FilesGetOutputCreatedBy struct {
	Type              string                                    `json:"type"`
	Name              string                                    `json:"name"`
	ImageUrl          *string                                   `json:"image_url,omitempty"`
	Email             *string                                   `json:"email,omitempty"`
	OrganizationActor *FilesGetOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *FilesGetOutputCreatedByConsumer          `json:"consumer,omitempty"`
}

// FilesGetOutput represents the files get output type.
type FilesGetOutput struct {
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
	Purpose   string                   `json:"purpose"`
	CreatedBy *FilesGetOutputCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - The files's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The files's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// MapFilesGetOutputFromJSON deserializes JSON data into a FilesGetOutput.
func MapFilesGetOutputFromJSON(data []byte) (*FilesGetOutput, error) {
	var v FilesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFilesGetOutputToJSON serializes a FilesGetOutput to JSON.
func MapFilesGetOutputToJSON(v *FilesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
