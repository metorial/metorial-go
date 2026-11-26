package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/skills/marketplaces"
)

// SkillsMarketplacesEndpoint provides access to a skill marketplace is a curated collection of skill plugins.
type SkillsMarketplacesEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMarketplacesEndpoint creates a new SkillsMarketplacesEndpoint.
func NewSkillsMarketplacesEndpoint(client *endpoint.Client) *SkillsMarketplacesEndpoint {
	return &SkillsMarketplacesEndpoint{client: client}
}

// SkillsMarketplacesEndpointListParams contains optional query parameters for List.
type SkillsMarketplacesEndpointListParams struct {
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
	// Id - Only return skill marketplaces with these IDs.
	Id *any `json:"id,omitempty"`
	// Status - Only return skill marketplaces with one of these statuses.
	Status *any `json:"status,omitempty"`
	// SkillConfigurationId - Only return skill marketplaces using one of these skill configurations.
	SkillConfigurationId *any `json:"skill_configuration_id,omitempty"`
	// SkillGroupId - Only return marketplaces that contain an active plugin synced from these skill groups.
	SkillGroupId *any `json:"skill_group_id,omitempty"`
	// Search - Only return skill marketplaces matching this search text.
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose skill marketplace creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill marketplace last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsMarketplacesEndpointCreateBody contains the request body for Create.
type SkillsMarketplacesEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// ImageFileId - Image file for the marketplace, by ID. `null` resets to the default image.
	ImageFileId *nullable.Nullable[string] `json:"image_file_id,omitempty"`
	// SkillConfigurationId - Skill configuration for the marketplace, by ID. `null` uses the default configuration.
	SkillConfigurationId *nullable.Nullable[string] `json:"skill_configuration_id,omitempty"`
	// RepositoryAccessMode - How syncs reach linked repositories. `pull_request` opens a pull request, `default_branch` pushes directly.
	RepositoryAccessMode *string `json:"repository_access_mode,omitempty"`
	// ForceMergeOrPush - Merge or push even when the repository changed unexpectedly.
	ForceMergeOrPush *bool `json:"force_merge_or_push,omitempty"`
	// MergeBeforeChecksPass - Merge pull requests before checks pass. Only applies to `pull_request` mode.
	MergeBeforeChecksPass *bool `json:"merge_before_checks_pass,omitempty"`
}

// SkillsMarketplacesEndpointUpdateBody contains the request body for Update.
type SkillsMarketplacesEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// ImageFileId - Image file for the marketplace, by ID. `null` resets to the default image.
	ImageFileId *nullable.Nullable[string] `json:"image_file_id,omitempty"`
	// SkillConfigurationId - Skill configuration for the marketplace, by ID. `null` uses the default configuration.
	SkillConfigurationId *nullable.Nullable[string] `json:"skill_configuration_id,omitempty"`
	// RepositoryAccessMode - How syncs reach linked repositories. `pull_request` opens a pull request, `default_branch` pushes directly.
	RepositoryAccessMode *string `json:"repository_access_mode,omitempty"`
	// ForceMergeOrPush - Merge or push even when the repository changed unexpectedly.
	ForceMergeOrPush *bool `json:"force_merge_or_push,omitempty"`
	// MergeBeforeChecksPass - Merge pull requests before checks pass. Only applies to `pull_request` mode.
	MergeBeforeChecksPass *bool `json:"merge_before_checks_pass,omitempty"`
}

// List returns a paginated list of skill marketplaces.
func (e *SkillsMarketplacesEndpoint) List(instanceId string, params *SkillsMarketplacesEndpointListParams) (*marketplaces.SkillsMarketplacesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skill-marketplaces"},
		Query: query,
	}
	var result marketplaces.SkillsMarketplacesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill marketplace.
func (e *SkillsMarketplacesEndpoint) Get(instanceId string, skillMarketplaceId string) (*marketplaces.SkillsMarketplacesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-marketplaces", skillMarketplaceId},
	}
	var result marketplaces.SkillsMarketplacesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a skill marketplace.
func (e *SkillsMarketplacesEndpoint) Create(instanceId string, body *SkillsMarketplacesEndpointCreateBody) (*marketplaces.SkillsMarketplacesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-marketplaces"},
		Body: body,
	}
	var result marketplaces.SkillsMarketplacesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a skill marketplace. Omitted fields keep their current value.
func (e *SkillsMarketplacesEndpoint) Update(instanceId string, skillMarketplaceId string, body *SkillsMarketplacesEndpointUpdateBody) (*marketplaces.SkillsMarketplacesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-marketplaces", skillMarketplaceId},
		Body: body,
	}
	var result marketplaces.SkillsMarketplacesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive archives a skill marketplace.
func (e *SkillsMarketplacesEndpoint) Archive(instanceId string, skillMarketplaceId string) (*marketplaces.SkillsMarketplacesArchiveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-marketplaces", skillMarketplaceId},
	}
	var result marketplaces.SkillsMarketplacesArchiveOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Sync starts a sync of the skill marketplace to its linked repositories, even if nothing changed.
func (e *SkillsMarketplacesEndpoint) Sync(instanceId string, skillMarketplaceId string) (*marketplaces.SkillsMarketplacesSyncOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-marketplaces", skillMarketplaceId, "sync"},
	}
	var result marketplaces.SkillsMarketplacesSyncOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
