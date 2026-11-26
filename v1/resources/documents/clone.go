package documents

import (
	"encoding/json"
	"time"
)

// DocumentsCloneOutputCreatedByOrganizationActorMember represents the documents clone output created by organization actor member type.
type DocumentsCloneOutputCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

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

// DocumentsCloneOutputCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
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
	Member   *DocumentsCloneOutputCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCloneOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCloneOutputCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsCloneOutputCreatedByConsumer struct {
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

// DocumentsCloneOutputCreatedBy - Actor that created the document.
type DocumentsCloneOutputCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsCloneOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsCloneOutputCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsCloneOutput represents the documents clone output type.
type DocumentsCloneOutput struct {
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
	CreatedBy *DocumentsCloneOutputCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
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
	// TargetDocumentId - ID to use for the cloned document.
	TargetDocumentId *string `json:"target_document_id,omitempty"`
	// Title - Title of the cloned document.
	Title *string `json:"title,omitempty"`
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
