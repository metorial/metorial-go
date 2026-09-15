package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/imports"
)

// SkillsImportsEndpoint provides access to import skills from public repositories or uploaded files.
type SkillsImportsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsImportsEndpoint creates a new SkillsImportsEndpoint.
func NewSkillsImportsEndpoint(client *endpoint.Client) *SkillsImportsEndpoint {
	return &SkillsImportsEndpoint{client: client}
}

// SkillsImportsEndpointListParams contains optional query parameters for List.
type SkillsImportsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Id     *any     `json:"id,omitempty"`
	Status *any     `json:"status,omitempty"`
}

// SkillsImportsEndpointCreateBody contains the request body for Create.
type SkillsImportsEndpointCreateBody struct {
	Source any `json:"source"`
}

// List returns a paginated list of skill imports.
func (e *SkillsImportsEndpoint) List(instanceId string, params *SkillsImportsEndpointListParams) (*imports.SkillsImportsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skill-imports"},
		Query: query,
	}
	var result imports.SkillsImportsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an individual skill import and its results.
func (e *SkillsImportsEndpoint) Get(instanceId string, skillImportId string) (*imports.SkillsImportsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-imports", skillImportId},
	}
	var result imports.SkillsImportsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create queues a skill import from a repository or uploaded file.
func (e *SkillsImportsEndpoint) Create(instanceId string, body *SkillsImportsEndpointCreateBody) (*imports.SkillsImportsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-imports"},
		Body: body,
	}
	var result imports.SkillsImportsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
