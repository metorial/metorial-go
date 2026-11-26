package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrationfirewallassignments"
)

// IntegrationFirewallAssignmentsEndpoint provides access to an integration firewall assignment is a firewall applied to one target. Targets are integration enclaves, providers, or the integration internet gateway.
type IntegrationFirewallAssignmentsEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationFirewallAssignmentsEndpoint creates a new IntegrationFirewallAssignmentsEndpoint.
func NewIntegrationFirewallAssignmentsEndpoint(client *endpoint.Client) *IntegrationFirewallAssignmentsEndpoint {
	return &IntegrationFirewallAssignmentsEndpoint{client: client}
}

// IntegrationFirewallAssignmentsEndpointListParams contains optional query parameters for List.
type IntegrationFirewallAssignmentsEndpointListParams struct {
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
	// Id - Only return integration firewall assignments with one of these IDs.
	Id *any `json:"id,omitempty"`
	// FirewallId - Only return assignments applying one of these firewalls.
	FirewallId *any `json:"firewall_id,omitempty"`
	// EnclaveId - Only return assignments targeting one of these enclaves.
	EnclaveId *any `json:"enclave_id,omitempty"`
	// ProviderId - Only return assignments targeting one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// InternetGatewayId - Only return assignments targeting one of these internet gateways.
	InternetGatewayId *any `json:"internet_gateway_id,omitempty"`
	// TargetType - Only return assignments with one of these target types.
	TargetType *any `json:"target_type,omitempty"`
	// CreatedAt - Only return results whose integration firewall assignment creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// IntegrationFirewallAssignmentsEndpointCreateBody contains the request body for Create.
type IntegrationFirewallAssignmentsEndpointCreateBody struct {
	// FirewallId - Firewall to apply to the target.
	FirewallId string `json:"firewall_id"`
	// TargetType - What the firewall applies to.
	TargetType string `json:"target_type"`
	// EnclaveId - Enclave to apply the firewall to. Required when `target_type` is `enclave`.
	EnclaveId *string `json:"enclave_id,omitempty"`
	// ProviderId - Provider to apply the firewall to. Required when `target_type` is `provider`.
	ProviderId *string `json:"provider_id,omitempty"`
}

// List returns a paginated list of integration firewall assignments in the instance.
func (e *IntegrationFirewallAssignmentsEndpoint) List(params *IntegrationFirewallAssignmentsEndpointListParams) (*integrationfirewallassignments.IntegrationFirewallAssignmentsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"integration-firewall-assignments"},
		Query: query,
	}
	var result integrationfirewallassignments.IntegrationFirewallAssignmentsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an integration firewall assignment.
func (e *IntegrationFirewallAssignmentsEndpoint) Get(integrationFirewallAssignmentId string) (*integrationfirewallassignments.IntegrationFirewallAssignmentsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-firewall-assignments", integrationFirewallAssignmentId},
	}
	var result integrationfirewallassignments.IntegrationFirewallAssignmentsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an integration firewall assignment.
func (e *IntegrationFirewallAssignmentsEndpoint) Create(body *IntegrationFirewallAssignmentsEndpointCreateBody) (*integrationfirewallassignments.IntegrationFirewallAssignmentsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-firewall-assignments"},
		Body: body,
	}
	var result integrationfirewallassignments.IntegrationFirewallAssignmentsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete permanently deletes an integration firewall assignment. The firewall stops applying to the target.
func (e *IntegrationFirewallAssignmentsEndpoint) Delete(integrationFirewallAssignmentId string) (*integrationfirewallassignments.IntegrationFirewallAssignmentsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-firewall-assignments", integrationFirewallAssignmentId},
	}
	var result integrationfirewallassignments.IntegrationFirewallAssignmentsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
