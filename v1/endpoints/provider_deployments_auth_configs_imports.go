package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerdeployments/authconfigs/imports"
)

// ProviderDeploymentsAuthConfigsImportsEndpoint provides access to a provider auth import is existing authentication credentials brought into Metorial from another system, so users skip re-authenticating.
type ProviderDeploymentsAuthConfigsImportsEndpoint struct {
	client *endpoint.Client
}

// NewProviderDeploymentsAuthConfigsImportsEndpoint creates a new ProviderDeploymentsAuthConfigsImportsEndpoint.
func NewProviderDeploymentsAuthConfigsImportsEndpoint(client *endpoint.Client) *ProviderDeploymentsAuthConfigsImportsEndpoint {
	return &ProviderDeploymentsAuthConfigsImportsEndpoint{client: client}
}

// ProviderDeploymentsAuthConfigsImportsEndpointListParams contains optional query parameters for List.
type ProviderDeploymentsAuthConfigsImportsEndpointListParams struct {
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
	// Id - Filter by import ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderAuthCredentialsId - Filter by auth credentials ID(s)
	ProviderAuthCredentialsId *any `json:"provider_auth_credentials_id,omitempty"`
	// ProviderAuthConfigId - Filter by auth config ID(s)
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// CreatedAt - Only return results whose provider auth import creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider auth import last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// ProviderDeploymentsAuthConfigsImportsEndpointCreateBody contains the request body for Create.
type ProviderDeploymentsAuthConfigsImportsEndpointCreateBody struct {
	// ProviderId - Provider the imported credentials belong to.
	ProviderId *string `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Deployment the imported credentials belong to.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderAuthConfigId - Auth config the imported credentials belong to.
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	// ProviderAuthMethodId - The authentication method used by these credentials
	ProviderAuthMethodId *string `json:"provider_auth_method_id,omitempty"`
	// Note - A note describing the import source or reason
	Note string `json:"note"`
	// Metadata - Custom data to identify or organize the provider auth import. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Value - The credential data to import
	Value map[string]any `json:"value"`
}

// ProviderDeploymentsAuthConfigsImportsEndpointGetSchemaParams contains optional query parameters for GetSchema.
type ProviderDeploymentsAuthConfigsImportsEndpointGetSchemaParams struct {
	// ProviderId - Provider to resolve the schema for.
	ProviderId *string `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Deployment to resolve the schema for.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderAuthConfigId - Auth config to resolve the schema for.
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	// ProviderAuthMethodId - Auth method to resolve the schema for.
	ProviderAuthMethodId *string `json:"provider_auth_method_id,omitempty"`
}

// List returns a paginated list of provider auth imports.
func (e *ProviderDeploymentsAuthConfigsImportsEndpoint) List(params *ProviderDeploymentsAuthConfigsImportsEndpointListParams) (*imports.ProviderDeploymentsAuthConfigsImportsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-auth-config-imports"},
		Query: query,
	}
	var result imports.ProviderDeploymentsAuthConfigsImportsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider auth import by ID.
func (e *ProviderDeploymentsAuthConfigsImportsEndpoint) Get(providerAuthImportId string) (*imports.ProviderDeploymentsAuthConfigsImportsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-auth-config-imports", providerAuthImportId},
	}
	var result imports.ProviderDeploymentsAuthConfigsImportsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create imports existing authentication credentials into Metorial.
func (e *ProviderDeploymentsAuthConfigsImportsEndpoint) Create(body *ProviderDeploymentsAuthConfigsImportsEndpointCreateBody) (*imports.ProviderDeploymentsAuthConfigsImportsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-auth-config-imports"},
		Body: body,
	}
	var result imports.ProviderDeploymentsAuthConfigsImportsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSchema retrieves the JSON Schema for importing authentication credentials.
func (e *ProviderDeploymentsAuthConfigsImportsEndpoint) GetSchema(params *ProviderDeploymentsAuthConfigsImportsEndpointGetSchemaParams) (*imports.ProviderDeploymentsAuthConfigsImportsGetSchemaOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-auth-config-imports", "schema"},
		Query: query,
	}
	var result imports.ProviderDeploymentsAuthConfigsImportsGetSchemaOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
