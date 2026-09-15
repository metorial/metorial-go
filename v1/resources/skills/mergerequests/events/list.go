package events

import (
	"encoding/json"
	"time"
)

// SkillsMergeRequestsEventsListOutputItemsActorOrganizationActorMember represents the skills merge requests events list output items actor organization actor member type.
type SkillsMergeRequestsEventsListOutputItemsActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsEventsListOutputItemsActorOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsEventsListOutputItemsActorOrganizationActorTeams struct {
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

// SkillsMergeRequestsEventsListOutputItemsActorOrganizationActor represents the skills merge requests events list output items actor organization actor type.
type SkillsMergeRequestsEventsListOutputItemsActorOrganizationActor struct {
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
	Member   *SkillsMergeRequestsEventsListOutputItemsActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsEventsListOutputItemsActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsEventsListOutputItemsActorConsumer represents the skills merge requests events list output items actor consumer type.
type SkillsMergeRequestsEventsListOutputItemsActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsEventsListOutputItemsActor represents the skills merge requests events list output items actor type.
type SkillsMergeRequestsEventsListOutputItemsActor struct {
	Type              string                                                          `json:"type"`
	Name              string                                                          `json:"name"`
	ImageUrl          *string                                                         `json:"image_url,omitempty"`
	Email             *string                                                         `json:"email,omitempty"`
	OrganizationActor *SkillsMergeRequestsEventsListOutputItemsActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsMergeRequestsEventsListOutputItemsActorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                 `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsEventsListOutputItemsCommentActorOrganizationActorMember represents the skills merge requests events list output items comment actor organization actor member type.
type SkillsMergeRequestsEventsListOutputItemsCommentActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsEventsListOutputItemsCommentActorOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsEventsListOutputItemsCommentActorOrganizationActorTeams struct {
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

// SkillsMergeRequestsEventsListOutputItemsCommentActorOrganizationActor represents the skills merge requests events list output items comment actor organization actor type.
type SkillsMergeRequestsEventsListOutputItemsCommentActorOrganizationActor struct {
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
	ImageUrl string                                                                       `json:"image_url"`
	Member   *SkillsMergeRequestsEventsListOutputItemsCommentActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsEventsListOutputItemsCommentActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsEventsListOutputItemsCommentActorConsumer represents the skills merge requests events list output items comment actor consumer type.
type SkillsMergeRequestsEventsListOutputItemsCommentActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsEventsListOutputItemsCommentActor represents the skills merge requests events list output items comment actor type.
type SkillsMergeRequestsEventsListOutputItemsCommentActor struct {
	Type              string                                                                 `json:"type"`
	Name              string                                                                 `json:"name"`
	ImageUrl          *string                                                                `json:"image_url,omitempty"`
	Email             *string                                                                `json:"email,omitempty"`
	OrganizationActor *SkillsMergeRequestsEventsListOutputItemsCommentActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsMergeRequestsEventsListOutputItemsCommentActorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                        `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsEventsListOutputItemsComment represents the skills merge requests events list output items comment type.
type SkillsMergeRequestsEventsListOutputItemsComment struct {
	Object                  string                                               `json:"object"`
	Id                      string                                               `json:"id"`
	SkillMergeRequestItemId *string                                              `json:"skill_merge_request_item_id,omitempty"`
	Actor                   SkillsMergeRequestsEventsListOutputItemsCommentActor `json:"actor"`
	Body                    string                                               `json:"body"`
	Path                    *string                                              `json:"path,omitempty"`
	InReplyToCommentId      *string                                              `json:"in_reply_to_comment_id,omitempty"`
	DeletedAt               *time.Time                                           `json:"deleted_at,omitempty"`
	CreatedAt               time.Time                                            `json:"created_at"`
	UpdatedAt               time.Time                                            `json:"updated_at"`
}

// SkillsMergeRequestsEventsListOutputItems represents the skills merge requests events list output items type.
type SkillsMergeRequestsEventsListOutputItems struct {
	Object       string                                           `json:"object"`
	Id           string                                           `json:"id"`
	Type         string                                           `json:"type"`
	Actor        *SkillsMergeRequestsEventsListOutputItemsActor   `json:"actor,omitempty"`
	Comment      *SkillsMergeRequestsEventsListOutputItemsComment `json:"comment,omitempty"`
	ErrorCode    *string                                          `json:"error_code,omitempty"`
	ErrorMessage *string                                          `json:"error_message,omitempty"`
	CreatedAt    time.Time                                        `json:"created_at"`
}

// SkillsMergeRequestsEventsListOutputPagination represents the skills merge requests events list output pagination type.
type SkillsMergeRequestsEventsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsMergeRequestsEventsListOutput represents the skills merge requests events list output type.
type SkillsMergeRequestsEventsListOutput struct {
	Items      []SkillsMergeRequestsEventsListOutputItems    `json:"items"`
	Pagination SkillsMergeRequestsEventsListOutputPagination `json:"pagination"`
}

// MapSkillsMergeRequestsEventsListOutputFromJSON deserializes JSON data into a SkillsMergeRequestsEventsListOutput.
func MapSkillsMergeRequestsEventsListOutputFromJSON(data []byte) (*SkillsMergeRequestsEventsListOutput, error) {
	var v SkillsMergeRequestsEventsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsEventsListOutputToJSON serializes a SkillsMergeRequestsEventsListOutput to JSON.
func MapSkillsMergeRequestsEventsListOutputToJSON(v *SkillsMergeRequestsEventsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMergeRequestsEventsListQueryCreatedAt - Filter skill merge request event creation time by date range
type SkillsMergeRequestsEventsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for skill merge request event creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill merge request event creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMergeRequestsEventsListQuery represents the skills merge requests events list query type.
type SkillsMergeRequestsEventsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Type   *any     `json:"type,omitempty"`
	// CreatedAt - Filter skill merge request event creation time by date range
	CreatedAt *SkillsMergeRequestsEventsListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapSkillsMergeRequestsEventsListQueryFromJSON deserializes JSON data into a SkillsMergeRequestsEventsListQuery.
func MapSkillsMergeRequestsEventsListQueryFromJSON(data []byte) (*SkillsMergeRequestsEventsListQuery, error) {
	var v SkillsMergeRequestsEventsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsEventsListQueryToJSON serializes a SkillsMergeRequestsEventsListQuery to JSON.
func MapSkillsMergeRequestsEventsListQueryToJSON(v *SkillsMergeRequestsEventsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
