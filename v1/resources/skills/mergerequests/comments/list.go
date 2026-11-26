package comments

import (
	"encoding/json"
	"time"
)

// SkillsMergeRequestsCommentsListOutputItemsActorOrganizationActorMember represents the skills merge requests comments list output items actor organization actor member type.
type SkillsMergeRequestsCommentsListOutputItemsActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsCommentsListOutputItemsActorOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsCommentsListOutputItemsActorOrganizationActorTeams struct {
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

// SkillsMergeRequestsCommentsListOutputItemsActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsMergeRequestsCommentsListOutputItemsActorOrganizationActor struct {
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
	ImageUrl string                                                                  `json:"image_url"`
	Member   *SkillsMergeRequestsCommentsListOutputItemsActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsCommentsListOutputItemsActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsCommentsListOutputItemsActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsMergeRequestsCommentsListOutputItemsActorConsumer struct {
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

// SkillsMergeRequestsCommentsListOutputItemsActor - Actor behind a participant.
type SkillsMergeRequestsCommentsListOutputItemsActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsMergeRequestsCommentsListOutputItemsActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsMergeRequestsCommentsListOutputItemsActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsCommentsListOutputItems represents the skills merge requests comments list output items type.
type SkillsMergeRequestsCommentsListOutputItems struct {
	Object                  string  `json:"object"`
	Id                      string  `json:"id"`
	SkillMergeRequestItemId *string `json:"skill_merge_request_item_id,omitempty"`
	// Actor - Actor behind a participant.
	Actor              SkillsMergeRequestsCommentsListOutputItemsActor `json:"actor"`
	Body               string                                          `json:"body"`
	Path               *string                                         `json:"path,omitempty"`
	InReplyToCommentId *string                                         `json:"in_reply_to_comment_id,omitempty"`
	DeletedAt          *time.Time                                      `json:"deleted_at,omitempty"`
	CreatedAt          time.Time                                       `json:"created_at"`
	UpdatedAt          time.Time                                       `json:"updated_at"`
}

// SkillsMergeRequestsCommentsListOutputPagination represents the skills merge requests comments list output pagination type.
type SkillsMergeRequestsCommentsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsMergeRequestsCommentsListOutput represents the skills merge requests comments list output type.
type SkillsMergeRequestsCommentsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                          `json:"object"`
	Items      []SkillsMergeRequestsCommentsListOutputItems    `json:"items"`
	Pagination SkillsMergeRequestsCommentsListOutputPagination `json:"pagination"`
}

// MapSkillsMergeRequestsCommentsListOutputFromJSON deserializes JSON data into a SkillsMergeRequestsCommentsListOutput.
func MapSkillsMergeRequestsCommentsListOutputFromJSON(data []byte) (*SkillsMergeRequestsCommentsListOutput, error) {
	var v SkillsMergeRequestsCommentsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsCommentsListOutputToJSON serializes a SkillsMergeRequestsCommentsListOutput to JSON.
func MapSkillsMergeRequestsCommentsListOutputToJSON(v *SkillsMergeRequestsCommentsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMergeRequestsCommentsListQuery represents the skills merge requests comments list query type.
type SkillsMergeRequestsCommentsListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// ItemId - Only return comments on this proposed change.
	ItemId *string `json:"item_id,omitempty"`
}

// MapSkillsMergeRequestsCommentsListQueryFromJSON deserializes JSON data into a SkillsMergeRequestsCommentsListQuery.
func MapSkillsMergeRequestsCommentsListQueryFromJSON(data []byte) (*SkillsMergeRequestsCommentsListQuery, error) {
	var v SkillsMergeRequestsCommentsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsCommentsListQueryToJSON serializes a SkillsMergeRequestsCommentsListQuery to JSON.
func MapSkillsMergeRequestsCommentsListQueryToJSON(v *SkillsMergeRequestsCommentsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
