package events

import (
	"encoding/json"
	"time"
)

// SkillsMergeRequestsEventsGetOutputActorOrganizationActorMember represents the skills merge requests events get output actor organization actor member type.
type SkillsMergeRequestsEventsGetOutputActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsEventsGetOutputActorOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsEventsGetOutputActorOrganizationActorTeams struct {
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

// SkillsMergeRequestsEventsGetOutputActorOrganizationActor represents the skills merge requests events get output actor organization actor type.
type SkillsMergeRequestsEventsGetOutputActorOrganizationActor struct {
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
	ImageUrl string                                                          `json:"image_url"`
	Member   *SkillsMergeRequestsEventsGetOutputActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsEventsGetOutputActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsEventsGetOutputActorConsumer represents the skills merge requests events get output actor consumer type.
type SkillsMergeRequestsEventsGetOutputActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsEventsGetOutputActor represents the skills merge requests events get output actor type.
type SkillsMergeRequestsEventsGetOutputActor struct {
	Type              string                                                    `json:"type"`
	Name              string                                                    `json:"name"`
	ImageUrl          *string                                                   `json:"image_url,omitempty"`
	Email             *string                                                   `json:"email,omitempty"`
	OrganizationActor *SkillsMergeRequestsEventsGetOutputActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsMergeRequestsEventsGetOutputActorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                           `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsEventsGetOutputCommentActorOrganizationActorMember represents the skills merge requests events get output comment actor organization actor member type.
type SkillsMergeRequestsEventsGetOutputCommentActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsEventsGetOutputCommentActorOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsEventsGetOutputCommentActorOrganizationActorTeams struct {
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

// SkillsMergeRequestsEventsGetOutputCommentActorOrganizationActor represents the skills merge requests events get output comment actor organization actor type.
type SkillsMergeRequestsEventsGetOutputCommentActorOrganizationActor struct {
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
	ImageUrl string                                                                 `json:"image_url"`
	Member   *SkillsMergeRequestsEventsGetOutputCommentActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsEventsGetOutputCommentActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsEventsGetOutputCommentActorConsumer represents the skills merge requests events get output comment actor consumer type.
type SkillsMergeRequestsEventsGetOutputCommentActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsEventsGetOutputCommentActor represents the skills merge requests events get output comment actor type.
type SkillsMergeRequestsEventsGetOutputCommentActor struct {
	Type              string                                                           `json:"type"`
	Name              string                                                           `json:"name"`
	ImageUrl          *string                                                          `json:"image_url,omitempty"`
	Email             *string                                                          `json:"email,omitempty"`
	OrganizationActor *SkillsMergeRequestsEventsGetOutputCommentActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsMergeRequestsEventsGetOutputCommentActorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                  `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsEventsGetOutputComment represents the skills merge requests events get output comment type.
type SkillsMergeRequestsEventsGetOutputComment struct {
	Object                  string                                         `json:"object"`
	Id                      string                                         `json:"id"`
	SkillMergeRequestItemId *string                                        `json:"skill_merge_request_item_id,omitempty"`
	Actor                   SkillsMergeRequestsEventsGetOutputCommentActor `json:"actor"`
	Body                    string                                         `json:"body"`
	Path                    *string                                        `json:"path,omitempty"`
	InReplyToCommentId      *string                                        `json:"in_reply_to_comment_id,omitempty"`
	DeletedAt               *time.Time                                     `json:"deleted_at,omitempty"`
	CreatedAt               time.Time                                      `json:"created_at"`
	UpdatedAt               time.Time                                      `json:"updated_at"`
}

// SkillsMergeRequestsEventsGetOutput represents the skills merge requests events get output type.
type SkillsMergeRequestsEventsGetOutput struct {
	Object       string                                     `json:"object"`
	Id           string                                     `json:"id"`
	Type         string                                     `json:"type"`
	Actor        *SkillsMergeRequestsEventsGetOutputActor   `json:"actor,omitempty"`
	Comment      *SkillsMergeRequestsEventsGetOutputComment `json:"comment,omitempty"`
	ErrorCode    *string                                    `json:"error_code,omitempty"`
	ErrorMessage *string                                    `json:"error_message,omitempty"`
	CreatedAt    time.Time                                  `json:"created_at"`
}

// MapSkillsMergeRequestsEventsGetOutputFromJSON deserializes JSON data into a SkillsMergeRequestsEventsGetOutput.
func MapSkillsMergeRequestsEventsGetOutputFromJSON(data []byte) (*SkillsMergeRequestsEventsGetOutput, error) {
	var v SkillsMergeRequestsEventsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsEventsGetOutputToJSON serializes a SkillsMergeRequestsEventsGetOutput to JSON.
func MapSkillsMergeRequestsEventsGetOutputToJSON(v *SkillsMergeRequestsEventsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
