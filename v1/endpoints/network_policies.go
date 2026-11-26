package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/networkpolicies"
)

// NetworkPoliciesEndpoint provides access to a network policy is a reusable set of allow and deny rules. Firewalls attach policies to enforce them.
type NetworkPoliciesEndpoint struct {
	client *endpoint.Client
}

// NewNetworkPoliciesEndpoint creates a new NetworkPoliciesEndpoint.
func NewNetworkPoliciesEndpoint(client *endpoint.Client) *NetworkPoliciesEndpoint {
	return &NetworkPoliciesEndpoint{client: client}
}

// NetworkPoliciesEndpointListParams contains optional query parameters for List.
type NetworkPoliciesEndpointListParams struct {
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
	// Id - Only return network policies with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Status - Only return network policies with one of these statuses.
	Status *any `json:"status,omitempty"`
	// FirewallId - Only return network policies attached to one of these firewalls.
	FirewallId *any `json:"firewall_id,omitempty"`
	// Search - Filter network policies by name or description.
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose network policy creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose network policy last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// NetworkPoliciesEndpointCreateBody contains the request body for Create.
type NetworkPoliciesEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Rules - Rules to create with the network policy.
	Rules *[]map[string]any `json:"rules,omitempty"`
}

// NetworkPoliciesEndpointUpdateBody contains the request body for Update.
type NetworkPoliciesEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Rules - Rules in the network policy. Replaces existing rules.
	Rules *[]map[string]any `json:"rules,omitempty"`
}

// List returns a paginated list of network policies in the instance.
func (e *NetworkPoliciesEndpoint) List(params *NetworkPoliciesEndpointListParams) (*networkpolicies.NetworkPoliciesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"network-policies"},
		Query: query,
	}
	var result networkpolicies.NetworkPoliciesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a network policy.
func (e *NetworkPoliciesEndpoint) Get(networkPolicyId string) (*networkpolicies.NetworkPoliciesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"network-policies", networkPolicyId},
	}
	var result networkpolicies.NetworkPoliciesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a network policy.
func (e *NetworkPoliciesEndpoint) Create(body *NetworkPoliciesEndpointCreateBody) (*networkpolicies.NetworkPoliciesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"network-policies"},
		Body: body,
	}
	var result networkpolicies.NetworkPoliciesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a network policy. Omitted fields keep their current value.
func (e *NetworkPoliciesEndpoint) Update(networkPolicyId string, body *NetworkPoliciesEndpointUpdateBody) (*networkpolicies.NetworkPoliciesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"network-policies", networkPolicyId},
		Body: body,
	}
	var result networkpolicies.NetworkPoliciesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a network policy. It stays listable with `status=archived`.
func (e *NetworkPoliciesEndpoint) Delete(networkPolicyId string) (*networkpolicies.NetworkPoliciesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"network-policies", networkPolicyId},
	}
	var result networkpolicies.NetworkPoliciesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
