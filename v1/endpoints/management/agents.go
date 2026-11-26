package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/agents"
)

// AgentsEndpoint provides access to an agent is an identity that acts in your instance. It links to the clients and instances it runs as.
type AgentsEndpoint struct {
	client *endpoint.Client
}

// NewAgentsEndpoint creates a new AgentsEndpoint.
func NewAgentsEndpoint(client *endpoint.Client) *AgentsEndpoint {
	return &AgentsEndpoint{client: client}
}

// AgentsEndpointListParams contains optional query parameters for List.
type AgentsEndpointListParams struct {
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
	// Search - Filter agents by name or description.
	Search *string `json:"search,omitempty"`
	// Status - Only return agents with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Type - Only return agents with one of these types.
	Type *any `json:"type,omitempty"`
	// Id - Only return agents with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CreatedAt - Only return results whose agent creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose agent last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of agents for the instance.
func (e *AgentsEndpoint) List(instanceId string, params *AgentsEndpointListParams) (*agents.AgentsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "agents"},
		Query: query,
	}
	var result agents.AgentsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an agent.
func (e *AgentsEndpoint) Get(instanceId string, agentId string) (*agents.AgentsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "agents", agentId},
	}
	var result agents.AgentsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
