package integrationfirewallassignments

import (
	"encoding/json"
	"time"
)

// IntegrationFirewallAssignmentsListOutputItemsFirewall - Firewall applied to the target.
type IntegrationFirewallAssignmentsListOutputItemsFirewall struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
}

// IntegrationFirewallAssignmentsListOutputItemsTarget - Target the firewall applies to. `null` when the target no longer exists.
type IntegrationFirewallAssignmentsListOutputItemsTarget struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Kind of target the firewall applies to.
	Type string `json:"type"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
}

// IntegrationFirewallAssignmentsListOutputItems represents the integration firewall assignments list output items type.
type IntegrationFirewallAssignmentsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// TargetType - What the firewall applies to.
	TargetType string `json:"target_type"`
	// Firewall - Firewall applied to the target.
	Firewall IntegrationFirewallAssignmentsListOutputItemsFirewall `json:"firewall"`
	// Target - Target the firewall applies to. `null` when the target no longer exists.
	Target *IntegrationFirewallAssignmentsListOutputItemsTarget `json:"target,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IntegrationFirewallAssignmentsListOutputPagination represents the integration firewall assignments list output pagination type.
type IntegrationFirewallAssignmentsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IntegrationFirewallAssignmentsListOutput represents the integration firewall assignments list output type.
type IntegrationFirewallAssignmentsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                             `json:"object"`
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

// IntegrationFirewallAssignmentsListQueryCreatedAt - Only return results whose integration firewall assignment creation time is in this range. Set `gt`, `lt`, or both.
type IntegrationFirewallAssignmentsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationFirewallAssignmentsListQuery represents the integration firewall assignments list query type.
type IntegrationFirewallAssignmentsListQuery struct {
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
