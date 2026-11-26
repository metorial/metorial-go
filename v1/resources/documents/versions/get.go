package versions

import (
	"encoding/json"
	"time"
)

// DocumentsVersionsGetOutputEditorsOrganizationActorMember represents the documents versions get output editors organization actor member type.
type DocumentsVersionsGetOutputEditorsOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

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

// DocumentsVersionsGetOutputEditorsOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
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
	Member   *DocumentsVersionsGetOutputEditorsOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsVersionsGetOutputEditorsOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsVersionsGetOutputEditorsConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsVersionsGetOutputEditorsConsumer struct {
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

// DocumentsVersionsGetOutputEditors - Actor behind a participant.
type DocumentsVersionsGetOutputEditors struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsVersionsGetOutputEditorsOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsVersionsGetOutputEditorsConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsVersionsGetOutput represents the documents versions get output type.
type DocumentsVersionsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// DocumentId - Document this version belongs to.
	DocumentId string `json:"document_id"`
	// VersionNumber - Sequential number of this version.
	VersionNumber float64 `json:"version_number"`
	// PreviousVersionId - Previous version, by ID. `null` for the first version.
	PreviousVersionId *string `json:"previous_version_id,omitempty"`
	// ListEditedAt - When the version content was last changed.
	ListEditedAt *time.Time `json:"list_edited_at,omitempty"`
	// Content - Text body of the document at this version.
	Content string `json:"content"`
	// Editors - Actors who edited this version.
	Editors []DocumentsVersionsGetOutputEditors `json:"editors"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
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
