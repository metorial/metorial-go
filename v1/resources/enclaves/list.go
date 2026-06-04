package enclaves

import (
	"encoding/json"
	"time"
)

// EnclavesListOutputItemsEnclaveEnvironment represents the enclaves list output items enclave environment type.
type EnclavesListOutputItemsEnclaveEnvironment struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// EnclavesListOutputItems represents the enclaves list output items type.
type EnclavesListOutputItems struct {
	Object               string                                    `json:"object"`
	Id                   string                                    `json:"id"`
	Slug                 string                                    `json:"slug"`
	Name                 string                                    `json:"name"`
	Description          *string                                   `json:"description,omitempty"`
	NetworkId            string                                    `json:"network_id"`
	ProviderDeploymentId string                                    `json:"provider_deployment_id"`
	EnclaveEnvironment   EnclavesListOutputItemsEnclaveEnvironment `json:"enclave_environment"`
	CreatedAt            time.Time                                 `json:"created_at"`
	LastUsedAt           *time.Time                                `json:"last_used_at,omitempty"`
}

// EnclavesListOutputPagination represents the enclaves list output pagination type.
type EnclavesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// EnclavesListOutput represents the enclaves list output type.
type EnclavesListOutput struct {
	Items      []EnclavesListOutputItems    `json:"items"`
	Pagination EnclavesListOutputPagination `json:"pagination"`
}

// MapEnclavesListOutputFromJSON deserializes JSON data into a EnclavesListOutput.
func MapEnclavesListOutputFromJSON(data []byte) (*EnclavesListOutput, error) {
	var v EnclavesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapEnclavesListOutputToJSON serializes a EnclavesListOutput to JSON.
func MapEnclavesListOutputToJSON(v *EnclavesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// EnclavesListQueryCreatedAt - Filter enclave creation time by date range
type EnclavesListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for enclave creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for enclave creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// EnclavesListQuery represents the enclaves list query type.
type EnclavesListQuery struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	Slug                 *any     `json:"slug,omitempty"`
	NetworkId            *any     `json:"network_id,omitempty"`
	ProviderDeploymentId *any     `json:"provider_deployment_id,omitempty"`
	ProviderId           *any     `json:"provider_id,omitempty"`
	FirewallId           *any     `json:"firewall_id,omitempty"`
	// CreatedAt - Filter enclave creation time by date range
	CreatedAt *EnclavesListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapEnclavesListQueryFromJSON deserializes JSON data into a EnclavesListQuery.
func MapEnclavesListQueryFromJSON(data []byte) (*EnclavesListQuery, error) {
	var v EnclavesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapEnclavesListQueryToJSON serializes a EnclavesListQuery to JSON.
func MapEnclavesListQueryToJSON(v *EnclavesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
