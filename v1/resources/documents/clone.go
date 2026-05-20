package documents

import (
	"encoding/json"
	"time"
)

// DocumentsCloneOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsCloneOutputCreatedByOrganizationActorTeams struct {
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

// DocumentsCloneOutputCreatedByOrganizationActor represents the documents clone output created by organization actor type.
type DocumentsCloneOutputCreatedByOrganizationActor struct {
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
	Teams    []DocumentsCloneOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCloneOutputCreatedByConsumer represents the documents clone output created by consumer type.
type DocumentsCloneOutputCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCloneOutputCreatedBy represents the documents clone output created by type.
type DocumentsCloneOutputCreatedBy struct {
	Type              string                                          `json:"type"`
	Name              string                                          `json:"name"`
	ImageUrl          *string                                         `json:"image_url,omitempty"`
	Email             *string                                         `json:"email,omitempty"`
	OrganizationActor *DocumentsCloneOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCloneOutputCreatedByConsumer          `json:"consumer,omitempty"`
}

// DocumentsCloneOutput represents the documents clone output type.
type DocumentsCloneOutput struct {
	// Object - String representing the object's type
	Object           string                         `json:"object"`
	Id               string                         `json:"id"`
	Status           string                         `json:"status"`
	Title            string                         `json:"title"`
	Content          string                         `json:"content"`
	FileId           string                         `json:"file_id"`
	ParentDocumentId *string                        `json:"parent_document_id,omitempty"`
	CurrentVersionId *string                        `json:"current_version_id,omitempty"`
	CreatedBy        *DocumentsCloneOutputCreatedBy `json:"created_by,omitempty"`
	CreatedAt        time.Time                      `json:"created_at"`
	UpdatedAt        time.Time                      `json:"updated_at"`
}

// MapDocumentsCloneOutputFromJSON deserializes JSON data into a DocumentsCloneOutput.
func MapDocumentsCloneOutputFromJSON(data []byte) (*DocumentsCloneOutput, error) {
	var v DocumentsCloneOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCloneOutputToJSON serializes a DocumentsCloneOutput to JSON.
func MapDocumentsCloneOutputToJSON(v *DocumentsCloneOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsCloneBody represents the documents clone body type.
type DocumentsCloneBody struct {
	TargetDocumentId *string `json:"target_document_id,omitempty"`
	Title            *string `json:"title,omitempty"`
}

// MapDocumentsCloneBodyFromJSON deserializes JSON data into a DocumentsCloneBody.
func MapDocumentsCloneBodyFromJSON(data []byte) (*DocumentsCloneBody, error) {
	var v DocumentsCloneBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCloneBodyToJSON serializes a DocumentsCloneBody to JSON.
func MapDocumentsCloneBodyToJSON(v *DocumentsCloneBody) ([]byte, error) {
	return json.Marshal(v)
}
