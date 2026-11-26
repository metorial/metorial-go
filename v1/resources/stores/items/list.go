package items

import (
	"encoding/json"
	"time"
)

// StoresItemsListOutputItemsFileCreatedByOrganizationActorMember represents the stores items list output items file created by organization actor member type.
type StoresItemsListOutputItemsFileCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// StoresItemsListOutputItemsFileCreatedByOrganizationActorTeams - The teams the actor belongs to
type StoresItemsListOutputItemsFileCreatedByOrganizationActorTeams struct {
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

// StoresItemsListOutputItemsFileCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type StoresItemsListOutputItemsFileCreatedByOrganizationActor struct {
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
	ImageUrl string                                                          `json:"image_url"`
	Member   *StoresItemsListOutputItemsFileCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []StoresItemsListOutputItemsFileCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsListOutputItemsFileCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type StoresItemsListOutputItemsFileCreatedByConsumer struct {
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

// StoresItemsListOutputItemsFileCreatedBy - Actor that uploaded the file.
type StoresItemsListOutputItemsFileCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *StoresItemsListOutputItemsFileCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *StoresItemsListOutputItemsFileCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// StoresItemsListOutputItemsFile - File at this item. `null` unless `kind` is `file`.
type StoresItemsListOutputItemsFile struct {
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
	CreatedBy *StoresItemsListOutputItemsFileCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsListOutputItemsDocumentCreatedByOrganizationActorMember represents the stores items list output items document created by organization actor member type.
type StoresItemsListOutputItemsDocumentCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// StoresItemsListOutputItemsDocumentCreatedByOrganizationActorTeams - The teams the actor belongs to
type StoresItemsListOutputItemsDocumentCreatedByOrganizationActorTeams struct {
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

// StoresItemsListOutputItemsDocumentCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type StoresItemsListOutputItemsDocumentCreatedByOrganizationActor struct {
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
	ImageUrl string                                                              `json:"image_url"`
	Member   *StoresItemsListOutputItemsDocumentCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []StoresItemsListOutputItemsDocumentCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsListOutputItemsDocumentCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type StoresItemsListOutputItemsDocumentCreatedByConsumer struct {
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

// StoresItemsListOutputItemsDocumentCreatedBy - Actor that created the document.
type StoresItemsListOutputItemsDocumentCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *StoresItemsListOutputItemsDocumentCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *StoresItemsListOutputItemsDocumentCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// StoresItemsListOutputItemsDocument - Document at this item. `null` unless `kind` is `document`.
type StoresItemsListOutputItemsDocument struct {
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
	CreatedBy *StoresItemsListOutputItemsDocumentCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsListOutputItems represents the stores items list output items type.
type StoresItemsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Kind - What the item holds: `file`, `document`, or `directory`.
	Kind string `json:"kind"`
	// Path - Path of the item within the store.
	Path string `json:"path"`
	// StoreId - Store the item belongs to.
	StoreId string `json:"store_id"`
	// DirectoryId - Parent directory, by ID. `null` for top-level items.
	DirectoryId *string `json:"directory_id,omitempty"`
	// File - File at this item. `null` unless `kind` is `file`.
	File *StoresItemsListOutputItemsFile `json:"file,omitempty"`
	// Document - Document at this item. `null` unless `kind` is `document`.
	Document *StoresItemsListOutputItemsDocument `json:"document,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsListOutputPagination represents the stores items list output pagination type.
type StoresItemsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// StoresItemsListOutput represents the stores items list output type.
type StoresItemsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                          `json:"object"`
	Items      []StoresItemsListOutputItems    `json:"items"`
	Pagination StoresItemsListOutputPagination `json:"pagination"`
}

// MapStoresItemsListOutputFromJSON deserializes JSON data into a StoresItemsListOutput.
func MapStoresItemsListOutputFromJSON(data []byte) (*StoresItemsListOutput, error) {
	var v StoresItemsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresItemsListOutputToJSON serializes a StoresItemsListOutput to JSON.
func MapStoresItemsListOutputToJSON(v *StoresItemsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// StoresItemsListQueryCreatedAt - Only return results whose item creation time is in this range. Set `gt`, `lt`, or both.
type StoresItemsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// StoresItemsListQueryUpdatedAt - Only return results whose item last update time is in this range. Set `gt`, `lt`, or both.
type StoresItemsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// StoresItemsListQuery represents the stores items list query type.
type StoresItemsListQuery struct {
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
	// Id - Only return items with one of these IDs.
	Id *any `json:"id,omitempty"`
	// FileId - Only return items holding one of these files.
	FileId *any `json:"file_id,omitempty"`
	// DocumentId - Only return items holding one of these documents.
	DocumentId *any `json:"document_id,omitempty"`
	// Type - Filter by store item type. Repeat `type` to include multiple values. Defaults to `file` and `document`.
	Type *any `json:"type,omitempty"`
	// CreatedAt - Only return results whose item creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *StoresItemsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose item last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *StoresItemsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapStoresItemsListQueryFromJSON deserializes JSON data into a StoresItemsListQuery.
func MapStoresItemsListQueryFromJSON(data []byte) (*StoresItemsListQuery, error) {
	var v StoresItemsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresItemsListQueryToJSON serializes a StoresItemsListQuery to JSON.
func MapStoresItemsListQueryToJSON(v *StoresItemsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
