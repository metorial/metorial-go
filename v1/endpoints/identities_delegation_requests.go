package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/identities/delegationrequests"
)

// IdentitiesDelegationRequestsEndpoint provides access to an identity delegation request is an approval workflow for a delegation that needs consent. Approving it creates the delegation.
type IdentitiesDelegationRequestsEndpoint struct {
	client *endpoint.Client
}

// NewIdentitiesDelegationRequestsEndpoint creates a new IdentitiesDelegationRequestsEndpoint.
func NewIdentitiesDelegationRequestsEndpoint(client *endpoint.Client) *IdentitiesDelegationRequestsEndpoint {
	return &IdentitiesDelegationRequestsEndpoint{client: client}
}

// IdentitiesDelegationRequestsEndpointListParams contains optional query parameters for List.
type IdentitiesDelegationRequestsEndpointListParams struct {
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
	// Status - Only return identity delegation requests with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return identity delegation requests with one of these IDs.
	Id *any `json:"id,omitempty"`
	// ActorId - Only return identity delegation requests from one of these requesters.
	ActorId *any `json:"actor_id,omitempty"`
	// IdentityId - Only return identity delegation requests for one of these identities.
	IdentityId *any `json:"identity_id,omitempty"`
	// CreatedAt - Only return results whose identity delegation request creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose identity delegation request last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// IdentitiesDelegationRequestsEndpointGetParams contains optional query parameters for Get.
type IdentitiesDelegationRequestsEndpointGetParams struct {
	// AllowDeleted - Return the request even if it has been deleted.
	AllowDeleted *bool `json:"allow_deleted,omitempty"`
}

// IdentitiesDelegationRequestsEndpointCreateBody contains the request body for Create.
type IdentitiesDelegationRequestsEndpointCreateBody struct {
	// IdentityId - Identity to request delegation for.
	IdentityId string `json:"identity_id"`
	// RequesterActorId - Actor requesting the delegation.
	RequesterActorId string `json:"requester_actor_id"`
	// DelegatorActorId - Actor submitting the request on behalf of the requester.
	DelegatorActorId *string `json:"delegator_actor_id,omitempty"`
	// Permissions - Permissions being requested.
	Permissions *[]string `json:"permissions,omitempty"`
	// ExpiresAt - Deadline for approving or denying the request.
	ExpiresAt string `json:"expires_at"`
	// DelegationConfigId - Delegation config to use for the resulting delegation.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// CredentialOverrides - Optional per-credential permission overrides.
	CredentialOverrides *[]map[string]any `json:"credential_overrides,omitempty"`
	// Note - Optional human-readable note for the request.
	Note *string `json:"note,omitempty"`
	// Metadata - Custom data to identify or organize the identity delegation request.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// IdentitiesDelegationRequestsEndpointApproveParams contains optional query parameters for Approve.
type IdentitiesDelegationRequestsEndpointApproveParams struct {
	// AllowDeleted - Allow approving a request that is already deleted.
	AllowDeleted *bool `json:"allow_deleted,omitempty"`
}

// IdentitiesDelegationRequestsEndpointDenyParams contains optional query parameters for Deny.
type IdentitiesDelegationRequestsEndpointDenyParams struct {
	// AllowDeleted - Allow denying a request that is already deleted.
	AllowDeleted *bool `json:"allow_deleted,omitempty"`
}

// List returns a paginated list of identity delegation requests.
func (e *IdentitiesDelegationRequestsEndpoint) List(params *IdentitiesDelegationRequestsEndpointListParams) (*delegationrequests.IdentitiesDelegationRequestsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"identity-delegation-requests"},
		Query: query,
	}
	var result delegationrequests.IdentitiesDelegationRequestsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an identity delegation request.
func (e *IdentitiesDelegationRequestsEndpoint) Get(identityDelegationRequestId string, params *IdentitiesDelegationRequestsEndpointGetParams) (*delegationrequests.IdentitiesDelegationRequestsGetOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"identity-delegation-requests", identityDelegationRequestId},
		Query: query,
	}
	var result delegationrequests.IdentitiesDelegationRequestsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an identity delegation request.
func (e *IdentitiesDelegationRequestsEndpoint) Create(body *IdentitiesDelegationRequestsEndpointCreateBody) (*delegationrequests.IdentitiesDelegationRequestsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"identity-delegation-requests"},
		Body: body,
	}
	var result delegationrequests.IdentitiesDelegationRequestsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Approve approves an identity delegation request.
func (e *IdentitiesDelegationRequestsEndpoint) Approve(identityDelegationRequestId string, params *IdentitiesDelegationRequestsEndpointApproveParams) (*delegationrequests.IdentitiesDelegationRequestsApproveOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"identity-delegation-requests", identityDelegationRequestId, "approve"},
		Query: query,
	}
	var result delegationrequests.IdentitiesDelegationRequestsApproveOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Deny denies an identity delegation request.
func (e *IdentitiesDelegationRequestsEndpoint) Deny(identityDelegationRequestId string, params *IdentitiesDelegationRequestsEndpointDenyParams) (*delegationrequests.IdentitiesDelegationRequestsDenyOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"identity-delegation-requests", identityDelegationRequestId, "deny"},
		Query: query,
	}
	var result delegationrequests.IdentitiesDelegationRequestsDenyOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
