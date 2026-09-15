package comments

import (
	"encoding/json"
	"time"
)

// SkillsMergeRequestsCommentsDeleteOutputActorOrganizationActorMember represents the skills merge requests comments delete output actor organization actor member type.
type SkillsMergeRequestsCommentsDeleteOutputActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsCommentsDeleteOutputActorOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsCommentsDeleteOutputActorOrganizationActorTeams struct {
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

// SkillsMergeRequestsCommentsDeleteOutputActorOrganizationActor represents the skills merge requests comments delete output actor organization actor type.
type SkillsMergeRequestsCommentsDeleteOutputActorOrganizationActor struct {
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
	Member   *SkillsMergeRequestsCommentsDeleteOutputActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsCommentsDeleteOutputActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsCommentsDeleteOutputActorConsumer represents the skills merge requests comments delete output actor consumer type.
type SkillsMergeRequestsCommentsDeleteOutputActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsCommentsDeleteOutputActor represents the skills merge requests comments delete output actor type.
type SkillsMergeRequestsCommentsDeleteOutputActor struct {
	Type              string                                                         `json:"type"`
	Name              string                                                         `json:"name"`
	ImageUrl          *string                                                        `json:"image_url,omitempty"`
	Email             *string                                                        `json:"email,omitempty"`
	OrganizationActor *SkillsMergeRequestsCommentsDeleteOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsMergeRequestsCommentsDeleteOutputActorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsCommentsDeleteOutput represents the skills merge requests comments delete output type.
type SkillsMergeRequestsCommentsDeleteOutput struct {
	Object                  string                                       `json:"object"`
	Id                      string                                       `json:"id"`
	SkillMergeRequestItemId *string                                      `json:"skill_merge_request_item_id,omitempty"`
	Actor                   SkillsMergeRequestsCommentsDeleteOutputActor `json:"actor"`
	Body                    string                                       `json:"body"`
	Path                    *string                                      `json:"path,omitempty"`
	InReplyToCommentId      *string                                      `json:"in_reply_to_comment_id,omitempty"`
	DeletedAt               *time.Time                                   `json:"deleted_at,omitempty"`
	CreatedAt               time.Time                                    `json:"created_at"`
	UpdatedAt               time.Time                                    `json:"updated_at"`
}

// MapSkillsMergeRequestsCommentsDeleteOutputFromJSON deserializes JSON data into a SkillsMergeRequestsCommentsDeleteOutput.
func MapSkillsMergeRequestsCommentsDeleteOutputFromJSON(data []byte) (*SkillsMergeRequestsCommentsDeleteOutput, error) {
	var v SkillsMergeRequestsCommentsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsCommentsDeleteOutputToJSON serializes a SkillsMergeRequestsCommentsDeleteOutput to JSON.
func MapSkillsMergeRequestsCommentsDeleteOutputToJSON(v *SkillsMergeRequestsCommentsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
