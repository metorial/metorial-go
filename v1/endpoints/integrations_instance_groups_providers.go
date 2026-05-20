package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrations/instancegroups/providers"
)

// IntegrationsInstanceGroupsProvidersEndpoint provides access to integration instance group providers define the effective routed provider set for an integration instance group.
type IntegrationsInstanceGroupsProvidersEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsInstanceGroupsProvidersEndpoint creates a new IntegrationsInstanceGroupsProvidersEndpoint.
func NewIntegrationsInstanceGroupsProvidersEndpoint(client *endpoint.Client) *IntegrationsInstanceGroupsProvidersEndpoint {
	return &IntegrationsInstanceGroupsProvidersEndpoint{client: client}
}

// IntegrationsInstanceGroupsProvidersEndpointListParams contains optional query parameters for List.
type IntegrationsInstanceGroupsProvidersEndpointListParams struct {
	Limit                         *float64 `json:"limit,omitempty"`
	After                         *string  `json:"after,omitempty"`
	Before                        *string  `json:"before,omitempty"`
	Cursor                        *string  `json:"cursor,omitempty"`
	Order                         *string  `json:"order,omitempty"`
	Status                        *any     `json:"status,omitempty"`
	Id                            *any     `json:"id,omitempty"`
	IntegrationInstanceGroupId    *any     `json:"integration_instance_group_id,omitempty"`
	IntegrationId                 *any     `json:"integration_id,omitempty"`
	IntegrationInstanceId         *any     `json:"integration_instance_id,omitempty"`
	IntegrationInstanceProviderId *any     `json:"integration_instance_provider_id,omitempty"`
	ProviderId                    *any     `json:"provider_id,omitempty"`
	IntegrationProviderId         *any     `json:"integration_provider_id,omitempty"`
	ProviderDeploymentId          *any     `json:"provider_deployment_id,omitempty"`
	ProviderConfigId              *any     `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId          *any     `json:"provider_auth_config_id,omitempty"`
	SessionTemplateId             *any     `json:"session_template_id,omitempty"`
	// CreatedAt - Filter integration instance group provider creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration instance group provider last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsInstanceGroupsProvidersEndpointSetBody contains the request body for Set.
type IntegrationsInstanceGroupsProvidersEndpointSetBody struct {
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

// Get retrieves a specific integration instance group provider.
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

// Set creates or updates the effective integration instance group provider materialization.
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

// Delete archives a specific integration instance group provider.
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
