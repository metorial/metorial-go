package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/templates"
)

// SkillsTemplatesEndpoint provides access to skill templates define reusable starting points for skills.
type SkillsTemplatesEndpoint struct {
	client *endpoint.Client
}

// NewSkillsTemplatesEndpoint creates a new SkillsTemplatesEndpoint.
func NewSkillsTemplatesEndpoint(client *endpoint.Client) *SkillsTemplatesEndpoint {
	return &SkillsTemplatesEndpoint{client: client}
}

// SkillsTemplatesEndpointListParams contains optional query parameters for List.
type SkillsTemplatesEndpointListParams struct {
	Limit         *float64 `json:"limit,omitempty"`
	After         *string  `json:"after,omitempty"`
	Before        *string  `json:"before,omitempty"`
	Cursor        *string  `json:"cursor,omitempty"`
	Order         *string  `json:"order,omitempty"`
	Search        *string  `json:"search,omitempty"`
	Status        *any     `json:"status,omitempty"`
	Owner         *any     `json:"owner,omitempty"`
	Id            *any     `json:"id,omitempty"`
	ProviderId    *any     `json:"provider_id,omitempty"`
	IntegrationId *any     `json:"integration_id,omitempty"`
	// CreatedAt - Filter skill template creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill template last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsTemplatesEndpointCreateBody contains the request body for Create.
type SkillsTemplatesEndpointCreateBody struct {
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	FromSkillId *string         `json:"from_skill_Id,omitempty"`
}

// SkillsTemplatesEndpointUpdateBody contains the request body for Update.
type SkillsTemplatesEndpointUpdateBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of skill templates.
func (e *SkillsTemplatesEndpoint) List(params *SkillsTemplatesEndpointListParams) (*templates.SkillsTemplatesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skill-templates"},
		Query: query,
	}
	var result templates.SkillsTemplatesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific skill template.
func (e *SkillsTemplatesEndpoint) Get(skillTemplateId string) (*templates.SkillsTemplatesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-templates", skillTemplateId},
	}
	var result templates.SkillsTemplatesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a skill template.
func (e *SkillsTemplatesEndpoint) Create(body *SkillsTemplatesEndpointCreateBody) (*templates.SkillsTemplatesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-templates"},
		Body: body,
	}
	var result templates.SkillsTemplatesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a skill template.
func (e *SkillsTemplatesEndpoint) Update(skillTemplateId string, body *SkillsTemplatesEndpointUpdateBody) (*templates.SkillsTemplatesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-templates", skillTemplateId},
		Body: body,
	}
	var result templates.SkillsTemplatesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a skill template.
func (e *SkillsTemplatesEndpoint) Delete(skillTemplateId string) (*templates.SkillsTemplatesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-templates", skillTemplateId},
	}
	var result templates.SkillsTemplatesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
