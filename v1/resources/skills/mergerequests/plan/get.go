package plan

import (
	"encoding/json"
	"time"
)

// SkillsMergeRequestsPlanGetOutputMergeRequestCreatedByOrganizationActorMember represents the skills merge requests plan get output merge request created by organization actor member type.
type SkillsMergeRequestsPlanGetOutputMergeRequestCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsPlanGetOutputMergeRequestCreatedByOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsPlanGetOutputMergeRequestCreatedByOrganizationActorTeams struct {
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

// SkillsMergeRequestsPlanGetOutputMergeRequestCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsMergeRequestsPlanGetOutputMergeRequestCreatedByOrganizationActor struct {
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
	ImageUrl string                                                                        `json:"image_url"`
	Member   *SkillsMergeRequestsPlanGetOutputMergeRequestCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsPlanGetOutputMergeRequestCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsPlanGetOutputMergeRequestCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsMergeRequestsPlanGetOutputMergeRequestCreatedByConsumer struct {
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

// SkillsMergeRequestsPlanGetOutputMergeRequestCreatedBy - Actor behind a participant.
type SkillsMergeRequestsPlanGetOutputMergeRequestCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsMergeRequestsPlanGetOutputMergeRequestCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsMergeRequestsPlanGetOutputMergeRequestCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsPlanGetOutputMergeRequest represents the skills merge requests plan get output merge request type.
type SkillsMergeRequestsPlanGetOutputMergeRequest struct {
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
	CreatedBy      *SkillsMergeRequestsPlanGetOutputMergeRequestCreatedBy `json:"created_by,omitempty"`
	ItemCount      float64                                                `json:"item_count"`
	CommentCount   float64                                                `json:"comment_count"`
	MergeStartedAt *time.Time                                             `json:"merge_started_at,omitempty"`
	MergedAt       *time.Time                                             `json:"merged_at,omitempty"`
	ClosedAt       *time.Time                                             `json:"closed_at,omitempty"`
	RolledBackAt   *time.Time                                             `json:"rolled_back_at,omitempty"`
	CreatedAt      time.Time                                              `json:"created_at"`
	UpdatedAt      time.Time                                              `json:"updated_at"`
}

// SkillsMergeRequestsPlanGetOutputItemsItemResolvedByOrganizationActorMember represents the skills merge requests plan get output items item resolved by organization actor member type.
type SkillsMergeRequestsPlanGetOutputItemsItemResolvedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsMergeRequestsPlanGetOutputItemsItemResolvedByOrganizationActorTeams - The teams the actor belongs to
type SkillsMergeRequestsPlanGetOutputItemsItemResolvedByOrganizationActorTeams struct {
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

// SkillsMergeRequestsPlanGetOutputItemsItemResolvedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsMergeRequestsPlanGetOutputItemsItemResolvedByOrganizationActor struct {
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
	ImageUrl string                                                                      `json:"image_url"`
	Member   *SkillsMergeRequestsPlanGetOutputItemsItemResolvedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsMergeRequestsPlanGetOutputItemsItemResolvedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsMergeRequestsPlanGetOutputItemsItemResolvedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsMergeRequestsPlanGetOutputItemsItemResolvedByConsumer struct {
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

// SkillsMergeRequestsPlanGetOutputItemsItemResolvedBy - Actor behind a participant.
type SkillsMergeRequestsPlanGetOutputItemsItemResolvedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsMergeRequestsPlanGetOutputItemsItemResolvedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsMergeRequestsPlanGetOutputItemsItemResolvedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsMergeRequestsPlanGetOutputItemsItem represents the skills merge requests plan get output items item type.
type SkillsMergeRequestsPlanGetOutputItemsItem struct {
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
	ResolvedBy *SkillsMergeRequestsPlanGetOutputItemsItemResolvedBy `json:"resolved_by,omitempty"`
	ResolvedAt *time.Time                                           `json:"resolved_at,omitempty"`
	AppliedAt  *time.Time                                           `json:"applied_at,omitempty"`
	CreatedAt  time.Time                                            `json:"created_at"`
	UpdatedAt  time.Time                                            `json:"updated_at"`
}

// SkillsMergeRequestsPlanGetOutputItemsBase represents the skills merge requests plan get output items base type.
type SkillsMergeRequestsPlanGetOutputItemsBase struct {
	Kind              string  `json:"kind"`
	Path              string  `json:"path"`
	FileId            *string `json:"file_id,omitempty"`
	DocumentId        *string `json:"document_id,omitempty"`
	DocumentTitle     *string `json:"document_title,omitempty"`
	DocumentVersionId *string `json:"document_version_id,omitempty"`
	Content           *string `json:"content,omitempty"`
}

// SkillsMergeRequestsPlanGetOutputItemsSource represents the skills merge requests plan get output items source type.
type SkillsMergeRequestsPlanGetOutputItemsSource struct {
	Kind              string  `json:"kind"`
	Path              string  `json:"path"`
	FileId            *string `json:"file_id,omitempty"`
	DocumentId        *string `json:"document_id,omitempty"`
	DocumentTitle     *string `json:"document_title,omitempty"`
	DocumentVersionId *string `json:"document_version_id,omitempty"`
	Content           *string `json:"content,omitempty"`
}

// SkillsMergeRequestsPlanGetOutputItemsTarget represents the skills merge requests plan get output items target type.
type SkillsMergeRequestsPlanGetOutputItemsTarget struct {
	Kind              string  `json:"kind"`
	Path              string  `json:"path"`
	FileId            *string `json:"file_id,omitempty"`
	DocumentId        *string `json:"document_id,omitempty"`
	DocumentTitle     *string `json:"document_title,omitempty"`
	DocumentVersionId *string `json:"document_version_id,omitempty"`
	Content           *string `json:"content,omitempty"`
}

// SkillsMergeRequestsPlanGetOutputItemsDocumentMerge represents the skills merge requests plan get output items document merge type.
type SkillsMergeRequestsPlanGetOutputItemsDocumentMerge struct {
	BaseContent   *string `json:"base_content,omitempty"`
	SourceContent *string `json:"source_content,omitempty"`
	TargetContent *string `json:"target_content,omitempty"`
	HasConflict   bool    `json:"has_conflict"`
}

// SkillsMergeRequestsPlanGetOutputItems represents the skills merge requests plan get output items type.
type SkillsMergeRequestsPlanGetOutputItems struct {
	Item          SkillsMergeRequestsPlanGetOutputItemsItem           `json:"item"`
	Base          *SkillsMergeRequestsPlanGetOutputItemsBase          `json:"base,omitempty"`
	Source        *SkillsMergeRequestsPlanGetOutputItemsSource        `json:"source,omitempty"`
	Target        *SkillsMergeRequestsPlanGetOutputItemsTarget        `json:"target,omitempty"`
	DocumentMerge *SkillsMergeRequestsPlanGetOutputItemsDocumentMerge `json:"document_merge,omitempty"`
}

// SkillsMergeRequestsPlanGetOutput represents the skills merge requests plan get output type.
type SkillsMergeRequestsPlanGetOutput struct {
	Object       string                                       `json:"object"`
	MergeRequest SkillsMergeRequestsPlanGetOutputMergeRequest `json:"merge_request"`
	Items        []SkillsMergeRequestsPlanGetOutputItems      `json:"items"`
}

// MapSkillsMergeRequestsPlanGetOutputFromJSON deserializes JSON data into a SkillsMergeRequestsPlanGetOutput.
func MapSkillsMergeRequestsPlanGetOutputFromJSON(data []byte) (*SkillsMergeRequestsPlanGetOutput, error) {
	var v SkillsMergeRequestsPlanGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMergeRequestsPlanGetOutputToJSON serializes a SkillsMergeRequestsPlanGetOutput to JSON.
func MapSkillsMergeRequestsPlanGetOutputToJSON(v *SkillsMergeRequestsPlanGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
