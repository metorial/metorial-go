package integrationfirewallassignments

import (
	"encoding/json"
	"time"
)

// IntegrationFirewallAssignmentsListOutputItemsFirewall represents the integration firewall assignments list output items firewall type.
type IntegrationFirewallAssignmentsListOutputItemsFirewall struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
}

// IntegrationFirewallAssignmentsListOutputItemsTarget represents the integration firewall assignments list output items target type.
type IntegrationFirewallAssignmentsListOutputItemsTarget struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

// IntegrationFirewallAssignmentsListOutputItems represents the integration firewall assignments list output items type.
type IntegrationFirewallAssignmentsListOutputItems struct {
	Object     string                                                `json:"object"`
	Id         string                                                `json:"id"`
	TargetType string                                                `json:"target_type"`
	Firewall   IntegrationFirewallAssignmentsListOutputItemsFirewall `json:"firewall"`
	Target     *IntegrationFirewallAssignmentsListOutputItemsTarget  `json:"target,omitempty"`
	CreatedAt  time.Time                                             `json:"created_at"`
}

// IntegrationFirewallAssignmentsListOutputPagination represents the integration firewall assignments list output pagination type.
type IntegrationFirewallAssignmentsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// IntegrationFirewallAssignmentsListOutput represents the integration firewall assignments list output type.
type IntegrationFirewallAssignmentsListOutput struct {
	Items      []IntegrationFirewallAssignmentsListOutputItems    `json:"items"`
	Pagination IntegrationFirewallAssignmentsListOutputPagination `json:"pagination"`
}

// MapIntegrationFirewallAssignmentsListOutputFromJSON deserializes JSON data into a IntegrationFirewallAssignmentsListOutput.
func MapIntegrationFirewallAssignmentsListOutputFromJSON(data []byte) (*IntegrationFirewallAssignmentsListOutput, error) {
	var v IntegrationFirewallAssignmentsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationFirewallAssignmentsListOutputToJSON serializes a IntegrationFirewallAssignmentsListOutput to JSON.
func MapIntegrationFirewallAssignmentsListOutputToJSON(v *IntegrationFirewallAssignmentsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationFirewallAssignmentsListQueryCreatedAt - Filter integration firewall assignment creation time by date range
type IntegrationFirewallAssignmentsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for integration firewall assignment creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration firewall assignment creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationFirewallAssignmentsListQuery represents the integration firewall assignments list query type.
type IntegrationFirewallAssignmentsListQuery struct {
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
	CreatedAt *IntegrationFirewallAssignmentsListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapIntegrationFirewallAssignmentsListQueryFromJSON deserializes JSON data into a IntegrationFirewallAssignmentsListQuery.
func MapIntegrationFirewallAssignmentsListQueryFromJSON(data []byte) (*IntegrationFirewallAssignmentsListQuery, error) {
	var v IntegrationFirewallAssignmentsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationFirewallAssignmentsListQueryToJSON serializes a IntegrationFirewallAssignmentsListQuery to JSON.
func MapIntegrationFirewallAssignmentsListQueryToJSON(v *IntegrationFirewallAssignmentsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
