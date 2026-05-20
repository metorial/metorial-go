package items

import (
	"encoding/json"
	"time"
)

// StoresItemsModifyOutputItemsFileCreatedByOrganizationActorTeams - The teams the actor belongs to
type StoresItemsModifyOutputItemsFileCreatedByOrganizationActorTeams struct {
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

// StoresItemsModifyOutputItemsFileCreatedByOrganizationActor represents the stores items modify output items file created by organization actor type.
type StoresItemsModifyOutputItemsFileCreatedByOrganizationActor struct {
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
	Teams    []StoresItemsModifyOutputItemsFileCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItemsFileCreatedByConsumer represents the stores items modify output items file created by consumer type.
type StoresItemsModifyOutputItemsFileCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItemsFileCreatedBy represents the stores items modify output items file created by type.
type StoresItemsModifyOutputItemsFileCreatedBy struct {
	Type              string                                                      `json:"type"`
	Name              string                                                      `json:"name"`
	ImageUrl          *string                                                     `json:"image_url,omitempty"`
	Email             *string                                                     `json:"email,omitempty"`
	OrganizationActor *StoresItemsModifyOutputItemsFileCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *StoresItemsModifyOutputItemsFileCreatedByConsumer          `json:"consumer,omitempty"`
}

// StoresItemsModifyOutputItemsFile represents the stores items modify output items file type.
type StoresItemsModifyOutputItemsFile struct {
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
	CreatedBy *StoresItemsModifyOutputItemsFileCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - The files's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The files's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActorTeams - The teams the actor belongs to
type StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActorTeams struct {
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

// StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActor represents the stores items modify output items document created by organization actor type.
type StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActor struct {
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
	ImageUrl string                                                                `json:"image_url"`
	Teams    []StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItemsDocumentCreatedByConsumer represents the stores items modify output items document created by consumer type.
type StoresItemsModifyOutputItemsDocumentCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItemsDocumentCreatedBy represents the stores items modify output items document created by type.
type StoresItemsModifyOutputItemsDocumentCreatedBy struct {
	Type              string                                                          `json:"type"`
	Name              string                                                          `json:"name"`
	ImageUrl          *string                                                         `json:"image_url,omitempty"`
	Email             *string                                                         `json:"email,omitempty"`
	OrganizationActor *StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *StoresItemsModifyOutputItemsDocumentCreatedByConsumer          `json:"consumer,omitempty"`
}

// StoresItemsModifyOutputItemsDocument represents the stores items modify output items document type.
type StoresItemsModifyOutputItemsDocument struct {
	// Object - String representing the object's type
	Object           string                                         `json:"object"`
	Id               string                                         `json:"id"`
	Status           string                                         `json:"status"`
	Title            string                                         `json:"title"`
	Content          string                                         `json:"content"`
	FileId           string                                         `json:"file_id"`
	ParentDocumentId *string                                        `json:"parent_document_id,omitempty"`
	CurrentVersionId *string                                        `json:"current_version_id,omitempty"`
	CreatedBy        *StoresItemsModifyOutputItemsDocumentCreatedBy `json:"created_by,omitempty"`
	CreatedAt        time.Time                                      `json:"created_at"`
	UpdatedAt        time.Time                                      `json:"updated_at"`
}

// StoresItemsModifyOutputItems represents the stores items modify output items type.
type StoresItemsModifyOutputItems struct {
	// Object - String representing the object's type
	Object      string                                `json:"object"`
	Id          string                                `json:"id"`
	Kind        string                                `json:"kind"`
	Path        string                                `json:"path"`
	StoreId     string                                `json:"store_id"`
	DirectoryId *string                               `json:"directory_id,omitempty"`
	File        *StoresItemsModifyOutputItemsFile     `json:"file,omitempty"`
	Document    *StoresItemsModifyOutputItemsDocument `json:"document,omitempty"`
	CreatedAt   time.Time                             `json:"created_at"`
	UpdatedAt   time.Time                             `json:"updated_at"`
}

// StoresItemsModifyOutput represents the stores items modify output type.
type StoresItemsModifyOutput struct {
	// Object - String representing the object's type
	Object string                         `json:"object"`
	Items  []StoresItemsModifyOutputItems `json:"items"`
}

// MapStoresItemsModifyOutputFromJSON deserializes JSON data into a StoresItemsModifyOutput.
func MapStoresItemsModifyOutputFromJSON(data []byte) (*StoresItemsModifyOutput, error) {
	var v StoresItemsModifyOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresItemsModifyOutputToJSON serializes a StoresItemsModifyOutput to JSON.
func MapStoresItemsModifyOutputToJSON(v *StoresItemsModifyOutput) ([]byte, error) {
	return json.Marshal(v)
}

// StoresItemsModifyBodyOperations represents the stores items modify body operations type.
type StoresItemsModifyBodyOperations struct {
	Type       *string `json:"type,omitempty"`
	ItemId     *string `json:"itemId,omitempty"`
	FileId     *string `json:"fileId,omitempty"`
	DocumentId *string `json:"documentId,omitempty"`
	Path       *string `json:"path,omitempty"`
}

// StoresItemsModifyBody represents the stores items modify body type.
type StoresItemsModifyBody struct {
	Operations []StoresItemsModifyBodyOperations `json:"operations"`
}

// MapStoresItemsModifyBodyFromJSON deserializes JSON data into a StoresItemsModifyBody.
func MapStoresItemsModifyBodyFromJSON(data []byte) (*StoresItemsModifyBody, error) {
	var v StoresItemsModifyBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresItemsModifyBodyToJSON serializes a StoresItemsModifyBody to JSON.
func MapStoresItemsModifyBodyToJSON(v *StoresItemsModifyBody) ([]byte, error) {
	return json.Marshal(v)
}
