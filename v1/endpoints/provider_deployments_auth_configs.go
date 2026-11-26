package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerdeployments/authconfigs"
)

// ProviderDeploymentsAuthConfigsEndpoint provides access to a provider auth config is a user's authenticated connection to a provider. It is created when a user completes OAuth or enters an API token.
type ProviderDeploymentsAuthConfigsEndpoint struct {
	client *endpoint.Client
}

// NewProviderDeploymentsAuthConfigsEndpoint creates a new ProviderDeploymentsAuthConfigsEndpoint.
func NewProviderDeploymentsAuthConfigsEndpoint(client *endpoint.Client) *ProviderDeploymentsAuthConfigsEndpoint {
	return &ProviderDeploymentsAuthConfigsEndpoint{client: client}
}

// ProviderDeploymentsAuthConfigsEndpointListParams contains optional query parameters for List.
type ProviderDeploymentsAuthConfigsEndpointListParams struct {
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
	// Ownership - Which resources to return under the resource ownership rules: shared (resources without an owner), usable (shared resources plus owned resources the caller may use), or all (everything the caller can see). Defaults to shared when ownership is enforced for the project, and all otherwise.
	Ownership *string `json:"ownership,omitempty"`
	// Status - Filter by status (active, archived)
	Status *any `json:"status,omitempty"`
	// Id - Filter by auth config ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// AvailableForUse - Only return auth configs that are not owned by another integration instance provider.
	AvailableForUse *bool `json:"available_for_use,omitempty"`
	// AvailableForProviderDeploymentId - Only return auth configs that are not locked to a different provider deployment.
	AvailableForProviderDeploymentId *string `json:"available_for_provider_deployment_id,omitempty"`
	// ProviderAuthCredentialsId - Filter by auth credentials ID(s)
	ProviderAuthCredentialsId *any `json:"provider_auth_credentials_id,omitempty"`
	// ProviderAuthMethodId - Filter by auth method ID(s)
	ProviderAuthMethodId *any `json:"provider_auth_method_id,omitempty"`
	// ActorId - Filter by actor ID(s)
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerId - Filter by consumer ID(s)
	ConsumerId *any `json:"consumer_id,omitempty"`
	// IdentityId - Filter by identity ID(s)
	IdentityId *any `json:"identity_id,omitempty"`
	// IdentityCredentialId - Filter by identity credential ID(s)
	IdentityCredentialId *any `json:"identity_credential_id,omitempty"`
	// Adapter - Only return auth configs whose auth method supports this adapter
	Adapter *string `json:"adapter,omitempty"`
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider auth config creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider auth config last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// ProviderDeploymentsAuthConfigsEndpointCreateBody contains the request body for Create.
type ProviderDeploymentsAuthConfigsEndpointCreateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider auth config. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// ProviderAuthMethodId - The authentication method this config uses (e.g., OAuth, API key)
	ProviderAuthMethodId string `json:"provider_auth_method_id"`
	// ProviderDeploymentId - The provider deployment this auth config is associated with (if applicable)
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// Value - Authentication config payload
	Value map[string]any `json:"value"`
}

// ProviderDeploymentsAuthConfigsEndpointUpdateBody contains the request body for Update.
type ProviderDeploymentsAuthConfigsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider auth config. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// List returns a paginated list of provider auth configs.
func (e *ProviderDeploymentsAuthConfigsEndpoint) List(params *ProviderDeploymentsAuthConfigsEndpointListParams) (*authconfigs.ProviderDeploymentsAuthConfigsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-auth-configs"},
		Query: query,
	}
	var result authconfigs.ProviderDeploymentsAuthConfigsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider auth config by ID.
func (e *ProviderDeploymentsAuthConfigsEndpoint) Get(providerAuthConfigId string) (*authconfigs.ProviderDeploymentsAuthConfigsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-auth-configs", providerAuthConfigId},
	}
	var result authconfigs.ProviderDeploymentsAuthConfigsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a provider auth config.
func (e *ProviderDeploymentsAuthConfigsEndpoint) Create(body *ProviderDeploymentsAuthConfigsEndpointCreateBody) (*authconfigs.ProviderDeploymentsAuthConfigsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-auth-configs"},
		Body: body,
	}
	var result authconfigs.ProviderDeploymentsAuthConfigsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific provider auth config.
func (e *ProviderDeploymentsAuthConfigsEndpoint) Update(providerAuthConfigId string, body *ProviderDeploymentsAuthConfigsEndpointUpdateBody) (*authconfigs.ProviderDeploymentsAuthConfigsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-auth-configs", providerAuthConfigId},
		Body: body,
	}
	var result authconfigs.ProviderDeploymentsAuthConfigsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a provider auth config. It stays listable with `status=archived`.
func (e *ProviderDeploymentsAuthConfigsEndpoint) Delete(providerAuthConfigId string) (*authconfigs.ProviderDeploymentsAuthConfigsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-auth-configs", providerAuthConfigId},
	}
	var result authconfigs.ProviderDeploymentsAuthConfigsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
