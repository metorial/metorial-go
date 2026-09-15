package providerspecificationchangenotifications

import (
	"encoding/json"
	"time"
)

// ProviderSpecificationChangeNotificationsListOutputItemsFromSpecification represents the provider specification change notifications list output items from specification type.
type ProviderSpecificationChangeNotificationsListOutputItemsFromSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsListOutputItemsToSpecification represents the provider specification change notifications list output items to specification type.
type ProviderSpecificationChangeNotificationsListOutputItemsToSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsListOutputItemsFromProviderVersion represents the provider specification change notifications list output items from provider version type.
type ProviderSpecificationChangeNotificationsListOutputItemsFromProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsListOutputItemsToProviderVersion represents the provider specification change notifications list output items to provider version type.
type ProviderSpecificationChangeNotificationsListOutputItemsToProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsListOutputItems represents the provider specification change notifications list output items type.
type ProviderSpecificationChangeNotificationsListOutputItems struct {
	Object              string                                                                      `json:"object"`
	Id                  string                                                                      `json:"id"`
	ProviderId          string                                                                      `json:"provider_id"`
	ProviderVersionId   string                                                                      `json:"provider_version_id"`
	FromSpecification   *ProviderSpecificationChangeNotificationsListOutputItemsFromSpecification   `json:"from_specification,omitempty"`
	ToSpecification     *ProviderSpecificationChangeNotificationsListOutputItemsToSpecification     `json:"to_specification,omitempty"`
	FromProviderVersion *ProviderSpecificationChangeNotificationsListOutputItemsFromProviderVersion `json:"from_provider_version,omitempty"`
	ToProviderVersion   *ProviderSpecificationChangeNotificationsListOutputItemsToProviderVersion   `json:"to_provider_version,omitempty"`
	CreatedAt           time.Time                                                                   `json:"created_at"`
}

// ProviderSpecificationChangeNotificationsListOutputPagination represents the provider specification change notifications list output pagination type.
type ProviderSpecificationChangeNotificationsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ProviderSpecificationChangeNotificationsListOutput represents the provider specification change notifications list output type.
type ProviderSpecificationChangeNotificationsListOutput struct {
	Items      []ProviderSpecificationChangeNotificationsListOutputItems    `json:"items"`
	Pagination ProviderSpecificationChangeNotificationsListOutputPagination `json:"pagination"`
}

// MapProviderSpecificationChangeNotificationsListOutputFromJSON deserializes JSON data into a ProviderSpecificationChangeNotificationsListOutput.
func MapProviderSpecificationChangeNotificationsListOutputFromJSON(data []byte) (*ProviderSpecificationChangeNotificationsListOutput, error) {
	var v ProviderSpecificationChangeNotificationsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderSpecificationChangeNotificationsListOutputToJSON serializes a ProviderSpecificationChangeNotificationsListOutput to JSON.
func MapProviderSpecificationChangeNotificationsListOutputToJSON(v *ProviderSpecificationChangeNotificationsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderSpecificationChangeNotificationsListQueryCreatedAt - Filter provider specification change notification time by date range
type ProviderSpecificationChangeNotificationsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for provider specification change notification time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for provider specification change notification time
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderSpecificationChangeNotificationsListQuery represents the provider specification change notifications list query type.
type ProviderSpecificationChangeNotificationsListQuery struct {
	Limit                   *float64 `json:"limit,omitempty"`
	After                   *string  `json:"after,omitempty"`
	Before                  *string  `json:"before,omitempty"`
	Cursor                  *string  `json:"cursor,omitempty"`
	Order                   *string  `json:"order,omitempty"`
	Id                      *any     `json:"id,omitempty"`
	Target                  *any     `json:"target,omitempty"`
	ProviderId              *any     `json:"provider_id,omitempty"`
	ProviderVersionId       *any     `json:"provider_version_id,omitempty"`
	ProviderSpecificationId *any     `json:"provider_specification_id,omitempty"`
	// CreatedAt - Filter provider specification change notification time by date range
	CreatedAt *ProviderSpecificationChangeNotificationsListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapProviderSpecificationChangeNotificationsListQueryFromJSON deserializes JSON data into a ProviderSpecificationChangeNotificationsListQuery.
func MapProviderSpecificationChangeNotificationsListQueryFromJSON(data []byte) (*ProviderSpecificationChangeNotificationsListQuery, error) {
	var v ProviderSpecificationChangeNotificationsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderSpecificationChangeNotificationsListQueryToJSON serializes a ProviderSpecificationChangeNotificationsListQuery to JSON.
func MapProviderSpecificationChangeNotificationsListQueryToJSON(v *ProviderSpecificationChangeNotificationsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
