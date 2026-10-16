package integrationenclaves

import (
	"encoding/json"
	"time"
)

// IntegrationEnclavesListOutputItemsEnclaveEnvironment represents the integration enclaves list output items enclave environment type.
type IntegrationEnclavesListOutputItemsEnclaveEnvironment struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// IntegrationEnclavesListOutputItems represents the integration enclaves list output items type.
type IntegrationEnclavesListOutputItems struct {
	Object               string                                               `json:"object"`
	Id                   string                                               `json:"id"`
	Slug                 string                                               `json:"slug"`
	Name                 string                                               `json:"name"`
	Description          *string                                              `json:"description,omitempty"`
	InternetGatewayId    string                                               `json:"internet_gateway_id"`
	ProviderDeploymentId string                                               `json:"provider_deployment_id"`
	EnclaveEnvironment   IntegrationEnclavesListOutputItemsEnclaveEnvironment `json:"enclave_environment"`
	CreatedAt            time.Time                                            `json:"created_at"`
	LastUsedAt           *time.Time                                           `json:"last_used_at,omitempty"`
}

// IntegrationEnclavesListOutputPagination represents the integration enclaves list output pagination type.
type IntegrationEnclavesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// IntegrationEnclavesListOutput represents the integration enclaves list output type.
type IntegrationEnclavesListOutput struct {
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

// IntegrationEnclavesListQueryCreatedAt - Filter integration enclave creation time by date range
type IntegrationEnclavesListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for integration enclave creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration enclave creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationEnclavesListQuery represents the integration enclaves list query type.
type IntegrationEnclavesListQuery struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	Slug                 *any     `json:"slug,omitempty"`
	InternetGatewayId    *any     `json:"internet_gateway_id,omitempty"`
	ProviderDeploymentId *any     `json:"provider_deployment_id,omitempty"`
	ProviderId           *any     `json:"provider_id,omitempty"`
	FirewallId           *any     `json:"firewall_id,omitempty"`
	// CreatedAt - Filter integration enclave creation time by date range
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
