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

// SkillsMergeRequestsCommentsUpdateOutputActorOrganizationActor represents the skills merge requests comments update output actor organization actor type.
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

// SkillsMergeRequestsCommentsUpdateOutputActorConsumer represents the skills merge requests comments update output actor consumer type.
type SkillsMergeRequestsCommentsUpdateOutputActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsCommentsUpdateOutputActor represents the skills merge requests comments update output actor type.
type SkillsMergeRequestsCommentsUpdateOutputActor struct {
	Type              string                                                         `json:"type"`
	Name              string                                                         `json:"name"`
	ImageUrl          *string                                                        `json:"image_url,omitempty"`
	Email             *string                                                        `json:"email,omitempty"`
	OrganizationActor *SkillsMergeRequestsCommentsUpdateOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsMergeRequestsCommentsUpdateOutputActorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsCommentsUpdateOutput represents the skills merge requests comments update output type.
type SkillsMergeRequestsCommentsUpdateOutput struct {
	Object                  string                                       `json:"object"`
	Id                      string                                       `json:"id"`
	SkillMergeRequestItemId *string                                      `json:"skill_merge_request_item_id,omitempty"`
	Actor                   SkillsMergeRequestsCommentsUpdateOutputActor `json:"actor"`
	Body                    string                                       `json:"body"`
	Path                    *string                                      `json:"path,omitempty"`
	InReplyToCommentId      *string                                      `json:"in_reply_to_comment_id,omitempty"`
	DeletedAt               *time.Time                                   `json:"deleted_at,omitempty"`
	CreatedAt               time.Time                                    `json:"created_at"`
	UpdatedAt               time.Time                                    `json:"updated_at"`
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
