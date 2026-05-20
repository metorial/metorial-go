package versions

import (
	"encoding/json"
	"time"
)

// DocumentsVersionsGetOutputEditorsOrganizationActorTeams - The teams the actor belongs to
type DocumentsVersionsGetOutputEditorsOrganizationActorTeams struct {
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

// DocumentsVersionsGetOutputEditorsOrganizationActor represents the documents versions get output editors organization actor type.
type DocumentsVersionsGetOutputEditorsOrganizationActor struct {
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
	Teams    []DocumentsVersionsGetOutputEditorsOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsVersionsGetOutputEditorsConsumer represents the documents versions get output editors consumer type.
type DocumentsVersionsGetOutputEditorsConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsVersionsGetOutputEditors represents the documents versions get output editors type.
type DocumentsVersionsGetOutputEditors struct {
	Type              string                                              `json:"type"`
	Name              string                                              `json:"name"`
	ImageUrl          *string                                             `json:"image_url,omitempty"`
	Email             *string                                             `json:"email,omitempty"`
	OrganizationActor *DocumentsVersionsGetOutputEditorsOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsVersionsGetOutputEditorsConsumer          `json:"consumer,omitempty"`
}

// DocumentsVersionsGetOutput represents the documents versions get output type.
type DocumentsVersionsGetOutput struct {
	// Object - String representing the object's type
	Object            string                              `json:"object"`
	Id                string                              `json:"id"`
	DocumentId        string                              `json:"document_id"`
	VersionNumber     float64                             `json:"version_number"`
	PreviousVersionId *string                             `json:"previous_version_id,omitempty"`
	ListEditedAt      *time.Time                          `json:"list_edited_at,omitempty"`
	Content           string                              `json:"content"`
	Editors           []DocumentsVersionsGetOutputEditors `json:"editors"`
	CreatedAt         time.Time                           `json:"created_at"`
}

// MapDocumentsVersionsGetOutputFromJSON deserializes JSON data into a DocumentsVersionsGetOutput.
func MapDocumentsVersionsGetOutputFromJSON(data []byte) (*DocumentsVersionsGetOutput, error) {
	var v DocumentsVersionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsVersionsGetOutputToJSON serializes a DocumentsVersionsGetOutput to JSON.
func MapDocumentsVersionsGetOutputToJSON(v *DocumentsVersionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
