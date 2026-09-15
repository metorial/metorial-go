package comments

import (
	"encoding/json"
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

// SkillsMergeRequestsCommentsCreateOutputActorOrganizationActor represents the skills merge requests comments create output actor organization actor type.
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

// SkillsMergeRequestsCommentsCreateOutputActorConsumer represents the skills merge requests comments create output actor consumer type.
type SkillsMergeRequestsCommentsCreateOutputActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsCommentsCreateOutputActor represents the skills merge requests comments create output actor type.
type SkillsMergeRequestsCommentsCreateOutputActor struct {
	Type              string                                                         `json:"type"`
	Name              string                                                         `json:"name"`
	ImageUrl          *string                                                        `json:"image_url,omitempty"`
	Email             *string                                                        `json:"email,omitempty"`
	OrganizationActor *SkillsMergeRequestsCommentsCreateOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsMergeRequestsCommentsCreateOutputActorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsCommentsCreateOutput represents the skills merge requests comments create output type.
type SkillsMergeRequestsCommentsCreateOutput struct {
	Object                  string                                       `json:"object"`
	Id                      string                                       `json:"id"`
	SkillMergeRequestItemId *string                                      `json:"skill_merge_request_item_id,omitempty"`
	Actor                   SkillsMergeRequestsCommentsCreateOutputActor `json:"actor"`
	Body                    string                                       `json:"body"`
	Path                    *string                                      `json:"path,omitempty"`
	InReplyToCommentId      *string                                      `json:"in_reply_to_comment_id,omitempty"`
	DeletedAt               *time.Time                                   `json:"deleted_at,omitempty"`
	CreatedAt               time.Time                                    `json:"created_at"`
	UpdatedAt               time.Time                                    `json:"updated_at"`
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
	ItemId             *string `json:"item_id,omitempty"`
	InReplyToCommentId *string `json:"in_reply_to_comment_id,omitempty"`
	Body               string  `json:"body"`
	Path               *string `json:"path,omitempty"`
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
