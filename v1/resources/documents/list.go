package documents

import (
	"encoding/json"
	"time"
)

// DocumentsListOutputItemsCreatedByOrganizationActorMember represents the documents list output items created by organization actor member type.
type DocumentsListOutputItemsCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsListOutputItemsCreatedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsListOutputItemsCreatedByOrganizationActorTeams struct {
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

// DocumentsListOutputItemsCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type DocumentsListOutputItemsCreatedByOrganizationActor struct {
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
	Member   *DocumentsListOutputItemsCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsListOutputItemsCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsListOutputItemsCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsListOutputItemsCreatedByConsumer struct {
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

// DocumentsListOutputItemsCreatedBy - Actor that created the document.
type DocumentsListOutputItemsCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsListOutputItemsCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsListOutputItemsCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsListOutputItems represents the documents list output items type.
type DocumentsListOutputItems struct {
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
	CreatedBy *DocumentsListOutputItemsCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsListOutputPagination represents the documents list output pagination type.
type DocumentsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// DocumentsListOutput represents the documents list output type.
type DocumentsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                        `json:"object"`
	Items      []DocumentsListOutputItems    `json:"items"`
	Pagination DocumentsListOutputPagination `json:"pagination"`
}

// MapDocumentsListOutputFromJSON deserializes JSON data into a DocumentsListOutput.
func MapDocumentsListOutputFromJSON(data []byte) (*DocumentsListOutput, error) {
	var v DocumentsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsListOutputToJSON serializes a DocumentsListOutput to JSON.
func MapDocumentsListOutputToJSON(v *DocumentsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsListQueryCreatedAt - Only return results whose document creation time is in this range. Set `gt`, `lt`, or both.
type DocumentsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsListQueryUpdatedAt - Only return results whose document last update time is in this range. Set `gt`, `lt`, or both.
type DocumentsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsListQuery represents the documents list query type.
type DocumentsListQuery struct {
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
	// Id - Only return documents with one of these IDs.
	Id *any `json:"id,omitempty"`
	// FileId - Only return documents stored in one of these files.
	FileId *any `json:"file_id,omitempty"`
	// StoreId - Only return documents in one of these stores.
	StoreId *any `json:"store_id,omitempty"`
	// ParentDocumentId - Only return documents cloned from one of these documents.
	ParentDocumentId *any `json:"parent_document_id,omitempty"`
	// CreatedAt - Only return results whose document creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *DocumentsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose document last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *DocumentsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapDocumentsListQueryFromJSON deserializes JSON data into a DocumentsListQuery.
func MapDocumentsListQueryFromJSON(data []byte) (*DocumentsListQuery, error) {
	var v DocumentsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsListQueryToJSON serializes a DocumentsListQuery to JSON.
func MapDocumentsListQueryToJSON(v *DocumentsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
