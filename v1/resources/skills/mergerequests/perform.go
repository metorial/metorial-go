package mergerequests

import (
	"encoding/json"
	"time"
)

// SkillsMergeRequestsPerformOutputCreatedByOrganizationActorMember represents the skills merge requests perform output created by organization actor member type.
type SkillsMergeRequestsPerformOutputCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsPerformOutputCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsPerformOutputCreatedByOrganizationActorTeams struct {
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

// SkillsMergeRequestsPerformOutputCreatedByOrganizationActor represents the skills merge requests perform output created by organization actor type.
type SkillsMergeRequestsPerformOutputCreatedByOrganizationActor struct {
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
	Member   *SkillsMergeRequestsPerformOutputCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsPerformOutputCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsPerformOutputCreatedByConsumer represents the skills merge requests perform output created by consumer type.
type SkillsMergeRequestsPerformOutputCreatedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsPerformOutputCreatedBy represents the skills merge requests perform output created by type.
type SkillsMergeRequestsPerformOutputCreatedBy struct {
	Type              string                                                      `json:"type"`
	Name              string                                                      `json:"name"`
	ImageUrl          *string                                                     `json:"image_url,omitempty"`
	Email             *string                                                     `json:"email,omitempty"`
	OrganizationActor *SkillsMergeRequestsPerformOutputCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsMergeRequestsPerformOutputCreatedByConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                             `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsPerformOutput represents the skills merge requests perform output type.
type SkillsMergeRequestsPerformOutput struct {
	Object                        string                                     `json:"object"`
	Id                            string                                     `json:"id"`
	Status                        string                                     `json:"status"`
	Direction                     string                                     `json:"direction"`
	BaseStrategy                  string                                     `json:"base_strategy"`
	Title                         string                                     `json:"title"`
	Description                   *string                                    `json:"description,omitempty"`
	MergeError                    *string                                    `json:"merge_error,omitempty"`
	MergeErrorCode                *string                                    `json:"merge_error_code,omitempty"`
	SourceSkillId                 string                                     `json:"source_skill_id"`
	TargetSkillId                 string                                     `json:"target_skill_id"`
	BaseTargetSkillVersionId      string                                     `json:"base_target_skill_version_id"`
	RequestedSourceSkillVersionId string                                     `json:"requested_source_skill_version_id"`
	RequestedTargetSkillVersionId string                                     `json:"requested_target_skill_version_id"`
	PreMergeTargetSkillVersionId  *string                                    `json:"pre_merge_target_skill_version_id,omitempty"`
	MergedTargetSkillVersionId    *string                                    `json:"merged_target_skill_version_id,omitempty"`
	RollbackTargetSkillVersionId  *string                                    `json:"rollback_target_skill_version_id,omitempty"`
	CreatedBy                     *SkillsMergeRequestsPerformOutputCreatedBy `json:"created_by,omitempty"`
	ItemCount                     float64                                    `json:"item_count"`
	CommentCount                  float64                                    `json:"comment_count"`
	MergeStartedAt                *time.Time                                 `json:"merge_started_at,omitempty"`
	MergedAt                      *time.Time                                 `json:"merged_at,omitempty"`
	ClosedAt                      *time.Time                                 `json:"closed_at,omitempty"`
	RolledBackAt                  *time.Time                                 `json:"rolled_back_at,omitempty"`
	CreatedAt                     time.Time                                  `json:"created_at"`
	UpdatedAt                     time.Time                                  `json:"updated_at"`
}

// MapSkillsMergeRequestsPerformOutputFromJSON deserializes JSON data into a SkillsMergeRequestsPerformOutput.
func MapSkillsMergeRequestsPerformOutputFromJSON(data []byte) (*SkillsMergeRequestsPerformOutput, error) {
	var v SkillsMergeRequestsPerformOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsPerformOutputToJSON serializes a SkillsMergeRequestsPerformOutput to JSON.
func MapSkillsMergeRequestsPerformOutputToJSON(v *SkillsMergeRequestsPerformOutput) ([]byte, error) {
	return json.Marshal(v)
}
