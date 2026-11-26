package comments

import (
	"encoding/json"
	"time"
)

// SkillsMergeRequestsCommentsUpdateOutputActorOrganizationActorMember represents the skills merge requests comments update output actor organization actor member type.
type SkillsMergeRequestsCommentsUpdateOutputActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsCommentsUpdateOutputActorOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsCommentsUpdateOutputActorOrganizationActorTeams struct {
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

// SkillsMergeRequestsCommentsUpdateOutputActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsMergeRequestsCommentsUpdateOutputActorOrganizationActor struct {
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
	Member   *SkillsMergeRequestsCommentsUpdateOutputActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsCommentsUpdateOutputActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsCommentsUpdateOutputActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsMergeRequestsCommentsUpdateOutputActorConsumer struct {
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

// SkillsMergeRequestsCommentsUpdateOutputActor - Actor behind a participant.
type SkillsMergeRequestsCommentsUpdateOutputActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsMergeRequestsCommentsUpdateOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsMergeRequestsCommentsUpdateOutputActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsCommentsUpdateOutput represents the skills merge requests comments update output type.
type SkillsMergeRequestsCommentsUpdateOutput struct {
	Object                  string  `json:"object"`
	Id                      string  `json:"id"`
	SkillMergeRequestItemId *string `json:"skill_merge_request_item_id,omitempty"`
	// Actor - Actor behind a participant.
	Actor              SkillsMergeRequestsCommentsUpdateOutputActor `json:"actor"`
	Body               string                                       `json:"body"`
	Path               *string                                      `json:"path,omitempty"`
	InReplyToCommentId *string                                      `json:"in_reply_to_comment_id,omitempty"`
	DeletedAt          *time.Time                                   `json:"deleted_at,omitempty"`
	CreatedAt          time.Time                                    `json:"created_at"`
	UpdatedAt          time.Time                                    `json:"updated_at"`
}

// MapSkillsMergeRequestsCommentsUpdateOutputFromJSON deserializes JSON data into a SkillsMergeRequestsCommentsUpdateOutput.
func MapSkillsMergeRequestsCommentsUpdateOutputFromJSON(data []byte) (*SkillsMergeRequestsCommentsUpdateOutput, error) {
	var v SkillsMergeRequestsCommentsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsCommentsUpdateOutputToJSON serializes a SkillsMergeRequestsCommentsUpdateOutput to JSON.
func MapSkillsMergeRequestsCommentsUpdateOutputToJSON(v *SkillsMergeRequestsCommentsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMergeRequestsCommentsUpdateBody represents the skills merge requests comments update body type.
type SkillsMergeRequestsCommentsUpdateBody struct {
	// Body - Comment text.
	Body string `json:"body"`
}

// MapSkillsMergeRequestsCommentsUpdateBodyFromJSON deserializes JSON data into a SkillsMergeRequestsCommentsUpdateBody.
func MapSkillsMergeRequestsCommentsUpdateBodyFromJSON(data []byte) (*SkillsMergeRequestsCommentsUpdateBody, error) {
	var v SkillsMergeRequestsCommentsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsCommentsUpdateBodyToJSON serializes a SkillsMergeRequestsCommentsUpdateBody to JSON.
func MapSkillsMergeRequestsCommentsUpdateBodyToJSON(v *SkillsMergeRequestsCommentsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
