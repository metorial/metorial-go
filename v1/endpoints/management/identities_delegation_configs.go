package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/identities/delegationconfigs"
)

// IdentitiesDelegationConfigsEndpoint provides access to an identity delegation config is the default policy for sub-delegation. It controls whether delegates can delegate further, and how deep.
type IdentitiesDelegationConfigsEndpoint struct {
	client *endpoint.Client
}

// NewIdentitiesDelegationConfigsEndpoint creates a new IdentitiesDelegationConfigsEndpoint.
func NewIdentitiesDelegationConfigsEndpoint(client *endpoint.Client) *IdentitiesDelegationConfigsEndpoint {
	return &IdentitiesDelegationConfigsEndpoint{client: client}
}

// IdentitiesDelegationConfigsEndpointListParams contains optional query parameters for List.
type IdentitiesDelegationConfigsEndpointListParams struct {
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
	// Search - Filter configs by name or description.
	Search *string `json:"search,omitempty"`
	// Status - Only return identity delegation configs with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return identity delegation configs with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CreatedAt - Only return results whose identity delegation config creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose identity delegation config last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IdentitiesDelegationConfigsEndpointCreateBody contains the request body for Create.
type IdentitiesDelegationConfigsEndpointCreateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the identity delegation config.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SubDelegationBehavior - Whether delegates can delegate further. `require_consent` asks the owner for approval first.
	SubDelegationBehavior string `json:"sub_delegation_behavior"`
	// SubDelegationDepth - Maximum sub-delegation depth. `0` blocks sub-delegation.
	SubDelegationDepth *float64 `json:"sub_delegation_depth,omitempty"`
}

// IdentitiesDelegationConfigsEndpointUpdateBody contains the request body for Update.
type IdentitiesDelegationConfigsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the identity delegation config. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SubDelegationBehavior - Whether delegates can delegate further. `require_consent` asks the owner for approval first.
	SubDelegationBehavior *string `json:"sub_delegation_behavior,omitempty"`
	// SubDelegationDepth - Maximum sub-delegation depth. `0` blocks sub-delegation.
	SubDelegationDepth *float64 `json:"sub_delegation_depth,omitempty"`
}

// List returns a paginated list of identity delegation configs.
func (e *IdentitiesDelegationConfigsEndpoint) List(instanceId string, params *IdentitiesDelegationConfigsEndpointListParams) (*delegationconfigs.IdentitiesDelegationConfigsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "identity-delegation-configs"},
		Query: query,
	}
	var result delegationconfigs.IdentitiesDelegationConfigsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an identity delegation config.
func (e *IdentitiesDelegationConfigsEndpoint) Get(instanceId string, identityDelegationConfigId string) (*delegationconfigs.IdentitiesDelegationConfigsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "identity-delegation-configs", identityDelegationConfigId},
	}
	var result delegationconfigs.IdentitiesDelegationConfigsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an identity delegation config.
func (e *IdentitiesDelegationConfigsEndpoint) Create(instanceId string, body *IdentitiesDelegationConfigsEndpointCreateBody) (*delegationconfigs.IdentitiesDelegationConfigsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "identity-delegation-configs"},
		Body: body,
	}
	var result delegationconfigs.IdentitiesDelegationConfigsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an identity delegation config. Omitted fields keep their current value.
func (e *IdentitiesDelegationConfigsEndpoint) Update(instanceId string, identityDelegationConfigId string, body *IdentitiesDelegationConfigsEndpointUpdateBody) (*delegationconfigs.IdentitiesDelegationConfigsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "identity-delegation-configs", identityDelegationConfigId},
		Body: body,
	}
	var result delegationconfigs.IdentitiesDelegationConfigsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives an identity delegation config. It stays listable with `status=archived`.
func (e *IdentitiesDelegationConfigsEndpoint) Delete(instanceId string, identityDelegationConfigId string) (*delegationconfigs.IdentitiesDelegationConfigsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "identity-delegation-configs", identityDelegationConfigId},
	}
	var result delegationconfigs.IdentitiesDelegationConfigsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
