package commentthreads

import (
	"encoding/json"
	"time"
)

// DocumentsCommentThreadsGetOutputCreatedByOrganizationActorMember represents the documents comment threads get output created by organization actor member type.
type DocumentsCommentThreadsGetOutputCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsGetOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsGetOutputCreatedByOrganizationActorTeams struct {
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

// DocumentsCommentThreadsGetOutputCreatedByOrganizationActor represents the documents comment threads get output created by organization actor type.
type DocumentsCommentThreadsGetOutputCreatedByOrganizationActor struct {
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
	Member   *DocumentsCommentThreadsGetOutputCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsGetOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsGetOutputCreatedByConsumer represents the documents comment threads get output created by consumer type.
type DocumentsCommentThreadsGetOutputCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsGetOutputCreatedBy represents the documents comment threads get output created by type.
type DocumentsCommentThreadsGetOutputCreatedBy struct {
	Type              string                                                      `json:"type"`
	Name              string                                                      `json:"name"`
	ImageUrl          *string                                                     `json:"image_url,omitempty"`
	Email             *string                                                     `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsGetOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsGetOutputCreatedByConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                             `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsGetOutputResolvedByOrganizationActorMember represents the documents comment threads get output resolved by organization actor member type.
type DocumentsCommentThreadsGetOutputResolvedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsGetOutputResolvedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsGetOutputResolvedByOrganizationActorTeams struct {
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

// DocumentsCommentThreadsGetOutputResolvedByOrganizationActor represents the documents comment threads get output resolved by organization actor type.
type DocumentsCommentThreadsGetOutputResolvedByOrganizationActor struct {
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
	ImageUrl string                                                             `json:"image_url"`
	Member   *DocumentsCommentThreadsGetOutputResolvedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsGetOutputResolvedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsGetOutputResolvedByConsumer represents the documents comment threads get output resolved by consumer type.
type DocumentsCommentThreadsGetOutputResolvedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsGetOutputResolvedBy represents the documents comment threads get output resolved by type.
type DocumentsCommentThreadsGetOutputResolvedBy struct {
	Type              string                                                       `json:"type"`
	Name              string                                                       `json:"name"`
	ImageUrl          *string                                                      `json:"image_url,omitempty"`
	Email             *string                                                      `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsGetOutputResolvedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsGetOutputResolvedByConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                              `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsGetOutputCommentsAuthorOrganizationActorMember represents the documents comment threads get output comments author organization actor member type.
type DocumentsCommentThreadsGetOutputCommentsAuthorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsGetOutputCommentsAuthorOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsGetOutputCommentsAuthorOrganizationActorTeams struct {
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

// DocumentsCommentThreadsGetOutputCommentsAuthorOrganizationActor represents the documents comment threads get output comments author organization actor type.
type DocumentsCommentThreadsGetOutputCommentsAuthorOrganizationActor struct {
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
	ImageUrl string                                                                 `json:"image_url"`
	Member   *DocumentsCommentThreadsGetOutputCommentsAuthorOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsGetOutputCommentsAuthorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsGetOutputCommentsAuthorConsumer represents the documents comment threads get output comments author consumer type.
type DocumentsCommentThreadsGetOutputCommentsAuthorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsGetOutputCommentsAuthor represents the documents comment threads get output comments author type.
type DocumentsCommentThreadsGetOutputCommentsAuthor struct {
	Type              string                                                           `json:"type"`
	Name              string                                                           `json:"name"`
	ImageUrl          *string                                                          `json:"image_url,omitempty"`
	Email             *string                                                          `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsGetOutputCommentsAuthorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsGetOutputCommentsAuthorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                  `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsGetOutputComments represents the documents comment threads get output comments type.
type DocumentsCommentThreadsGetOutputComments struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The unique identifier of the comment
	Id string `json:"id"`
	// ThreadId - The comment thread this comment belongs to
	ThreadId string `json:"thread_id"`
	// Body - The plain-text body of the comment
	Body string `json:"body"`
	// AuthorId - The resource actor id of the comment author
	AuthorId string                                         `json:"author_id"`
	Author   DocumentsCommentThreadsGetOutputCommentsAuthor `json:"author"`
	// EditedAt - When the comment body was last edited
	EditedAt  *time.Time `json:"edited_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// DocumentsCommentThreadsGetOutput represents the documents comment threads get output type.
type DocumentsCommentThreadsGetOutput struct {
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
	CreatedById string                                    `json:"created_by_id"`
	CreatedBy   DocumentsCommentThreadsGetOutputCreatedBy `json:"created_by"`
	// ResolvedById - The resource actor id of whoever resolved the thread
	ResolvedById *string                                     `json:"resolved_by_id,omitempty"`
	ResolvedBy   *DocumentsCommentThreadsGetOutputResolvedBy `json:"resolved_by,omitempty"`
	// ResolvedAt - When the thread was resolved
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	// Comments - The comments in this thread, oldest first
	Comments  []DocumentsCommentThreadsGetOutputComments `json:"comments"`
	CreatedAt time.Time                                  `json:"created_at"`
	UpdatedAt time.Time                                  `json:"updated_at"`
}

// MapDocumentsCommentThreadsGetOutputFromJSON deserializes JSON data into a DocumentsCommentThreadsGetOutput.
func MapDocumentsCommentThreadsGetOutputFromJSON(data []byte) (*DocumentsCommentThreadsGetOutput, error) {
	var v DocumentsCommentThreadsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCommentThreadsGetOutputToJSON serializes a DocumentsCommentThreadsGetOutput to JSON.
func MapDocumentsCommentThreadsGetOutputToJSON(v *DocumentsCommentThreadsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
