package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/agents/instances"
)

// AgentsInstancesEndpoint provides access to inspect agents and their linked clients and instances.
type AgentsInstancesEndpoint struct {
	client *endpoint.Client
}

// NewAgentsInstancesEndpoint creates a new AgentsInstancesEndpoint.
func NewAgentsInstancesEndpoint(client *endpoint.Client) *AgentsInstancesEndpoint {
	return &AgentsInstancesEndpoint{client: client}
}

// AgentsInstancesEndpointListParams contains optional query parameters for List.
type AgentsInstancesEndpointListParams struct {
	Limit         *float64 `json:"limit,omitempty"`
	After         *string  `json:"after,omitempty"`
	Before        *string  `json:"before,omitempty"`
	Cursor        *string  `json:"cursor,omitempty"`
	Order         *string  `json:"order,omitempty"`
	Type          *any     `json:"type,omitempty"`
	Id            *any     `json:"id,omitempty"`
	AgentClientId *any     `json:"agent_client_id,omitempty"`
	// CreatedAt - Filter agent instance creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter agent instance last update time by date range
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

// Get retrieves a specific agent instance by ID.
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
