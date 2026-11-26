package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerdeployments/configs"
)

// ProviderDeploymentsConfigsEndpoint provides access to a provider config is a set of configuration values for a provider. It holds values directly or references a config vault.
type ProviderDeploymentsConfigsEndpoint struct {
	client *endpoint.Client
}

// NewProviderDeploymentsConfigsEndpoint creates a new ProviderDeploymentsConfigsEndpoint.
func NewProviderDeploymentsConfigsEndpoint(client *endpoint.Client) *ProviderDeploymentsConfigsEndpoint {
	return &ProviderDeploymentsConfigsEndpoint{client: client}
}

// ProviderDeploymentsConfigsEndpointListParams contains optional query parameters for List.
type ProviderDeploymentsConfigsEndpointListParams struct {
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
	// Ownership - Which resources to return under the resource ownership rules. 'shared' returns only resources that aren't owned, 'usable' returns shared resources plus owned resources the caller can use, 'all' returns every resource the caller can see. Defaults to 'usable' for consumers, and otherwise to 'shared' when ownership is enforced for the project and 'all' when it isn't.
	Ownership *string `json:"ownership,omitempty"`
	// Status - Filter by status (active, archived)
	Status *any `json:"status,omitempty"`
	// Id - Filter by config ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderSpecificationId - Filter by provider specification ID(s)
	ProviderSpecificationId *any `json:"provider_specification_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// AvailableForUse - Only return configs that are not owned by another integration instance provider.
	AvailableForUse *bool `json:"available_for_use,omitempty"`
	// AvailableForProviderDeploymentId - Only return configs that are not locked to a different provider deployment.
	AvailableForProviderDeploymentId *string `json:"available_for_provider_deployment_id,omitempty"`
	// ProviderConfigVaultId - Filter by config vault ID(s)
	ProviderConfigVaultId *any `json:"provider_config_vault_id,omitempty"`
	// ActorId - Filter by actor ID(s)
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerId - Filter by consumer ID(s)
	ConsumerId *any `json:"consumer_id,omitempty"`
	// IdentityId - Filter by identity ID(s)
	IdentityId *any `json:"identity_id,omitempty"`
	// IdentityCredentialId - Filter by identity credential ID(s)
	IdentityCredentialId *any `json:"identity_credential_id,omitempty"`
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider config creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider config last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// ProviderDeploymentsConfigsEndpointCreateBody contains the request body for Create.
type ProviderDeploymentsConfigsEndpointCreateBody struct {
	// ProviderId - Provider to create the config for.
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Deployment the config is locked to.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider config. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// Value - Provider-specific configuration values
	Value *map[string]any `json:"value,omitempty"`
	// ProviderConfigVaultId - Config vault ID to use as template
	ProviderConfigVaultId *string `json:"provider_config_vault_id,omitempty"`
}

// ProviderDeploymentsConfigsEndpointUpdateBody contains the request body for Update.
type ProviderDeploymentsConfigsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider config. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// ProviderDeploymentsConfigsEndpointGetConfigSchemaParams contains optional query parameters for GetConfigSchema.
type ProviderDeploymentsConfigsEndpointGetConfigSchemaParams struct {
	// ProviderId - Provider to resolve the schema for.
	ProviderId *string `json:"provider_id,omitempty"`
	// ProviderConfigId - Config to resolve the schema for.
	ProviderConfigId *string `json:"provider_config_id,omitempty"`
	// ProviderVersionId - Version to resolve the schema for.
	ProviderVersionId *string `json:"provider_version_id,omitempty"`
	// ProviderDeploymentId - Deployment to resolve the schema for.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
}

// List returns a paginated list of provider configs.
func (e *ProviderDeploymentsConfigsEndpoint) List(instanceId string, params *ProviderDeploymentsConfigsEndpointListParams) (*configs.ProviderDeploymentsConfigsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "provider-configs"},
		Query: query,
	}
	var result configs.ProviderDeploymentsConfigsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider config by ID.
func (e *ProviderDeploymentsConfigsEndpoint) Get(instanceId string, providerConfigId string) (*configs.ProviderDeploymentsConfigsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-configs", providerConfigId},
	}
	var result configs.ProviderDeploymentsConfigsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a provider config.
func (e *ProviderDeploymentsConfigsEndpoint) Create(instanceId string, body *ProviderDeploymentsConfigsEndpointCreateBody) (*configs.ProviderDeploymentsConfigsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-configs"},
		Body: body,
	}
	var result configs.ProviderDeploymentsConfigsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific provider config.
func (e *ProviderDeploymentsConfigsEndpoint) Update(instanceId string, providerConfigId string, body *ProviderDeploymentsConfigsEndpointUpdateBody) (*configs.ProviderDeploymentsConfigsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-configs", providerConfigId},
		Body: body,
	}
	var result configs.ProviderDeploymentsConfigsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a provider config. It stays listable with `status=archived`.
func (e *ProviderDeploymentsConfigsEndpoint) Delete(instanceId string, providerConfigId string) (*configs.ProviderDeploymentsConfigsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-configs", providerConfigId},
	}
	var result configs.ProviderDeploymentsConfigsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetConfigSchema returns the JSON Schema for provider configuration, resolved from the given provider, version, deployment, or config.
func (e *ProviderDeploymentsConfigsEndpoint) GetConfigSchema(instanceId string, params *ProviderDeploymentsConfigsEndpointGetConfigSchemaParams) (*configs.ProviderDeploymentsConfigsGetConfigSchemaOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "provider-config-schema"},
		Query: query,
	}
	var result configs.ProviderDeploymentsConfigsGetConfigSchemaOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
