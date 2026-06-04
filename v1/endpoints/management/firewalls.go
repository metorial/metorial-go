package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/firewalls"
)

// FirewallsEndpoint provides access to manage firewalls and their attached network policies.
type FirewallsEndpoint struct {
	client *endpoint.Client
}

// NewFirewallsEndpoint creates a new FirewallsEndpoint.
func NewFirewallsEndpoint(client *endpoint.Client) *FirewallsEndpoint {
	return &FirewallsEndpoint{client: client}
}

// FirewallsEndpointListParams contains optional query parameters for List.
type FirewallsEndpointListParams struct {
	Limit           *float64 `json:"limit,omitempty"`
	After           *string  `json:"after,omitempty"`
	Before          *string  `json:"before,omitempty"`
	Cursor          *string  `json:"cursor,omitempty"`
	Order           *string  `json:"order,omitempty"`
	Id              *any     `json:"id,omitempty"`
	Slug            *any     `json:"slug,omitempty"`
	Status          *any     `json:"status,omitempty"`
	NetworkId       *any     `json:"network_id,omitempty"`
	EnclaveId       *any     `json:"enclave_id,omitempty"`
	ProviderId      *any     `json:"provider_id,omitempty"`
	NetworkPolicyId *any     `json:"network_policy_id,omitempty"`
	// CreatedAt - Filter firewall creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter firewall last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// FirewallsEndpointCreateBody contains the request body for Create.
type FirewallsEndpointCreateBody struct {
	Name             string            `json:"name"`
	Description      *string           `json:"description,omitempty"`
	Slug             *string           `json:"slug,omitempty"`
	NetworkId        string            `json:"network_id"`
	Bindings         *[]map[string]any `json:"bindings,omitempty"`
	NetworkPolicyIds *[]string         `json:"network_policy_ids,omitempty"`
}

// FirewallsEndpointUpdateBody contains the request body for Update.
type FirewallsEndpointUpdateBody struct {
	Name             *string   `json:"name,omitempty"`
	Description      *string   `json:"description,omitempty"`
	Slug             *string   `json:"slug,omitempty"`
	NetworkPolicyIds *[]string `json:"network_policy_ids,omitempty"`
}

// List returns a paginated list of firewalls.
func (e *FirewallsEndpoint) List(instanceId string, params *FirewallsEndpointListParams) (*firewalls.FirewallsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "firewalls"},
		Query: query,
	}
	var result firewalls.FirewallsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific firewall by ID.
func (e *FirewallsEndpoint) Get(instanceId string, firewallId string) (*firewalls.FirewallsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "firewalls", firewallId},
	}
	var result firewalls.FirewallsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new firewall.
func (e *FirewallsEndpoint) Create(instanceId string, body *FirewallsEndpointCreateBody) (*firewalls.FirewallsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "firewalls"},
		Body: body,
	}
	var result firewalls.FirewallsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a firewall definition.
func (e *FirewallsEndpoint) Update(instanceId string, firewallId string, body *FirewallsEndpointUpdateBody) (*firewalls.FirewallsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "firewalls", firewallId},
		Body: body,
	}
	var result firewalls.FirewallsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a firewall.
func (e *FirewallsEndpoint) Delete(instanceId string, firewallId string) (*firewalls.FirewallsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "firewalls", firewallId},
	}
	var result firewalls.FirewallsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
