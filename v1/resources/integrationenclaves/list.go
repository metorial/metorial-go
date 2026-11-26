package integrationenclaves

import (
	"encoding/json"
	"time"
)

// IntegrationEnclavesListOutputItemsEnclaveEnvironment - Environment the enclave runs in.
type IntegrationEnclavesListOutputItemsEnclaveEnvironment struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Type - `metorial` for Metorial-hosted environments. `outpost` for environments in your own infrastructure.
	Type string `json:"type"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IntegrationEnclavesListOutputItems represents the integration enclaves list output items type.
type IntegrationEnclavesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// InternetGatewayId - Internet gateway the enclave uses.
	InternetGatewayId string `json:"internet_gateway_id"`
	// ProviderDeploymentId - Provider deployment running in the enclave.
	ProviderDeploymentId string `json:"provider_deployment_id"`
	// EnclaveEnvironment - Environment the enclave runs in.
	EnclaveEnvironment IntegrationEnclavesListOutputItemsEnclaveEnvironment `json:"enclave_environment"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// LastUsedAt - When the enclave was last used. `null` when never used.
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
}

// IntegrationEnclavesListOutputPagination represents the integration enclaves list output pagination type.
type IntegrationEnclavesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IntegrationEnclavesListOutput represents the integration enclaves list output type.
type IntegrationEnclavesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                  `json:"object"`
	Items      []IntegrationEnclavesListOutputItems    `json:"items"`
	Pagination IntegrationEnclavesListOutputPagination `json:"pagination"`
}

// MapIntegrationEnclavesListOutputFromJSON deserializes JSON data into a IntegrationEnclavesListOutput.
func MapIntegrationEnclavesListOutputFromJSON(data []byte) (*IntegrationEnclavesListOutput, error) {
	var v IntegrationEnclavesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationEnclavesListOutputToJSON serializes a IntegrationEnclavesListOutput to JSON.
func MapIntegrationEnclavesListOutputToJSON(v *IntegrationEnclavesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationEnclavesListQueryCreatedAt - Only return results whose integration enclave creation time is in this range. Set `gt`, `lt`, or both.
type IntegrationEnclavesListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationEnclavesListQuery represents the integration enclaves list query type.
type IntegrationEnclavesListQuery struct {
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
	// Id - Only return integration enclaves with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Slug - Only return integration enclaves with one of these slugs.
	Slug *any `json:"slug,omitempty"`
	// InternetGatewayId - Only return integration enclaves using one of these internet gateways.
	InternetGatewayId *any `json:"internet_gateway_id,omitempty"`
	// ProviderDeploymentId - Only return integration enclaves running one of these provider deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderId - Only return integration enclaves running one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// FirewallId - Only return integration enclaves protected by one of these firewalls.
	FirewallId *any `json:"firewall_id,omitempty"`
	// CreatedAt - Only return results whose integration enclave creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IntegrationEnclavesListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapIntegrationEnclavesListQueryFromJSON deserializes JSON data into a IntegrationEnclavesListQuery.
func MapIntegrationEnclavesListQueryFromJSON(data []byte) (*IntegrationEnclavesListQuery, error) {
	var v IntegrationEnclavesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationEnclavesListQueryToJSON serializes a IntegrationEnclavesListQuery to JSON.
func MapIntegrationEnclavesListQueryToJSON(v *IntegrationEnclavesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
