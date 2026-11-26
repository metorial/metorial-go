package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/customproviders"
)

// CustomProvidersEndpoint provides access to a custom provider is an MCP server you deploy and use through your instance. It runs from a container image, a remote URL, or serverless function code.
type CustomProvidersEndpoint struct {
	client *endpoint.Client
}

// NewCustomProvidersEndpoint creates a new CustomProvidersEndpoint.
func NewCustomProvidersEndpoint(client *endpoint.Client) *CustomProvidersEndpoint {
	return &CustomProvidersEndpoint{client: client}
}

// CustomProvidersEndpointListParams contains optional query parameters for List.
type CustomProvidersEndpointListParams struct {
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
	// Status - Only return custom providers with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Type - Only return custom providers with one of these types.
	Type *any `json:"type,omitempty"`
	// Id - Only return custom providers with one of these IDs.
	Id *any `json:"id,omitempty"`
	// ProviderId - Only return custom providers linked to one of these provider IDs (matches providers connected to sessions).
	ProviderId *any `json:"provider_id,omitempty"`
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose custom provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose custom provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// CustomProvidersEndpointCreateBody contains the request body for Create.
type CustomProvidersEndpointCreateBody struct {
	// Name - Display name for the custom provider.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the custom provider.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// From - Source the provider runs from.
	From any `json:"from"`
	// Config - Configuration schema and transformer for the provider.
	Config *map[string]any `json:"config,omitempty"`
}

// CustomProvidersEndpointUpdateBody contains the request body for Update.
type CustomProvidersEndpointUpdateBody struct {
	// Name - Display name for the custom provider.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the custom provider.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Readme - README content in markdown format
	Readme *string `json:"readme,omitempty"`
}

// List returns a paginated list of custom providers.
func (e *CustomProvidersEndpoint) List(params *CustomProvidersEndpointListParams) (*customproviders.CustomProvidersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"custom-providers"},
		Query: query,
	}
	var result customproviders.CustomProvidersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a custom provider.
func (e *CustomProvidersEndpoint) Get(customProviderId string) (*customproviders.CustomProvidersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"custom-providers", customProviderId},
	}
	var result customproviders.CustomProvidersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetEnv retrieves the environment variables for a custom provider.
func (e *CustomProvidersEndpoint) GetEnv(customProviderId string) (*customproviders.CustomProvidersGetEnvOutput, error) {
	req := &endpoint.Request{
		Path: []string{"custom-providers", customProviderId, "env"},
	}
	var result customproviders.CustomProvidersGetEnvOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new custom provider.
func (e *CustomProvidersEndpoint) Create(body *CustomProvidersEndpointCreateBody) (*customproviders.CustomProvidersCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"custom-providers"},
		Body: body,
	}
	var result customproviders.CustomProvidersCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a custom provider.
func (e *CustomProvidersEndpoint) Update(customProviderId string, body *CustomProvidersEndpointUpdateBody) (*customproviders.CustomProvidersUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"custom-providers", customProviderId},
		Body: body,
	}
	var result customproviders.CustomProvidersUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive archives a custom provider and disables new connections to it.
func (e *CustomProvidersEndpoint) Archive(customProviderId string) (*customproviders.CustomProvidersArchiveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"custom-providers", customProviderId, "archive"},
	}
	var result customproviders.CustomProvidersArchiveOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
