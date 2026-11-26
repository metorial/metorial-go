package exports

import (
	"encoding/json"
	"time"
)

// SkillsExportsListOutputItemsFileCreatedByOrganizationActorMember represents the skills exports list output items file created by organization actor member type.
type SkillsExportsListOutputItemsFileCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsExportsListOutputItemsFileCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsExportsListOutputItemsFileCreatedByOrganizationActorTeams struct {
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

// SkillsExportsListOutputItemsFileCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsExportsListOutputItemsFileCreatedByOrganizationActor struct {
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
	ImageUrl string                                                            `json:"image_url"`
	Member   *SkillsExportsListOutputItemsFileCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsExportsListOutputItemsFileCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsListOutputItemsFileCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsExportsListOutputItemsFileCreatedByConsumer struct {
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

// SkillsExportsListOutputItemsFileCreatedBy - Actor that uploaded the file.
type SkillsExportsListOutputItemsFileCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsExportsListOutputItemsFileCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsExportsListOutputItemsFileCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsExportsListOutputItemsFile represents the skills exports list output items file type.
type SkillsExportsListOutputItemsFile struct {
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
	CreatedBy *SkillsExportsListOutputItemsFileCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsListOutputItemsFileLink represents the skills exports list output items file link type.
type SkillsExportsListOutputItemsFileLink struct {
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

// SkillsExportsListOutputItemsCreatedByOrganizationActorMember represents the skills exports list output items created by organization actor member type.
type SkillsExportsListOutputItemsCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsExportsListOutputItemsCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsExportsListOutputItemsCreatedByOrganizationActorTeams struct {
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

// SkillsExportsListOutputItemsCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsExportsListOutputItemsCreatedByOrganizationActor struct {
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
	ImageUrl string                                                        `json:"image_url"`
	Member   *SkillsExportsListOutputItemsCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsExportsListOutputItemsCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsExportsListOutputItemsCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsExportsListOutputItemsCreatedByConsumer struct {
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

// SkillsExportsListOutputItemsCreatedBy - Actor behind a participant.
type SkillsExportsListOutputItemsCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsExportsListOutputItemsCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsExportsListOutputItemsCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsExportsListOutputItems represents the skills exports list output items type.
type SkillsExportsListOutputItems struct {
	Object   string                                `json:"object"`
	Id       string                                `json:"id"`
	Target   string                                `json:"target"`
	Status   string                                `json:"status"`
	File     *SkillsExportsListOutputItemsFile     `json:"file,omitempty"`
	FileLink *SkillsExportsListOutputItemsFileLink `json:"file_link,omitempty"`
	// CreatedBy - Actor behind a participant.
	CreatedBy   *SkillsExportsListOutputItemsCreatedBy `json:"created_by,omitempty"`
	CreatedAt   time.Time                              `json:"created_at"`
	StartedAt   *time.Time                             `json:"started_at,omitempty"`
	CompletedAt *time.Time                             `json:"completed_at,omitempty"`
}

// SkillsExportsListOutputPagination represents the skills exports list output pagination type.
type SkillsExportsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsExportsListOutput represents the skills exports list output type.
type SkillsExportsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                            `json:"object"`
	Items      []SkillsExportsListOutputItems    `json:"items"`
	Pagination SkillsExportsListOutputPagination `json:"pagination"`
}

// MapSkillsExportsListOutputFromJSON deserializes JSON data into a SkillsExportsListOutput.
func MapSkillsExportsListOutputFromJSON(data []byte) (*SkillsExportsListOutput, error) {
	var v SkillsExportsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsExportsListOutputToJSON serializes a SkillsExportsListOutput to JSON.
func MapSkillsExportsListOutputToJSON(v *SkillsExportsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsExportsListQuery represents the skills exports list query type.
type SkillsExportsListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Id - Only return skill exports with these IDs.
	Id *any `json:"id,omitempty"`
	// Target - Only return skill exports of one of these types.
	Target *any `json:"target,omitempty"`
	// Status - Only return skill exports with one of these statuses.
	Status *any `json:"status,omitempty"`
}

// MapSkillsExportsListQueryFromJSON deserializes JSON data into a SkillsExportsListQuery.
func MapSkillsExportsListQueryFromJSON(data []byte) (*SkillsExportsListQuery, error) {
	var v SkillsExportsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsExportsListQueryToJSON serializes a SkillsExportsListQuery to JSON.
func MapSkillsExportsListQueryToJSON(v *SkillsExportsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
