package documents

import (
	"encoding/json"
	"time"
)

// DocumentsGetOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsGetOutputCreatedByOrganizationActorTeams struct {
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

// DocumentsGetOutputCreatedByOrganizationActor represents the documents get output created by organization actor type.
type DocumentsGetOutputCreatedByOrganizationActor struct {
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
	ImageUrl string                                              `json:"image_url"`
	Teams    []DocumentsGetOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsGetOutputCreatedByConsumer represents the documents get output created by consumer type.
type DocumentsGetOutputCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsGetOutputCreatedBy represents the documents get output created by type.
type DocumentsGetOutputCreatedBy struct {
	Type              string                                        `json:"type"`
	Name              string                                        `json:"name"`
	ImageUrl          *string                                       `json:"image_url,omitempty"`
	Email             *string                                       `json:"email,omitempty"`
	OrganizationActor *DocumentsGetOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsGetOutputCreatedByConsumer          `json:"consumer,omitempty"`
}

// DocumentsGetOutput represents the documents get output type.
type DocumentsGetOutput struct {
	// Object - String representing the object's type
	Object           string                       `json:"object"`
	Id               string                       `json:"id"`
	Status           string                       `json:"status"`
	Title            string                       `json:"title"`
	Content          string                       `json:"content"`
	FileId           string                       `json:"file_id"`
	ParentDocumentId *string                      `json:"parent_document_id,omitempty"`
	CurrentVersionId *string                      `json:"current_version_id,omitempty"`
	CreatedBy        *DocumentsGetOutputCreatedBy `json:"created_by,omitempty"`
	CreatedAt        time.Time                    `json:"created_at"`
	UpdatedAt        time.Time                    `json:"updated_at"`
}

// MapDocumentsGetOutputFromJSON deserializes JSON data into a DocumentsGetOutput.
func MapDocumentsGetOutputFromJSON(data []byte) (*DocumentsGetOutput, error) {
	var v DocumentsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsGetOutputToJSON serializes a DocumentsGetOutput to JSON.
func MapDocumentsGetOutputToJSON(v *DocumentsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
