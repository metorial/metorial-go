package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/agents"
)

// AgentsEndpoint provides access to inspect agents and their linked clients and instances.
type AgentsEndpoint struct {
	client *endpoint.Client
}

// NewAgentsEndpoint creates a new AgentsEndpoint.
func NewAgentsEndpoint(client *endpoint.Client) *AgentsEndpoint {
	return &AgentsEndpoint{client: client}
}

// AgentsEndpointListParams contains optional query parameters for List.
type AgentsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Search *string  `json:"search,omitempty"`
	Status *any     `json:"status,omitempty"`
	Type   *any     `json:"type,omitempty"`
	Id     *any     `json:"id,omitempty"`
	// CreatedAt - Filter agent creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter agent last update time by date range
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

// Get retrieves a specific agent by ID.
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
