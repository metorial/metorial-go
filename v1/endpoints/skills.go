package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/skills"
)

// SkillsEndpoint provides access to a skill is a reusable composition of provider and integration capabilities.
type SkillsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsEndpoint creates a new SkillsEndpoint.
func NewSkillsEndpoint(client *endpoint.Client) *SkillsEndpoint {
	return &SkillsEndpoint{client: client}
}

// SkillsEndpointListParams contains optional query parameters for List.
type SkillsEndpointListParams struct {
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
	// Search - Only return skills matching this search text.
	Search *string `json:"search,omitempty"`
	// Status - Only return skills with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return skills with these IDs.
	Id *any `json:"id,omitempty"`
	// SkillGroupId - Only return skills in one of these skill groups.
	SkillGroupId *any `json:"skill_group_id,omitempty"`
	// IntegrationId - Only return skills using one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// ProviderId - Only return skills using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// OwnedByMe - Only return skills the caller created or owns, including ones created through their other member or consumer identities.
	OwnedByMe *bool `json:"owned_by_me,omitempty"`
	// CreatedAt - Only return results whose skill creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsEndpointCreateBody contains the request body for Create.
type SkillsEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the skill.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ClientName - Client-facing skill name. Defaults to `name`. Lowercase letters, numbers, and hyphens.
	ClientName *string `json:"client_name,omitempty"`
	// ClientDescription - Client-facing description.
	ClientDescription *string `json:"client_description,omitempty"`
	// License - License terms for the skill.
	License *string `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill.
	Compatibility *string `json:"compatibility,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill.
	ClientMetadata *map[string]any `json:"client_metadata,omitempty"`
	// ImageFileId - Image file for the skill, by ID.
	ImageFileId *string `json:"image_file_id,omitempty"`
	// TemplateId - Skill template to create the skill from. Defaults to the default template.
	TemplateId *string `json:"template_id,omitempty"`
	// SkillGroupId - Skill group to add the skill to.
	SkillGroupId *string `json:"skill_group_id,omitempty"`
}

// SkillsEndpointUpdateBody contains the request body for Update.
type SkillsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// ClientName - Client-facing skill name. Lowercase letters, numbers, and hyphens.
	ClientName *string `json:"client_name,omitempty"`
	// ClientDescription - Client-facing description.
	ClientDescription *string `json:"client_description,omitempty"`
	// License - License terms for the skill. `null` clears it.
	License *nullable.Nullable[string] `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill. `null` clears it.
	Compatibility *nullable.Nullable[string] `json:"compatibility,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill. `null` clears it.
	ClientMetadata *nullable.Nullable[map[string]any] `json:"client_metadata,omitempty"`
	// Metadata - Custom data to identify or organize the skill. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
	// ImageFileId - Image file for the skill, by ID. `null` resets to the default image. Cannot be combined with `image_icon`.
	ImageFileId *nullable.Nullable[string] `json:"image_file_id,omitempty"`
	// ImageIcon - Use a library icon with a tile color as the skill image. Pass null to reset to the default image. Cannot be combined with image_file_id.
	ImageIcon *nullable.Nullable[map[string]any] `json:"image_icon,omitempty"`
}

// SkillsEndpointForkBody contains the request body for Fork.
type SkillsEndpointForkBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ClientName - Client-facing skill name. Defaults to `name`. Lowercase letters, numbers, and hyphens.
	ClientName *string `json:"client_name,omitempty"`
	// ClientDescription - Client-facing description.
	ClientDescription *string `json:"client_description,omitempty"`
	// License - License terms for the skill.
	License *string `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill.
	Compatibility *string `json:"compatibility,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill.
	ClientMetadata *map[string]any `json:"client_metadata,omitempty"`
	// Metadata - Custom data to identify or organize the skill.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ImageFileId - Image file for the skill, by ID.
	ImageFileId *string `json:"image_file_id,omitempty"`
}

// SkillsEndpointShareBody contains the request body for Share.
type SkillsEndpointShareBody struct {
	// ConsumerProfileIds - Consumers to share the skill with, by ID.
	ConsumerProfileIds *[]string `json:"consumer_profile_ids,omitempty"`
	// OrganizationMemberIds - Must be omitted. Organization member access cannot be changed.
	OrganizationMemberIds *[]string `json:"organization_member_ids,omitempty"`
	// Permission - For consumers with inherited edit and share access, `write` and `owner` add portal visibility. Consumers must own the skill to grant ownership. Only admins can remove or downgrade an owner.
	Permission string `json:"permission"`
}

// SkillsEndpointDuplicateBody contains the request body for Duplicate.
type SkillsEndpointDuplicateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ClientName - Client-facing skill name. Defaults to `name`. Lowercase letters, numbers, and hyphens.
	ClientName *string `json:"client_name,omitempty"`
	// ClientDescription - Client-facing description.
	ClientDescription *string `json:"client_description,omitempty"`
	// License - License terms for the skill.
	License *string `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill.
	Compatibility *string `json:"compatibility,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill.
	ClientMetadata *map[string]any `json:"client_metadata,omitempty"`
	// Metadata - Custom data to identify or organize the skill.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of skills.
func (e *SkillsEndpoint) List(params *SkillsEndpointListParams) (*skills.SkillsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skills"},
		Query: query,
	}
	var result skills.SkillsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill.
func (e *SkillsEndpoint) Get(skillId string) (*skills.SkillsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId},
	}
	var result skills.SkillsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new skill.
func (e *SkillsEndpoint) Create(body *SkillsEndpointCreateBody) (*skills.SkillsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills"},
		Body: body,
	}
	var result skills.SkillsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a skill. Omitted fields keep their current value.
func (e *SkillsEndpoint) Update(skillId string, body *SkillsEndpointUpdateBody) (*skills.SkillsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId},
		Body: body,
	}
	var result skills.SkillsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a skill.
func (e *SkillsEndpoint) Delete(skillId string) (*skills.SkillsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId},
	}
	var result skills.SkillsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fork forks a skill for the current consumer. Non-consumer callers duplicate the skill instead.
func (e *SkillsEndpoint) Fork(skillId string, body *SkillsEndpointForkBody) (*skills.SkillsForkOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "fork"},
		Body: body,
	}
	var result skills.SkillsForkOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PublishConsumerSkill publishes a consumer-owned skill to the consumer groups they belong to.
func (e *SkillsEndpoint) PublishConsumerSkill(skillId string) (*skills.SkillsPublishConsumerSkillOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "publish"},
	}
	var result skills.SkillsPublishConsumerSkillOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Share shares a skill with consumers.
func (e *SkillsEndpoint) Share(skillId string, body *SkillsEndpointShareBody) (*skills.SkillsShareOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "shares"},
		Body: body,
	}
	var result skills.SkillsShareOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Duplicate duplicates a skill.
func (e *SkillsEndpoint) Duplicate(skillId string, body *SkillsEndpointDuplicateBody) (*skills.SkillsDuplicateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "duplicate"},
		Body: body,
	}
	var result skills.SkillsDuplicateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
