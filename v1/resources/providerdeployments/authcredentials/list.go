package authcredentials

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsAuthCredentialsListOutputItems represents the provider deployments auth credentials list output items type.
type ProviderDeploymentsAuthCredentialsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - Must be `oauth`.
	Type string `json:"type"`
	// Status - Archived once the credentials are archived. Archived credentials stay listable.
	Status string `json:"status"`
	// IsDefault - Used when a request doesn't name credentials.
	IsDefault bool `json:"is_default"`
	// IsManaged - Whether Metorial or the installation administrator manages these OAuth credentials.
	IsManaged bool `json:"is_managed"`
	// ManagedBy - Who manages these credentials: Metorial, or the administrator of this self-hosted installation. Null for credentials created in the project.
	ManagedBy *string `json:"managed_by,omitempty"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Scopes - OAuth scopes requested by this credential
	Scopes *[]string `json:"scopes,omitempty"`
	// ProviderId - Provider these credentials belong to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsAuthCredentialsListOutputPagination represents the provider deployments auth credentials list output pagination type.
type ProviderDeploymentsAuthCredentialsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderDeploymentsAuthCredentialsListOutput represents the provider deployments auth credentials list output type.
type ProviderDeploymentsAuthCredentialsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                                 `json:"object"`
	Items      []ProviderDeploymentsAuthCredentialsListOutputItems    `json:"items"`
	Pagination ProviderDeploymentsAuthCredentialsListOutputPagination `json:"pagination"`
}

// MapProviderDeploymentsAuthCredentialsListOutputFromJSON deserializes JSON data into a ProviderDeploymentsAuthCredentialsListOutput.
func MapProviderDeploymentsAuthCredentialsListOutputFromJSON(data []byte) (*ProviderDeploymentsAuthCredentialsListOutput, error) {
	var v ProviderDeploymentsAuthCredentialsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthCredentialsListOutputToJSON serializes a ProviderDeploymentsAuthCredentialsListOutput to JSON.
func MapProviderDeploymentsAuthCredentialsListOutputToJSON(v *ProviderDeploymentsAuthCredentialsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsAuthCredentialsListQueryCreatedAt - Only return results whose provider auth credentials creation time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsAuthCredentialsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsAuthCredentialsListQueryUpdatedAt - Only return results whose provider auth credentials last update time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsAuthCredentialsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsAuthCredentialsListQuery represents the provider deployments auth credentials list query type.
type ProviderDeploymentsAuthCredentialsListQuery struct {
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
	// Status - Filter by status (active, archived)
	Status *any `json:"status,omitempty"`
	// Id - Filter by credential ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderAuthMethodId - Filter by provider auth method ID(s)
	ProviderAuthMethodId *any `json:"provider_auth_method_id,omitempty"`
	// Origin - Filter by credential origin (custom, managed)
	Origin *any `json:"origin,omitempty"`
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider auth credentials creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProviderDeploymentsAuthCredentialsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider auth credentials last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ProviderDeploymentsAuthCredentialsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapProviderDeploymentsAuthCredentialsListQueryFromJSON deserializes JSON data into a ProviderDeploymentsAuthCredentialsListQuery.
func MapProviderDeploymentsAuthCredentialsListQueryFromJSON(data []byte) (*ProviderDeploymentsAuthCredentialsListQuery, error) {
	var v ProviderDeploymentsAuthCredentialsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthCredentialsListQueryToJSON serializes a ProviderDeploymentsAuthCredentialsListQuery to JSON.
func MapProviderDeploymentsAuthCredentialsListQueryToJSON(v *ProviderDeploymentsAuthCredentialsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
