package comments

import (
	"encoding/json"
	"time"
)

// SkillsMergeRequestsCommentsGetOutputActorOrganizationActorMember represents the skills merge requests comments get output actor organization actor member type.
type SkillsMergeRequestsCommentsGetOutputActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsCommentsGetOutputActorOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsCommentsGetOutputActorOrganizationActorTeams struct {
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

// SkillsMergeRequestsCommentsGetOutputActorOrganizationActor represents the skills merge requests comments get output actor organization actor type.
type SkillsMergeRequestsCommentsGetOutputActorOrganizationActor struct {
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
	ImageUrl string                                                            `json:"image_url"`
	Member   *SkillsMergeRequestsCommentsGetOutputActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsCommentsGetOutputActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsCommentsGetOutputActorConsumer represents the skills merge requests comments get output actor consumer type.
type SkillsMergeRequestsCommentsGetOutputActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsCommentsGetOutputActor represents the skills merge requests comments get output actor type.
type SkillsMergeRequestsCommentsGetOutputActor struct {
	Type              string                                                      `json:"type"`
	Name              string                                                      `json:"name"`
	ImageUrl          *string                                                     `json:"image_url,omitempty"`
	Email             *string                                                     `json:"email,omitempty"`
	OrganizationActor *SkillsMergeRequestsCommentsGetOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsMergeRequestsCommentsGetOutputActorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                             `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsCommentsGetOutput represents the skills merge requests comments get output type.
type SkillsMergeRequestsCommentsGetOutput struct {
	Object                  string                                    `json:"object"`
	Id                      string                                    `json:"id"`
	SkillMergeRequestItemId *string                                   `json:"skill_merge_request_item_id,omitempty"`
	Actor                   SkillsMergeRequestsCommentsGetOutputActor `json:"actor"`
	Body                    string                                    `json:"body"`
	Path                    *string                                   `json:"path,omitempty"`
	InReplyToCommentId      *string                                   `json:"in_reply_to_comment_id,omitempty"`
	DeletedAt               *time.Time                                `json:"deleted_at,omitempty"`
	CreatedAt               time.Time                                 `json:"created_at"`
	UpdatedAt               time.Time                                 `json:"updated_at"`
}

// MapSkillsMergeRequestsCommentsGetOutputFromJSON deserializes JSON data into a SkillsMergeRequestsCommentsGetOutput.
func MapSkillsMergeRequestsCommentsGetOutputFromJSON(data []byte) (*SkillsMergeRequestsCommentsGetOutput, error) {
	var v SkillsMergeRequestsCommentsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsCommentsGetOutputToJSON serializes a SkillsMergeRequestsCommentsGetOutput to JSON.
func MapSkillsMergeRequestsCommentsGetOutputToJSON(v *SkillsMergeRequestsCommentsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
