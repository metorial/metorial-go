package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerdeployments/configvaults"
)

// ProviderDeploymentsConfigVaultsEndpoint provides access to a provider config vault is a saved, reusable set of configuration values. Configs can reference a vault instead of holding values directly.
type ProviderDeploymentsConfigVaultsEndpoint struct {
	client *endpoint.Client
}

// NewProviderDeploymentsConfigVaultsEndpoint creates a new ProviderDeploymentsConfigVaultsEndpoint.
func NewProviderDeploymentsConfigVaultsEndpoint(client *endpoint.Client) *ProviderDeploymentsConfigVaultsEndpoint {
	return &ProviderDeploymentsConfigVaultsEndpoint{client: client}
}

// ProviderDeploymentsConfigVaultsEndpointListParams contains optional query parameters for List.
type ProviderDeploymentsConfigVaultsEndpointListParams struct {
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
	// Status - Filter by status (active, archived)
	Status *any `json:"status,omitempty"`
	// Id - Filter by config vault ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter by provider config ID(s)
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderConfigVaultId - Filter by config vault ID(s)
	ProviderConfigVaultId *any `json:"provider_config_vault_id,omitempty"`
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider config vault creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider config vault last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// ProviderDeploymentsConfigVaultsEndpointCreateBody contains the request body for Create.
type ProviderDeploymentsConfigVaultsEndpointCreateBody struct {
	// ProviderId - Provider to create the vault for.
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Deployment the vault is locked to.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider config vault. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Value - Secure configuration values to store in the vault
	Value map[string]any `json:"value"`
}

// ProviderDeploymentsConfigVaultsEndpointUpdateBody contains the request body for Update.
type ProviderDeploymentsConfigVaultsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider config vault. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of provider config vaults.
func (e *ProviderDeploymentsConfigVaultsEndpoint) List(params *ProviderDeploymentsConfigVaultsEndpointListParams) (*configvaults.ProviderDeploymentsConfigVaultsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-config-vaults"},
		Query: query,
	}
	var result configvaults.ProviderDeploymentsConfigVaultsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider config vault by ID.
func (e *ProviderDeploymentsConfigVaultsEndpoint) Get(providerConfigVaultId string) (*configvaults.ProviderDeploymentsConfigVaultsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-config-vaults", providerConfigVaultId},
	}
	var result configvaults.ProviderDeploymentsConfigVaultsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new provider config vault.
func (e *ProviderDeploymentsConfigVaultsEndpoint) Create(body *ProviderDeploymentsConfigVaultsEndpointCreateBody) (*configvaults.ProviderDeploymentsConfigVaultsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-config-vaults"},
		Body: body,
	}
	var result configvaults.ProviderDeploymentsConfigVaultsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific provider config vault.
func (e *ProviderDeploymentsConfigVaultsEndpoint) Update(providerConfigVaultId string, body *ProviderDeploymentsConfigVaultsEndpointUpdateBody) (*configvaults.ProviderDeploymentsConfigVaultsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-config-vaults", providerConfigVaultId},
		Body: body,
	}
	var result configvaults.ProviderDeploymentsConfigVaultsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a provider config vault. It stays listable with `status=archived`.
func (e *ProviderDeploymentsConfigVaultsEndpoint) Delete(providerConfigVaultId string) (*configvaults.ProviderDeploymentsConfigVaultsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-config-vaults", providerConfigVaultId},
	}
	var result configvaults.ProviderDeploymentsConfigVaultsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
