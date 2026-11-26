package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sessions"
)

// SessionsEndpoint provides access to a session is a connection between clients and providers over MCP. Each session connects to one or more providers.
type SessionsEndpoint struct {
	client *endpoint.Client
}

// NewSessionsEndpoint creates a new SessionsEndpoint.
func NewSessionsEndpoint(client *endpoint.Client) *SessionsEndpoint {
	return &SessionsEndpoint{client: client}
}

// SessionsEndpointListParams contains optional query parameters for List.
type SessionsEndpointListParams struct {
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
	// Id - Filter by session ID(s)
	Id *any `json:"id,omitempty"`
	// AgentId - Filter by connected agent ID(s)
	AgentId *any `json:"agent_id,omitempty"`
	// ActorId - Filter by connected identity actor ID(s)
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerId - Filter by connected consumer ID(s)
	ConsumerId *any `json:"consumer_id,omitempty"`
	// IdentityId - Filter by connected identity ID(s)
	IdentityId *any `json:"identity_id,omitempty"`
	// SessionTemplateId - Filter by session template ID(s)
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// SessionProviderId - Filter by session provider ID(s)
	SessionProviderId *any `json:"session_provider_id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter by provider config ID(s)
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Filter by provider auth config ID(s)
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Only return results whose session creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
	// Search - Search by session ID, or by the name of a related provider, auth method, deployment, config, auth config, agent, actor, identity, or consumer
	Search *string `json:"search,omitempty"`
}

// SessionsEndpointCreateBody contains the request body for Create.
type SessionsEndpointCreateBody struct {
	// Name - Display name. Generated when omitted.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the session.
	Metadata  *map[string]any  `json:"metadata,omitempty"`
	Providers []map[string]any `json:"providers"`
}

// SessionsEndpointUpdateBody contains the request body for Update.
type SessionsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the session. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of sessions.
func (e *SessionsEndpoint) List(instanceId string, params *SessionsEndpointListParams) (*sessions.SessionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "sessions"},
		Query: query,
	}
	var result sessions.SessionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a session.
func (e *SessionsEndpoint) Get(instanceId string, sessionId string) (*sessions.SessionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sessions", sessionId},
	}
	var result sessions.SessionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a session with one or more providers, from a template or inline definitions.
func (e *SessionsEndpoint) Create(instanceId string, body *SessionsEndpointCreateBody) (*sessions.SessionsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sessions"},
		Body: body,
	}
	var result sessions.SessionsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a session. Omitted fields keep their current value.
func (e *SessionsEndpoint) Update(instanceId string, sessionId string, body *SessionsEndpointUpdateBody) (*sessions.SessionsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sessions", sessionId},
		Body: body,
	}
	var result sessions.SessionsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a session. Its connections disconnect and its providers are archived. It stays listable with `status=archived`.
func (e *SessionsEndpoint) Delete(instanceId string, sessionId string) (*sessions.SessionsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sessions", sessionId},
	}
	var result sessions.SessionsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
