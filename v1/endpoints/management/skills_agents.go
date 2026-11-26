package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/skills/agents"
)

// SkillsAgentsEndpoint provides access to a skill agent is a reusable sub-agent document stored in a skill.
type SkillsAgentsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsAgentsEndpoint creates a new SkillsAgentsEndpoint.
func NewSkillsAgentsEndpoint(client *endpoint.Client) *SkillsAgentsEndpoint {
	return &SkillsAgentsEndpoint{client: client}
}

// SkillsAgentsEndpointCreateBody contains the request body for Create.
type SkillsAgentsEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Content - Agent document content.
	Content *string `json:"content,omitempty"`
}

// SkillsAgentsEndpointListParams contains optional query parameters for List.
type SkillsAgentsEndpointListParams struct {
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
	// IncludeArchived - Include archived agents in the results.
	IncludeArchived *bool `json:"include_archived,omitempty"`
}

// SkillsAgentsEndpointUpdateBody contains the request body for Update.
type SkillsAgentsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
}

// Create creates a new agent document in the skill agents directory.
func (e *SkillsAgentsEndpoint) Create(instanceId string, skillId string, body *SkillsAgentsEndpointCreateBody) (*agents.SkillsAgentsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "agents"},
		Body: body,
	}
	var result agents.SkillsAgentsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List returns a paginated list of skill agents.
func (e *SkillsAgentsEndpoint) List(instanceId string, skillId string, params *SkillsAgentsEndpointListParams) (*agents.SkillsAgentsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skills", skillId, "agents"},
		Query: query,
	}
	var result agents.SkillsAgentsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill agent.
func (e *SkillsAgentsEndpoint) Get(instanceId string, skillId string, skillAgentId string) (*agents.SkillsAgentsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "agents", skillAgentId},
	}
	var result agents.SkillsAgentsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a skill agent. Omitted fields keep their current value.
func (e *SkillsAgentsEndpoint) Update(instanceId string, skillId string, skillAgentId string, body *SkillsAgentsEndpointUpdateBody) (*agents.SkillsAgentsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "agents", skillAgentId},
		Body: body,
	}
	var result agents.SkillsAgentsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a skill agent and removes its linked store item.
func (e *SkillsAgentsEndpoint) Delete(instanceId string, skillId string, skillAgentId string) (*agents.SkillsAgentsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "agents", skillAgentId},
	}
	var result agents.SkillsAgentsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
