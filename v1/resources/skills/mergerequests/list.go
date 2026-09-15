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

// SkillsMergeRequestsListOutputItemsCreatedByOrganizationActor represents the skills merge requests list output items created by organization actor type.
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

// SkillsMergeRequestsListOutputItemsCreatedByConsumer represents the skills merge requests list output items created by consumer type.
type SkillsMergeRequestsListOutputItemsCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsListOutputItemsCreatedBy represents the skills merge requests list output items created by type.
type SkillsMergeRequestsListOutputItemsCreatedBy struct {
	Type              string                                                        `json:"type"`
	Name              string                                                        `json:"name"`
	ImageUrl          *string                                                       `json:"image_url,omitempty"`
	Email             *string                                                       `json:"email,omitempty"`
	OrganizationActor *SkillsMergeRequestsListOutputItemsCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsMergeRequestsListOutputItemsCreatedByConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                               `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsListOutputItems represents the skills merge requests list output items type.
type SkillsMergeRequestsListOutputItems struct {
	Object                        string                                       `json:"object"`
	Id                            string                                       `json:"id"`
	Status                        string                                       `json:"status"`
	Direction                     string                                       `json:"direction"`
	BaseStrategy                  string                                       `json:"base_strategy"`
	Title                         string                                       `json:"title"`
	Description                   *string                                      `json:"description,omitempty"`
	MergeError                    *string                                      `json:"merge_error,omitempty"`
	MergeErrorCode                *string                                      `json:"merge_error_code,omitempty"`
	SourceSkillId                 string                                       `json:"source_skill_id"`
	TargetSkillId                 string                                       `json:"target_skill_id"`
	BaseTargetSkillVersionId      string                                       `json:"base_target_skill_version_id"`
	RequestedSourceSkillVersionId string                                       `json:"requested_source_skill_version_id"`
	RequestedTargetSkillVersionId string                                       `json:"requested_target_skill_version_id"`
	PreMergeTargetSkillVersionId  *string                                      `json:"pre_merge_target_skill_version_id,omitempty"`
	MergedTargetSkillVersionId    *string                                      `json:"merged_target_skill_version_id,omitempty"`
	RollbackTargetSkillVersionId  *string                                      `json:"rollback_target_skill_version_id,omitempty"`
	CreatedBy                     *SkillsMergeRequestsListOutputItemsCreatedBy `json:"created_by,omitempty"`
	ItemCount                     float64                                      `json:"item_count"`
	CommentCount                  float64                                      `json:"comment_count"`
	MergeStartedAt                *time.Time                                   `json:"merge_started_at,omitempty"`
	MergedAt                      *time.Time                                   `json:"merged_at,omitempty"`
	ClosedAt                      *time.Time                                   `json:"closed_at,omitempty"`
	RolledBackAt                  *time.Time                                   `json:"rolled_back_at,omitempty"`
	CreatedAt                     time.Time                                    `json:"created_at"`
	UpdatedAt                     time.Time                                    `json:"updated_at"`
}

// SkillsMergeRequestsListOutputPagination represents the skills merge requests list output pagination type.
type SkillsMergeRequestsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsMergeRequestsListOutput represents the skills merge requests list output type.
type SkillsMergeRequestsListOutput struct {
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

// SkillsMergeRequestsListQueryCreatedAt - Filter skill merge request creation time by date range
type SkillsMergeRequestsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for skill merge request creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill merge request creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMergeRequestsListQuery represents the skills merge requests list query type.
type SkillsMergeRequestsListQuery struct {
	Limit            *float64 `json:"limit,omitempty"`
	After            *string  `json:"after,omitempty"`
	Before           *string  `json:"before,omitempty"`
	Cursor           *string  `json:"cursor,omitempty"`
	Order            *string  `json:"order,omitempty"`
	Id               *any     `json:"id,omitempty"`
	SourceSkillId    *any     `json:"source_skill_id,omitempty"`
	TargetSkillId    *any     `json:"target_skill_id,omitempty"`
	Status           *any     `json:"status,omitempty"`
	CreatedByActorId *any     `json:"created_by_actor_id,omitempty"`
	// CreatedAt - Filter skill merge request creation time by date range
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
