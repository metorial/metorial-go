package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/enclaves"
)

// EnclavesEndpoint provides access to read enclave records for provider deployments in an instance.
type EnclavesEndpoint struct {
	client *endpoint.Client
}

// NewEnclavesEndpoint creates a new EnclavesEndpoint.
func NewEnclavesEndpoint(client *endpoint.Client) *EnclavesEndpoint {
	return &EnclavesEndpoint{client: client}
}

// EnclavesEndpointListParams contains optional query parameters for List.
type EnclavesEndpointListParams struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	Slug                 *any     `json:"slug,omitempty"`
	NetworkId            *any     `json:"network_id,omitempty"`
	ProviderDeploymentId *any     `json:"provider_deployment_id,omitempty"`
	ProviderId           *any     `json:"provider_id,omitempty"`
	FirewallId           *any     `json:"firewall_id,omitempty"`
	// CreatedAt - Filter enclave creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// List returns a paginated list of enclaves.
func (e *EnclavesEndpoint) List(params *EnclavesEndpointListParams) (*enclaves.EnclavesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"enclaves"},
		Query: query,
	}
	var result enclaves.EnclavesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific enclave by ID.
func (e *EnclavesEndpoint) Get(enclaveId string) (*enclaves.EnclavesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"enclaves", enclaveId},
	}
	var result enclaves.EnclavesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
