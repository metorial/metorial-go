package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/magicmcpendpoints"
)

// MagicMcpEndpointsEndpoint provides access to magic MCP endpoints combine multiple Magic MCP servers behind one routed connection target.
type MagicMcpEndpointsEndpoint struct {
	client *endpoint.Client
}

// NewMagicMcpEndpointsEndpoint creates a new MagicMcpEndpointsEndpoint.
func NewMagicMcpEndpointsEndpoint(client *endpoint.Client) *MagicMcpEndpointsEndpoint {
	return &MagicMcpEndpointsEndpoint{client: client}
}

// MagicMcpEndpointsEndpointListParams contains optional query parameters for List.
type MagicMcpEndpointsEndpointListParams struct {
	Limit            *float64 `json:"limit,omitempty"`
	After            *string  `json:"after,omitempty"`
	Before           *string  `json:"before,omitempty"`
	Cursor           *string  `json:"cursor,omitempty"`
	Order            *string  `json:"order,omitempty"`
	Status           *any     `json:"status,omitempty"`
	MagicMcpServerId *any     `json:"magic_mcp_server_id,omitempty"`
	Search           *string  `json:"search,omitempty"`
}

// MagicMcpEndpointsEndpointCreateBody contains the request body for Create.
type MagicMcpEndpointsEndpointCreateBody struct {
	Name              *string           `json:"name,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Metadata          *map[string]any   `json:"metadata,omitempty"`
	ConsumerProfileId *string           `json:"consumer_profile_id,omitempty"`
	MagicMcpServerIds *[]string         `json:"magic_mcp_server_ids,omitempty"`
	Servers           *[]map[string]any `json:"servers,omitempty"`
}

// MagicMcpEndpointsEndpointUpdateBody contains the request body for Update.
type MagicMcpEndpointsEndpointUpdateBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
}

// MagicMcpEndpointsEndpointAddServersBody contains the request body for AddServers.
type MagicMcpEndpointsEndpointAddServersBody struct {
	MagicMcpServerIds *[]string         `json:"magic_mcp_server_ids,omitempty"`
	Servers           *[]map[string]any `json:"servers,omitempty"`
}

// MagicMcpEndpointsEndpointRemoveServersBody contains the request body for RemoveServers.
type MagicMcpEndpointsEndpointRemoveServersBody struct {
	MagicMcpServerIds []string `json:"magic_mcp_server_ids"`
}

// List returns a paginated list of magic MCP endpoints.
func (e *MagicMcpEndpointsEndpoint) List(params *MagicMcpEndpointsEndpointListParams) (*magicmcpendpoints.MagicMcpEndpointsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"magic-mcp-endpoints"},
		Query: query,
	}
	var result magicmcpendpoints.MagicMcpEndpointsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific magic MCP endpoint.
func (e *MagicMcpEndpointsEndpoint) Get(magicMcpEndpointId string) (*magicmcpendpoints.MagicMcpEndpointsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-endpoints", magicMcpEndpointId},
	}
	var result magicmcpendpoints.MagicMcpEndpointsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a magic MCP endpoint.
func (e *MagicMcpEndpointsEndpoint) Create(body *MagicMcpEndpointsEndpointCreateBody) (*magicmcpendpoints.MagicMcpEndpointsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-endpoints"},
		Body: body,
	}
	var result magicmcpendpoints.MagicMcpEndpointsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a magic MCP endpoint.
func (e *MagicMcpEndpointsEndpoint) Delete(magicMcpEndpointId string) (*magicmcpendpoints.MagicMcpEndpointsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-endpoints", magicMcpEndpointId},
	}
	var result magicmcpendpoints.MagicMcpEndpointsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a magic MCP endpoint.
func (e *MagicMcpEndpointsEndpoint) Update(magicMcpEndpointId string, body *MagicMcpEndpointsEndpointUpdateBody) (*magicmcpendpoints.MagicMcpEndpointsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-endpoints", magicMcpEndpointId},
		Body: body,
	}
	var result magicmcpendpoints.MagicMcpEndpointsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AddServers adds magic MCP servers to a magic MCP endpoint.
func (e *MagicMcpEndpointsEndpoint) AddServers(magicMcpEndpointId string, body *MagicMcpEndpointsEndpointAddServersBody) (*magicmcpendpoints.MagicMcpEndpointsAddServersOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-endpoints", magicMcpEndpointId, "add-servers"},
		Body: body,
	}
	var result magicmcpendpoints.MagicMcpEndpointsAddServersOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RemoveServers removes magic MCP servers from a magic MCP endpoint.
func (e *MagicMcpEndpointsEndpoint) RemoveServers(magicMcpEndpointId string, body *MagicMcpEndpointsEndpointRemoveServersBody) (*magicmcpendpoints.MagicMcpEndpointsRemoveServersOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-endpoints", magicMcpEndpointId, "remove-servers"},
		Body: body,
	}
	var result magicmcpendpoints.MagicMcpEndpointsRemoveServersOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
