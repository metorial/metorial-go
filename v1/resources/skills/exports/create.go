package exports

import (
	"encoding/json"
	"time"
)

// SkillsExportsCreateOutputFileCreatedByOrganizationActorMember represents the skills exports create output file created by organization actor member type.
type SkillsExportsCreateOutputFileCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsExportsCreateOutputFileCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsExportsCreateOutputFileCreatedByOrganizationActorTeams struct {
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

// SkillsExportsCreateOutputFileCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsExportsCreateOutputFileCreatedByOrganizationActor struct {
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
	ImageUrl string                                                         `json:"image_url"`
	Member   *SkillsExportsCreateOutputFileCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsExportsCreateOutputFileCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsCreateOutputFileCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsExportsCreateOutputFileCreatedByConsumer struct {
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

// SkillsExportsCreateOutputFileCreatedBy - Actor that uploaded the file.
type SkillsExportsCreateOutputFileCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsExportsCreateOutputFileCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsExportsCreateOutputFileCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsExportsCreateOutputFile represents the skills exports create output file type.
type SkillsExportsCreateOutputFile struct {
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
	CreatedBy *SkillsExportsCreateOutputFileCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsCreateOutputFileLink represents the skills exports create output file link type.
type SkillsExportsCreateOutputFileLink struct {
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

// SkillsExportsCreateOutputCreatedByOrganizationActorMember represents the skills exports create output created by organization actor member type.
type SkillsExportsCreateOutputCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsExportsCreateOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsExportsCreateOutputCreatedByOrganizationActorTeams struct {
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

// SkillsExportsCreateOutputCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsExportsCreateOutputCreatedByOrganizationActor struct {
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
	ImageUrl string                                                     `json:"image_url"`
	Member   *SkillsExportsCreateOutputCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsExportsCreateOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsCreateOutputCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsExportsCreateOutputCreatedByConsumer struct {
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

// SkillsExportsCreateOutputCreatedBy - Actor behind a participant.
type SkillsExportsCreateOutputCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsExportsCreateOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsExportsCreateOutputCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsExportsCreateOutput represents the skills exports create output type.
type SkillsExportsCreateOutput struct {
	Object   string                             `json:"object"`
	Id       string                             `json:"id"`
	Target   string                             `json:"target"`
	Status   string                             `json:"status"`
	File     *SkillsExportsCreateOutputFile     `json:"file,omitempty"`
	FileLink *SkillsExportsCreateOutputFileLink `json:"file_link,omitempty"`
	// CreatedBy - Actor behind a participant.
	CreatedBy   *SkillsExportsCreateOutputCreatedBy `json:"created_by,omitempty"`
	CreatedAt   time.Time                           `json:"created_at"`
	StartedAt   *time.Time                          `json:"started_at,omitempty"`
	CompletedAt *time.Time                          `json:"completed_at,omitempty"`
}

// MapSkillsExportsCreateOutputFromJSON deserializes JSON data into a SkillsExportsCreateOutput.
func MapSkillsExportsCreateOutputFromJSON(data []byte) (*SkillsExportsCreateOutput, error) {
	var v SkillsExportsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsExportsCreateOutputToJSON serializes a SkillsExportsCreateOutput to JSON.
func MapSkillsExportsCreateOutputToJSON(v *SkillsExportsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsExportsCreateBody represents the skills exports create body type.
type SkillsExportsCreateBody struct {
	// Target - What to export: a skill, plugin, or marketplace.
	Target string `json:"target"`
	// SkillId - Skill to export, by ID. Required when `target` is `skill`.
	SkillId *string `json:"skill_id,omitempty"`
	// SkillPluginId - Skill plugin to export, by ID. Required when `target` is `plugin`.
	SkillPluginId *string `json:"skill_plugin_id,omitempty"`
	// SkillMarketplaceId - Skill marketplace to export, by ID. Required when `target` is `marketplace`.
	SkillMarketplaceId *string `json:"skill_marketplace_id,omitempty"`
}

// MapSkillsExportsCreateBodyFromJSON deserializes JSON data into a SkillsExportsCreateBody.
func MapSkillsExportsCreateBodyFromJSON(data []byte) (*SkillsExportsCreateBody, error) {
	var v SkillsExportsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsExportsCreateBodyToJSON serializes a SkillsExportsCreateBody to JSON.
func MapSkillsExportsCreateBodyToJSON(v *SkillsExportsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
