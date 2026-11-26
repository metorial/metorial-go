package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sessiontemplates"
)

// SessionTemplatesEndpoint provides access to a session template is a reusable configuration for creating sessions with the same providers.
type SessionTemplatesEndpoint struct {
	client *endpoint.Client
}

// NewSessionTemplatesEndpoint creates a new SessionTemplatesEndpoint.
func NewSessionTemplatesEndpoint(client *endpoint.Client) *SessionTemplatesEndpoint {
	return &SessionTemplatesEndpoint{client: client}
}

// SessionTemplatesEndpointListParams contains optional query parameters for List.
type SessionTemplatesEndpointListParams struct {
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
	// Ownership - Which resources to return under the resource ownership rules: shared (resources without an owner), usable (shared resources plus resources the caller can use), or all. Defaults to shared when ownership is enforced for the project and to all otherwise.
	Ownership *string `json:"ownership,omitempty"`
	// Status - Filter by status (active, archived)
	Status *any `json:"status,omitempty"`
	// Id - Filter by session template ID(s)
	Id *any `json:"id,omitempty"`
	// SessionId - Filter templates that include sessions with these IDs
	SessionId *any `json:"session_id,omitempty"`
	// SessionProviderId - Filter templates that include session providers with these IDs
	SessionProviderId *any `json:"session_provider_id,omitempty"`
	// ProviderId - Filter templates that include providers with these IDs
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter templates that include provider deployments with these IDs
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter templates that include provider configs with these IDs
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Filter templates that include provider auth configs with these IDs
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Only return results whose session template creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session template last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SessionTemplatesEndpointCreateBody contains the request body for Create.
type SessionTemplatesEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the session template.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Providers - Optional list of providers to include in the template
	Providers *[]map[string]any `json:"providers,omitempty"`
}

// SessionTemplatesEndpointUpdateBody contains the request body for Update.
type SessionTemplatesEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the session template. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of session templates.
func (e *SessionTemplatesEndpoint) List(instanceId string, params *SessionTemplatesEndpointListParams) (*sessiontemplates.SessionTemplatesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "session-templates"},
		Query: query,
	}
	var result sessiontemplates.SessionTemplatesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a session template.
func (e *SessionTemplatesEndpoint) Get(instanceId string, sessionTemplateId string) (*sessiontemplates.SessionTemplatesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "session-templates", sessionTemplateId},
	}
	var result sessiontemplates.SessionTemplatesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a session template.
func (e *SessionTemplatesEndpoint) Create(instanceId string, body *SessionTemplatesEndpointCreateBody) (*sessiontemplates.SessionTemplatesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "session-templates"},
		Body: body,
	}
	var result sessiontemplates.SessionTemplatesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a session template. Omitted fields keep their current value.
func (e *SessionTemplatesEndpoint) Update(instanceId string, sessionTemplateId string, body *SessionTemplatesEndpointUpdateBody) (*sessiontemplates.SessionTemplatesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "session-templates", sessionTemplateId},
		Body: body,
	}
	var result sessiontemplates.SessionTemplatesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a session template. It stays listable with `status=archived`.
func (e *SessionTemplatesEndpoint) Delete(instanceId string, sessionTemplateId string) (*sessiontemplates.SessionTemplatesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "session-templates", sessionTemplateId},
	}
	var result sessiontemplates.SessionTemplatesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListTools returns the effective set of tools available through the providers in a session template, filtered by the tool filters of each provider, deployment, config, and auth config.
func (e *SessionTemplatesEndpoint) ListTools(instanceId string, sessionTemplateId string) (*sessiontemplates.SessionTemplatesListToolsOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "session-templates", sessionTemplateId, "tools"},
	}
	var result sessiontemplates.SessionTemplatesListToolsOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
