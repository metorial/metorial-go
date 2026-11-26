package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/magicmcpsessions"
)

// MagicMcpSessionsEndpoint provides access to a magic MCP session is the live connection behind a magic MCP server or endpoint. Sessions are created on demand when a client connects.
type MagicMcpSessionsEndpoint struct {
	client *endpoint.Client
}

// NewMagicMcpSessionsEndpoint creates a new MagicMcpSessionsEndpoint.
func NewMagicMcpSessionsEndpoint(client *endpoint.Client) *MagicMcpSessionsEndpoint {
	return &MagicMcpSessionsEndpoint{client: client}
}

// MagicMcpSessionsEndpointListParams contains optional query parameters for List.
type MagicMcpSessionsEndpointListParams struct {
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
	// MagicMcpServerId - Only return sessions for one of these servers.
	MagicMcpServerId *any `json:"magic_mcp_server_id,omitempty"`
}

// List returns a paginated list of magic MCP sessions.
func (e *MagicMcpSessionsEndpoint) List(params *MagicMcpSessionsEndpointListParams) (*magicmcpsessions.MagicMcpSessionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"magic-mcp-sessions"},
		Query: query,
	}
	var result magicmcpsessions.MagicMcpSessionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific magic MCP session.
func (e *MagicMcpSessionsEndpoint) Get(magicMcpSessionId string) (*magicmcpsessions.MagicMcpSessionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"magic-mcp-sessions", magicMcpSessionId},
	}
	var result magicmcpsessions.MagicMcpSessionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
