package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/agents"
)

// SkillsAgentsEndpoint provides access to manage sub-agents attached to a skill.
type SkillsAgentsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsAgentsEndpoint creates a new SkillsAgentsEndpoint.
func NewSkillsAgentsEndpoint(client *endpoint.Client) *SkillsAgentsEndpoint {
	return &SkillsAgentsEndpoint{client: client}
}

// SkillsAgentsEndpointCreateBody contains the request body for Create.
type SkillsAgentsEndpointCreateBody struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Content     *string `json:"content,omitempty"`
}

// SkillsAgentsEndpointListParams contains optional query parameters for List.
type SkillsAgentsEndpointListParams struct {
	Limit           *float64 `json:"limit,omitempty"`
	After           *string  `json:"after,omitempty"`
	Before          *string  `json:"before,omitempty"`
	Cursor          *string  `json:"cursor,omitempty"`
	Order           *string  `json:"order,omitempty"`
	IncludeArchived *bool    `json:"include_archived,omitempty"`
}

// SkillsAgentsEndpointUpdateBody contains the request body for Update.
type SkillsAgentsEndpointUpdateBody struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
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

// List returns a paginated list of agents for a specific skill.
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

// Get retrieves a specific agent within a skill.
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

// Update updates the name or description for a specific skill agent.
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

// Delete archives a specific skill agent and removes its linked store item.
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
