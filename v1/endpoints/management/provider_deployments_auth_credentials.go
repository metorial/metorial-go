package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerdeployments/authcredentials"
)

// ProviderDeploymentsAuthCredentialsEndpoint provides access to provider auth credentials are an OAuth app registration used with a provider: a client ID, a client secret, and scopes.
type ProviderDeploymentsAuthCredentialsEndpoint struct {
	client *endpoint.Client
}

// NewProviderDeploymentsAuthCredentialsEndpoint creates a new ProviderDeploymentsAuthCredentialsEndpoint.
func NewProviderDeploymentsAuthCredentialsEndpoint(client *endpoint.Client) *ProviderDeploymentsAuthCredentialsEndpoint {
	return &ProviderDeploymentsAuthCredentialsEndpoint{client: client}
}

// ProviderDeploymentsAuthCredentialsEndpointListParams contains optional query parameters for List.
type ProviderDeploymentsAuthCredentialsEndpointListParams struct {
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
	// Id - Filter by credential ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderAuthMethodId - Filter by provider auth method ID(s)
	ProviderAuthMethodId *any `json:"provider_auth_method_id,omitempty"`
	// Origin - Filter by credential origin (custom, managed)
	Origin *any `json:"origin,omitempty"`
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider auth credentials creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider auth credentials last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// ProviderDeploymentsAuthCredentialsEndpointCreateBody contains the request body for Create.
type ProviderDeploymentsAuthCredentialsEndpointCreateBody struct {
	// ProviderId - Provider to create credentials for.
	ProviderId string `json:"provider_id"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider auth credentials. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Config - OAuth app credentials to store.
	Config map[string]any `json:"config"`
}

// ProviderDeploymentsAuthCredentialsEndpointUpdateBody contains the request body for Update.
type ProviderDeploymentsAuthCredentialsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider auth credentials. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ClientId - Client ID of your OAuth app.
	ClientId *string `json:"client_id,omitempty"`
	// ClientSecret - Client secret of your OAuth app.
	ClientSecret *string `json:"client_secret,omitempty"`
	// Scopes - OAuth scopes to request.
	Scopes *[]string `json:"scopes,omitempty"`
}

// List returns a paginated list of provider auth credentials.
func (e *ProviderDeploymentsAuthCredentialsEndpoint) List(instanceId string, params *ProviderDeploymentsAuthCredentialsEndpointListParams) (*authcredentials.ProviderDeploymentsAuthCredentialsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "provider-auth-credentials"},
		Query: query,
	}
	var result authcredentials.ProviderDeploymentsAuthCredentialsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves provider auth credentials.
func (e *ProviderDeploymentsAuthCredentialsEndpoint) Get(instanceId string, providerAuthCredentialsId string) (*authcredentials.ProviderDeploymentsAuthCredentialsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-auth-credentials", providerAuthCredentialsId},
	}
	var result authcredentials.ProviderDeploymentsAuthCredentialsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates provider auth credentials.
func (e *ProviderDeploymentsAuthCredentialsEndpoint) Create(instanceId string, body *ProviderDeploymentsAuthCredentialsEndpointCreateBody) (*authcredentials.ProviderDeploymentsAuthCredentialsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-auth-credentials"},
		Body: body,
	}
	var result authcredentials.ProviderDeploymentsAuthCredentialsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates provider auth credentials.
func (e *ProviderDeploymentsAuthCredentialsEndpoint) Update(instanceId string, providerAuthCredentialsId string, body *ProviderDeploymentsAuthCredentialsEndpointUpdateBody) (*authcredentials.ProviderDeploymentsAuthCredentialsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-auth-credentials", providerAuthCredentialsId},
		Body: body,
	}
	var result authcredentials.ProviderDeploymentsAuthCredentialsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives provider auth credentials. They stay listable with `status=archived`.
func (e *ProviderDeploymentsAuthCredentialsEndpoint) Delete(instanceId string, providerAuthCredentialsId string) (*authcredentials.ProviderDeploymentsAuthCredentialsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-auth-credentials", providerAuthCredentialsId},
	}
	var result authcredentials.ProviderDeploymentsAuthCredentialsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
