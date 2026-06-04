package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/networkpolicies"
)

// NetworkPoliciesEndpoint provides access to manage reusable network policy definitions and their rules.
type NetworkPoliciesEndpoint struct {
	client *endpoint.Client
}

// NewNetworkPoliciesEndpoint creates a new NetworkPoliciesEndpoint.
func NewNetworkPoliciesEndpoint(client *endpoint.Client) *NetworkPoliciesEndpoint {
	return &NetworkPoliciesEndpoint{client: client}
}

// NetworkPoliciesEndpointListParams contains optional query parameters for List.
type NetworkPoliciesEndpointListParams struct {
	Limit      *float64 `json:"limit,omitempty"`
	After      *string  `json:"after,omitempty"`
	Before     *string  `json:"before,omitempty"`
	Cursor     *string  `json:"cursor,omitempty"`
	Order      *string  `json:"order,omitempty"`
	Id         *any     `json:"id,omitempty"`
	Status     *any     `json:"status,omitempty"`
	FirewallId *any     `json:"firewall_id,omitempty"`
	Search     *string  `json:"search,omitempty"`
	// CreatedAt - Filter network policy creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter network policy last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// NetworkPoliciesEndpointCreateBody contains the request body for Create.
type NetworkPoliciesEndpointCreateBody struct {
	Name        string            `json:"name"`
	Description *string           `json:"description,omitempty"`
	Rules       *[]map[string]any `json:"rules,omitempty"`
}

// NetworkPoliciesEndpointUpdateBody contains the request body for Update.
type NetworkPoliciesEndpointUpdateBody struct {
	Name        *string           `json:"name,omitempty"`
	Description *string           `json:"description,omitempty"`
	Rules       *[]map[string]any `json:"rules,omitempty"`
}

// List returns a paginated list of network policies.
func (e *NetworkPoliciesEndpoint) List(instanceId string, params *NetworkPoliciesEndpointListParams) (*networkpolicies.NetworkPoliciesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "network-policies"},
		Query: query,
	}
	var result networkpolicies.NetworkPoliciesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific network policy by ID.
func (e *NetworkPoliciesEndpoint) Get(instanceId string, networkPolicyId string) (*networkpolicies.NetworkPoliciesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "network-policies", networkPolicyId},
	}
	var result networkpolicies.NetworkPoliciesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new network policy.
func (e *NetworkPoliciesEndpoint) Create(instanceId string, body *NetworkPoliciesEndpointCreateBody) (*networkpolicies.NetworkPoliciesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "network-policies"},
		Body: body,
	}
	var result networkpolicies.NetworkPoliciesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a network policy definition.
func (e *NetworkPoliciesEndpoint) Update(instanceId string, networkPolicyId string, body *NetworkPoliciesEndpointUpdateBody) (*networkpolicies.NetworkPoliciesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "network-policies", networkPolicyId},
		Body: body,
	}
	var result networkpolicies.NetworkPoliciesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a network policy.
func (e *NetworkPoliciesEndpoint) Delete(instanceId string, networkPolicyId string) (*networkpolicies.NetworkPoliciesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "network-policies", networkPolicyId},
	}
	var result networkpolicies.NetworkPoliciesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
