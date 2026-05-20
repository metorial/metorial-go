package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/marketplaces"
)

// SkillsMarketplacesEndpoint provides access to manage skill marketplaces for an instance.
type SkillsMarketplacesEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMarketplacesEndpoint creates a new SkillsMarketplacesEndpoint.
func NewSkillsMarketplacesEndpoint(client *endpoint.Client) *SkillsMarketplacesEndpoint {
	return &SkillsMarketplacesEndpoint{client: client}
}

// SkillsMarketplacesEndpointListParams contains optional query parameters for List.
type SkillsMarketplacesEndpointListParams struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	Status               *any     `json:"status,omitempty"`
	SkillConfigurationId *any     `json:"skill_configuration_id,omitempty"`
	Search               *string  `json:"search,omitempty"`
	// CreatedAt - Filter skill marketplace creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill marketplace last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsMarketplacesEndpointCreateBody contains the request body for Create.
type SkillsMarketplacesEndpointCreateBody struct {
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	ImageFileId          *string `json:"image_file_id,omitempty"`
	SkillConfigurationId *string `json:"skill_configuration_id,omitempty"`
}

// SkillsMarketplacesEndpointUpdateBody contains the request body for Update.
type SkillsMarketplacesEndpointUpdateBody struct {
	Name                 *string `json:"name,omitempty"`
	Description          *string `json:"description,omitempty"`
	ImageFileId          *string `json:"image_file_id,omitempty"`
	SkillConfigurationId *string `json:"skill_configuration_id,omitempty"`
}

// List returns a paginated list of skill marketplaces.
func (e *SkillsMarketplacesEndpoint) List(params *SkillsMarketplacesEndpointListParams) (*marketplaces.SkillsMarketplacesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skill-marketplaces"},
		Query: query,
	}
	var result marketplaces.SkillsMarketplacesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill marketplace.
func (e *SkillsMarketplacesEndpoint) Get(skillMarketplaceId string) (*marketplaces.SkillsMarketplacesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-marketplaces", skillMarketplaceId},
	}
	var result marketplaces.SkillsMarketplacesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a skill marketplace.
func (e *SkillsMarketplacesEndpoint) Create(body *SkillsMarketplacesEndpointCreateBody) (*marketplaces.SkillsMarketplacesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-marketplaces"},
		Body: body,
	}
	var result marketplaces.SkillsMarketplacesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a skill marketplace.
func (e *SkillsMarketplacesEndpoint) Update(skillMarketplaceId string, body *SkillsMarketplacesEndpointUpdateBody) (*marketplaces.SkillsMarketplacesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-marketplaces", skillMarketplaceId},
		Body: body,
	}
	var result marketplaces.SkillsMarketplacesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive archives a skill marketplace.
func (e *SkillsMarketplacesEndpoint) Archive(skillMarketplaceId string) (*marketplaces.SkillsMarketplacesArchiveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-marketplaces", skillMarketplaceId},
	}
	var result marketplaces.SkillsMarketplacesArchiveOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Sync forces a skill marketplace sync.
func (e *SkillsMarketplacesEndpoint) Sync(skillMarketplaceId string) (*marketplaces.SkillsMarketplacesSyncOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-marketplaces", skillMarketplaceId, "sync"},
	}
	var result marketplaces.SkillsMarketplacesSyncOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
