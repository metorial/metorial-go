package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrationfirewallassignments"
)

// IntegrationFirewallAssignmentsEndpoint provides access to manage assignments that apply firewalls to integration enclaves, providers, or the integration internet gateway.
type IntegrationFirewallAssignmentsEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationFirewallAssignmentsEndpoint creates a new IntegrationFirewallAssignmentsEndpoint.
func NewIntegrationFirewallAssignmentsEndpoint(client *endpoint.Client) *IntegrationFirewallAssignmentsEndpoint {
	return &IntegrationFirewallAssignmentsEndpoint{client: client}
}

// IntegrationFirewallAssignmentsEndpointListParams contains optional query parameters for List.
type IntegrationFirewallAssignmentsEndpointListParams struct {
	Limit             *float64 `json:"limit,omitempty"`
	After             *string  `json:"after,omitempty"`
	Before            *string  `json:"before,omitempty"`
	Cursor            *string  `json:"cursor,omitempty"`
	Order             *string  `json:"order,omitempty"`
	Id                *any     `json:"id,omitempty"`
	FirewallId        *any     `json:"firewall_id,omitempty"`
	EnclaveId         *any     `json:"enclave_id,omitempty"`
	ProviderId        *any     `json:"provider_id,omitempty"`
	InternetGatewayId *any     `json:"internet_gateway_id,omitempty"`
	TargetType        *any     `json:"target_type,omitempty"`
	// CreatedAt - Filter integration firewall assignment creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// IntegrationFirewallAssignmentsEndpointCreateBody contains the request body for Create.
type IntegrationFirewallAssignmentsEndpointCreateBody struct {
	FirewallId string  `json:"firewall_id"`
	TargetType string  `json:"target_type"`
	EnclaveId  *string `json:"enclave_id,omitempty"`
	ProviderId *string `json:"provider_id,omitempty"`
}

// List returns a paginated list of integration firewall assignments.
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

// Get retrieves a specific integration firewall assignment by ID.
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

// Create creates an assignment that applies a firewall to a target.
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

// Delete deletes an integration firewall assignment.
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
