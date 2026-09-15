package items

import (
	"encoding/json"
	"time"
)

// SkillsMergeRequestsItemsResolveOutputResolvedByOrganizationActorMember represents the skills merge requests items resolve output resolved by organization actor member type.
type SkillsMergeRequestsItemsResolveOutputResolvedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsItemsResolveOutputResolvedByOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsItemsResolveOutputResolvedByOrganizationActorTeams struct {
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

// SkillsMergeRequestsItemsResolveOutputResolvedByOrganizationActor represents the skills merge requests items resolve output resolved by organization actor type.
type SkillsMergeRequestsItemsResolveOutputResolvedByOrganizationActor struct {
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
	Member   *SkillsMergeRequestsItemsResolveOutputResolvedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsItemsResolveOutputResolvedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsItemsResolveOutputResolvedByConsumer represents the skills merge requests items resolve output resolved by consumer type.
type SkillsMergeRequestsItemsResolveOutputResolvedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsItemsResolveOutputResolvedBy represents the skills merge requests items resolve output resolved by type.
type SkillsMergeRequestsItemsResolveOutputResolvedBy struct {
	Type              string                                                            `json:"type"`
	Name              string                                                            `json:"name"`
	ImageUrl          *string                                                           `json:"image_url,omitempty"`
	Email             *string                                                           `json:"email,omitempty"`
	OrganizationActor *SkillsMergeRequestsItemsResolveOutputResolvedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsMergeRequestsItemsResolveOutputResolvedByConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                   `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsItemsResolveOutput represents the skills merge requests items resolve output type.
type SkillsMergeRequestsItemsResolveOutput struct {
	Object              string                                           `json:"object"`
	Id                  string                                           `json:"id"`
	SkillMergeRequestId string                                           `json:"skill_merge_request_id"`
	Path                string                                           `json:"path"`
	Kind                string                                           `json:"kind"`
	ChangeType          string                                           `json:"change_type"`
	Status              string                                           `json:"status"`
	ResolutionType      *string                                          `json:"resolution_type,omitempty"`
	ConflictReason      *string                                          `json:"conflict_reason,omitempty"`
	Resolution          *map[string]any                                  `json:"resolution,omitempty"`
	ResolvedBy          *SkillsMergeRequestsItemsResolveOutputResolvedBy `json:"resolved_by,omitempty"`
	ResolvedAt          *time.Time                                       `json:"resolved_at,omitempty"`
	AppliedAt           *time.Time                                       `json:"applied_at,omitempty"`
	CreatedAt           time.Time                                        `json:"created_at"`
	UpdatedAt           time.Time                                        `json:"updated_at"`
}

// MapSkillsMergeRequestsItemsResolveOutputFromJSON deserializes JSON data into a SkillsMergeRequestsItemsResolveOutput.
func MapSkillsMergeRequestsItemsResolveOutputFromJSON(data []byte) (*SkillsMergeRequestsItemsResolveOutput, error) {
	var v SkillsMergeRequestsItemsResolveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsItemsResolveOutputToJSON serializes a SkillsMergeRequestsItemsResolveOutput to JSON.
func MapSkillsMergeRequestsItemsResolveOutputToJSON(v *SkillsMergeRequestsItemsResolveOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMergeRequestsItemsResolveBodyResolution represents the skills merge requests items resolve body resolution type.
type SkillsMergeRequestsItemsResolveBodyResolution struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
	FileId  *string `json:"fileId,omitempty"`
}

// SkillsMergeRequestsItemsResolveBody represents the skills merge requests items resolve body type.
type SkillsMergeRequestsItemsResolveBody struct {
	ResolutionType string                                         `json:"resolution_type"`
	Resolution     *SkillsMergeRequestsItemsResolveBodyResolution `json:"resolution,omitempty"`
}

// MapSkillsMergeRequestsItemsResolveBodyFromJSON deserializes JSON data into a SkillsMergeRequestsItemsResolveBody.
func MapSkillsMergeRequestsItemsResolveBodyFromJSON(data []byte) (*SkillsMergeRequestsItemsResolveBody, error) {
	var v SkillsMergeRequestsItemsResolveBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsItemsResolveBodyToJSON serializes a SkillsMergeRequestsItemsResolveBody to JSON.
func MapSkillsMergeRequestsItemsResolveBodyToJSON(v *SkillsMergeRequestsItemsResolveBody) ([]byte, error) {
	return json.Marshal(v)
}
