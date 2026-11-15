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

// DocumentsCommentThreadsListOutputItemsCreatedByOrganizationActor represents the documents comment threads list output items created by organization actor type.
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

// DocumentsCommentThreadsListOutputItemsCreatedByConsumer represents the documents comment threads list output items created by consumer type.
type DocumentsCommentThreadsListOutputItemsCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsListOutputItemsCreatedBy represents the documents comment threads list output items created by type.
type DocumentsCommentThreadsListOutputItemsCreatedBy struct {
	Type              string                                                            `json:"type"`
	Name              string                                                            `json:"name"`
	ImageUrl          *string                                                           `json:"image_url,omitempty"`
	Email             *string                                                           `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsListOutputItemsCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsListOutputItemsCreatedByConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                   `json:"consumer_profile,omitempty"`
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

// DocumentsCommentThreadsListOutputItemsResolvedByOrganizationActor represents the documents comment threads list output items resolved by organization actor type.
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

// DocumentsCommentThreadsListOutputItemsResolvedByConsumer represents the documents comment threads list output items resolved by consumer type.
type DocumentsCommentThreadsListOutputItemsResolvedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsListOutputItemsResolvedBy represents the documents comment threads list output items resolved by type.
type DocumentsCommentThreadsListOutputItemsResolvedBy struct {
	Type              string                                                             `json:"type"`
	Name              string                                                             `json:"name"`
	ImageUrl          *string                                                            `json:"image_url,omitempty"`
	Email             *string                                                            `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsListOutputItemsResolvedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsListOutputItemsResolvedByConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                    `json:"consumer_profile,omitempty"`
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

// DocumentsCommentThreadsListOutputItemsCommentsAuthorOrganizationActor represents the documents comment threads list output items comments author organization actor type.
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

// DocumentsCommentThreadsListOutputItemsCommentsAuthorConsumer represents the documents comment threads list output items comments author consumer type.
type DocumentsCommentThreadsListOutputItemsCommentsAuthorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsListOutputItemsCommentsAuthor represents the documents comment threads list output items comments author type.
type DocumentsCommentThreadsListOutputItemsCommentsAuthor struct {
	Type              string                                                                 `json:"type"`
	Name              string                                                                 `json:"name"`
	ImageUrl          *string                                                                `json:"image_url,omitempty"`
	Email             *string                                                                `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsListOutputItemsCommentsAuthorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsListOutputItemsCommentsAuthorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                        `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsListOutputItemsComments represents the documents comment threads list output items comments type.
type DocumentsCommentThreadsListOutputItemsComments struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The unique identifier of the comment
	Id string `json:"id"`
	// ThreadId - The comment thread this comment belongs to
	ThreadId string `json:"thread_id"`
	// Body - The plain-text body of the comment
	Body string `json:"body"`
	// AuthorId - The resource actor id of the comment author
	AuthorId string                                               `json:"author_id"`
	Author   DocumentsCommentThreadsListOutputItemsCommentsAuthor `json:"author"`
	// EditedAt - When the comment body was last edited
	EditedAt  *time.Time `json:"edited_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// DocumentsCommentThreadsListOutputItems represents the documents comment threads list output items type.
type DocumentsCommentThreadsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The unique identifier of the comment thread
	Id string `json:"id"`
	// DocumentId - The document the thread is attached to
	DocumentId string `json:"document_id"`
	// Status - Whether the discussion is still open
	Status string `json:"status"`
	// AnchorStatus - Whether the commented text still exists in the document. Detached threads lost their text through edits.
	AnchorStatus string `json:"anchor_status"`
	// Quote - The text that was selected when the thread was created
	Quote string `json:"quote"`
	// CreatedById - The resource actor id of the thread author
	CreatedById string                                          `json:"created_by_id"`
	CreatedBy   DocumentsCommentThreadsListOutputItemsCreatedBy `json:"created_by"`
	// ResolvedById - The resource actor id of whoever resolved the thread
	ResolvedById *string                                           `json:"resolved_by_id,omitempty"`
	ResolvedBy   *DocumentsCommentThreadsListOutputItemsResolvedBy `json:"resolved_by,omitempty"`
	// ResolvedAt - When the thread was resolved
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	// Comments - The comments in this thread, oldest first
	Comments  []DocumentsCommentThreadsListOutputItemsComments `json:"comments"`
	CreatedAt time.Time                                        `json:"created_at"`
	UpdatedAt time.Time                                        `json:"updated_at"`
}

// DocumentsCommentThreadsListOutputPagination represents the documents comment threads list output pagination type.
type DocumentsCommentThreadsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// DocumentsCommentThreadsListOutput represents the documents comment threads list output type.
type DocumentsCommentThreadsListOutput struct {
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

// DocumentsCommentThreadsListQueryCreatedAt - Filter Filter by creation time by date range
type DocumentsCommentThreadsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for Filter by creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// DocumentsCommentThreadsListQuery represents the documents comment threads list query type.
type DocumentsCommentThreadsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by comment thread ID
	Id *any `json:"id,omitempty"`
	// Status - Filter by thread status
	Status *any `json:"status,omitempty"`
	// AnchorStatus - Filter by whether the commented text still exists
	AnchorStatus *any `json:"anchor_status,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
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
