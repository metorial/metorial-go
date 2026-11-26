package exports

import (
	"encoding/json"
	"time"
)

// SkillsExportsGetOutputFileCreatedByOrganizationActorMember represents the skills exports get output file created by organization actor member type.
type SkillsExportsGetOutputFileCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsExportsGetOutputFileCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsExportsGetOutputFileCreatedByOrganizationActorTeams struct {
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

// SkillsExportsGetOutputFileCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsExportsGetOutputFileCreatedByOrganizationActor struct {
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
	ImageUrl string                                                      `json:"image_url"`
	Member   *SkillsExportsGetOutputFileCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsExportsGetOutputFileCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsGetOutputFileCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsExportsGetOutputFileCreatedByConsumer struct {
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

// SkillsExportsGetOutputFileCreatedBy - Actor that uploaded the file.
type SkillsExportsGetOutputFileCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsExportsGetOutputFileCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsExportsGetOutputFileCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsExportsGetOutputFile represents the skills exports get output file type.
type SkillsExportsGetOutputFile struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `deleted` after the file is deleted.
	Status string `json:"status"`
	// FileName - Name of the uploaded file.
	FileName string `json:"file_name"`
	// FileSize - Size of the file in bytes.
	FileSize float64 `json:"file_size"`
	// FileType - MIME type of the file.
	FileType string `json:"file_type"`
	// Title - Label shown for the file.
	Title string `json:"title"`
	// Purpose - Why the file was uploaded. Controls what it can be used for.
	Purpose string `json:"purpose"`
	// CreatedBy - Actor that uploaded the file.
	CreatedBy *SkillsExportsGetOutputFileCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsGetOutputFileLink represents the skills exports get output file link type.
type SkillsExportsGetOutputFileLink struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// FileId - File this link downloads.
	FileId string `json:"file_id"`
	// Url - Public URL for downloading the file.
	Url string `json:"url"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - When the link stops working. `null` if it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// SkillsExportsGetOutputCreatedByOrganizationActorMember represents the skills exports get output created by organization actor member type.
type SkillsExportsGetOutputCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsExportsGetOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsExportsGetOutputCreatedByOrganizationActorTeams struct {
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

// SkillsExportsGetOutputCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsExportsGetOutputCreatedByOrganizationActor struct {
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
	ImageUrl string                                                  `json:"image_url"`
	Member   *SkillsExportsGetOutputCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsExportsGetOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsGetOutputCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsExportsGetOutputCreatedByConsumer struct {
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

// SkillsExportsGetOutputCreatedBy - Actor behind a participant.
type SkillsExportsGetOutputCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsExportsGetOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsExportsGetOutputCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsExportsGetOutput represents the skills exports get output type.
type SkillsExportsGetOutput struct {
	Object   string                          `json:"object"`
	Id       string                          `json:"id"`
	Target   string                          `json:"target"`
	Status   string                          `json:"status"`
	File     *SkillsExportsGetOutputFile     `json:"file,omitempty"`
	FileLink *SkillsExportsGetOutputFileLink `json:"file_link,omitempty"`
	// CreatedBy - Actor behind a participant.
	CreatedBy   *SkillsExportsGetOutputCreatedBy `json:"created_by,omitempty"`
	CreatedAt   time.Time                        `json:"created_at"`
	StartedAt   *time.Time                       `json:"started_at,omitempty"`
	CompletedAt *time.Time                       `json:"completed_at,omitempty"`
}

// MapSkillsExportsGetOutputFromJSON deserializes JSON data into a SkillsExportsGetOutput.
func MapSkillsExportsGetOutputFromJSON(data []byte) (*SkillsExportsGetOutput, error) {
	var v SkillsExportsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsExportsGetOutputToJSON serializes a SkillsExportsGetOutput to JSON.
func MapSkillsExportsGetOutputToJSON(v *SkillsExportsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
