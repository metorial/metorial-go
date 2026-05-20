package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/versions"
)

// SkillsVersionsEndpoint provides access to inspect version history and snapshots for a skill.
type SkillsVersionsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsVersionsEndpoint creates a new SkillsVersionsEndpoint.
func NewSkillsVersionsEndpoint(client *endpoint.Client) *SkillsVersionsEndpoint {
	return &SkillsVersionsEndpoint{client: client}
}

// SkillsVersionsEndpointListParams contains optional query parameters for List.
type SkillsVersionsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// List returns a paginated list of versions for a specific skill.
func (e *SkillsVersionsEndpoint) List(instanceId string, skillId string, params *SkillsVersionsEndpointListParams) (*versions.SkillsVersionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skills", skillId, "versions"},
		Query: query,
	}
	var result versions.SkillsVersionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific skill version by its ID.
func (e *SkillsVersionsEndpoint) Get(instanceId string, skillId string, skillVersionId string) (*versions.SkillsVersionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "versions", skillVersionId},
	}
	var result versions.SkillsVersionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
