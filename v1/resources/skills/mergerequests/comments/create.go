package comments

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// SkillsMergeRequestsCommentsCreateOutputActorOrganizationActorMember represents the skills merge requests comments create output actor organization actor member type.
type SkillsMergeRequestsCommentsCreateOutputActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsCommentsCreateOutputActorOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsCommentsCreateOutputActorOrganizationActorTeams struct {
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

// SkillsMergeRequestsCommentsCreateOutputActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsMergeRequestsCommentsCreateOutputActorOrganizationActor struct {
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
	Member   *SkillsMergeRequestsCommentsCreateOutputActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsCommentsCreateOutputActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsCommentsCreateOutputActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsMergeRequestsCommentsCreateOutputActorConsumer struct {
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

// SkillsMergeRequestsCommentsCreateOutputActor - Actor behind a participant.
type SkillsMergeRequestsCommentsCreateOutputActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsMergeRequestsCommentsCreateOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsMergeRequestsCommentsCreateOutputActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsCommentsCreateOutput represents the skills merge requests comments create output type.
type SkillsMergeRequestsCommentsCreateOutput struct {
	Object                  string  `json:"object"`
	Id                      string  `json:"id"`
	SkillMergeRequestItemId *string `json:"skill_merge_request_item_id,omitempty"`
	// Actor - Actor behind a participant.
	Actor              SkillsMergeRequestsCommentsCreateOutputActor `json:"actor"`
	Body               string                                       `json:"body"`
	Path               *string                                      `json:"path,omitempty"`
	InReplyToCommentId *string                                      `json:"in_reply_to_comment_id,omitempty"`
	DeletedAt          *time.Time                                   `json:"deleted_at,omitempty"`
	CreatedAt          time.Time                                    `json:"created_at"`
	UpdatedAt          time.Time                                    `json:"updated_at"`
}

// MapSkillsMergeRequestsCommentsCreateOutputFromJSON deserializes JSON data into a SkillsMergeRequestsCommentsCreateOutput.
func MapSkillsMergeRequestsCommentsCreateOutputFromJSON(data []byte) (*SkillsMergeRequestsCommentsCreateOutput, error) {
	var v SkillsMergeRequestsCommentsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsCommentsCreateOutputToJSON serializes a SkillsMergeRequestsCommentsCreateOutput to JSON.
func MapSkillsMergeRequestsCommentsCreateOutputToJSON(v *SkillsMergeRequestsCommentsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMergeRequestsCommentsCreateBody represents the skills merge requests comments create body type.
type SkillsMergeRequestsCommentsCreateBody struct {
	// ItemId - Proposed change to comment on. Omit to comment on the merge request.
	ItemId *string `json:"item_id,omitempty"`
	// InReplyToCommentId - Comment this replies to, by ID.
	InReplyToCommentId *string `json:"in_reply_to_comment_id,omitempty"`
	// Body - Comment text.
	Body string `json:"body"`
	// Path - File path the comment refers to. `null` for general comments.
	Path *nullable.Nullable[string] `json:"path,omitempty"`
}

// MapSkillsMergeRequestsCommentsCreateBodyFromJSON deserializes JSON data into a SkillsMergeRequestsCommentsCreateBody.
func MapSkillsMergeRequestsCommentsCreateBodyFromJSON(data []byte) (*SkillsMergeRequestsCommentsCreateBody, error) {
	var v SkillsMergeRequestsCommentsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsCommentsCreateBodyToJSON serializes a SkillsMergeRequestsCommentsCreateBody to JSON.
func MapSkillsMergeRequestsCommentsCreateBodyToJSON(v *SkillsMergeRequestsCommentsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
