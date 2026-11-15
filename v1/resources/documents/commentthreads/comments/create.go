package comments

import (
	"encoding/json"
	"time"
)

// DocumentsCommentThreadsCommentsCreateOutputAuthorOrganizationActorMember represents the documents comment threads comments create output author organization actor member type.
type DocumentsCommentThreadsCommentsCreateOutputAuthorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsCommentsCreateOutputAuthorOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsCommentsCreateOutputAuthorOrganizationActorTeams struct {
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

// DocumentsCommentThreadsCommentsCreateOutputAuthorOrganizationActor represents the documents comment threads comments create output author organization actor type.
type DocumentsCommentThreadsCommentsCreateOutputAuthorOrganizationActor struct {
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
	Member   *DocumentsCommentThreadsCommentsCreateOutputAuthorOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsCommentsCreateOutputAuthorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsCommentsCreateOutputAuthorConsumer represents the documents comment threads comments create output author consumer type.
type DocumentsCommentThreadsCommentsCreateOutputAuthorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsCommentsCreateOutputAuthor represents the documents comment threads comments create output author type.
type DocumentsCommentThreadsCommentsCreateOutputAuthor struct {
	Type              string                                                              `json:"type"`
	Name              string                                                              `json:"name"`
	ImageUrl          *string                                                             `json:"image_url,omitempty"`
	Email             *string                                                             `json:"email,omitempty"`
	OrganizationActor *DocumentsCommentThreadsCommentsCreateOutputAuthorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *DocumentsCommentThreadsCommentsCreateOutputAuthorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                     `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsCommentsCreateOutput represents the documents comment threads comments create output type.
type DocumentsCommentThreadsCommentsCreateOutput struct {
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
	Author   DocumentsCommentThreadsCommentsCreateOutputAuthor `json:"author"`
	// EditedAt - When the comment body was last edited
	EditedAt  *time.Time `json:"edited_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// MapDocumentsCommentThreadsCommentsCreateOutputFromJSON deserializes JSON data into a DocumentsCommentThreadsCommentsCreateOutput.
func MapDocumentsCommentThreadsCommentsCreateOutputFromJSON(data []byte) (*DocumentsCommentThreadsCommentsCreateOutput, error) {
	var v DocumentsCommentThreadsCommentsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCommentThreadsCommentsCreateOutputToJSON serializes a DocumentsCommentThreadsCommentsCreateOutput to JSON.
func MapDocumentsCommentThreadsCommentsCreateOutputToJSON(v *DocumentsCommentThreadsCommentsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsCommentThreadsCommentsCreateBody represents the documents comment threads comments create body type.
type DocumentsCommentThreadsCommentsCreateBody struct {
	// Body - The comment body
	Body string `json:"body"`
}

// MapDocumentsCommentThreadsCommentsCreateBodyFromJSON deserializes JSON data into a DocumentsCommentThreadsCommentsCreateBody.
func MapDocumentsCommentThreadsCommentsCreateBodyFromJSON(data []byte) (*DocumentsCommentThreadsCommentsCreateBody, error) {
	var v DocumentsCommentThreadsCommentsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCommentThreadsCommentsCreateBodyToJSON serializes a DocumentsCommentThreadsCommentsCreateBody to JSON.
func MapDocumentsCommentThreadsCommentsCreateBodyToJSON(v *DocumentsCommentThreadsCommentsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
