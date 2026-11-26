package connections

import (
	"encoding/json"
	"time"
)

// ChatConnectionsListOutputItemsProvidersProvider represents the chat connections list output items providers provider type.
type ChatConnectionsListOutputItemsProvidersProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatConnectionsListOutputItemsProvidersDeployment represents the chat connections list output items providers deployment type.
type ChatConnectionsListOutputItemsProvidersDeployment struct {
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

// ChatConnectionsListOutputItemsProvidersConfig represents the chat connections list output items providers config type.
type ChatConnectionsListOutputItemsProvidersConfig struct {
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

// ChatConnectionsListOutputItemsProviders represents the chat connections list output items providers type.
type ChatConnectionsListOutputItemsProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the connection is archived.
	Status     string                                             `json:"status"`
	Provider   ChatConnectionsListOutputItemsProvidersProvider    `json:"provider"`
	Deployment *ChatConnectionsListOutputItemsProvidersDeployment `json:"deployment,omitempty"`
	Config     *ChatConnectionsListOutputItemsProvidersConfig     `json:"config,omitempty"`
	// AuthMethodId - The auth method this connection authenticates with, if any
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - The auth credentials this connection authenticates with, if any
	AuthCredentialsId *string `json:"auth_credentials_id,omitempty"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Timestamp when the provider was linked
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider link was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatConnectionsListOutputItems represents the chat connections list output items type.
type ChatConnectionsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the connection. Archived connections stay listable.
	Status string `json:"status"`
	// Slug - URL-safe name, unique across Metorial.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// Providers - The provider linked to this chat connection. Currently always zero or one items.
	Providers []ChatConnectionsListOutputItemsProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatConnectionsListOutputPagination represents the chat connections list output pagination type.
type ChatConnectionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ChatConnectionsListOutput represents the chat connections list output type.
type ChatConnectionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                              `json:"object"`
	Items      []ChatConnectionsListOutputItems    `json:"items"`
	Pagination ChatConnectionsListOutputPagination `json:"pagination"`
}

// MapChatConnectionsListOutputFromJSON deserializes JSON data into a ChatConnectionsListOutput.
func MapChatConnectionsListOutputFromJSON(data []byte) (*ChatConnectionsListOutput, error) {
	var v ChatConnectionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatConnectionsListOutputToJSON serializes a ChatConnectionsListOutput to JSON.
func MapChatConnectionsListOutputToJSON(v *ChatConnectionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatConnectionsListQueryCreatedAt - Only return results whose chat connection creation time is in this range. Set `gt`, `lt`, or both.
type ChatConnectionsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatConnectionsListQueryUpdatedAt - Only return results whose chat connection last update time is in this range. Set `gt`, `lt`, or both.
type ChatConnectionsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatConnectionsListQuery represents the chat connections list query type.
type ChatConnectionsListQuery struct {
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
	// Search - Search by name
	Search *string `json:"search,omitempty"`
	// Status - Only return chat connections with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return chat connections with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CreatedAt - Only return results whose chat connection creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ChatConnectionsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose chat connection last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ChatConnectionsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapChatConnectionsListQueryFromJSON deserializes JSON data into a ChatConnectionsListQuery.
func MapChatConnectionsListQueryFromJSON(data []byte) (*ChatConnectionsListQuery, error) {
	var v ChatConnectionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatConnectionsListQueryToJSON serializes a ChatConnectionsListQuery to JSON.
func MapChatConnectionsListQueryToJSON(v *ChatConnectionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
