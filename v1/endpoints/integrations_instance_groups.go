package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrations/instancegroups"
)

// IntegrationsInstanceGroupsEndpoint provides access to integration instance groups combine instance providers into a grouped routed configuration.
type IntegrationsInstanceGroupsEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsInstanceGroupsEndpoint creates a new IntegrationsInstanceGroupsEndpoint.
func NewIntegrationsInstanceGroupsEndpoint(client *endpoint.Client) *IntegrationsInstanceGroupsEndpoint {
	return &IntegrationsInstanceGroupsEndpoint{client: client}
}

// IntegrationsInstanceGroupsEndpointListParams contains optional query parameters for List.
type IntegrationsInstanceGroupsEndpointListParams struct {
	Limit                         *float64 `json:"limit,omitempty"`
	After                         *string  `json:"after,omitempty"`
	Before                        *string  `json:"before,omitempty"`
	Cursor                        *string  `json:"cursor,omitempty"`
	Order                         *string  `json:"order,omitempty"`
	Status                        *any     `json:"status,omitempty"`
	Id                            *any     `json:"id,omitempty"`
	IntegrationId                 *any     `json:"integration_id,omitempty"`
	IntegrationInstanceId         *any     `json:"integration_instance_id,omitempty"`
	IntegrationInstanceProviderId *any     `json:"integration_instance_provider_id,omitempty"`
	ProviderId                    *any     `json:"provider_id,omitempty"`
	IntegrationProviderId         *any     `json:"integration_provider_id,omitempty"`
	ProviderDeploymentId          *any     `json:"provider_deployment_id,omitempty"`
	ProviderConfigId              *any     `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId          *any     `json:"provider_auth_config_id,omitempty"`
	SessionTemplateId             *any     `json:"session_template_id,omitempty"`
	// CreatedAt - Filter integration instance group creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration instance group last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsInstanceGroupsEndpointCreateSessionTemplateBody contains the request body for CreateSessionTemplate.
type IntegrationsInstanceGroupsEndpointCreateSessionTemplateBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
}

// IntegrationsInstanceGroupsEndpointCreateSessionBody contains the request body for CreateSession.
type IntegrationsInstanceGroupsEndpointCreateSessionBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
}

// IntegrationsInstanceGroupsEndpointCreateBody contains the request body for Create.
type IntegrationsInstanceGroupsEndpointCreateBody struct {
	Name        string            `json:"name"`
	Description *string           `json:"description,omitempty"`
	Metadata    *map[string]any   `json:"metadata,omitempty"`
	Providers   *[]map[string]any `json:"providers,omitempty"`
}

// IntegrationsInstanceGroupsEndpointUpdateBody contains the request body for Update.
type IntegrationsInstanceGroupsEndpointUpdateBody struct {
	Name        *string           `json:"name,omitempty"`
	Description *string           `json:"description,omitempty"`
	Metadata    *map[string]any   `json:"metadata,omitempty"`
	Providers   *[]map[string]any `json:"providers,omitempty"`
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

// Get retrieves a specific integration instance group.
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

// CreateSessionTemplate creates or updates the shared session template for a specific integration instance group.
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

// CreateSession creates a session from the shared session template of a specific integration instance group.
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

// Create creates a new integration instance group.
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

// Update updates a specific integration instance group.
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

// Delete archives a specific integration instance group.
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
