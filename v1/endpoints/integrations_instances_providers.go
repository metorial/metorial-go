package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrations/instances/providers"
)

// IntegrationsInstancesProvidersEndpoint provides access to integration instance providers resolve the effective per-instance provider materialization for an integration.
type IntegrationsInstancesProvidersEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsInstancesProvidersEndpoint creates a new IntegrationsInstancesProvidersEndpoint.
func NewIntegrationsInstancesProvidersEndpoint(client *endpoint.Client) *IntegrationsInstancesProvidersEndpoint {
	return &IntegrationsInstancesProvidersEndpoint{client: client}
}

// IntegrationsInstancesProvidersEndpointListParams contains optional query parameters for List.
type IntegrationsInstancesProvidersEndpointListParams struct {
	Limit                 *float64 `json:"limit,omitempty"`
	After                 *string  `json:"after,omitempty"`
	Before                *string  `json:"before,omitempty"`
	Cursor                *string  `json:"cursor,omitempty"`
	Order                 *string  `json:"order,omitempty"`
	Search                *string  `json:"search,omitempty"`
	Status                *any     `json:"status,omitempty"`
	Id                    *any     `json:"id,omitempty"`
	IntegrationId         *any     `json:"integration_id,omitempty"`
	IntegrationInstanceId *any     `json:"integration_instance_id,omitempty"`
	ProviderId            *any     `json:"provider_id,omitempty"`
	IntegrationProviderId *any     `json:"integration_provider_id,omitempty"`
	ProviderDeploymentId  *any     `json:"provider_deployment_id,omitempty"`
	ProviderConfigId      *any     `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId  *any     `json:"provider_auth_config_id,omitempty"`
	SessionTemplateId     *any     `json:"session_template_id,omitempty"`
	// CreatedAt - Filter integration instance provider creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration instance provider last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsInstancesProvidersEndpointSetBody contains the request body for Set.
type IntegrationsInstancesProvidersEndpointSetBody struct {
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	ProviderConfigId     *string `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	ToolFilters          *any    `json:"tool_filters,omitempty"`
	IsOverrideToolFilter *bool   `json:"is_override_tool_filter,omitempty"`
}

// List returns a paginated list of integration instance providers.
func (e *IntegrationsInstancesProvidersEndpoint) List(params *IntegrationsInstancesProvidersEndpointListParams) (*providers.IntegrationsInstancesProvidersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"integration-instance-providers"},
		Query: query,
	}
	var result providers.IntegrationsInstancesProvidersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific integration instance provider.
func (e *IntegrationsInstancesProvidersEndpoint) Get(integrationInstanceProviderId string) (*providers.IntegrationsInstancesProvidersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instance-providers", integrationInstanceProviderId},
	}
	var result providers.IntegrationsInstancesProvidersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Set creates or updates the effective integration instance provider materialization.
func (e *IntegrationsInstancesProvidersEndpoint) Set(integrationInstanceId string, providerId string, body *IntegrationsInstancesProvidersEndpointSetBody) (*providers.IntegrationsInstancesProvidersSetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instances", integrationInstanceId, "providers", providerId},
		Body: body,
	}
	var result providers.IntegrationsInstancesProvidersSetOutput
	if err := e.client.Put(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
