package mergerequests

import (
	"encoding/json"
	"time"
)

// SkillsMergeRequestsCloseOutputCreatedByOrganizationActorMember represents the skills merge requests close output created by organization actor member type.
type SkillsMergeRequestsCloseOutputCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsCloseOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsCloseOutputCreatedByOrganizationActorTeams struct {
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

// SkillsMergeRequestsCloseOutputCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsMergeRequestsCloseOutputCreatedByOrganizationActor struct {
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
	Member   *SkillsMergeRequestsCloseOutputCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsCloseOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsCloseOutputCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsMergeRequestsCloseOutputCreatedByConsumer struct {
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

// SkillsMergeRequestsCloseOutputCreatedBy - Actor behind a participant.
type SkillsMergeRequestsCloseOutputCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsMergeRequestsCloseOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsMergeRequestsCloseOutputCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsCloseOutput represents the skills merge requests close output type.
type SkillsMergeRequestsCloseOutput struct {
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
	CreatedBy      *SkillsMergeRequestsCloseOutputCreatedBy `json:"created_by,omitempty"`
	ItemCount      float64                                  `json:"item_count"`
	CommentCount   float64                                  `json:"comment_count"`
	MergeStartedAt *time.Time                               `json:"merge_started_at,omitempty"`
	MergedAt       *time.Time                               `json:"merged_at,omitempty"`
	ClosedAt       *time.Time                               `json:"closed_at,omitempty"`
	RolledBackAt   *time.Time                               `json:"rolled_back_at,omitempty"`
	CreatedAt      time.Time                                `json:"created_at"`
	UpdatedAt      time.Time                                `json:"updated_at"`
}

// MapSkillsMergeRequestsCloseOutputFromJSON deserializes JSON data into a SkillsMergeRequestsCloseOutput.
func MapSkillsMergeRequestsCloseOutputFromJSON(data []byte) (*SkillsMergeRequestsCloseOutput, error) {
	var v SkillsMergeRequestsCloseOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsCloseOutputToJSON serializes a SkillsMergeRequestsCloseOutput to JSON.
func MapSkillsMergeRequestsCloseOutputToJSON(v *SkillsMergeRequestsCloseOutput) ([]byte, error) {
	return json.Marshal(v)
}
