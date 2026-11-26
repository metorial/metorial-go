package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/skills/templates"
)

// SkillsTemplatesEndpoint provides access to a skill template is a reusable starting point for creating skills.
type SkillsTemplatesEndpoint struct {
	client *endpoint.Client
}

// NewSkillsTemplatesEndpoint creates a new SkillsTemplatesEndpoint.
func NewSkillsTemplatesEndpoint(client *endpoint.Client) *SkillsTemplatesEndpoint {
	return &SkillsTemplatesEndpoint{client: client}
}

// SkillsTemplatesEndpointListParams contains optional query parameters for List.
type SkillsTemplatesEndpointListParams struct {
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
	// Search - Only return skill templates matching this search text.
	Search *string `json:"search,omitempty"`
	// Status - Only return skill templates with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Owner - Only return skill templates with one of these owners.
	Owner *any `json:"owner,omitempty"`
	// Id - Only return skill templates with these IDs.
	Id *any `json:"id,omitempty"`
	// ProviderId - Only return skill templates using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationId - Only return skill templates using one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// CreatedAt - Only return results whose skill template creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill template last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsTemplatesEndpointCreateBody contains the request body for Create.
type SkillsTemplatesEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the skill template.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// FromSkillId - Skill to copy content and items from, by ID.
	FromSkillId *string `json:"from_skill_Id,omitempty"`
}

// SkillsTemplatesEndpointUpdateBody contains the request body for Update.
type SkillsTemplatesEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the skill template. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
}

// List returns a paginated list of skill templates.
func (e *SkillsTemplatesEndpoint) List(instanceId string, params *SkillsTemplatesEndpointListParams) (*templates.SkillsTemplatesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skill-templates"},
		Query: query,
	}
	var result templates.SkillsTemplatesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill template.
func (e *SkillsTemplatesEndpoint) Get(instanceId string, skillTemplateId string) (*templates.SkillsTemplatesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-templates", skillTemplateId},
	}
	var result templates.SkillsTemplatesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a skill template.
func (e *SkillsTemplatesEndpoint) Create(instanceId string, body *SkillsTemplatesEndpointCreateBody) (*templates.SkillsTemplatesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-templates"},
		Body: body,
	}
	var result templates.SkillsTemplatesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a skill template. Omitted fields keep their current value.
func (e *SkillsTemplatesEndpoint) Update(instanceId string, skillTemplateId string, body *SkillsTemplatesEndpointUpdateBody) (*templates.SkillsTemplatesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-templates", skillTemplateId},
		Body: body,
	}
	var result templates.SkillsTemplatesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a skill template.
func (e *SkillsTemplatesEndpoint) Delete(instanceId string, skillTemplateId string) (*templates.SkillsTemplatesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-templates", skillTemplateId},
	}
	var result templates.SkillsTemplatesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
