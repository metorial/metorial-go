package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/integrations/providers"
)

// IntegrationsProvidersEndpoint provides access to an integration provider is one provider in an integration, with the deployment, configuration, credentials, and tool filters it runs with.
type IntegrationsProvidersEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsProvidersEndpoint creates a new IntegrationsProvidersEndpoint.
func NewIntegrationsProvidersEndpoint(client *endpoint.Client) *IntegrationsProvidersEndpoint {
	return &IntegrationsProvidersEndpoint{client: client}
}

// IntegrationsProvidersEndpointListParams contains optional query parameters for List.
type IntegrationsProvidersEndpointListParams struct {
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
	// Search - Search integration providers by name, description, ID, or provider.
	Search *string `json:"search,omitempty"`
	// Status - Only return integration providers with one of these statuses. `deleted` providers are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return integration providers with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return integration providers for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// ProviderId - Only return integration providers using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Only return integration providers running on one of these deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderAuthMethodId - Only return integration providers using one of these auth methods.
	ProviderAuthMethodId *any `json:"provider_auth_method_id,omitempty"`
	// ProviderAuthCredentialsId - Only return integration providers using one of these credentials.
	ProviderAuthCredentialsId *any `json:"provider_auth_credentials_id,omitempty"`
	// ProviderConfigId - Only return integration providers using one of these configs.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// CreatedAt - Only return results whose integration provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsProvidersEndpointCreateBody contains the request body for Create.
type IntegrationsProvidersEndpointCreateBody struct {
	// IntegrationId - Integration the provider belongs to.
	IntegrationId string `json:"integration_id"`
	// ProviderId - Provider to add to the integration.
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Deployment the provider runs with.
	ProviderDeploymentId string `json:"provider_deployment_id"`
	// ProviderAuthMethodId - Auth method the provider runs with. Omitted leaves it unset.
	ProviderAuthMethodId *string `json:"provider_auth_method_id,omitempty"`
	// ProviderAuthCredentialsId - Credentials the provider runs with. Omitted leaves them unset.
	ProviderAuthCredentialsId *string `json:"provider_auth_credentials_id,omitempty"`
	// ProviderConfigId - Provider config the provider runs with. Omitted leaves it unset.
	ProviderConfigId *string `json:"provider_config_id,omitempty"`
	// Name - Display name for the integration provider. Defaults to the provider name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration provider.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// Callbacks - Provider callback state for this integration provider
	Callbacks *map[string]any `json:"callbacks,omitempty"`
}

// IntegrationsProvidersEndpointUpdateBody contains the request body for Update.
type IntegrationsProvidersEndpointUpdateBody struct {
	// ProviderDeploymentId - Deployment the provider runs with. Omitted keeps the current one.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderAuthMethodId - Auth method the provider runs with. `null` clears it. Omitted keeps the current one.
	ProviderAuthMethodId *nullable.Nullable[string] `json:"provider_auth_method_id,omitempty"`
	// ProviderAuthCredentialsId - Credentials the provider runs with. `null` clears them. Omitted keeps the current ones.
	ProviderAuthCredentialsId *nullable.Nullable[string] `json:"provider_auth_credentials_id,omitempty"`
	// ProviderConfigId - Provider config the provider runs with. `null` clears it. Omitted keeps the current one.
	ProviderConfigId *nullable.Nullable[string] `json:"provider_config_id,omitempty"`
	// Name - Display name for the integration provider. Omitted keeps the current one.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration provider. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// Callbacks - Provider callback state for this integration provider
	Callbacks *map[string]any `json:"callbacks,omitempty"`
}

// List returns a paginated list of integration providers.
func (e *IntegrationsProvidersEndpoint) List(instanceId string, params *IntegrationsProvidersEndpointListParams) (*providers.IntegrationsProvidersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "integration-providers"},
		Query: query,
	}
	var result providers.IntegrationsProvidersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an integration provider.
func (e *IntegrationsProvidersEndpoint) Get(instanceId string, integrationProviderId string) (*providers.IntegrationsProvidersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "integration-providers", integrationProviderId},
	}
	var result providers.IntegrationsProvidersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an integration provider.
func (e *IntegrationsProvidersEndpoint) Create(instanceId string, body *IntegrationsProvidersEndpointCreateBody) (*providers.IntegrationsProvidersCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "integration-providers"},
		Body: body,
	}
	var result providers.IntegrationsProvidersCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an integration provider. Omitted fields keep their current value.
func (e *IntegrationsProvidersEndpoint) Update(instanceId string, integrationProviderId string, body *IntegrationsProvidersEndpointUpdateBody) (*providers.IntegrationsProvidersUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "integration-providers", integrationProviderId},
		Body: body,
	}
	var result providers.IntegrationsProvidersUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives an integration provider. It stays retrievable by ID and listable with `status=archived`.
func (e *IntegrationsProvidersEndpoint) Delete(instanceId string, integrationProviderId string) (*providers.IntegrationsProvidersDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "integration-providers", integrationProviderId},
	}
	var result providers.IntegrationsProvidersDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
