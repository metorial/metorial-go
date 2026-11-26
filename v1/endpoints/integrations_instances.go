package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/integrations/instances"
)

// IntegrationsInstancesEndpoint provides access to an integration instance is an integration set up for a specific actor or identity. It binds each of the integration’s providers to the deployment, config, credentials, and tool filters it runs with.
type IntegrationsInstancesEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsInstancesEndpoint creates a new IntegrationsInstancesEndpoint.
func NewIntegrationsInstancesEndpoint(client *endpoint.Client) *IntegrationsInstancesEndpoint {
	return &IntegrationsInstancesEndpoint{client: client}
}

// IntegrationsInstancesEndpointListParams contains optional query parameters for List.
type IntegrationsInstancesEndpointListParams struct {
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
	// Ownership - Ownership scope: `shared` (not owned by someone else), `usable` (shared plus ones you may use), or `all`. Defaults to `shared` when ownership is enforced, `all` otherwise.
	Ownership *string `json:"ownership,omitempty"`
	// Search - Search integration instances by name or description.
	Search *string `json:"search,omitempty"`
	// Status - Only return integration instances with one of these statuses. `deleted` instances are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return integration instances with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return integration instances for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// ProviderId - Only return integration instances using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationProviderId - Only return integration instances using one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// IdentityId - Only return integration instances for one of these identities.
	IdentityId *any `json:"identity_id,omitempty"`
	// IdentityCredentialId - Only return integration instances using one of these identity credentials.
	IdentityCredentialId *any `json:"identity_credential_id,omitempty"`
	// IdentityActorId - Only return integration instances for one of these identity actors.
	IdentityActorId *any `json:"identity_actor_id,omitempty"`
	// ProviderDeploymentId - Only return integration instances running on one of these deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Only return integration instances using one of these configs.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Only return integration instances using one of these credentials.
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// SessionTemplateId - Only return integration instances using one of these session templates.
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// CreatedAt - Only return results whose integration instance creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration instance last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsInstancesEndpointCreateSessionTemplateBody contains the request body for CreateSessionTemplate.
type IntegrationsInstancesEndpointCreateSessionTemplateBody struct {
	// Name - Display name for the session template.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the session template.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// IntegrationsInstancesEndpointCreateSessionBody contains the request body for CreateSession.
type IntegrationsInstancesEndpointCreateSessionBody struct {
	// Name - Display name for the session. Defaults to `Session` plus the current time.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the session.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// IntegrationsInstancesEndpointCreateBody contains the request body for Create.
type IntegrationsInstancesEndpointCreateBody struct {
	// IntegrationId - Integration to create the instance for.
	IntegrationId string `json:"integration_id"`
	// Name - Display name for the integration instance.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration instance.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IdentityActorId - Identity actor the instance connects as. Omitted leaves the instance without one.
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// IdentityId - Identity the instance connects as. Omitted leaves the instance without one.
	IdentityId *string `json:"identity_id,omitempty"`
	// Providers - Providers to set up on the instance.
	Providers *[]map[string]any `json:"providers,omitempty"`
}

// IntegrationsInstancesEndpointUpdateBody contains the request body for Update.
type IntegrationsInstancesEndpointUpdateBody struct {
	// Name - Display name for the integration instance.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration instance. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
	// IdentityActorId - Identity actor the instance connects as. `null` clears it. Omitted keeps the current one.
	IdentityActorId *nullable.Nullable[string] `json:"identity_actor_id,omitempty"`
	// IdentityId - Identity the instance connects as. `null` clears it. Omitted keeps the current one.
	IdentityId *nullable.Nullable[string] `json:"identity_id,omitempty"`
	// Providers - Providers to create or update. Existing providers not listed are left alone.
	Providers *[]map[string]any `json:"providers,omitempty"`
}

// List returns a paginated list of integration instances.
func (e *IntegrationsInstancesEndpoint) List(params *IntegrationsInstancesEndpointListParams) (*instances.IntegrationsInstancesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"integration-instances"},
		Query: query,
	}
	var result instances.IntegrationsInstancesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an integration instance.
func (e *IntegrationsInstancesEndpoint) Get(integrationInstanceId string) (*instances.IntegrationsInstancesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instances", integrationInstanceId},
	}
	var result instances.IntegrationsInstancesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateSessionTemplate creates or updates the shared session template for the integration instance.
func (e *IntegrationsInstancesEndpoint) CreateSessionTemplate(integrationInstanceId string, body *IntegrationsInstancesEndpointCreateSessionTemplateBody) (*instances.IntegrationsInstancesCreateSessionTemplateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instances", integrationInstanceId, "session-template"},
		Body: body,
	}
	var result instances.IntegrationsInstancesCreateSessionTemplateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateSession creates a session from the shared session template of the integration instance.
func (e *IntegrationsInstancesEndpoint) CreateSession(integrationInstanceId string, body *IntegrationsInstancesEndpointCreateSessionBody) (*instances.IntegrationsInstancesCreateSessionOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instances", integrationInstanceId, "session"},
		Body: body,
	}
	var result instances.IntegrationsInstancesCreateSessionOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an integration instance.
func (e *IntegrationsInstancesEndpoint) Create(body *IntegrationsInstancesEndpointCreateBody) (*instances.IntegrationsInstancesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instances"},
		Body: body,
	}
	var result instances.IntegrationsInstancesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an integration instance. Omitted fields keep their current value.
func (e *IntegrationsInstancesEndpoint) Update(integrationInstanceId string, body *IntegrationsInstancesEndpointUpdateBody) (*instances.IntegrationsInstancesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instances", integrationInstanceId},
		Body: body,
	}
	var result instances.IntegrationsInstancesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives an integration instance. It stays retrievable by ID and listable with `status=archived`.
func (e *IntegrationsInstancesEndpoint) Delete(integrationInstanceId string) (*instances.IntegrationsInstancesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instances", integrationInstanceId},
	}
	var result instances.IntegrationsInstancesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
