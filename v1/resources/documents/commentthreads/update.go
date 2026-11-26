package commentthreads

import (
	"encoding/json"
	"time"
)

// DocumentsCommentThreadsUpdateOutputCreatedByOrganizationActorMember represents the documents comment threads update output created by organization actor member type.
type DocumentsCommentThreadsUpdateOutputCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsUpdateOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsUpdateOutputCreatedByOrganizationActorTeams struct {
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

// DocumentsCommentThreadsUpdateOutputCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type DocumentsCommentThreadsUpdateOutputCreatedByOrganizationActor struct {
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
	Member   *DocumentsCommentThreadsUpdateOutputCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsUpdateOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsUpdateOutputCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsCommentThreadsUpdateOutputCreatedByConsumer struct {
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

// DocumentsCommentThreadsUpdateOutputCreatedBy - Actor behind a participant.
type DocumentsCommentThreadsUpdateOutputCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsCommentThreadsUpdateOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsCommentThreadsUpdateOutputCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsUpdateOutputResolvedByOrganizationActorMember represents the documents comment threads update output resolved by organization actor member type.
type DocumentsCommentThreadsUpdateOutputResolvedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsUpdateOutputResolvedByOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsUpdateOutputResolvedByOrganizationActorTeams struct {
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

// DocumentsCommentThreadsUpdateOutputResolvedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type DocumentsCommentThreadsUpdateOutputResolvedByOrganizationActor struct {
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
	Member   *DocumentsCommentThreadsUpdateOutputResolvedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsUpdateOutputResolvedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsUpdateOutputResolvedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsCommentThreadsUpdateOutputResolvedByConsumer struct {
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

// DocumentsCommentThreadsUpdateOutputResolvedBy - Actor behind a participant.
type DocumentsCommentThreadsUpdateOutputResolvedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsCommentThreadsUpdateOutputResolvedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsCommentThreadsUpdateOutputResolvedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsUpdateOutputCommentsAuthorOrganizationActorMember represents the documents comment threads update output comments author organization actor member type.
type DocumentsCommentThreadsUpdateOutputCommentsAuthorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// DocumentsCommentThreadsUpdateOutputCommentsAuthorOrganizationActorTeams - The teams the actor belongs to
type DocumentsCommentThreadsUpdateOutputCommentsAuthorOrganizationActorTeams struct {
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

// DocumentsCommentThreadsUpdateOutputCommentsAuthorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type DocumentsCommentThreadsUpdateOutputCommentsAuthorOrganizationActor struct {
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
	Member   *DocumentsCommentThreadsUpdateOutputCommentsAuthorOrganizationActorMember `json:"member,omitempty"`
	Teams    []DocumentsCommentThreadsUpdateOutputCommentsAuthorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsUpdateOutputCommentsAuthorConsumer - Consumer details. `null` unless `type` is `consumer`.
type DocumentsCommentThreadsUpdateOutputCommentsAuthorConsumer struct {
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

// DocumentsCommentThreadsUpdateOutputCommentsAuthor - Actor behind a participant.
type DocumentsCommentThreadsUpdateOutputCommentsAuthor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *DocumentsCommentThreadsUpdateOutputCommentsAuthorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *DocumentsCommentThreadsUpdateOutputCommentsAuthorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// DocumentsCommentThreadsUpdateOutputComments represents the documents comment threads update output comments type.
type DocumentsCommentThreadsUpdateOutputComments struct {
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
	Author DocumentsCommentThreadsUpdateOutputCommentsAuthor `json:"author"`
	// EditedAt - When the comment body was last edited.
	EditedAt *time.Time `json:"edited_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentsCommentThreadsUpdateOutput represents the documents comment threads update output type.
type DocumentsCommentThreadsUpdateOutput struct {
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
	CreatedBy DocumentsCommentThreadsUpdateOutputCreatedBy `json:"created_by"`
	// ResolvedById - The resource actor id of whoever resolved the thread.
	ResolvedById *string `json:"resolved_by_id,omitempty"`
	// ResolvedBy - Actor behind a participant.
	ResolvedBy *DocumentsCommentThreadsUpdateOutputResolvedBy `json:"resolved_by,omitempty"`
	// ResolvedAt - When the thread was marked resolved.
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	// Comments - The comments in this thread, oldest first.
	Comments []DocumentsCommentThreadsUpdateOutputComments `json:"comments"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapDocumentsCommentThreadsUpdateOutputFromJSON deserializes JSON data into a DocumentsCommentThreadsUpdateOutput.
func MapDocumentsCommentThreadsUpdateOutputFromJSON(data []byte) (*DocumentsCommentThreadsUpdateOutput, error) {
	var v DocumentsCommentThreadsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCommentThreadsUpdateOutputToJSON serializes a DocumentsCommentThreadsUpdateOutput to JSON.
func MapDocumentsCommentThreadsUpdateOutputToJSON(v *DocumentsCommentThreadsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// DocumentsCommentThreadsUpdateBody represents the documents comment threads update body type.
type DocumentsCommentThreadsUpdateBody struct {
	// Status - New status for the thread.
	Status *string `json:"status,omitempty"`
}

// MapDocumentsCommentThreadsUpdateBodyFromJSON deserializes JSON data into a DocumentsCommentThreadsUpdateBody.
func MapDocumentsCommentThreadsUpdateBodyFromJSON(data []byte) (*DocumentsCommentThreadsUpdateBody, error) {
	var v DocumentsCommentThreadsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsCommentThreadsUpdateBodyToJSON serializes a DocumentsCommentThreadsUpdateBody to JSON.
func MapDocumentsCommentThreadsUpdateBodyToJSON(v *DocumentsCommentThreadsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
