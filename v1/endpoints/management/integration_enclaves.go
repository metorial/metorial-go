package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrationenclaves"
)

// IntegrationEnclavesEndpoint provides access to read integration enclave records for provider deployments in an instance.
type IntegrationEnclavesEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationEnclavesEndpoint creates a new IntegrationEnclavesEndpoint.
func NewIntegrationEnclavesEndpoint(client *endpoint.Client) *IntegrationEnclavesEndpoint {
	return &IntegrationEnclavesEndpoint{client: client}
}

// IntegrationEnclavesEndpointListParams contains optional query parameters for List.
type IntegrationEnclavesEndpointListParams struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	Slug                 *any     `json:"slug,omitempty"`
	InternetGatewayId    *any     `json:"internet_gateway_id,omitempty"`
	ProviderDeploymentId *any     `json:"provider_deployment_id,omitempty"`
	ProviderId           *any     `json:"provider_id,omitempty"`
	FirewallId           *any     `json:"firewall_id,omitempty"`
	// CreatedAt - Filter integration enclave creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// List returns a paginated list of integration enclaves.
func (e *IntegrationEnclavesEndpoint) List(instanceId string, params *IntegrationEnclavesEndpointListParams) (*integrationenclaves.IntegrationEnclavesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "integration-enclaves"},
		Query: query,
	}
	var result integrationenclaves.IntegrationEnclavesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific integration enclave by ID.
func (e *IntegrationEnclavesEndpoint) Get(instanceId string, integrationEnclaveId string) (*integrationenclaves.IntegrationEnclavesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "integration-enclaves", integrationEnclaveId},
	}
	var result integrationenclaves.IntegrationEnclavesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
