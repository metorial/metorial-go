package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrations"
)

// IntegrationsEndpoint provides access to integrations define reusable provider contracts that can then be materialized into integration instances.
type IntegrationsEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsEndpoint creates a new IntegrationsEndpoint.
func NewIntegrationsEndpoint(client *endpoint.Client) *IntegrationsEndpoint {
	return &IntegrationsEndpoint{client: client}
}

// IntegrationsEndpointListParams contains optional query parameters for List.
type IntegrationsEndpointListParams struct {
	Limit                 *float64 `json:"limit,omitempty"`
	After                 *string  `json:"after,omitempty"`
	Before                *string  `json:"before,omitempty"`
	Cursor                *string  `json:"cursor,omitempty"`
	Order                 *string  `json:"order,omitempty"`
	Search                *string  `json:"search,omitempty"`
	Status                *any     `json:"status,omitempty"`
	Id                    *any     `json:"id,omitempty"`
	ProviderId            *any     `json:"provider_id,omitempty"`
	IntegrationProviderId *any     `json:"integration_provider_id,omitempty"`
	// CreatedAt - Filter integration creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsEndpointCreateBody contains the request body for Create.
type IntegrationsEndpointCreateBody struct {
	Name                          string          `json:"name"`
	Description                   *string         `json:"description,omitempty"`
	Metadata                      *map[string]any `json:"metadata,omitempty"`
	CanAttachCustomToolFilters    *bool           `json:"can_attach_custom_tool_filters,omitempty"`
	CanAttachCustomProviderConfig *bool           `json:"can_attach_custom_provider_config,omitempty"`
	CanOverrideToolFilters        *bool           `json:"can_override_tool_filters,omitempty"`
}

// IntegrationsEndpointUpdateBody contains the request body for Update.
type IntegrationsEndpointUpdateBody struct {
	Name                          *string         `json:"name,omitempty"`
	Description                   *string         `json:"description,omitempty"`
	Metadata                      *map[string]any `json:"metadata,omitempty"`
	CanAttachCustomToolFilters    *bool           `json:"can_attach_custom_tool_filters,omitempty"`
	CanAttachCustomProviderConfig *bool           `json:"can_attach_custom_provider_config,omitempty"`
	CanOverrideToolFilters        *bool           `json:"can_override_tool_filters,omitempty"`
}

// List returns a paginated list of integrations.
func (e *IntegrationsEndpoint) List(params *IntegrationsEndpointListParams) (*integrations.IntegrationsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"integrations"},
		Query: query,
	}
	var result integrations.IntegrationsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific integration.
func (e *IntegrationsEndpoint) Get(integrationId string) (*integrations.IntegrationsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integrations", integrationId},
	}
	var result integrations.IntegrationsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new integration.
func (e *IntegrationsEndpoint) Create(body *IntegrationsEndpointCreateBody) (*integrations.IntegrationsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integrations"},
		Body: body,
	}
	var result integrations.IntegrationsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific integration.
func (e *IntegrationsEndpoint) Update(integrationId string, body *IntegrationsEndpointUpdateBody) (*integrations.IntegrationsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integrations", integrationId},
		Body: body,
	}
	var result integrations.IntegrationsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a specific integration.
func (e *IntegrationsEndpoint) Delete(integrationId string) (*integrations.IntegrationsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integrations", integrationId},
	}
	var result integrations.IntegrationsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
