package comments

import (
	"encoding/json"
	"time"
)

// DocumentsCommentThreadsCommentsUpdateOutputAuthorOrganizationActorMember represents the documents comment threads comments update output author organization actor member type.
type DocumentsCommentThreadsCommentsUpdateOutputAuthorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsCommentsUpdateOutputAuthorOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsCommentsUpdateOutputAuthorOrganizationActorTeams struct {
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

// DocumentsCommentThreadsCommentsUpdateOutputAuthorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type DocumentsCommentThreadsCommentsUpdateOutputAuthorOrganizationActor struct {
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
	Member   *DocumentsCommentThreadsCommentsUpdateOutputAuthorOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsCommentsUpdateOutputAuthorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsCommentsUpdateOutputAuthorConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsCommentThreadsCommentsUpdateOutputAuthorConsumer struct {
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

// DocumentsCommentThreadsCommentsUpdateOutputAuthor - Actor behind a participant.
type DocumentsCommentThreadsCommentsUpdateOutputAuthor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsCommentThreadsCommentsUpdateOutputAuthorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsCommentThreadsCommentsUpdateOutputAuthorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsCommentsUpdateOutput represents the documents comment threads comments update output type.
type DocumentsCommentThreadsCommentsUpdateOutput struct {
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
	Author DocumentsCommentThreadsCommentsUpdateOutputAuthor `json:"author"`
	// EditedAt - When the comment body was last edited.
	EditedAt *time.Time `json:"edited_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapDocumentsCommentThreadsCommentsUpdateOutputFromJSON deserializes JSON data into a DocumentsCommentThreadsCommentsUpdateOutput.
func MapDocumentsCommentThreadsCommentsUpdateOutputFromJSON(data []byte) (*DocumentsCommentThreadsCommentsUpdateOutput, error) {
	var v DocumentsCommentThreadsCommentsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCommentThreadsCommentsUpdateOutputToJSON serializes a DocumentsCommentThreadsCommentsUpdateOutput to JSON.
func MapDocumentsCommentThreadsCommentsUpdateOutputToJSON(v *DocumentsCommentThreadsCommentsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsCommentThreadsCommentsUpdateBody represents the documents comment threads comments update body type.
type DocumentsCommentThreadsCommentsUpdateBody struct {
	// Body - The new comment body
	Body string `json:"body"`
}

// MapDocumentsCommentThreadsCommentsUpdateBodyFromJSON deserializes JSON data into a DocumentsCommentThreadsCommentsUpdateBody.
func MapDocumentsCommentThreadsCommentsUpdateBodyFromJSON(data []byte) (*DocumentsCommentThreadsCommentsUpdateBody, error) {
	var v DocumentsCommentThreadsCommentsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCommentThreadsCommentsUpdateBodyToJSON serializes a DocumentsCommentThreadsCommentsUpdateBody to JSON.
func MapDocumentsCommentThreadsCommentsUpdateBodyToJSON(v *DocumentsCommentThreadsCommentsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
