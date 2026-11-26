package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/integrations/instancegroups"
)

// IntegrationsInstanceGroupsEndpoint provides access to an integration instance group combines providers from multiple integration instances into one set, with a shared session template.
type IntegrationsInstanceGroupsEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsInstanceGroupsEndpoint creates a new IntegrationsInstanceGroupsEndpoint.
func NewIntegrationsInstanceGroupsEndpoint(client *endpoint.Client) *IntegrationsInstanceGroupsEndpoint {
	return &IntegrationsInstanceGroupsEndpoint{client: client}
}

// IntegrationsInstanceGroupsEndpointListParams contains optional query parameters for List.
type IntegrationsInstanceGroupsEndpointListParams struct {
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
	// Status - Only return integration instance groups with one of these statuses. `deleted` groups are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return integration instance groups with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return integration instance groups for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationInstanceId - Only return integration instance groups containing one of these integration instances.
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// IntegrationInstanceProviderId - Only return integration instance groups containing one of these integration instance providers.
	IntegrationInstanceProviderId *any `json:"integration_instance_provider_id,omitempty"`
	// ProviderId - Only return integration instance groups using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationProviderId - Only return integration instance groups using one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// ProviderDeploymentId - Only return integration instance groups running on one of these deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Only return integration instance groups using one of these configs.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Only return integration instance groups using one of these credentials.
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// SessionTemplateId - Only return integration instance groups using one of these session templates.
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// CreatedAt - Only return results whose integration instance group creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration instance group last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsInstanceGroupsEndpointCreateSessionTemplateBody contains the request body for CreateSessionTemplate.
type IntegrationsInstanceGroupsEndpointCreateSessionTemplateBody struct {
	// Name - Display name for the session template.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the session template.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// IntegrationsInstanceGroupsEndpointCreateSessionBody contains the request body for CreateSession.
type IntegrationsInstanceGroupsEndpointCreateSessionBody struct {
	// Name - Display name for the session. Defaults to `Session` plus the current time.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the session.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// IntegrationsInstanceGroupsEndpointCreateBody contains the request body for Create.
type IntegrationsInstanceGroupsEndpointCreateBody struct {
	// Name - Display name for the integration instance group.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration instance group.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Providers - Providers to link into the group.
	Providers *[]map[string]any `json:"providers,omitempty"`
}

// IntegrationsInstanceGroupsEndpointUpdateBody contains the request body for Update.
type IntegrationsInstanceGroupsEndpointUpdateBody struct {
	// Name - Display name for the integration instance group.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration instance group. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
	// Providers - Providers to link or update. Existing providers not listed are left alone.
	Providers *[]map[string]any `json:"providers,omitempty"`
}

// List returns a paginated list of integration instance groups.
func (e *IntegrationsInstanceGroupsEndpoint) List(params *IntegrationsInstanceGroupsEndpointListParams) (*instancegroups.IntegrationsInstanceGroupsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"integration-instance-groups"},
		Query: query,
	}
	var result instancegroups.IntegrationsInstanceGroupsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an integration instance group.
func (e *IntegrationsInstanceGroupsEndpoint) Get(integrationInstanceGroupId string) (*instancegroups.IntegrationsInstanceGroupsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instance-groups", integrationInstanceGroupId},
	}
	var result instancegroups.IntegrationsInstanceGroupsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateSessionTemplate creates or updates the shared session template for the integration instance group.
func (e *IntegrationsInstanceGroupsEndpoint) CreateSessionTemplate(integrationInstanceGroupId string, body *IntegrationsInstanceGroupsEndpointCreateSessionTemplateBody) (*instancegroups.IntegrationsInstanceGroupsCreateSessionTemplateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instance-groups", integrationInstanceGroupId, "session-template"},
		Body: body,
	}
	var result instancegroups.IntegrationsInstanceGroupsCreateSessionTemplateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateSession creates a session from the shared session template of the integration instance group.
func (e *IntegrationsInstanceGroupsEndpoint) CreateSession(integrationInstanceGroupId string, body *IntegrationsInstanceGroupsEndpointCreateSessionBody) (*instancegroups.IntegrationsInstanceGroupsCreateSessionOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instance-groups", integrationInstanceGroupId, "session"},
		Body: body,
	}
	var result instancegroups.IntegrationsInstanceGroupsCreateSessionOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an integration instance group.
func (e *IntegrationsInstanceGroupsEndpoint) Create(body *IntegrationsInstanceGroupsEndpointCreateBody) (*instancegroups.IntegrationsInstanceGroupsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instance-groups"},
		Body: body,
	}
	var result instancegroups.IntegrationsInstanceGroupsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an integration instance group. Omitted fields keep their current value.
func (e *IntegrationsInstanceGroupsEndpoint) Update(integrationInstanceGroupId string, body *IntegrationsInstanceGroupsEndpointUpdateBody) (*instancegroups.IntegrationsInstanceGroupsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instance-groups", integrationInstanceGroupId},
		Body: body,
	}
	var result instancegroups.IntegrationsInstanceGroupsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives an integration instance group. It stays retrievable by ID and listable with `status=archived`.
func (e *IntegrationsInstanceGroupsEndpoint) Delete(integrationInstanceGroupId string) (*instancegroups.IntegrationsInstanceGroupsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-instance-groups", integrationInstanceGroupId},
	}
	var result instancegroups.IntegrationsInstanceGroupsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
