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

// SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
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

// SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByConsumer struct {
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

// SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedBy - Actor behind a participant.
type SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsItemsBulkResolveOutputItems represents the skills merge requests items bulk resolve output items type.
type SkillsMergeRequestsItemsBulkResolveOutputItems struct {
	Object              string          `json:"object"`
	Id                  string          `json:"id"`
	SkillMergeRequestId string          `json:"skill_merge_request_id"`
	Path                string          `json:"path"`
	Kind                string          `json:"kind"`
	ChangeType          string          `json:"change_type"`
	Status              string          `json:"status"`
	ResolutionType      *string         `json:"resolution_type,omitempty"`
	ConflictReason      *string         `json:"conflict_reason,omitempty"`
	Resolution          *map[string]any `json:"resolution,omitempty"`
	// ResolvedBy - Actor behind a participant.
	ResolvedBy *SkillsMergeRequestsItemsBulkResolveOutputItemsResolvedBy `json:"resolved_by,omitempty"`
	ResolvedAt *time.Time                                                `json:"resolved_at,omitempty"`
	AppliedAt  *time.Time                                                `json:"applied_at,omitempty"`
	CreatedAt  time.Time                                                 `json:"created_at"`
	UpdatedAt  time.Time                                                 `json:"updated_at"`
}

// SkillsMergeRequestsItemsBulkResolveOutputPagination represents the skills merge requests items bulk resolve output pagination type.
type SkillsMergeRequestsItemsBulkResolveOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsMergeRequestsItemsBulkResolveOutput represents the skills merge requests items bulk resolve output type.
type SkillsMergeRequestsItemsBulkResolveOutput struct {
	// Object - String representing the object’s type.
	Object     string                                              `json:"object"`
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

// SkillsMergeRequestsItemsBulkResolveBodyItemsResolution - Custom content for the resolution. Required for `edit_document` and `replace_file`.
type SkillsMergeRequestsItemsBulkResolveBodyItemsResolution struct {
	// Title - Document title for the resolution.
	Title *string `json:"title,omitempty"`
	// Content - Document or file content for the resolution.
	Content *string `json:"content,omitempty"`
	// FileId - Replacement file for the resolution, by ID.
	FileId *string `json:"fileId,omitempty"`
}

// SkillsMergeRequestsItemsBulkResolveBodyItems represents the skills merge requests items bulk resolve body items type.
type SkillsMergeRequestsItemsBulkResolveBodyItems struct {
	// ItemId - Proposed change to resolve, by ID.
	ItemId string `json:"item_id"`
	// ResolutionType - How to resolve the proposed change.
	ResolutionType string `json:"resolution_type"`
	// Resolution - Custom content for the resolution. Required for `edit_document` and `replace_file`.
	Resolution *SkillsMergeRequestsItemsBulkResolveBodyItemsResolution `json:"resolution,omitempty"`
}

// SkillsMergeRequestsItemsBulkResolveBody represents the skills merge requests items bulk resolve body type.
type SkillsMergeRequestsItemsBulkResolveBody struct {
	// Items - Resolutions to save.
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
