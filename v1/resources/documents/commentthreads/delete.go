package commentthreads

import (
	"encoding/json"
	"time"
)

// DocumentsCommentThreadsDeleteOutputCreatedByOrganizationActorMember represents the documents comment threads delete output created by organization actor member type.
type DocumentsCommentThreadsDeleteOutputCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsDeleteOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsDeleteOutputCreatedByOrganizationActorTeams struct {
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

// DocumentsCommentThreadsDeleteOutputCreatedByOrganizationActor represents the documents comment threads delete output created by organization actor type.
type DocumentsCommentThreadsDeleteOutputCreatedByOrganizationActor struct {
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
	ImageUrl string                                                               `json:"image_url"`
	Member   *DocumentsCommentThreadsDeleteOutputCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsDeleteOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsDeleteOutputCreatedByConsumer represents the documents comment threads delete output created by consumer type.
type DocumentsCommentThreadsDeleteOutputCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsDeleteOutputCreatedBy represents the documents comment threads delete output created by type.
type DocumentsCommentThreadsDeleteOutputCreatedBy struct {
	Type              string                                                         `json:"type"`
	Name              string                                                         `json:"name"`
	ImageUrl          *string                                                        `json:"image_url,omitempty"`
	Email             *string                                                        `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsDeleteOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsDeleteOutputCreatedByConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsDeleteOutputResolvedByOrganizationActorMember represents the documents comment threads delete output resolved by organization actor member type.
type DocumentsCommentThreadsDeleteOutputResolvedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsDeleteOutputResolvedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsDeleteOutputResolvedByOrganizationActorTeams struct {
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

// DocumentsCommentThreadsDeleteOutputResolvedByOrganizationActor represents the documents comment threads delete output resolved by organization actor type.
type DocumentsCommentThreadsDeleteOutputResolvedByOrganizationActor struct {
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
	ImageUrl string                                                                `json:"image_url"`
	Member   *DocumentsCommentThreadsDeleteOutputResolvedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsDeleteOutputResolvedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsDeleteOutputResolvedByConsumer represents the documents comment threads delete output resolved by consumer type.
type DocumentsCommentThreadsDeleteOutputResolvedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsDeleteOutputResolvedBy represents the documents comment threads delete output resolved by type.
type DocumentsCommentThreadsDeleteOutputResolvedBy struct {
	Type              string                                                          `json:"type"`
	Name              string                                                          `json:"name"`
	ImageUrl          *string                                                         `json:"image_url,omitempty"`
	Email             *string                                                         `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsDeleteOutputResolvedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsDeleteOutputResolvedByConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                 `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsDeleteOutputCommentsAuthorOrganizationActorMember represents the documents comment threads delete output comments author organization actor member type.
type DocumentsCommentThreadsDeleteOutputCommentsAuthorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsDeleteOutputCommentsAuthorOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsDeleteOutputCommentsAuthorOrganizationActorTeams struct {
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

// DocumentsCommentThreadsDeleteOutputCommentsAuthorOrganizationActor represents the documents comment threads delete output comments author organization actor type.
type DocumentsCommentThreadsDeleteOutputCommentsAuthorOrganizationActor struct {
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
	ImageUrl string                                                                    `json:"image_url"`
	Member   *DocumentsCommentThreadsDeleteOutputCommentsAuthorOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsDeleteOutputCommentsAuthorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsDeleteOutputCommentsAuthorConsumer represents the documents comment threads delete output comments author consumer type.
type DocumentsCommentThreadsDeleteOutputCommentsAuthorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsDeleteOutputCommentsAuthor represents the documents comment threads delete output comments author type.
type DocumentsCommentThreadsDeleteOutputCommentsAuthor struct {
	Type              string                                                              `json:"type"`
	Name              string                                                              `json:"name"`
	ImageUrl          *string                                                             `json:"image_url,omitempty"`
	Email             *string                                                             `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsDeleteOutputCommentsAuthorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsDeleteOutputCommentsAuthorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                     `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsDeleteOutputComments represents the documents comment threads delete output comments type.
type DocumentsCommentThreadsDeleteOutputComments struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The unique identifier of the comment
	Id string `json:"id"`
	// ThreadId - The comment thread this comment belongs to
	ThreadId string `json:"thread_id"`
	// Body - The plain-text body of the comment
	Body string `json:"body"`
	// AuthorId - The resource actor id of the comment author
	AuthorId string                                            `json:"author_id"`
	Author   DocumentsCommentThreadsDeleteOutputCommentsAuthor `json:"author"`
	// EditedAt - When the comment body was last edited
	EditedAt  *time.Time `json:"edited_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// DocumentsCommentThreadsDeleteOutput represents the documents comment threads delete output type.
type DocumentsCommentThreadsDeleteOutput struct {
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
	CreatedById string                                       `json:"created_by_id"`
	CreatedBy   DocumentsCommentThreadsDeleteOutputCreatedBy `json:"created_by"`
	// ResolvedById - The resource actor id of whoever resolved the thread
	ResolvedById *string                                        `json:"resolved_by_id,omitempty"`
	ResolvedBy   *DocumentsCommentThreadsDeleteOutputResolvedBy `json:"resolved_by,omitempty"`
	// ResolvedAt - When the thread was resolved
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	// Comments - The comments in this thread, oldest first
	Comments  []DocumentsCommentThreadsDeleteOutputComments `json:"comments"`
	CreatedAt time.Time                                     `json:"created_at"`
	UpdatedAt time.Time                                     `json:"updated_at"`
}

// MapDocumentsCommentThreadsDeleteOutputFromJSON deserializes JSON data into a DocumentsCommentThreadsDeleteOutput.
func MapDocumentsCommentThreadsDeleteOutputFromJSON(data []byte) (*DocumentsCommentThreadsDeleteOutput, error) {
	var v DocumentsCommentThreadsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCommentThreadsDeleteOutputToJSON serializes a DocumentsCommentThreadsDeleteOutput to JSON.
func MapDocumentsCommentThreadsDeleteOutputToJSON(v *DocumentsCommentThreadsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
