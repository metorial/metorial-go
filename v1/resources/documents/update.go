package documents

import (
	"encoding/json"
	"time"
)

// DocumentsUpdateOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsUpdateOutputCreatedByOrganizationActorTeams struct {
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

// DocumentsUpdateOutputCreatedByOrganizationActor represents the documents update output created by organization actor type.
type DocumentsUpdateOutputCreatedByOrganizationActor struct {
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
	Teams    []DocumentsUpdateOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsUpdateOutputCreatedByConsumer represents the documents update output created by consumer type.
type DocumentsUpdateOutputCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsUpdateOutputCreatedBy represents the documents update output created by type.
type DocumentsUpdateOutputCreatedBy struct {
	Type              string                                           `json:"type"`
	Name              string                                           `json:"name"`
	ImageUrl          *string                                          `json:"image_url,omitempty"`
	Email             *string                                          `json:"email,omitempty"`
	OrganizationActor *DocumentsUpdateOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsUpdateOutputCreatedByConsumer          `json:"consumer,omitempty"`
}

// DocumentsUpdateOutput represents the documents update output type.
type DocumentsUpdateOutput struct {
	// Object - String representing the object's type
	Object           string                          `json:"object"`
	Id               string                          `json:"id"`
	Status           string                          `json:"status"`
	Title            string                          `json:"title"`
	Content          string                          `json:"content"`
	FileId           string                          `json:"file_id"`
	ParentDocumentId *string                         `json:"parent_document_id,omitempty"`
	CurrentVersionId *string                         `json:"current_version_id,omitempty"`
	CreatedBy        *DocumentsUpdateOutputCreatedBy `json:"created_by,omitempty"`
	CreatedAt        time.Time                       `json:"created_at"`
	UpdatedAt        time.Time                       `json:"updated_at"`
}

// MapDocumentsUpdateOutputFromJSON deserializes JSON data into a DocumentsUpdateOutput.
func MapDocumentsUpdateOutputFromJSON(data []byte) (*DocumentsUpdateOutput, error) {
	var v DocumentsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsUpdateOutputToJSON serializes a DocumentsUpdateOutput to JSON.
func MapDocumentsUpdateOutputToJSON(v *DocumentsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsUpdateBody represents the documents update body type.
type DocumentsUpdateBody struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

// MapDocumentsUpdateBodyFromJSON deserializes JSON data into a DocumentsUpdateBody.
func MapDocumentsUpdateBodyFromJSON(data []byte) (*DocumentsUpdateBody, error) {
	var v DocumentsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsUpdateBodyToJSON serializes a DocumentsUpdateBody to JSON.
func MapDocumentsUpdateBodyToJSON(v *DocumentsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
