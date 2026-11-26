package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/agents/instances"
)

// AgentsInstancesEndpoint provides access to an agent is an identity that acts in your instance. It links to the clients and instances it runs as.
type AgentsInstancesEndpoint struct {
	client *endpoint.Client
}

// NewAgentsInstancesEndpoint creates a new AgentsInstancesEndpoint.
func NewAgentsInstancesEndpoint(client *endpoint.Client) *AgentsInstancesEndpoint {
	return &AgentsInstancesEndpoint{client: client}
}

// AgentsInstancesEndpointListParams contains optional query parameters for List.
type AgentsInstancesEndpointListParams struct {
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
	// Type - Only return agent instances with one of these types.
	Type *any `json:"type,omitempty"`
	// Id - Only return agent instances with one of these IDs.
	Id *any `json:"id,omitempty"`
	// AgentClientId - Only return agent instances linked to one of these agent clients.
	AgentClientId *any `json:"agent_client_id,omitempty"`
	// CreatedAt - Only return results whose agent instance creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose agent instance last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of instances for an agent.
func (e *AgentsInstancesEndpoint) List(instanceId string, agentId string, params *AgentsInstancesEndpointListParams) (*instances.AgentsInstancesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "agents", agentId, "instances"},
		Query: query,
	}
	var result instances.AgentsInstancesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an agent instance.
func (e *AgentsInstancesEndpoint) Get(instanceId string, agentId string, agentInstanceId string) (*instances.AgentsInstancesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "agents", agentId, "instances", agentInstanceId},
	}
	var result instances.AgentsInstancesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
