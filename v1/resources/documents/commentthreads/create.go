package commentthreads

import (
	"encoding/json"
	"time"
)

// DocumentsCommentThreadsCreateOutputCreatedByOrganizationActorMember represents the documents comment threads create output created by organization actor member type.
type DocumentsCommentThreadsCreateOutputCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsCreateOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsCreateOutputCreatedByOrganizationActorTeams struct {
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

// DocumentsCommentThreadsCreateOutputCreatedByOrganizationActor represents the documents comment threads create output created by organization actor type.
type DocumentsCommentThreadsCreateOutputCreatedByOrganizationActor struct {
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
	Member   *DocumentsCommentThreadsCreateOutputCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsCreateOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsCreateOutputCreatedByConsumer represents the documents comment threads create output created by consumer type.
type DocumentsCommentThreadsCreateOutputCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsCreateOutputCreatedBy represents the documents comment threads create output created by type.
type DocumentsCommentThreadsCreateOutputCreatedBy struct {
	Type              string                                                         `json:"type"`
	Name              string                                                         `json:"name"`
	ImageUrl          *string                                                        `json:"image_url,omitempty"`
	Email             *string                                                        `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsCreateOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsCreateOutputCreatedByConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsCreateOutputResolvedByOrganizationActorMember represents the documents comment threads create output resolved by organization actor member type.
type DocumentsCommentThreadsCreateOutputResolvedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsCreateOutputResolvedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsCreateOutputResolvedByOrganizationActorTeams struct {
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

// DocumentsCommentThreadsCreateOutputResolvedByOrganizationActor represents the documents comment threads create output resolved by organization actor type.
type DocumentsCommentThreadsCreateOutputResolvedByOrganizationActor struct {
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
	Member   *DocumentsCommentThreadsCreateOutputResolvedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsCreateOutputResolvedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsCreateOutputResolvedByConsumer represents the documents comment threads create output resolved by consumer type.
type DocumentsCommentThreadsCreateOutputResolvedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsCreateOutputResolvedBy represents the documents comment threads create output resolved by type.
type DocumentsCommentThreadsCreateOutputResolvedBy struct {
	Type              string                                                          `json:"type"`
	Name              string                                                          `json:"name"`
	ImageUrl          *string                                                         `json:"image_url,omitempty"`
	Email             *string                                                         `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsCreateOutputResolvedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsCreateOutputResolvedByConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                 `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsCreateOutputCommentsAuthorOrganizationActorMember represents the documents comment threads create output comments author organization actor member type.
type DocumentsCommentThreadsCreateOutputCommentsAuthorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsCreateOutputCommentsAuthorOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsCreateOutputCommentsAuthorOrganizationActorTeams struct {
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

// DocumentsCommentThreadsCreateOutputCommentsAuthorOrganizationActor represents the documents comment threads create output comments author organization actor type.
type DocumentsCommentThreadsCreateOutputCommentsAuthorOrganizationActor struct {
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
	Member   *DocumentsCommentThreadsCreateOutputCommentsAuthorOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsCreateOutputCommentsAuthorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsCreateOutputCommentsAuthorConsumer represents the documents comment threads create output comments author consumer type.
type DocumentsCommentThreadsCreateOutputCommentsAuthorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsCreateOutputCommentsAuthor represents the documents comment threads create output comments author type.
type DocumentsCommentThreadsCreateOutputCommentsAuthor struct {
	Type              string                                                              `json:"type"`
	Name              string                                                              `json:"name"`
	ImageUrl          *string                                                             `json:"image_url,omitempty"`
	Email             *string                                                             `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsCreateOutputCommentsAuthorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsCreateOutputCommentsAuthorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                     `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsCreateOutputComments represents the documents comment threads create output comments type.
type DocumentsCommentThreadsCreateOutputComments struct {
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
	Author   DocumentsCommentThreadsCreateOutputCommentsAuthor `json:"author"`
	// EditedAt - When the comment body was last edited
	EditedAt  *time.Time `json:"edited_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// DocumentsCommentThreadsCreateOutput represents the documents comment threads create output type.
type DocumentsCommentThreadsCreateOutput struct {
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
	CreatedBy   DocumentsCommentThreadsCreateOutputCreatedBy `json:"created_by"`
	// ResolvedById - The resource actor id of whoever resolved the thread
	ResolvedById *string                                        `json:"resolved_by_id,omitempty"`
	ResolvedBy   *DocumentsCommentThreadsCreateOutputResolvedBy `json:"resolved_by,omitempty"`
	// ResolvedAt - When the thread was resolved
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	// Comments - The comments in this thread, oldest first
	Comments  []DocumentsCommentThreadsCreateOutputComments `json:"comments"`
	CreatedAt time.Time                                     `json:"created_at"`
	UpdatedAt time.Time                                     `json:"updated_at"`
}

// MapDocumentsCommentThreadsCreateOutputFromJSON deserializes JSON data into a DocumentsCommentThreadsCreateOutput.
func MapDocumentsCommentThreadsCreateOutputFromJSON(data []byte) (*DocumentsCommentThreadsCreateOutput, error) {
	var v DocumentsCommentThreadsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCommentThreadsCreateOutputToJSON serializes a DocumentsCommentThreadsCreateOutput to JSON.
func MapDocumentsCommentThreadsCreateOutputToJSON(v *DocumentsCommentThreadsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsCommentThreadsCreateBodyAnchor - The text selection the thread is attached to
type DocumentsCommentThreadsCreateBodyAnchor struct {
	// StartBlockId - Block id where the selection starts
	StartBlockId string `json:"start_block_id"`
	// StartOffset - Text offset inside the start block
	StartOffset float64 `json:"start_offset"`
	// EndBlockId - Block id where the selection ends
	EndBlockId string `json:"end_block_id"`
	// EndOffset - Text offset inside the end block
	EndOffset float64 `json:"end_offset"`
	// Quote - The selected text
	Quote string `json:"quote"`
}

// DocumentsCommentThreadsCreateBody represents the documents comment threads create body type.
type DocumentsCommentThreadsCreateBody struct {
	// Anchor - The text selection the thread is attached to
	Anchor DocumentsCommentThreadsCreateBodyAnchor `json:"anchor"`
	// Body - The first comment in the thread
	Body string `json:"body"`
}

// MapDocumentsCommentThreadsCreateBodyFromJSON deserializes JSON data into a DocumentsCommentThreadsCreateBody.
func MapDocumentsCommentThreadsCreateBodyFromJSON(data []byte) (*DocumentsCommentThreadsCreateBody, error) {
	var v DocumentsCommentThreadsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCommentThreadsCreateBodyToJSON serializes a DocumentsCommentThreadsCreateBody to JSON.
func MapDocumentsCommentThreadsCreateBodyToJSON(v *DocumentsCommentThreadsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
