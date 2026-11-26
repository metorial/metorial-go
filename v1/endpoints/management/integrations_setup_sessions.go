package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrations/setupsessions"
)

// IntegrationsSetupSessionsEndpoint provides access to an integration setup session is a shareable flow that connects each provider an integration instance needs. It creates a draft instance unless you pass one, and expires 30 minutes after creation unless `expires_at` is set.
type IntegrationsSetupSessionsEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsSetupSessionsEndpoint creates a new IntegrationsSetupSessionsEndpoint.
func NewIntegrationsSetupSessionsEndpoint(client *endpoint.Client) *IntegrationsSetupSessionsEndpoint {
	return &IntegrationsSetupSessionsEndpoint{client: client}
}

// IntegrationsSetupSessionsEndpointListParams contains optional query parameters for List.
type IntegrationsSetupSessionsEndpointListParams struct {
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
	// Status - Only return setup sessions with one of these statuses. `deleted` sessions are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return setup sessions with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return setup sessions for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationInstanceId - Only return setup sessions for one of these integration instances.
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// CreatedAt - Only return results whose integration setup session creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration setup session last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsSetupSessionsEndpointCreateBody contains the request body for Create.
type IntegrationsSetupSessionsEndpointCreateBody struct {
	// IntegrationId - Integration to set up.
	IntegrationId string `json:"integration_id"`
	// IntegrationInstanceId - Existing instance to reuse. Omitted creates a draft instance.
	IntegrationInstanceId *string `json:"integration_instance_id,omitempty"`
	// Name - Display name for the setup session.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the setup session.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IdentityActorId - Identity actor the instance is set up for. Cannot be set when `integration_instance_id` is given.
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// IdentityId - Identity the instance is set up for. Cannot be set when `integration_instance_id` is given.
	IdentityId *string `json:"identity_id,omitempty"`
	// ExpiresAt - When the session expires. Defaults to 30 minutes after creation.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// RedirectUrl - Where to send the user once setup completes.
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// Configuration - Configuration for the setup flow.
	Configuration *map[string]any `json:"configuration,omitempty"`
}

// List returns a paginated list of integration setup sessions.
func (e *IntegrationsSetupSessionsEndpoint) List(instanceId string, params *IntegrationsSetupSessionsEndpointListParams) (*setupsessions.IntegrationsSetupSessionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "integration-setup-sessions"},
		Query: query,
	}
	var result setupsessions.IntegrationsSetupSessionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an integration setup session.
func (e *IntegrationsSetupSessionsEndpoint) Get(instanceId string, integrationSetupSessionId string) (*setupsessions.IntegrationsSetupSessionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "integration-setup-sessions", integrationSetupSessionId},
	}
	var result setupsessions.IntegrationsSetupSessionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an integration setup session. Creates a draft integration instance unless `integration_instance_id` is given, in which case that instance is reused and its providers are configured on completion.
func (e *IntegrationsSetupSessionsEndpoint) Create(instanceId string, body *IntegrationsSetupSessionsEndpointCreateBody) (*setupsessions.IntegrationsSetupSessionsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "integration-setup-sessions"},
		Body: body,
	}
	var result setupsessions.IntegrationsSetupSessionsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
