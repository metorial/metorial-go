package documents

import (
	"encoding/json"
	"time"
)

// DocumentsCreateOutputCreatedByOrganizationActorMember represents the documents create output created by organization actor member type.
type DocumentsCreateOutputCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

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

// DocumentsCreateOutputCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
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
	Member   *DocumentsCreateOutputCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCreateOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCreateOutputCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsCreateOutputCreatedByConsumer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address.
	Email string `json:"email"`
	// ImageUrl - Avatar image URL.
	ImageUrl string `json:"image_url"`
	// UserId - User this consumer is linked to. `null` when unlinked.
	UserId *string `json:"user_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCreateOutputCreatedBy - Actor that created the document.
type DocumentsCreateOutputCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsCreateOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsCreateOutputCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsCreateOutput represents the documents create output type.
type DocumentsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `deleted` after the document is deleted.
	Status string `json:"status"`
	// Title - Label shown for the document.
	Title string `json:"title"`
	// Content - Text body of the document.
	Content string `json:"content"`
	// FileId - File holding the document content.
	FileId string `json:"file_id"`
	// ParentDocumentId - Document this was cloned from. `null` for top-level documents.
	ParentDocumentId *string `json:"parent_document_id,omitempty"`
	// CurrentVersionId - Latest version of the document, by ID.
	CurrentVersionId *string `json:"current_version_id,omitempty"`
	// CreatedBy - Actor that created the document.
	CreatedBy *DocumentsCreateOutputCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
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
	// Title - Title of the document.
	Title string `json:"title"`
	// Content - Content of the document.
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
