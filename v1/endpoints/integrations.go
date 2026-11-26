package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/integrations"
)

// IntegrationsEndpoint provides access to an integration is a reusable set of providers with shared configuration. Integration instances connect it for a specific actor or identity.
type IntegrationsEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsEndpoint creates a new IntegrationsEndpoint.
func NewIntegrationsEndpoint(client *endpoint.Client) *IntegrationsEndpoint {
	return &IntegrationsEndpoint{client: client}
}

// IntegrationsEndpointListParams contains optional query parameters for List.
type IntegrationsEndpointListParams struct {
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
	// Search - Search integrations by name, description, ID, or slug.
	Search *string `json:"search,omitempty"`
	// Status - Only return integrations with one of these statuses. `deleted` integrations are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return integrations with one of these IDs.
	Id *any `json:"id,omitempty"`
	// ProviderId - Only return integrations using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationProviderId - Only return integrations using one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// CreatedAt - Only return results whose integration creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsEndpointCreateBody contains the request body for Create.
type IntegrationsEndpointCreateBody struct {
	// Name - Display name for the integration.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// UseIntegrationNameInToolNames - Names session providers after the integration instead of the provider. `null` uses the default.
	UseIntegrationNameInToolNames *nullable.Nullable[bool] `json:"use_integration_name_in_tool_names,omitempty"`
	// CanAttachCustomToolFilters - Lets integration instances attach their own tool filters.
	CanAttachCustomToolFilters *bool `json:"can_attach_custom_tool_filters,omitempty"`
	// CanAttachCustomProviderConfig - Lets integration instances use their own provider configs.
	CanAttachCustomProviderConfig *bool `json:"can_attach_custom_provider_config,omitempty"`
	// CanOverrideToolFilters - Lets integration instances ignore the integration tool filters.
	CanOverrideToolFilters *bool `json:"can_override_tool_filters,omitempty"`
	// EnableCallbackTools - Makes callback event tools available for this integration.
	EnableCallbackTools *bool `json:"enable_callback_tools,omitempty"`
}

// IntegrationsEndpointUpdateBody contains the request body for Update.
type IntegrationsEndpointUpdateBody struct {
	// Name - Display name for the integration.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
	// UseIntegrationNameInToolNames - Names session providers after the integration instead of the provider. `null` uses the default.
	UseIntegrationNameInToolNames *nullable.Nullable[bool] `json:"use_integration_name_in_tool_names,omitempty"`
	// CanAttachCustomToolFilters - Lets integration instances attach their own tool filters.
	CanAttachCustomToolFilters *bool `json:"can_attach_custom_tool_filters,omitempty"`
	// CanAttachCustomProviderConfig - Lets integration instances use their own provider configs.
	CanAttachCustomProviderConfig *bool `json:"can_attach_custom_provider_config,omitempty"`
	// CanOverrideToolFilters - Lets integration instances ignore the integration tool filters.
	CanOverrideToolFilters *bool `json:"can_override_tool_filters,omitempty"`
	// EnableCallbackTools - Makes callback event tools available for this integration.
	EnableCallbackTools *bool `json:"enable_callback_tools,omitempty"`
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

// Get retrieves an integration.
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

// Create creates an integration.
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

// Update updates an integration. Omitted fields keep their current value.
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

// Delete archives an integration. It stays retrievable by ID and listable with `status=archived`.
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
