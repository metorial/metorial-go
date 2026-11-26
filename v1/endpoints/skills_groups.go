package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/skills/groups"
)

// SkillsGroupsEndpoint provides access to a skill group is a reusable collection of skills.
type SkillsGroupsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsGroupsEndpoint creates a new SkillsGroupsEndpoint.
func NewSkillsGroupsEndpoint(client *endpoint.Client) *SkillsGroupsEndpoint {
	return &SkillsGroupsEndpoint{client: client}
}

// SkillsGroupsEndpointListParams contains optional query parameters for List.
type SkillsGroupsEndpointListParams struct {
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
	// Search - Only return skill groups matching this search text.
	Search *string `json:"search,omitempty"`
	// Status - Only return skill groups with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return skill groups with these IDs.
	Id *any `json:"id,omitempty"`
	// SkillId - Only return skill groups containing one of these skills.
	SkillId *any `json:"skill_id,omitempty"`
	// CreatedAt - Only return results whose skill group creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill group last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsGroupsEndpointCreateBody contains the request body for Create.
type SkillsGroupsEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the skill group.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SkillIds - Skills to add to the group, by ID.
	SkillIds *[]string `json:"skill_ids,omitempty"`
	// AllowConsumerSkillAssignment - Allow consumers to assign their skills to the group.
	AllowConsumerSkillAssignment *bool `json:"allow_consumer_skill_assignment,omitempty"`
}

// SkillsGroupsEndpointUpdateBody contains the request body for Update.
type SkillsGroupsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the skill group. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
	// SkillIds - Skills in the group, by ID. Replaces existing membership.
	SkillIds *[]string `json:"skill_ids,omitempty"`
	// AllowConsumerSkillAssignment - Allow consumers to assign their skills to the group.
	AllowConsumerSkillAssignment *bool `json:"allow_consumer_skill_assignment,omitempty"`
	// ImageFileId - Image file for the skill group, by ID. `null` resets to the default image. Cannot be combined with `image_icon`.
	ImageFileId *nullable.Nullable[string] `json:"image_file_id,omitempty"`
	// ImageIcon - Use a library icon with a tile color as the skill group image. Pass null to reset to the default image. Cannot be combined with image_file_id.
	ImageIcon *nullable.Nullable[map[string]any] `json:"image_icon,omitempty"`
}

// List returns a paginated list of skill groups.
func (e *SkillsGroupsEndpoint) List(params *SkillsGroupsEndpointListParams) (*groups.SkillsGroupsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skill-groups"},
		Query: query,
	}
	var result groups.SkillsGroupsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill group.
func (e *SkillsGroupsEndpoint) Get(skillGroupId string) (*groups.SkillsGroupsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-groups", skillGroupId},
	}
	var result groups.SkillsGroupsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a skill group.
func (e *SkillsGroupsEndpoint) Create(body *SkillsGroupsEndpointCreateBody) (*groups.SkillsGroupsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-groups"},
		Body: body,
	}
	var result groups.SkillsGroupsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a skill group. Omitted fields keep their current value.
func (e *SkillsGroupsEndpoint) Update(skillGroupId string, body *SkillsGroupsEndpointUpdateBody) (*groups.SkillsGroupsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-groups", skillGroupId},
		Body: body,
	}
	var result groups.SkillsGroupsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a skill group.
func (e *SkillsGroupsEndpoint) Delete(skillGroupId string) (*groups.SkillsGroupsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-groups", skillGroupId},
	}
	var result groups.SkillsGroupsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
