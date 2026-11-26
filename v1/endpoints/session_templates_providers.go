package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sessiontemplates/providers"
)

// SessionTemplatesProvidersEndpoint provides access to a session template provider is one provider in a session template, with its deployment, config, auth config, and tool filters.
type SessionTemplatesProvidersEndpoint struct {
	client *endpoint.Client
}

// NewSessionTemplatesProvidersEndpoint creates a new SessionTemplatesProvidersEndpoint.
func NewSessionTemplatesProvidersEndpoint(client *endpoint.Client) *SessionTemplatesProvidersEndpoint {
	return &SessionTemplatesProvidersEndpoint{client: client}
}

// SessionTemplatesProvidersEndpointListParams contains optional query parameters for List.
type SessionTemplatesProvidersEndpointListParams struct {
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
	// Status - Only return session template providers with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Filter by session template provider ID(s)
	Id *any `json:"id,omitempty"`
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
	// CreatedAt - Only return results whose session template provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session template provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SessionTemplatesProvidersEndpointCreateBody contains the request body for Create.
type SessionTemplatesProvidersEndpointCreateBody struct {
	// SessionTemplateId - Session template to add the provider to.
	SessionTemplateId string `json:"session_template_id"`
	// ProviderDeploymentId - Deployment the template provider runs with.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Config the template provider runs with.
	ProviderConfigId      *string `json:"provider_config_id,omitempty"`
	ProviderConfigVaultId *string `json:"provider_config_vault_id,omitempty"`
	// ProviderAuthConfigId - Auth config the template provider runs with.
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// SessionTemplatesProvidersEndpointUpdateBody contains the request body for Update.
type SessionTemplatesProvidersEndpointUpdateBody struct {
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// List returns a paginated list of session template providers.
func (e *SessionTemplatesProvidersEndpoint) List(params *SessionTemplatesProvidersEndpointListParams) (*providers.SessionTemplatesProvidersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"session-template-providers"},
		Query: query,
	}
	var result providers.SessionTemplatesProvidersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a session template provider.
func (e *SessionTemplatesProvidersEndpoint) Get(sessionTemplateProviderId string) (*providers.SessionTemplatesProvidersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"session-template-providers", sessionTemplateProviderId},
	}
	var result providers.SessionTemplatesProvidersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a session template provider.
func (e *SessionTemplatesProvidersEndpoint) Create(body *SessionTemplatesProvidersEndpointCreateBody) (*providers.SessionTemplatesProvidersCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"session-template-providers"},
		Body: body,
	}
	var result providers.SessionTemplatesProvidersCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a session template provider.
func (e *SessionTemplatesProvidersEndpoint) Update(sessionTemplateProviderId string, body *SessionTemplatesProvidersEndpointUpdateBody) (*providers.SessionTemplatesProvidersUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"session-template-providers", sessionTemplateProviderId},
		Body: body,
	}
	var result providers.SessionTemplatesProvidersUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a session template provider. It stays listable with `status=archived`.
func (e *SessionTemplatesProvidersEndpoint) Delete(sessionTemplateProviderId string) (*providers.SessionTemplatesProvidersDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"session-template-providers", sessionTemplateProviderId},
	}
	var result providers.SessionTemplatesProvidersDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
