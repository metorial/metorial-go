package mergerequests

import (
	"encoding/json"
	"time"
)

// SkillsMergeRequestsListOutputItemsCreatedByOrganizationActorMember represents the skills merge requests list output items created by organization actor member type.
type SkillsMergeRequestsListOutputItemsCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsListOutputItemsCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsListOutputItemsCreatedByOrganizationActorTeams struct {
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

// SkillsMergeRequestsListOutputItemsCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsMergeRequestsListOutputItemsCreatedByOrganizationActor struct {
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
	ImageUrl string                                                              `json:"image_url"`
	Member   *SkillsMergeRequestsListOutputItemsCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsListOutputItemsCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsListOutputItemsCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsMergeRequestsListOutputItemsCreatedByConsumer struct {
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

// SkillsMergeRequestsListOutputItemsCreatedBy - Actor behind a participant.
type SkillsMergeRequestsListOutputItemsCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsMergeRequestsListOutputItemsCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsMergeRequestsListOutputItemsCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsListOutputItems represents the skills merge requests list output items type.
type SkillsMergeRequestsListOutputItems struct {
	Object                        string  `json:"object"`
	Id                            string  `json:"id"`
	Status                        string  `json:"status"`
	Direction                     string  `json:"direction"`
	BaseStrategy                  string  `json:"base_strategy"`
	Title                         string  `json:"title"`
	Description                   *string `json:"description,omitempty"`
	MergeError                    *string `json:"merge_error,omitempty"`
	MergeErrorCode                *string `json:"merge_error_code,omitempty"`
	SourceSkillId                 string  `json:"source_skill_id"`
	TargetSkillId                 string  `json:"target_skill_id"`
	BaseTargetSkillVersionId      string  `json:"base_target_skill_version_id"`
	RequestedSourceSkillVersionId string  `json:"requested_source_skill_version_id"`
	RequestedTargetSkillVersionId string  `json:"requested_target_skill_version_id"`
	PreMergeTargetSkillVersionId  *string `json:"pre_merge_target_skill_version_id,omitempty"`
	MergedTargetSkillVersionId    *string `json:"merged_target_skill_version_id,omitempty"`
	RollbackTargetSkillVersionId  *string `json:"rollback_target_skill_version_id,omitempty"`
	// CreatedBy - Actor behind a participant.
	CreatedBy      *SkillsMergeRequestsListOutputItemsCreatedBy `json:"created_by,omitempty"`
	ItemCount      float64                                      `json:"item_count"`
	CommentCount   float64                                      `json:"comment_count"`
	MergeStartedAt *time.Time                                   `json:"merge_started_at,omitempty"`
	MergedAt       *time.Time                                   `json:"merged_at,omitempty"`
	ClosedAt       *time.Time                                   `json:"closed_at,omitempty"`
	RolledBackAt   *time.Time                                   `json:"rolled_back_at,omitempty"`
	CreatedAt      time.Time                                    `json:"created_at"`
	UpdatedAt      time.Time                                    `json:"updated_at"`
}

// SkillsMergeRequestsListOutputPagination represents the skills merge requests list output pagination type.
type SkillsMergeRequestsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsMergeRequestsListOutput represents the skills merge requests list output type.
type SkillsMergeRequestsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                  `json:"object"`
	Items      []SkillsMergeRequestsListOutputItems    `json:"items"`
	Pagination SkillsMergeRequestsListOutputPagination `json:"pagination"`
}

// MapSkillsMergeRequestsListOutputFromJSON deserializes JSON data into a SkillsMergeRequestsListOutput.
func MapSkillsMergeRequestsListOutputFromJSON(data []byte) (*SkillsMergeRequestsListOutput, error) {
	var v SkillsMergeRequestsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsListOutputToJSON serializes a SkillsMergeRequestsListOutput to JSON.
func MapSkillsMergeRequestsListOutputToJSON(v *SkillsMergeRequestsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMergeRequestsListQueryCreatedAt - Only return results whose skill merge request creation time is in this range. Set `gt`, `lt`, or both.
type SkillsMergeRequestsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMergeRequestsListQuery represents the skills merge requests list query type.
type SkillsMergeRequestsListQuery struct {
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
	// Id - Only return merge requests with these IDs.
	Id *any `json:"id,omitempty"`
	// SourceSkillId - Only return merge requests from one of these source skills.
	SourceSkillId *any `json:"source_skill_id,omitempty"`
	// TargetSkillId - Only return merge requests into one of these target skills.
	TargetSkillId *any `json:"target_skill_id,omitempty"`
	// Status - Only return merge requests with one of these statuses.
	Status *any `json:"status,omitempty"`
	// CreatedByActorId - Only return merge requests created by one of these actors.
	CreatedByActorId *any `json:"created_by_actor_id,omitempty"`
	// CreatedAt - Only return results whose skill merge request creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SkillsMergeRequestsListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapSkillsMergeRequestsListQueryFromJSON deserializes JSON data into a SkillsMergeRequestsListQuery.
func MapSkillsMergeRequestsListQueryFromJSON(data []byte) (*SkillsMergeRequestsListQuery, error) {
	var v SkillsMergeRequestsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsListQueryToJSON serializes a SkillsMergeRequestsListQuery to JSON.
func MapSkillsMergeRequestsListQueryToJSON(v *SkillsMergeRequestsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
