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

// DocumentsCommentThreadsCommentsCreateOutputAuthorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
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

// DocumentsCommentThreadsCommentsCreateOutputAuthorConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsCommentThreadsCommentsCreateOutputAuthorConsumer struct {
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

// DocumentsCommentThreadsCommentsCreateOutputAuthor - Actor behind a participant.
type DocumentsCommentThreadsCommentsCreateOutputAuthor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsCommentThreadsCommentsCreateOutputAuthorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsCommentThreadsCommentsCreateOutputAuthorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsCommentsCreateOutput represents the documents comment threads comments create output type.
type DocumentsCommentThreadsCommentsCreateOutput struct {
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
	Author DocumentsCommentThreadsCommentsCreateOutputAuthor `json:"author"`
	// EditedAt - When the comment body was last edited.
	EditedAt *time.Time `json:"edited_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
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
