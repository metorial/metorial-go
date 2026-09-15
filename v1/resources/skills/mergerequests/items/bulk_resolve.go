package items

import (
	"encoding/json"
	"time"
)

// SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByOrganizationActorMember represents the skills merge requests items bulk resolve output items resolved by organization actor member type.
type SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByOrganizationActorTeams struct {
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

// SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByOrganizationActor represents the skills merge requests items bulk resolve output items resolved by organization actor type.
type SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByOrganizationActor struct {
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
	ImageUrl string                                                                           `json:"image_url"`
	Member   *SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByConsumer represents the skills merge requests items bulk resolve output items resolved by consumer type.
type SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedBy represents the skills merge requests items bulk resolve output items resolved by type.
type SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedBy struct {
	Type              string                                                                     `json:"type"`
	Name              string                                                                     `json:"name"`
	ImageUrl          *string                                                                    `json:"image_url,omitempty"`
	Email             *string                                                                    `json:"email,omitempty"`
	OrganizationActor *SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                            `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsItemsBulkResolveOutputItems represents the skills merge requests items bulk resolve output items type.
type SkillsMergeRequestsItemsBulkResolveOutputItems struct {
	Object              string                                                    `json:"object"`
	Id                  string                                                    `json:"id"`
	SkillMergeRequestId string                                                    `json:"skill_merge_request_id"`
	Path                string                                                    `json:"path"`
	Kind                string                                                    `json:"kind"`
	ChangeType          string                                                    `json:"change_type"`
	Status              string                                                    `json:"status"`
	ResolutionType      *string                                                   `json:"resolution_type,omitempty"`
	ConflictReason      *string                                                   `json:"conflict_reason,omitempty"`
	Resolution          *map[string]any                                           `json:"resolution,omitempty"`
	ResolvedBy          *SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedBy `json:"resolved_by,omitempty"`
	ResolvedAt          *time.Time                                                `json:"resolved_at,omitempty"`
	AppliedAt           *time.Time                                                `json:"applied_at,omitempty"`
	CreatedAt           time.Time                                                 `json:"created_at"`
	UpdatedAt           time.Time                                                 `json:"updated_at"`
}

// SkillsMergeRequestsItemsBulkResolveOutputPagination represents the skills merge requests items bulk resolve output pagination type.
type SkillsMergeRequestsItemsBulkResolveOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsMergeRequestsItemsBulkResolveOutput represents the skills merge requests items bulk resolve output type.
type SkillsMergeRequestsItemsBulkResolveOutput struct {
	Items      []SkillsMergeRequestsItemsBulkResolveOutputItems    `json:"items"`
	Pagination SkillsMergeRequestsItemsBulkResolveOutputPagination `json:"pagination"`
}

// MapSkillsMergeRequestsItemsBulkResolveOutputFromJSON deserializes JSON data into a SkillsMergeRequestsItemsBulkResolveOutput.
func MapSkillsMergeRequestsItemsBulkResolveOutputFromJSON(data []byte) (*SkillsMergeRequestsItemsBulkResolveOutput, error) {
	var v SkillsMergeRequestsItemsBulkResolveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsItemsBulkResolveOutputToJSON serializes a SkillsMergeRequestsItemsBulkResolveOutput to JSON.
func MapSkillsMergeRequestsItemsBulkResolveOutputToJSON(v *SkillsMergeRequestsItemsBulkResolveOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMergeRequestsItemsBulkResolveBodyItemsResolution represents the skills merge requests items bulk resolve body items resolution type.
type SkillsMergeRequestsItemsBulkResolveBodyItemsResolution struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
	FileId  *string `json:"fileId,omitempty"`
}

// SkillsMergeRequestsItemsBulkResolveBodyItems represents the skills merge requests items bulk resolve body items type.
type SkillsMergeRequestsItemsBulkResolveBodyItems struct {
	ItemId         string                                                  `json:"item_id"`
	ResolutionType string                                                  `json:"resolution_type"`
	Resolution     *SkillsMergeRequestsItemsBulkResolveBodyItemsResolution `json:"resolution,omitempty"`
}

// SkillsMergeRequestsItemsBulkResolveBody represents the skills merge requests items bulk resolve body type.
type SkillsMergeRequestsItemsBulkResolveBody struct {
	Items []SkillsMergeRequestsItemsBulkResolveBodyItems `json:"items"`
}

// MapSkillsMergeRequestsItemsBulkResolveBodyFromJSON deserializes JSON data into a SkillsMergeRequestsItemsBulkResolveBody.
func MapSkillsMergeRequestsItemsBulkResolveBodyFromJSON(data []byte) (*SkillsMergeRequestsItemsBulkResolveBody, error) {
	var v SkillsMergeRequestsItemsBulkResolveBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsItemsBulkResolveBodyToJSON serializes a SkillsMergeRequestsItemsBulkResolveBody to JSON.
func MapSkillsMergeRequestsItemsBulkResolveBodyToJSON(v *SkillsMergeRequestsItemsBulkResolveBody) ([]byte, error) {
	return json.Marshal(v)
}
