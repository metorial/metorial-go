package items

import (
	"encoding/json"
	"time"
)

// StoresItemsGetOutputFileCreatedByOrganizationActorTeams - The teams the actor belongs to
type StoresItemsGetOutputFileCreatedByOrganizationActorTeams struct {
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

// StoresItemsGetOutputFileCreatedByOrganizationActor represents the stores items get output file created by organization actor type.
type StoresItemsGetOutputFileCreatedByOrganizationActor struct {
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
	Teams    []StoresItemsGetOutputFileCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsGetOutputFileCreatedByConsumer represents the stores items get output file created by consumer type.
type StoresItemsGetOutputFileCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsGetOutputFileCreatedBy represents the stores items get output file created by type.
type StoresItemsGetOutputFileCreatedBy struct {
	Type              string                                              `json:"type"`
	Name              string                                              `json:"name"`
	ImageUrl          *string                                             `json:"image_url,omitempty"`
	Email             *string                                             `json:"email,omitempty"`
	OrganizationActor *StoresItemsGetOutputFileCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *StoresItemsGetOutputFileCreatedByConsumer          `json:"consumer,omitempty"`
}

// StoresItemsGetOutputFile represents the stores items get output file type.
type StoresItemsGetOutputFile struct {
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
	Purpose   string                             `json:"purpose"`
	CreatedBy *StoresItemsGetOutputFileCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - The files's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The files's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsGetOutputDocumentCreatedByOrganizationActorTeams - The teams the actor belongs to
type StoresItemsGetOutputDocumentCreatedByOrganizationActorTeams struct {
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

// StoresItemsGetOutputDocumentCreatedByOrganizationActor represents the stores items get output document created by organization actor type.
type StoresItemsGetOutputDocumentCreatedByOrganizationActor struct {
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
	Teams    []StoresItemsGetOutputDocumentCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsGetOutputDocumentCreatedByConsumer represents the stores items get output document created by consumer type.
type StoresItemsGetOutputDocumentCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsGetOutputDocumentCreatedBy represents the stores items get output document created by type.
type StoresItemsGetOutputDocumentCreatedBy struct {
	Type              string                                                  `json:"type"`
	Name              string                                                  `json:"name"`
	ImageUrl          *string                                                 `json:"image_url,omitempty"`
	Email             *string                                                 `json:"email,omitempty"`
	OrganizationActor *StoresItemsGetOutputDocumentCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *StoresItemsGetOutputDocumentCreatedByConsumer          `json:"consumer,omitempty"`
}

// StoresItemsGetOutputDocument represents the stores items get output document type.
type StoresItemsGetOutputDocument struct {
	// Object - String representing the object's type
	Object           string                                 `json:"object"`
	Id               string                                 `json:"id"`
	Status           string                                 `json:"status"`
	Title            string                                 `json:"title"`
	Content          string                                 `json:"content"`
	FileId           string                                 `json:"file_id"`
	ParentDocumentId *string                                `json:"parent_document_id,omitempty"`
	CurrentVersionId *string                                `json:"current_version_id,omitempty"`
	CreatedBy        *StoresItemsGetOutputDocumentCreatedBy `json:"created_by,omitempty"`
	CreatedAt        time.Time                              `json:"created_at"`
	UpdatedAt        time.Time                              `json:"updated_at"`
}

// StoresItemsGetOutput represents the stores items get output type.
type StoresItemsGetOutput struct {
	// Object - String representing the object's type
	Object      string                        `json:"object"`
	Id          string                        `json:"id"`
	Kind        string                        `json:"kind"`
	Path        string                        `json:"path"`
	StoreId     string                        `json:"store_id"`
	DirectoryId *string                       `json:"directory_id,omitempty"`
	File        *StoresItemsGetOutputFile     `json:"file,omitempty"`
	Document    *StoresItemsGetOutputDocument `json:"document,omitempty"`
	CreatedAt   time.Time                     `json:"created_at"`
	UpdatedAt   time.Time                     `json:"updated_at"`
}

// MapStoresItemsGetOutputFromJSON deserializes JSON data into a StoresItemsGetOutput.
func MapStoresItemsGetOutputFromJSON(data []byte) (*StoresItemsGetOutput, error) {
	var v StoresItemsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresItemsGetOutputToJSON serializes a StoresItemsGetOutput to JSON.
func MapStoresItemsGetOutputToJSON(v *StoresItemsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
