package documents

import (
	"encoding/json"
	"time"
)

// DocumentsCreateOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsCreateOutputCreatedByOrganizationActorTeams struct {
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

// DocumentsCreateOutputCreatedByOrganizationActor represents the documents create output created by organization actor type.
type DocumentsCreateOutputCreatedByOrganizationActor struct {
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
	ImageUrl string                                                 `json:"image_url"`
	Teams    []DocumentsCreateOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCreateOutputCreatedByConsumer represents the documents create output created by consumer type.
type DocumentsCreateOutputCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCreateOutputCreatedBy represents the documents create output created by type.
type DocumentsCreateOutputCreatedBy struct {
	Type              string                                           `json:"type"`
	Name              string                                           `json:"name"`
	ImageUrl          *string                                          `json:"image_url,omitempty"`
	Email             *string                                          `json:"email,omitempty"`
	OrganizationActor *DocumentsCreateOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCreateOutputCreatedByConsumer          `json:"consumer,omitempty"`
}

// DocumentsCreateOutput represents the documents create output type.
type DocumentsCreateOutput struct {
	// Object - String representing the object's type
	Object           string                          `json:"object"`
	Id               string                          `json:"id"`
	Status           string                          `json:"status"`
	Title            string                          `json:"title"`
	Content          string                          `json:"content"`
	FileId           string                          `json:"file_id"`
	ParentDocumentId *string                         `json:"parent_document_id,omitempty"`
	CurrentVersionId *string                         `json:"current_version_id,omitempty"`
	CreatedBy        *DocumentsCreateOutputCreatedBy `json:"created_by,omitempty"`
	CreatedAt        time.Time                       `json:"created_at"`
	UpdatedAt        time.Time                       `json:"updated_at"`
}

// MapDocumentsCreateOutputFromJSON deserializes JSON data into a DocumentsCreateOutput.
func MapDocumentsCreateOutputFromJSON(data []byte) (*DocumentsCreateOutput, error) {
	var v DocumentsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCreateOutputToJSON serializes a DocumentsCreateOutput to JSON.
func MapDocumentsCreateOutputToJSON(v *DocumentsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsCreateBody represents the documents create body type.
type DocumentsCreateBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// MapDocumentsCreateBodyFromJSON deserializes JSON data into a DocumentsCreateBody.
func MapDocumentsCreateBodyFromJSON(data []byte) (*DocumentsCreateBody, error) {
	var v DocumentsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCreateBodyToJSON serializes a DocumentsCreateBody to JSON.
func MapDocumentsCreateBodyToJSON(v *DocumentsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
