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

// SkillsMergeRequestsEventsListOutputItemsActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
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

// SkillsMergeRequestsEventsListOutputItemsActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsMergeRequestsEventsListOutputItemsActorConsumer struct {
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

// SkillsMergeRequestsEventsListOutputItemsActor - Actor behind a participant.
type SkillsMergeRequestsEventsListOutputItemsActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsMergeRequestsEventsListOutputItemsActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsMergeRequestsEventsListOutputItemsActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
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

// SkillsMergeRequestsEventsListOutputItemsCommentActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
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

// SkillsMergeRequestsEventsListOutputItemsCommentActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsMergeRequestsEventsListOutputItemsCommentActorConsumer struct {
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

// SkillsMergeRequestsEventsListOutputItemsCommentActor - Actor behind a participant.
type SkillsMergeRequestsEventsListOutputItemsCommentActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsMergeRequestsEventsListOutputItemsCommentActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsMergeRequestsEventsListOutputItemsCommentActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsEventsListOutputItemsComment represents the skills merge requests events list output items comment type.
type SkillsMergeRequestsEventsListOutputItemsComment struct {
	Object                  string  `json:"object"`
	Id                      string  `json:"id"`
	SkillMergeRequestItemId *string `json:"skill_merge_request_item_id,omitempty"`
	// Actor - Actor behind a participant.
	Actor              SkillsMergeRequestsEventsListOutputItemsCommentActor `json:"actor"`
	Body               string                                               `json:"body"`
	Path               *string                                              `json:"path,omitempty"`
	InReplyToCommentId *string                                              `json:"in_reply_to_comment_id,omitempty"`
	DeletedAt          *time.Time                                           `json:"deleted_at,omitempty"`
	CreatedAt          time.Time                                            `json:"created_at"`
	UpdatedAt          time.Time                                            `json:"updated_at"`
}

// SkillsMergeRequestsEventsListOutputItems represents the skills merge requests events list output items type.
type SkillsMergeRequestsEventsListOutputItems struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Type   string `json:"type"`
	// Actor - Actor behind a participant.
	Actor        *SkillsMergeRequestsEventsListOutputItemsActor   `json:"actor,omitempty"`
	Comment      *SkillsMergeRequestsEventsListOutputItemsComment `json:"comment,omitempty"`
	ErrorCode    *string                                          `json:"error_code,omitempty"`
	ErrorMessage *string                                          `json:"error_message,omitempty"`
	CreatedAt    time.Time                                        `json:"created_at"`
}

// SkillsMergeRequestsEventsListOutputPagination represents the skills merge requests events list output pagination type.
type SkillsMergeRequestsEventsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsMergeRequestsEventsListOutput represents the skills merge requests events list output type.
type SkillsMergeRequestsEventsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                        `json:"object"`
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

// SkillsMergeRequestsEventsListQueryCreatedAt - Only return results whose skill merge request event creation time is in this range. Set `gt`, `lt`, or both.
type SkillsMergeRequestsEventsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMergeRequestsEventsListQuery represents the skills merge requests events list query type.
type SkillsMergeRequestsEventsListQuery struct {
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
	// Type - Only return events of one of these types.
	Type *any `json:"type,omitempty"`
	// CreatedAt - Only return results whose skill merge request event creation time is in this range. Set `gt`, `lt`, or both.
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
