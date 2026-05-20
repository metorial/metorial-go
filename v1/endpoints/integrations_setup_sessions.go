package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrations/setupsessions"
)

// IntegrationsSetupSessionsEndpoint provides access to integration setup sessions orchestrate configuring every provider required by an integration instance.
type IntegrationsSetupSessionsEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationsSetupSessionsEndpoint creates a new IntegrationsSetupSessionsEndpoint.
func NewIntegrationsSetupSessionsEndpoint(client *endpoint.Client) *IntegrationsSetupSessionsEndpoint {
	return &IntegrationsSetupSessionsEndpoint{client: client}
}

// IntegrationsSetupSessionsEndpointListParams contains optional query parameters for List.
type IntegrationsSetupSessionsEndpointListParams struct {
	Limit                 *float64 `json:"limit,omitempty"`
	After                 *string  `json:"after,omitempty"`
	Before                *string  `json:"before,omitempty"`
	Cursor                *string  `json:"cursor,omitempty"`
	Order                 *string  `json:"order,omitempty"`
	Status                *any     `json:"status,omitempty"`
	Id                    *any     `json:"id,omitempty"`
	IntegrationId         *any     `json:"integration_id,omitempty"`
	IntegrationInstanceId *any     `json:"integration_instance_id,omitempty"`
	// CreatedAt - Filter integration setup session creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration setup session last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IntegrationsSetupSessionsEndpointCreateBody contains the request body for Create.
type IntegrationsSetupSessionsEndpointCreateBody struct {
	IntegrationId   string          `json:"integration_id"`
	Name            string          `json:"name"`
	Description     *string         `json:"description,omitempty"`
	Metadata        *map[string]any `json:"metadata,omitempty"`
	IdentityActorId *string         `json:"identity_actor_id,omitempty"`
	IdentityId      *string         `json:"identity_id,omitempty"`
	ExpiresAt       *string         `json:"expires_at,omitempty"`
	RedirectUrl     *string         `json:"redirect_url,omitempty"`
	Configuration   *map[string]any `json:"configuration,omitempty"`
}

// List returns a paginated list of integration setup sessions.
func (e *IntegrationsSetupSessionsEndpoint) List(params *IntegrationsSetupSessionsEndpointListParams) (*setupsessions.IntegrationsSetupSessionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"integration-setup-sessions"},
		Query: query,
	}
	var result setupsessions.IntegrationsSetupSessionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific integration setup session.
func (e *IntegrationsSetupSessionsEndpoint) Get(integrationSetupSessionId string) (*setupsessions.IntegrationsSetupSessionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-setup-sessions", integrationSetupSessionId},
	}
	var result setupsessions.IntegrationsSetupSessionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new integration setup session and draft integration instance.
func (e *IntegrationsSetupSessionsEndpoint) Create(body *IntegrationsSetupSessionsEndpointCreateBody) (*setupsessions.IntegrationsSetupSessionsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-setup-sessions"},
		Body: body,
	}
	var result setupsessions.IntegrationsSetupSessionsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
