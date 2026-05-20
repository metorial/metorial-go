package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrations/providers"
)

// IntegrationsProvidersEndpoint provides access to integration providers define the shared provider-level contract for a given integration.
type IntegrationsProvidersEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsProvidersEndpoint creates a new IntegrationsProvidersEndpoint.
func NewIntegrationsProvidersEndpoint(client *endpoint.Client) *IntegrationsProvidersEndpoint {
	return &IntegrationsProvidersEndpoint{client: client}
}

// IntegrationsProvidersEndpointListParams contains optional query parameters for List.
type IntegrationsProvidersEndpointListParams struct {
	Limit                     *float64 `json:"limit,omitempty"`
	After                     *string  `json:"after,omitempty"`
	Before                    *string  `json:"before,omitempty"`
	Cursor                    *string  `json:"cursor,omitempty"`
	Order                     *string  `json:"order,omitempty"`
	Search                    *string  `json:"search,omitempty"`
	Status                    *any     `json:"status,omitempty"`
	Id                        *any     `json:"id,omitempty"`
	IntegrationId             *any     `json:"integration_id,omitempty"`
	ProviderId                *any     `json:"provider_id,omitempty"`
	ProviderDeploymentId      *any     `json:"provider_deployment_id,omitempty"`
	ProviderAuthMethodId      *any     `json:"provider_auth_method_id,omitempty"`
	ProviderAuthCredentialsId *any     `json:"provider_auth_credentials_id,omitempty"`
	ProviderConfigId          *any     `json:"provider_config_id,omitempty"`
	// CreatedAt - Filter integration provider creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration provider last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsProvidersEndpointCreateBody contains the request body for Create.
type IntegrationsProvidersEndpointCreateBody struct {
	IntegrationId             string          `json:"integration_id"`
	ProviderId                string          `json:"provider_id"`
	ProviderDeploymentId      string          `json:"provider_deployment_id"`
	ProviderAuthMethodId      *string         `json:"provider_auth_method_id,omitempty"`
	ProviderAuthCredentialsId *string         `json:"provider_auth_credentials_id,omitempty"`
	ProviderConfigId          *string         `json:"provider_config_id,omitempty"`
	Name                      *string         `json:"name,omitempty"`
	Description               *string         `json:"description,omitempty"`
	Metadata                  *map[string]any `json:"metadata,omitempty"`
	ToolFilters               *any            `json:"tool_filters,omitempty"`
}

// IntegrationsProvidersEndpointUpdateBody contains the request body for Update.
type IntegrationsProvidersEndpointUpdateBody struct {
	ProviderDeploymentId      *string         `json:"provider_deployment_id,omitempty"`
	ProviderAuthMethodId      *string         `json:"provider_auth_method_id,omitempty"`
	ProviderAuthCredentialsId *string         `json:"provider_auth_credentials_id,omitempty"`
	ProviderConfigId          *string         `json:"provider_config_id,omitempty"`
	Name                      *string         `json:"name,omitempty"`
	Description               *string         `json:"description,omitempty"`
	Metadata                  *map[string]any `json:"metadata,omitempty"`
	ToolFilters               *any            `json:"tool_filters,omitempty"`
}

// List returns a paginated list of integration providers.
func (e *IntegrationsProvidersEndpoint) List(params *IntegrationsProvidersEndpointListParams) (*providers.IntegrationsProvidersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"integration-providers"},
		Query: query,
	}
	var result providers.IntegrationsProvidersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific integration provider.
func (e *IntegrationsProvidersEndpoint) Get(integrationProviderId string) (*providers.IntegrationsProvidersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-providers", integrationProviderId},
	}
	var result providers.IntegrationsProvidersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new integration provider.
func (e *IntegrationsProvidersEndpoint) Create(body *IntegrationsProvidersEndpointCreateBody) (*providers.IntegrationsProvidersCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-providers"},
		Body: body,
	}
	var result providers.IntegrationsProvidersCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific integration provider.
func (e *IntegrationsProvidersEndpoint) Update(integrationProviderId string, body *IntegrationsProvidersEndpointUpdateBody) (*providers.IntegrationsProvidersUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-providers", integrationProviderId},
		Body: body,
	}
	var result providers.IntegrationsProvidersUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a specific integration provider.
func (e *IntegrationsProvidersEndpoint) Delete(integrationProviderId string) (*providers.IntegrationsProvidersDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-providers", integrationProviderId},
	}
	var result providers.IntegrationsProvidersDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
