package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrations/instances/providers"
)

// IntegrationsInstancesProvidersEndpoint provides access to an integration instance provider binds one provider of an integration instance to the deployment, config, credentials, and tool filters it runs with.
type IntegrationsInstancesProvidersEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsInstancesProvidersEndpoint creates a new IntegrationsInstancesProvidersEndpoint.
func NewIntegrationsInstancesProvidersEndpoint(client *endpoint.Client) *IntegrationsInstancesProvidersEndpoint {
	return &IntegrationsInstancesProvidersEndpoint{client: client}
}

// IntegrationsInstancesProvidersEndpointListParams contains optional query parameters for List.
type IntegrationsInstancesProvidersEndpointListParams struct {
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
	// Search - Search integration instance providers by name, description, or ID.
	Search *string `json:"search,omitempty"`
	// Status - Only return integration instance providers with one of these statuses. `deleted` providers are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return integration instance providers with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return integration instance providers for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationInstanceId - Only return integration instance providers for one of these integration instances.
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// ProviderId - Only return integration instance providers using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationProviderId - Only return integration instance providers using one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// ProviderDeploymentId - Only return integration instance providers running on one of these deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Only return integration instance providers using one of these configs.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Only return integration instance providers using one of these credentials.
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// SessionTemplateId - Only return integration instance providers using one of these session templates.
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// CreatedAt - Only return results whose integration instance provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration instance provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsInstancesProvidersEndpointSetBody contains the request body for Set.
type IntegrationsInstancesProvidersEndpointSetBody struct {
	// ProviderDeploymentId - Deployment to run with. Omitted uses the integration provider’s deployment.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Provider config to run with. Omitted keeps the current one, or the shared config when none exists.
	ProviderConfigId *string `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Credentials to run with. Omitted keeps the current ones.
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter *bool `json:"is_override_tool_filter,omitempty"`
}

// List returns a paginated list of integration instance providers.
func (e *IntegrationsInstancesProvidersEndpoint) List(instanceId string, params *IntegrationsInstancesProvidersEndpointListParams) (*providers.IntegrationsInstancesProvidersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "integration-instance-providers"},
		Query: query,
	}
	var result providers.IntegrationsInstancesProvidersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an integration instance provider.
func (e *IntegrationsInstancesProvidersEndpoint) Get(instanceId string, integrationInstanceProviderId string) (*providers.IntegrationsInstancesProvidersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "integration-instance-providers", integrationInstanceProviderId},
	}
	var result providers.IntegrationsInstancesProvidersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Set sets the deployment, config, credentials, and tool filters one provider of an integration instance runs with. Creates the entry when none exists, otherwise updates it.
func (e *IntegrationsInstancesProvidersEndpoint) Set(instanceId string, integrationInstanceId string, providerId string, body *IntegrationsInstancesProvidersEndpointSetBody) (*providers.IntegrationsInstancesProvidersSetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "integration-instances", integrationInstanceId, "providers", providerId},
		Body: body,
	}
	var result providers.IntegrationsInstancesProvidersSetOutput
	if err := e.client.Put(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
