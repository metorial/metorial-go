package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrations/instances"
)

// IntegrationsInstancesEndpoint provides access to integration instances materialize an integration for a specific actor, identity, or runtime configuration.
type IntegrationsInstancesEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsInstancesEndpoint creates a new IntegrationsInstancesEndpoint.
func NewIntegrationsInstancesEndpoint(client *endpoint.Client) *IntegrationsInstancesEndpoint {
	return &IntegrationsInstancesEndpoint{client: client}
}

// IntegrationsInstancesEndpointListParams contains optional query parameters for List.
type IntegrationsInstancesEndpointListParams struct {
	Limit                 *float64 `json:"limit,omitempty"`
	After                 *string  `json:"after,omitempty"`
	Before                *string  `json:"before,omitempty"`
	Cursor                *string  `json:"cursor,omitempty"`
	Order                 *string  `json:"order,omitempty"`
	Search                *string  `json:"search,omitempty"`
	Status                *any     `json:"status,omitempty"`
	Id                    *any     `json:"id,omitempty"`
	IntegrationId         *any     `json:"integration_id,omitempty"`
	ProviderId            *any     `json:"provider_id,omitempty"`
	IntegrationProviderId *any     `json:"integration_provider_id,omitempty"`
	IdentityId            *any     `json:"identity_id,omitempty"`
	IdentityCredentialId  *any     `json:"identity_credential_id,omitempty"`
	IdentityActorId       *any     `json:"identity_actor_id,omitempty"`
	ProviderDeploymentId  *any     `json:"provider_deployment_id,omitempty"`
	ProviderConfigId      *any     `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId  *any     `json:"provider_auth_config_id,omitempty"`
	SessionTemplateId     *any     `json:"session_template_id,omitempty"`
	// CreatedAt - Filter integration instance creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration instance last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsInstancesEndpointCreateSessionTemplateBody contains the request body for CreateSessionTemplate.
type IntegrationsInstancesEndpointCreateSessionTemplateBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
}

// IntegrationsInstancesEndpointCreateSessionBody contains the request body for CreateSession.
type IntegrationsInstancesEndpointCreateSessionBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
}

// IntegrationsInstancesEndpointCreateBody contains the request body for Create.
type IntegrationsInstancesEndpointCreateBody struct {
	IntegrationId   string            `json:"integration_id"`
	Name            string            `json:"name"`
	Description     *string           `json:"description,omitempty"`
	Metadata        *map[string]any   `json:"metadata,omitempty"`
	IdentityActorId *string           `json:"identity_actor_id,omitempty"`
	IdentityId      *string           `json:"identity_id,omitempty"`
	Providers       *[]map[string]any `json:"providers,omitempty"`
}

// IntegrationsInstancesEndpointUpdateBody contains the request body for Update.
type IntegrationsInstancesEndpointUpdateBody struct {
	Name            *string           `json:"name,omitempty"`
	Description     *string           `json:"description,omitempty"`
	Metadata        *map[string]any   `json:"metadata,omitempty"`
	IdentityActorId *string           `json:"identity_actor_id,omitempty"`
	IdentityId      *string           `json:"identity_id,omitempty"`
	Providers       *[]map[string]any `json:"providers,omitempty"`
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

// Get retrieves a specific integration instance.
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

// CreateSessionTemplate creates or updates the shared session template for a specific integration instance.
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

// CreateSession creates a session from the shared session template of a specific integration instance.
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

// Create creates a new integration instance.
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

// Update updates a specific integration instance.
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

// Delete archives a specific integration instance.
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
