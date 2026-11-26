package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sessions/providers"
)

// SessionsProvidersEndpoint provides access to a session provider is one provider connected to a session, with the deployment, config, and credentials it runs with.
type SessionsProvidersEndpoint struct {
	client *endpoint.Client
}

// NewSessionsProvidersEndpoint creates a new SessionsProvidersEndpoint.
func NewSessionsProvidersEndpoint(client *endpoint.Client) *SessionsProvidersEndpoint {
	return &SessionsProvidersEndpoint{client: client}
}

// SessionsProvidersEndpointListParams contains optional query parameters for List.
type SessionsProvidersEndpointListParams struct {
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
	// Id - Filter by session provider ID(s)
	Id *any `json:"id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// SessionTemplateId - Filter by session template ID(s)
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter by provider config ID(s)
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Filter by provider auth config ID(s)
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// Status - Only return session providers with one of these statuses.
	Status *any `json:"status,omitempty"`
	// CreatedAt - Only return results whose session provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SessionsProvidersEndpointCreateBody contains the request body for Create.
type SessionsProvidersEndpointCreateBody struct {
	SessionId string `json:"session_id"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// SessionsProvidersEndpointUpdateBody contains the request body for Update.
type SessionsProvidersEndpointUpdateBody struct {
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// List returns a paginated list of session providers.
func (e *SessionsProvidersEndpoint) List(instanceId string, params *SessionsProvidersEndpointListParams) (*providers.SessionsProvidersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "session-providers"},
		Query: query,
	}
	var result providers.SessionsProvidersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a session provider.
func (e *SessionsProvidersEndpoint) Get(instanceId string, sessionProviderId string) (*providers.SessionsProvidersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "session-providers", sessionProviderId},
	}
	var result providers.SessionsProvidersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a session provider. The session must be active.
func (e *SessionsProvidersEndpoint) Create(instanceId string, body *SessionsProvidersEndpointCreateBody) (*providers.SessionsProvidersCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "session-providers"},
		Body: body,
	}
	var result providers.SessionsProvidersCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a session provider.
func (e *SessionsProvidersEndpoint) Update(instanceId string, sessionProviderId string, body *SessionsProvidersEndpointUpdateBody) (*providers.SessionsProvidersUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "session-providers", sessionProviderId},
		Body: body,
	}
	var result providers.SessionsProvidersUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a session provider. It stays listable with `status=archived`.
func (e *SessionsProvidersEndpoint) Delete(instanceId string, sessionProviderId string) (*providers.SessionsProvidersDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "session-providers", sessionProviderId},
	}
	var result providers.SessionsProvidersDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
