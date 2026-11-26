package providers

import (
	"encoding/json"
	"time"
)

// SessionsProvidersListOutputItemsUsage - Usage statistics
type SessionsProvidersListOutputItemsUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsProvidersListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionsProvidersListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionsProvidersListOutputItemsDeployment represents the sessions providers list output items deployment type.
type SessionsProvidersListOutputItemsDeployment struct {
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

// SessionsProvidersListOutputItemsConfig represents the sessions providers list output items config type.
type SessionsProvidersListOutputItemsConfig struct {
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

// SessionsProvidersListOutputItemsAuthConfig - Auth config the provider runs with. `null` when there is none.
type SessionsProvidersListOutputItemsAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionsProvidersListOutputItems represents the sessions providers list output items type.
type SessionsProvidersListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the session provider is deleted. Archived session providers stay listable.
	Status string `json:"status"`
	// Usage - Usage statistics
	Usage SessionsProvidersListOutputItemsUsage `json:"usage"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionsProvidersListOutputItemsToolFilter `json:"tool_filter"`
	// ProviderId - Provider this session provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionId - Session this provider is connected to.
	SessionId string `json:"session_id"`
	// FromTemplateId - Template this provider was created from. `null` when it was added directly.
	FromTemplateId *string `json:"from_template_id,omitempty"`
	// FromTemplateProviderId - Template provider this was created from. `null` when added directly.
	FromTemplateProviderId *string                                    `json:"from_template_provider_id,omitempty"`
	Deployment             SessionsProvidersListOutputItemsDeployment `json:"deployment"`
	Config                 SessionsProvidersListOutputItemsConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *SessionsProvidersListOutputItemsAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionsProvidersListOutputPagination represents the sessions providers list output pagination type.
type SessionsProvidersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SessionsProvidersListOutput represents the sessions providers list output type.
type SessionsProvidersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                `json:"object"`
	Items      []SessionsProvidersListOutputItems    `json:"items"`
	Pagination SessionsProvidersListOutputPagination `json:"pagination"`
}

// MapSessionsProvidersListOutputFromJSON deserializes JSON data into a SessionsProvidersListOutput.
func MapSessionsProvidersListOutputFromJSON(data []byte) (*SessionsProvidersListOutput, error) {
	var v SessionsProvidersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsProvidersListOutputToJSON serializes a SessionsProvidersListOutput to JSON.
func MapSessionsProvidersListOutputToJSON(v *SessionsProvidersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionsProvidersListQueryCreatedAt - Only return results whose session provider creation time is in this range. Set `gt`, `lt`, or both.
type SessionsProvidersListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsProvidersListQueryUpdatedAt - Only return results whose session provider last update time is in this range. Set `gt`, `lt`, or both.
type SessionsProvidersListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsProvidersListQuery represents the sessions providers list query type.
type SessionsProvidersListQuery struct {
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
	// Id - Filter by session provider ID(s)
	Id *any `json:"id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
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
	// Status - Only return session providers with one of these statuses.
	Status *any `json:"status,omitempty"`
	// CreatedAt - Only return results whose session provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SessionsProvidersListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SessionsProvidersListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSessionsProvidersListQueryFromJSON deserializes JSON data into a SessionsProvidersListQuery.
func MapSessionsProvidersListQueryFromJSON(data []byte) (*SessionsProvidersListQuery, error) {
	var v SessionsProvidersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsProvidersListQueryToJSON serializes a SessionsProvidersListQuery to JSON.
func MapSessionsProvidersListQueryToJSON(v *SessionsProvidersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
