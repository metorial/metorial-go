package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrationenclaves"
)

// IntegrationEnclavesEndpoint provides access to an integration enclave is the isolated network a provider deployment runs in. It connects the deployment to its internet gateway and firewalls.
type IntegrationEnclavesEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationEnclavesEndpoint creates a new IntegrationEnclavesEndpoint.
func NewIntegrationEnclavesEndpoint(client *endpoint.Client) *IntegrationEnclavesEndpoint {
	return &IntegrationEnclavesEndpoint{client: client}
}

// IntegrationEnclavesEndpointListParams contains optional query parameters for List.
type IntegrationEnclavesEndpointListParams struct {
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
	// Id - Only return integration enclaves with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Slug - Only return integration enclaves with one of these slugs.
	Slug *any `json:"slug,omitempty"`
	// InternetGatewayId - Only return integration enclaves using one of these internet gateways.
	InternetGatewayId *any `json:"internet_gateway_id,omitempty"`
	// ProviderDeploymentId - Only return integration enclaves running one of these provider deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderId - Only return integration enclaves running one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// FirewallId - Only return integration enclaves protected by one of these firewalls.
	FirewallId *any `json:"firewall_id,omitempty"`
	// CreatedAt - Only return results whose integration enclave creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// List returns a paginated list of integration enclaves in the instance.
func (e *IntegrationEnclavesEndpoint) List(params *IntegrationEnclavesEndpointListParams) (*integrationenclaves.IntegrationEnclavesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"integration-enclaves"},
		Query: query,
	}
	var result integrationenclaves.IntegrationEnclavesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an integration enclave.
func (e *IntegrationEnclavesEndpoint) Get(integrationEnclaveId string) (*integrationenclaves.IntegrationEnclavesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-enclaves", integrationEnclaveId},
	}
	var result integrationenclaves.IntegrationEnclavesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
