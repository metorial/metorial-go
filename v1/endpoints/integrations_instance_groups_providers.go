package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrations/instancegroups/providers"
)

// IntegrationsInstanceGroupsProvidersEndpoint provides access to an integration instance group provider links one integration instance provider into a group, with optional tool filters.
type IntegrationsInstanceGroupsProvidersEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsInstanceGroupsProvidersEndpoint creates a new IntegrationsInstanceGroupsProvidersEndpoint.
func NewIntegrationsInstanceGroupsProvidersEndpoint(client *endpoint.Client) *IntegrationsInstanceGroupsProvidersEndpoint {
	return &IntegrationsInstanceGroupsProvidersEndpoint{client: client}
}

// IntegrationsInstanceGroupsProvidersEndpointListParams contains optional query parameters for List.
type IntegrationsInstanceGroupsProvidersEndpointListParams struct {
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
	// Status - Only return group providers with one of these statuses. `deleted` providers are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return group providers with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationInstanceGroupId - Only return group providers for one of these groups.
	IntegrationInstanceGroupId *any `json:"integration_instance_group_id,omitempty"`
	// IntegrationId - Only return group providers for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationInstanceId - Only return group providers for one of these integration instances.
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// IntegrationInstanceProviderId - Only return group providers linking one of these integration instance providers.
	IntegrationInstanceProviderId *any `json:"integration_instance_provider_id,omitempty"`
	// ProviderId - Only return group providers using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationProviderId - Only return group providers using one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// ProviderDeploymentId - Only return group providers running on one of these deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Only return group providers using one of these configs.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Only return group providers using one of these credentials.
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// SessionTemplateId - Only return group providers using one of these session templates.
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// CreatedAt - Only return results whose integration instance group provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration instance group provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsInstanceGroupsProvidersEndpointSetBody contains the request body for Set.
type IntegrationsInstanceGroupsProvidersEndpointSetBody struct {
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// List returns a paginated list of integration instance group providers.
func (e *IntegrationsInstanceGroupsProvidersEndpoint) List(params *IntegrationsInstanceGroupsProvidersEndpointListParams) (*providers.IntegrationsInstanceGroupsProvidersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"integration-instance-group-providers"},
		Query: query,
	}
	var result providers.IntegrationsInstanceGroupsProvidersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an integration instance group provider.
func (e *IntegrationsInstanceGroupsProvidersEndpoint) Get(integrationInstanceGroupProviderId string) (*providers.IntegrationsInstanceGroupsProvidersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instance-group-providers", integrationInstanceGroupProviderId},
	}
	var result providers.IntegrationsInstanceGroupsProvidersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Set links an existing integration instance provider into the group, with optional tool filters. Creates the link when none exists, otherwise updates its tool filters.
func (e *IntegrationsInstanceGroupsProvidersEndpoint) Set(integrationInstanceGroupId string, integrationInstanceProviderId string, body *IntegrationsInstanceGroupsProvidersEndpointSetBody) (*providers.IntegrationsInstanceGroupsProvidersSetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instance-groups", integrationInstanceGroupId, "providers", integrationInstanceProviderId},
		Body: body,
	}
	var result providers.IntegrationsInstanceGroupsProvidersSetOutput
	if err := e.client.Put(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives an integration instance group provider. It stays retrievable by ID and listable with `status=archived`.
func (e *IntegrationsInstanceGroupsProvidersEndpoint) Delete(integrationInstanceGroupProviderId string) (*providers.IntegrationsInstanceGroupsProvidersDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instance-group-providers", integrationInstanceGroupProviderId},
	}
	var result providers.IntegrationsInstanceGroupsProvidersDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
