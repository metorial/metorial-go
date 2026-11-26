package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/aigateway/providerconnections"
)

// AiGatewayProviderConnectionsEndpoint provides access to your own credentials for an upstream provider such as OpenAI or Anthropic. The gateway routes requests through them, so the provider bills you directly.
type AiGatewayProviderConnectionsEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayProviderConnectionsEndpoint creates a new AiGatewayProviderConnectionsEndpoint.
func NewAiGatewayProviderConnectionsEndpoint(client *endpoint.Client) *AiGatewayProviderConnectionsEndpoint {
	return &AiGatewayProviderConnectionsEndpoint{client: client}
}

// AiGatewayProviderConnectionsEndpointListParams contains optional query parameters for List.
type AiGatewayProviderConnectionsEndpointListParams struct {
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
	// Provider - Only connections to the provider with this slug, e.g. openai
	Provider *string `json:"provider,omitempty"`
	// Status - Only connections with one of these statuses
	Status *any `json:"status,omitempty"`
}

// AiGatewayProviderConnectionsEndpointCreateBody contains the request body for Create.
type AiGatewayProviderConnectionsEndpointCreateBody struct {
	// Provider - The slug of the provider to connect, e.g. openai or anthropic
	Provider string `json:"provider"`
	// Name - A name for the connection
	Name string `json:"name"`
	// Description - A description of the connection
	Description *string `json:"description,omitempty"`
	// Priority - Connections with a higher priority are tried first when several connections can serve a request. Between -1000 and 1000, defaults to 0.
	Priority *float64 `json:"priority,omitempty"`
	// Credentials - The credentials for the provider, e.g. { "apiKey": "sk-..." }. Checked with the provider before they are saved, and never returned.
	Credentials map[string]string `json:"credentials"`
	// Config - Provider-specific settings that are not secret, e.g. { "resourceName": "my-resource" } for Azure
	Config *map[string]any `json:"config,omitempty"`
	// Features - Capabilities of the provider account behind the connection. Replaces the existing features when updated.
	Features *map[string]any `json:"features,omitempty"`
	// PricingOverride - Prices you pay this provider when they differ from the catalog, e.g. because of a negotiated discount. Used to calculate the cost of requests routed through the connection. Replaces the existing override when updated; pass null to use catalog prices.
	PricingOverride *nullable.Nullable[map[string]any] `json:"pricing_override,omitempty"`
}

// AiGatewayProviderConnectionsEndpointUpdateBody contains the request body for Update.
type AiGatewayProviderConnectionsEndpointUpdateBody struct {
	// Name - A name for the connection
	Name *string `json:"name,omitempty"`
	// Description - A description of the connection
	Description *string `json:"description,omitempty"`
	// Priority - Connections with a higher priority are tried first when several connections can serve a request. Between -1000 and 1000, defaults to 0.
	Priority *float64 `json:"priority,omitempty"`
	// Credentials - New credentials that replace the stored ones. Checked with the provider before they are saved, and never returned.
	Credentials *map[string]string `json:"credentials,omitempty"`
	// Config - Provider-specific settings that are not secret, e.g. { "resourceName": "my-resource" } for Azure
	Config *map[string]any `json:"config,omitempty"`
	// Features - Capabilities of the provider account behind the connection. Replaces the existing features when updated.
	Features *map[string]any `json:"features,omitempty"`
	// PricingOverride - Prices you pay this provider when they differ from the catalog, e.g. because of a negotiated discount. Used to calculate the cost of requests routed through the connection. Replaces the existing override when updated; pass null to use catalog prices.
	PricingOverride *nullable.Nullable[map[string]any] `json:"pricing_override,omitempty"`
}

// List returns a paginated list of provider connections. Deleted connections are not included.
func (e *AiGatewayProviderConnectionsEndpoint) List(params *AiGatewayProviderConnectionsEndpointListParams) (*providerconnections.AiGatewayProviderConnectionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"ai-gateway", "provider-connections"},
		Query: query,
	}
	var result providerconnections.AiGatewayProviderConnectionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a provider connection by ID. The credentials are never returned.
func (e *AiGatewayProviderConnectionsEndpoint) Get(connectionId string) (*providerconnections.AiGatewayProviderConnectionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "provider-connections", connectionId},
	}
	var result providerconnections.AiGatewayProviderConnectionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create connects a provider with your own credentials. The credentials are checked with the provider first, and the request fails if the provider rejects them.
func (e *AiGatewayProviderConnectionsEndpoint) Create(body *AiGatewayProviderConnectionsEndpointCreateBody) (*providerconnections.AiGatewayProviderConnectionsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "provider-connections"},
		Body: body,
	}
	var result providerconnections.AiGatewayProviderConnectionsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a provider connection. New credentials or config are checked with the provider first, and the request fails if the provider rejects them.
func (e *AiGatewayProviderConnectionsEndpoint) Update(connectionId string, body *AiGatewayProviderConnectionsEndpointUpdateBody) (*providerconnections.AiGatewayProviderConnectionsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "provider-connections", connectionId},
		Body: body,
	}
	var result providerconnections.AiGatewayProviderConnectionsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a provider connection. Requests are no longer routed through it and its credentials are discarded.
func (e *AiGatewayProviderConnectionsEndpoint) Delete(connectionId string) (*providerconnections.AiGatewayProviderConnectionsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "provider-connections", connectionId},
	}
	var result providerconnections.AiGatewayProviderConnectionsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Check checks the credentials with the provider now and updates the status. Disabled connections cannot be checked.
func (e *AiGatewayProviderConnectionsEndpoint) Check(connectionId string) (*providerconnections.AiGatewayProviderConnectionsCheckOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "provider-connections", connectionId, "check"},
	}
	var result providerconnections.AiGatewayProviderConnectionsCheckOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Enable turns a disabled connection back on and checks its credentials with the provider.
func (e *AiGatewayProviderConnectionsEndpoint) Enable(connectionId string) (*providerconnections.AiGatewayProviderConnectionsEnableOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "provider-connections", connectionId, "enable"},
	}
	var result providerconnections.AiGatewayProviderConnectionsEnableOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Disable turns a connection off without deleting it. Requests are not routed through it and its credentials are not checked until it is enabled again.
func (e *AiGatewayProviderConnectionsEndpoint) Disable(connectionId string) (*providerconnections.AiGatewayProviderConnectionsDisableOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "provider-connections", connectionId, "disable"},
	}
	var result providerconnections.AiGatewayProviderConnectionsDisableOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
