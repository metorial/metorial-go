package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/provider/tools"
)

// ProviderToolsEndpoint provides access to a tool is a single action a provider can perform like 'search_issues' or 'send_message'. Tools are what AI agents call via MCP. By default, tools from the latest provider version are returned. Use the optional version filter to get tools for a specific version.
type ProviderToolsEndpoint struct {
	client *endpoint.Client
}

// NewProviderToolsEndpoint creates a new ProviderToolsEndpoint.
func NewProviderToolsEndpoint(client *endpoint.Client) *ProviderToolsEndpoint {
	return &ProviderToolsEndpoint{client: client}
}

// ProviderToolsEndpointListParams contains optional query parameters for List.
type ProviderToolsEndpointListParams struct {
	Limit             *float64 `json:"limit,omitempty"`
	After             *string  `json:"after,omitempty"`
	Before            *string  `json:"before,omitempty"`
	Cursor            *string  `json:"cursor,omitempty"`
	Order             *string  `json:"order,omitempty"`
	ProviderVersionId string   `json:"provider_version_id"`
}

// List returns a paginated list of provider tools. By default returns tools from the latest version. Use optional filters to get tools for a specific version.
func (e *ProviderToolsEndpoint) List(params *ProviderToolsEndpointListParams) (*tools.ProviderToolsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-tools"},
		Query: query,
	}
	var result tools.ProviderToolsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider tool by ID.
func (e *ProviderToolsEndpoint) Get(providerToolId string) (*tools.ProviderToolsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-tools", providerToolId},
	}
	var result tools.ProviderToolsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
