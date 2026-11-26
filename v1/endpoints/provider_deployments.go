package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/providerdeployments"
)

// ProviderDeploymentsEndpoint provides access to a provider deployment is a provider set up for your instance, with its own configuration and credentials. It follows the provider's current version unless you lock it to one.
type ProviderDeploymentsEndpoint struct {
	client *endpoint.Client
}

// NewProviderDeploymentsEndpoint creates a new ProviderDeploymentsEndpoint.
func NewProviderDeploymentsEndpoint(client *endpoint.Client) *ProviderDeploymentsEndpoint {
	return &ProviderDeploymentsEndpoint{client: client}
}

// ProviderDeploymentsEndpointListParams contains optional query parameters for List.
type ProviderDeploymentsEndpointListParams struct {
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
	// Id - Filter by deployment ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderVersionId - Filter by version ID(s)
	ProviderVersionId *any `json:"provider_version_id,omitempty"`
	// ActorId - Filter by actor ID(s)
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerId - Filter by consumer ID(s)
	ConsumerId *any `json:"consumer_id,omitempty"`
	// IdentityId - Filter by identity ID(s)
	IdentityId *any `json:"identity_id,omitempty"`
	// IdentityCredentialId - Filter by identity credential ID(s)
	IdentityCredentialId *any `json:"identity_credential_id,omitempty"`
	// Status - Filter by status (active, archived)
	Status *any `json:"status,omitempty"`
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider deployment creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider deployment last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// ProviderDeploymentsEndpointCreateBody contains the request body for Create.
type ProviderDeploymentsEndpointCreateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider deployment. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// ProviderId - The provider to deploy
	ProviderId string `json:"provider_id"`
	// LockedProviderVersionId - Pin this deployment to a specific provider version
	LockedProviderVersionId *string `json:"locked_provider_version_id,omitempty"`
	// ProviderConfigId - ID of an existing config to use.
	ProviderConfigId *string `json:"provider_config_id,omitempty"`
	ProviderConfig   *any    `json:"provider_config,omitempty"`
}

// ProviderDeploymentsEndpointUpdateBody contains the request body for Update.
type ProviderDeploymentsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider deployment. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// LockedProviderVersionId - Pin this deployment to a specific provider version, or null to follow latest
	LockedProviderVersionId *nullable.Nullable[string] `json:"locked_provider_version_id,omitempty"`
}

// List returns a paginated list of provider deployments.
func (e *ProviderDeploymentsEndpoint) List(params *ProviderDeploymentsEndpointListParams) (*providerdeployments.ProviderDeploymentsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-deployments"},
		Query: query,
	}
	var result providerdeployments.ProviderDeploymentsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider deployment by ID.
func (e *ProviderDeploymentsEndpoint) Get(providerDeploymentId string) (*providerdeployments.ProviderDeploymentsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-deployments", providerDeploymentId},
	}
	var result providerdeployments.ProviderDeploymentsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a provider deployment.
func (e *ProviderDeploymentsEndpoint) Create(body *ProviderDeploymentsEndpointCreateBody) (*providerdeployments.ProviderDeploymentsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-deployments"},
		Body: body,
	}
	var result providerdeployments.ProviderDeploymentsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific provider deployment.
func (e *ProviderDeploymentsEndpoint) Update(providerDeploymentId string, body *ProviderDeploymentsEndpointUpdateBody) (*providerdeployments.ProviderDeploymentsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-deployments", providerDeploymentId},
		Body: body,
	}
	var result providerdeployments.ProviderDeploymentsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a provider deployment. It stays listable with `status=archived`.
func (e *ProviderDeploymentsEndpoint) Delete(providerDeploymentId string) (*providerdeployments.ProviderDeploymentsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-deployments", providerDeploymentId},
	}
	var result providerdeployments.ProviderDeploymentsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
