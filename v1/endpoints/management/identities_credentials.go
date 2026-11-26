package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/identities/credentials"
)

// IdentitiesCredentialsEndpoint provides access to an identity credential is one provider attachment on an identity. It links the identity to a provider deployment, config, or auth config.
type IdentitiesCredentialsEndpoint struct {
	client *endpoint.Client
}

// NewIdentitiesCredentialsEndpoint creates a new IdentitiesCredentialsEndpoint.
func NewIdentitiesCredentialsEndpoint(client *endpoint.Client) *IdentitiesCredentialsEndpoint {
	return &IdentitiesCredentialsEndpoint{client: client}
}

// IdentitiesCredentialsEndpointListParams contains optional query parameters for List.
type IdentitiesCredentialsEndpointListParams struct {
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
	// Status - Only return identity credentials with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return identity credentials with one of these IDs.
	Id *any `json:"id,omitempty"`
	// AgentId - Only return identity credentials owned by one of these agents.
	AgentId *any `json:"agent_id,omitempty"`
	// ActorId - Only return identity credentials owned by one of these identity actors.
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerProfileId - Only return identity credentials owned by one of these consumer profiles.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// IdentityId - Only return identity credentials attached to one of these identities.
	IdentityId *any `json:"identity_id,omitempty"`
	// ProviderId - Only return identity credentials for one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Only return identity credentials attached to one of these provider deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Only return identity credentials attached to one of these provider configs.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Only return identity credentials attached to one of these provider auth configs.
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Only return results whose identity credential creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose identity credential last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IdentitiesCredentialsEndpointCreateBody contains the request body for Create.
type IdentitiesCredentialsEndpointCreateBody struct {
	// IdentityId - Identity that will own the new credential.
	IdentityId string `json:"identity_id"`
	// DeploymentId - Provider deployment to attach to the credential.
	DeploymentId *string `json:"deployment_id,omitempty"`
	// ConfigId - Provider config to attach to the credential.
	ConfigId *string `json:"config_id,omitempty"`
	// AuthConfigId - Provider auth config to attach to the credential.
	AuthConfigId *string `json:"auth_config_id,omitempty"`
	// DelegationConfigId - Delegation config to apply to the credential.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
}

// IdentitiesCredentialsEndpointUpdateBody contains the request body for Update.
type IdentitiesCredentialsEndpointUpdateBody struct {
	// DelegationConfigId - Delegation config to apply to the credential.
	DelegationConfigId string `json:"delegation_config_id"`
}

// List returns a paginated list of identity credentials for the instance.
func (e *IdentitiesCredentialsEndpoint) List(instanceId string, params *IdentitiesCredentialsEndpointListParams) (*credentials.IdentitiesCredentialsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "identity-credentials"},
		Query: query,
	}
	var result credentials.IdentitiesCredentialsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an identity credential.
func (e *IdentitiesCredentialsEndpoint) Get(instanceId string, identityCredentialId string) (*credentials.IdentitiesCredentialsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "identity-credentials", identityCredentialId},
	}
	var result credentials.IdentitiesCredentialsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an identity credential.
func (e *IdentitiesCredentialsEndpoint) Create(instanceId string, body *IdentitiesCredentialsEndpointCreateBody) (*credentials.IdentitiesCredentialsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "identity-credentials"},
		Body: body,
	}
	var result credentials.IdentitiesCredentialsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates the delegation config attached to an identity credential.
func (e *IdentitiesCredentialsEndpoint) Update(instanceId string, identityCredentialId string, body *IdentitiesCredentialsEndpointUpdateBody) (*credentials.IdentitiesCredentialsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "identity-credentials", identityCredentialId},
		Body: body,
	}
	var result credentials.IdentitiesCredentialsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives an identity credential. It stays listable with `status=archived`.
func (e *IdentitiesCredentialsEndpoint) Delete(instanceId string, identityCredentialId string) (*credentials.IdentitiesCredentialsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "identity-credentials", identityCredentialId},
	}
	var result credentials.IdentitiesCredentialsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
