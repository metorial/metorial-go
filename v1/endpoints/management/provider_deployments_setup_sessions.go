package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerdeployments/setupsessions"
)

// ProviderDeploymentsSetupSessionsEndpoint provides access to a provider setup session is a guided flow that collects authentication, configuration, or both for a provider. `admin` sessions create a consent link for someone else to complete instead.
type ProviderDeploymentsSetupSessionsEndpoint struct {
	client *endpoint.Client
}

// NewProviderDeploymentsSetupSessionsEndpoint creates a new ProviderDeploymentsSetupSessionsEndpoint.
func NewProviderDeploymentsSetupSessionsEndpoint(client *endpoint.Client) *ProviderDeploymentsSetupSessionsEndpoint {
	return &ProviderDeploymentsSetupSessionsEndpoint{client: client}
}

// ProviderDeploymentsSetupSessionsEndpointListParams contains optional query parameters for List.
type ProviderDeploymentsSetupSessionsEndpointListParams struct {
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
	// Id - Filter by setup session ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderAuthMethodId - Filter by auth method ID(s)
	ProviderAuthMethodId *any `json:"provider_auth_method_id,omitempty"`
	// ProviderAuthConfigId - Filter by provider auth config ID(s)
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// ProviderAuthCredentialsId - Filter by provider auth credentials ID(s)
	ProviderAuthCredentialsId *any `json:"provider_auth_credentials_id,omitempty"`
	// Status - Filter by session status (archived, failed, completed, expired, pending)
	Status *any `json:"status,omitempty"`
	// Type - Filter by session type (auth_only, config_only, auth_and_config, admin), matching either the requested or the resolved type
	Type *any `json:"type,omitempty"`
	// CreatedAt - Only return results whose provider setup session creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider setup session last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// ProviderDeploymentsSetupSessionsEndpointCreateBody contains the request body for Create.
type ProviderDeploymentsSetupSessionsEndpointCreateBody struct {
	// ProviderId - Provider to set up.
	ProviderId *string `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Deployment to set up.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider setup session. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderAuthMethodId - The authentication method to use (e.g., OAuth flow)
	ProviderAuthMethodId *string `json:"provider_auth_method_id,omitempty"`
	// ProviderAuthCredentialsId - Optional OAuth app credentials to use instead of defaults
	ProviderAuthCredentialsId *string `json:"provider_auth_credentials_id,omitempty"`
	// IdentityId - Optional identity to link this setup session to
	IdentityId *string `json:"identity_id,omitempty"`
	// ConsumerId - Optional consumer to link this setup session to
	ConsumerId *string `json:"consumer_id,omitempty"`
	// RedirectUrl - URL to redirect to after setup.
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// Type - The type of setup session, determining the flow and outcome of the session. `admin` creates an admin consent link for auth methods that support it, never creates credentials or configs, and cannot be bound to an identity, consumer, or config.
	Type *string `json:"type,omitempty"`
	// Configuration - Options shaping the session flow.
	Configuration *map[string]any `json:"configuration,omitempty"`
}

// ProviderDeploymentsSetupSessionsEndpointUpdateBody contains the request body for Update.
type ProviderDeploymentsSetupSessionsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider setup session. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IdentityId - Optional identity to link this setup session to
	IdentityId *string `json:"identity_id,omitempty"`
}

// List returns a paginated list of provider setup sessions.
func (e *ProviderDeploymentsSetupSessionsEndpoint) List(instanceId string, params *ProviderDeploymentsSetupSessionsEndpointListParams) (*setupsessions.ProviderDeploymentsSetupSessionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "provider-setup-sessions"},
		Query: query,
	}
	var result setupsessions.ProviderDeploymentsSetupSessionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider setup session by ID.
func (e *ProviderDeploymentsSetupSessionsEndpoint) Get(instanceId string, providerSetupSessionId string) (*setupsessions.ProviderDeploymentsSetupSessionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-setup-sessions", providerSetupSessionId},
	}
	var result setupsessions.ProviderDeploymentsSetupSessionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a provider setup session.
func (e *ProviderDeploymentsSetupSessionsEndpoint) Create(instanceId string, body *ProviderDeploymentsSetupSessionsEndpointCreateBody) (*setupsessions.ProviderDeploymentsSetupSessionsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-setup-sessions"},
		Body: body,
	}
	var result setupsessions.ProviderDeploymentsSetupSessionsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific provider setup session.
func (e *ProviderDeploymentsSetupSessionsEndpoint) Update(instanceId string, providerSetupSessionId string, body *ProviderDeploymentsSetupSessionsEndpointUpdateBody) (*setupsessions.ProviderDeploymentsSetupSessionsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-setup-sessions", providerSetupSessionId},
		Body: body,
	}
	var result setupsessions.ProviderDeploymentsSetupSessionsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a provider setup session. It stays listable with `status=archived`.
func (e *ProviderDeploymentsSetupSessionsEndpoint) Delete(instanceId string, providerSetupSessionId string) (*setupsessions.ProviderDeploymentsSetupSessionsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-setup-sessions", providerSetupSessionId},
	}
	var result setupsessions.ProviderDeploymentsSetupSessionsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
