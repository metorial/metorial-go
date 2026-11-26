package commentthreads

import (
	"encoding/json"
	"time"
)

// DocumentsCommentThreadsListOutputItemsCreatedByOrganizationActorMember represents the documents comment threads list output items created by organization actor member type.
type DocumentsCommentThreadsListOutputItemsCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsListOutputItemsCreatedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsListOutputItemsCreatedByOrganizationActorTeams struct {
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

// DocumentsCommentThreadsListOutputItemsCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type DocumentsCommentThreadsListOutputItemsCreatedByOrganizationActor struct {
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
	ImageUrl string                                                                  `json:"image_url"`
	Member   *DocumentsCommentThreadsListOutputItemsCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsListOutputItemsCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsListOutputItemsCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsCommentThreadsListOutputItemsCreatedByConsumer struct {
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

// DocumentsCommentThreadsListOutputItemsCreatedBy - Actor behind a participant.
type DocumentsCommentThreadsListOutputItemsCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsCommentThreadsListOutputItemsCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsCommentThreadsListOutputItemsCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsListOutputItemsResolvedByOrganizationActorMember represents the documents comment threads list output items resolved by organization actor member type.
type DocumentsCommentThreadsListOutputItemsResolvedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsListOutputItemsResolvedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsListOutputItemsResolvedByOrganizationActorTeams struct {
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

// DocumentsCommentThreadsListOutputItemsResolvedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type DocumentsCommentThreadsListOutputItemsResolvedByOrganizationActor struct {
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
	ImageUrl string                                                                   `json:"image_url"`
	Member   *DocumentsCommentThreadsListOutputItemsResolvedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsListOutputItemsResolvedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsListOutputItemsResolvedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsCommentThreadsListOutputItemsResolvedByConsumer struct {
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

// DocumentsCommentThreadsListOutputItemsResolvedBy - Actor behind a participant.
type DocumentsCommentThreadsListOutputItemsResolvedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsCommentThreadsListOutputItemsResolvedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsCommentThreadsListOutputItemsResolvedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsListOutputItemsCommentsAuthorOrganizationActorMember represents the documents comment threads list output items comments author organization actor member type.
type DocumentsCommentThreadsListOutputItemsCommentsAuthorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsListOutputItemsCommentsAuthorOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsListOutputItemsCommentsAuthorOrganizationActorTeams struct {
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

// DocumentsCommentThreadsListOutputItemsCommentsAuthorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type DocumentsCommentThreadsListOutputItemsCommentsAuthorOrganizationActor struct {
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
	ImageUrl string                                                                       `json:"image_url"`
	Member   *DocumentsCommentThreadsListOutputItemsCommentsAuthorOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsListOutputItemsCommentsAuthorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsListOutputItemsCommentsAuthorConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsCommentThreadsListOutputItemsCommentsAuthorConsumer struct {
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

// DocumentsCommentThreadsListOutputItemsCommentsAuthor - Actor behind a participant.
type DocumentsCommentThreadsListOutputItemsCommentsAuthor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsCommentThreadsListOutputItemsCommentsAuthorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsCommentThreadsListOutputItemsCommentsAuthorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsListOutputItemsComments represents the documents comment threads list output items comments type.
type DocumentsCommentThreadsListOutputItemsComments struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ThreadId - Thread this comment belongs to.
	ThreadId string `json:"thread_id"`
	// Body - The plain-text body of the comment.
	Body string `json:"body"`
	// AuthorId - The resource actor id of the comment author.
	AuthorId string `json:"author_id"`
	// Author - Actor behind a participant.
	Author DocumentsCommentThreadsListOutputItemsCommentsAuthor `json:"author"`
	// EditedAt - When the comment body was last edited.
	EditedAt *time.Time `json:"edited_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsListOutputItems represents the documents comment threads list output items type.
type DocumentsCommentThreadsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// DocumentId - Document this thread is attached to.
	DocumentId string `json:"document_id"`
	// Status - Whether the discussion is still open.
	Status string `json:"status"`
	// AnchorStatus - Whether the commented text still exists in the document. Detached threads lost their text through edits.
	AnchorStatus string `json:"anchor_status"`
	// Quote - The text that was selected when the thread was created.
	Quote string `json:"quote"`
	// CreatedById - The resource actor id of the thread author.
	CreatedById string `json:"created_by_id"`
	// CreatedBy - Actor behind a participant.
	CreatedBy DocumentsCommentThreadsListOutputItemsCreatedBy `json:"created_by"`
	// ResolvedById - The resource actor id of whoever resolved the thread.
	ResolvedById *string `json:"resolved_by_id,omitempty"`
	// ResolvedBy - Actor behind a participant.
	ResolvedBy *DocumentsCommentThreadsListOutputItemsResolvedBy `json:"resolved_by,omitempty"`
	// ResolvedAt - When the thread was marked resolved.
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	// Comments - The comments in this thread, oldest first.
	Comments []DocumentsCommentThreadsListOutputItemsComments `json:"comments"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsListOutputPagination represents the documents comment threads list output pagination type.
type DocumentsCommentThreadsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// DocumentsCommentThreadsListOutput represents the documents comment threads list output type.
type DocumentsCommentThreadsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                      `json:"object"`
	Items      []DocumentsCommentThreadsListOutputItems    `json:"items"`
	Pagination DocumentsCommentThreadsListOutputPagination `json:"pagination"`
}

// MapDocumentsCommentThreadsListOutputFromJSON deserializes JSON data into a DocumentsCommentThreadsListOutput.
func MapDocumentsCommentThreadsListOutputFromJSON(data []byte) (*DocumentsCommentThreadsListOutput, error) {
	var v DocumentsCommentThreadsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCommentThreadsListOutputToJSON serializes a DocumentsCommentThreadsListOutput to JSON.
func MapDocumentsCommentThreadsListOutputToJSON(v *DocumentsCommentThreadsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsCommentThreadsListQueryCreatedAt - Only return results whose thread creation time is in this range. Set `gt`, `lt`, or both.
type DocumentsCommentThreadsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsCommentThreadsListQuery represents the documents comment threads list query type.
type DocumentsCommentThreadsListQuery struct {
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
	// Id - Only return threads with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Status - Only return threads with one of these statuses.
	Status *any `json:"status,omitempty"`
	// AnchorStatus - Filter by whether the commented text still exists
	AnchorStatus *any `json:"anchor_status,omitempty"`
	// CreatedAt - Only return results whose thread creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *DocumentsCommentThreadsListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapDocumentsCommentThreadsListQueryFromJSON deserializes JSON data into a DocumentsCommentThreadsListQuery.
func MapDocumentsCommentThreadsListQueryFromJSON(data []byte) (*DocumentsCommentThreadsListQuery, error) {
	var v DocumentsCommentThreadsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCommentThreadsListQueryToJSON serializes a DocumentsCommentThreadsListQuery to JSON.
func MapDocumentsCommentThreadsListQueryToJSON(v *DocumentsCommentThreadsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
