package providers

import (
	"encoding/json"
	"time"
)

// SessionTemplatesProvidersListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionTemplatesProvidersListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionTemplatesProvidersListOutputItemsDeployment represents the session templates providers list output items deployment type.
type SessionTemplatesProvidersListOutputItemsDeployment struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Used when a request doesn't name a deployment.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this deployment belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionTemplatesProvidersListOutputItemsConfig represents the session templates providers list output items config type.
type SessionTemplatesProvidersListOutputItemsConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - The deployment falls back to this config when none is specified.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this config belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionTemplatesProvidersListOutputItemsAuthConfig - Auth config the provider runs with. `null` when there is none.
type SessionTemplatesProvidersListOutputItemsAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionTemplatesProvidersListOutputItems represents the session templates providers list output items type.
type SessionTemplatesProvidersListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the template provider is removed. Removed template providers stay listable.
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionTemplatesProvidersListOutputItemsToolFilter `json:"tool_filter"`
	// ProviderId - Provider this template provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionTemplateId - Template this provider belongs to.
	SessionTemplateId string                                             `json:"session_template_id"`
	Deployment        SessionTemplatesProvidersListOutputItemsDeployment `json:"deployment"`
	Config            SessionTemplatesProvidersListOutputItemsConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *SessionTemplatesProvidersListOutputItemsAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionTemplatesProvidersListOutputPagination represents the session templates providers list output pagination type.
type SessionTemplatesProvidersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SessionTemplatesProvidersListOutput represents the session templates providers list output type.
type SessionTemplatesProvidersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                        `json:"object"`
	Items      []SessionTemplatesProvidersListOutputItems    `json:"items"`
	Pagination SessionTemplatesProvidersListOutputPagination `json:"pagination"`
}

// MapSessionTemplatesProvidersListOutputFromJSON deserializes JSON data into a SessionTemplatesProvidersListOutput.
func MapSessionTemplatesProvidersListOutputFromJSON(data []byte) (*SessionTemplatesProvidersListOutput, error) {
	var v SessionTemplatesProvidersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesProvidersListOutputToJSON serializes a SessionTemplatesProvidersListOutput to JSON.
func MapSessionTemplatesProvidersListOutputToJSON(v *SessionTemplatesProvidersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionTemplatesProvidersListQueryCreatedAt - Only return results whose session template provider creation time is in this range. Set `gt`, `lt`, or both.
type SessionTemplatesProvidersListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionTemplatesProvidersListQueryUpdatedAt - Only return results whose session template provider last update time is in this range. Set `gt`, `lt`, or both.
type SessionTemplatesProvidersListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionTemplatesProvidersListQuery represents the session templates providers list query type.
type SessionTemplatesProvidersListQuery struct {
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
	// Status - Only return session template providers with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Filter by session template provider ID(s)
	Id *any `json:"id,omitempty"`
	// SessionTemplateId - Filter by session template ID(s)
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter by provider config ID(s)
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Filter by provider auth config ID(s)
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Only return results whose session template provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SessionTemplatesProvidersListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session template provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SessionTemplatesProvidersListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSessionTemplatesProvidersListQueryFromJSON deserializes JSON data into a SessionTemplatesProvidersListQuery.
func MapSessionTemplatesProvidersListQueryFromJSON(data []byte) (*SessionTemplatesProvidersListQuery, error) {
	var v SessionTemplatesProvidersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesProvidersListQueryToJSON serializes a SessionTemplatesProvidersListQuery to JSON.
func MapSessionTemplatesProvidersListQueryToJSON(v *SessionTemplatesProvidersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
