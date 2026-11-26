package files

import (
	"encoding/json"
	"time"
)

// FilesListOutputItemsCreatedByOrganizationActorMember represents the files list output items created by organization actor member type.
type FilesListOutputItemsCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// FilesListOutputItemsCreatedByOrganizationActorTeams - The teams the actor belongs to
type FilesListOutputItemsCreatedByOrganizationActorTeams struct {
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

// FilesListOutputItemsCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type FilesListOutputItemsCreatedByOrganizationActor struct {
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
	Member   *FilesListOutputItemsCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []FilesListOutputItemsCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// FilesListOutputItemsCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type FilesListOutputItemsCreatedByConsumer struct {
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

// FilesListOutputItemsCreatedBy - Actor that uploaded the file.
type FilesListOutputItemsCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *FilesListOutputItemsCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *FilesListOutputItemsCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// FilesListOutputItems represents the files list output items type.
type FilesListOutputItems struct {
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
	CreatedBy *FilesListOutputItemsCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// FilesListOutputPagination represents the files list output pagination type.
type FilesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// FilesListOutput represents the files list output type.
type FilesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                    `json:"object"`
	Items      []FilesListOutputItems    `json:"items"`
	Pagination FilesListOutputPagination `json:"pagination"`
}

// MapFilesListOutputFromJSON deserializes JSON data into a FilesListOutput.
func MapFilesListOutputFromJSON(data []byte) (*FilesListOutput, error) {
	var v FilesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFilesListOutputToJSON serializes a FilesListOutput to JSON.
func MapFilesListOutputToJSON(v *FilesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// FilesListQueryCreatedAt - Only return results whose file creation time is in this range. Set `gt`, `lt`, or both.
type FilesListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// FilesListQueryUpdatedAt - Only return results whose file last update time is in this range. Set `gt`, `lt`, or both.
type FilesListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// FilesListQuery represents the files list query type.
type FilesListQuery struct {
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
	// Id - Only return files with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Purpose - Only return files with one of these purposes.
	Purpose *any `json:"purpose,omitempty"`
	// StoreId - Only return files in one of these stores.
	StoreId *any `json:"store_id,omitempty"`
	// DocumentId - Only return files attached to one of these documents.
	DocumentId *any `json:"document_id,omitempty"`
	// FileLinkId - Only return files shared through one of these file links.
	FileLinkId *any `json:"file_link_id,omitempty"`
	// CreatedAt - Only return results whose file creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *FilesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose file last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *FilesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapFilesListQueryFromJSON deserializes JSON data into a FilesListQuery.
func MapFilesListQueryFromJSON(data []byte) (*FilesListQuery, error) {
	var v FilesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFilesListQueryToJSON serializes a FilesListQuery to JSON.
func MapFilesListQueryToJSON(v *FilesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
