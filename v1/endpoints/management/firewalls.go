package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/firewalls"
)

// FirewallsEndpoint provides access to a firewall is a set of network policies applied to targets. Firewalls apply to integration enclaves, providers, and sandbox machines.
type FirewallsEndpoint struct {
	client *endpoint.Client
}

// NewFirewallsEndpoint creates a new FirewallsEndpoint.
func NewFirewallsEndpoint(client *endpoint.Client) *FirewallsEndpoint {
	return &FirewallsEndpoint{client: client}
}

// FirewallsEndpointListParams contains optional query parameters for List.
type FirewallsEndpointListParams struct {
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
	// Search - Filter firewalls by name or description.
	Search *string `json:"search,omitempty"`
	// Id - Only return firewalls with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Slug - Only return firewalls with one of these slugs.
	Slug *any `json:"slug,omitempty"`
	// Status - Only return firewalls with one of these statuses.
	Status *any `json:"status,omitempty"`
	// EnclaveId - Only return firewalls applied to one of these enclaves.
	EnclaveId *any `json:"enclave_id,omitempty"`
	// ProviderId - Only return firewalls applied to one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// NetworkPolicyId - Only return firewalls with one of these network policies attached.
	NetworkPolicyId *any `json:"network_policy_id,omitempty"`
	// MachineId - Only return firewalls applied to one of these sandbox machines.
	MachineId *any `json:"machine_id,omitempty"`
	// CreatedAt - Only return results whose firewall creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose firewall last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// FirewallsEndpointCreateBody contains the request body for Create.
type FirewallsEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name, unique within the instance. Generated from the name when omitted.
	Slug *string `json:"slug,omitempty"`
	// NetworkPolicyIds - Network policies to attach to the firewall, in order.
	NetworkPolicyIds *[]string `json:"network_policy_ids,omitempty"`
}

// FirewallsEndpointUpdateBody contains the request body for Update.
type FirewallsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name, unique within the instance.
	Slug *string `json:"slug,omitempty"`
	// NetworkPolicyIds - Network policies attached to the firewall, in order. Replaces existing attachments.
	NetworkPolicyIds *[]string `json:"network_policy_ids,omitempty"`
}

// List returns a paginated list of firewalls in the instance.
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

// Get retrieves a firewall.
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

// Create creates a firewall.
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

// Update updates a firewall. Omitted fields keep their current value.
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

// Delete archives a firewall. It stays listable with `status=archived`.
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
