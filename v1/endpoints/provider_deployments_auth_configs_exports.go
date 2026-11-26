package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerdeployments/authconfigs/exports"
)

// ProviderDeploymentsAuthConfigsExportsEndpoint provides access to a provider auth export is a copy of authentication credentials taken out of Metorial for use in another system.
type ProviderDeploymentsAuthConfigsExportsEndpoint struct {
	client *endpoint.Client
}

// NewProviderDeploymentsAuthConfigsExportsEndpoint creates a new ProviderDeploymentsAuthConfigsExportsEndpoint.
func NewProviderDeploymentsAuthConfigsExportsEndpoint(client *endpoint.Client) *ProviderDeploymentsAuthConfigsExportsEndpoint {
	return &ProviderDeploymentsAuthConfigsExportsEndpoint{client: client}
}

// ProviderDeploymentsAuthConfigsExportsEndpointListParams contains optional query parameters for List.
type ProviderDeploymentsAuthConfigsExportsEndpointListParams struct {
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
	// Id - Filter by export ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderAuthCredentialsId - Filter by auth credentials ID(s)
	ProviderAuthCredentialsId *any `json:"provider_auth_credentials_id,omitempty"`
	// ProviderAuthConfigId - Filter by auth config ID(s)
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Only return results whose provider auth export creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider auth export last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// ProviderDeploymentsAuthConfigsExportsEndpointCreateBody contains the request body for Create.
type ProviderDeploymentsAuthConfigsExportsEndpointCreateBody struct {
	// ProviderAuthConfigId - Auth config to export credentials from.
	ProviderAuthConfigId string `json:"provider_auth_config_id"`
	// Note - Why the credentials are exported.
	Note string `json:"note"`
	// Metadata - Custom data to identify or organize the provider auth export. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of provider auth exports.
func (e *ProviderDeploymentsAuthConfigsExportsEndpoint) List(params *ProviderDeploymentsAuthConfigsExportsEndpointListParams) (*exports.ProviderDeploymentsAuthConfigsExportsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-auth-config-exports"},
		Query: query,
	}
	var result exports.ProviderDeploymentsAuthConfigsExportsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider auth export by ID.
func (e *ProviderDeploymentsAuthConfigsExportsEndpoint) Get(providerAuthExportId string) (*exports.ProviderDeploymentsAuthConfigsExportsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-auth-config-exports", providerAuthExportId},
	}
	var result exports.ProviderDeploymentsAuthConfigsExportsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create exports the credentials of a provider auth config for use outside Metorial.
func (e *ProviderDeploymentsAuthConfigsExportsEndpoint) Create(body *ProviderDeploymentsAuthConfigsExportsEndpointCreateBody) (*exports.ProviderDeploymentsAuthConfigsExportsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-auth-config-exports"},
		Body: body,
	}
	var result exports.ProviderDeploymentsAuthConfigsExportsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
