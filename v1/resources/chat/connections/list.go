package connections

import (
	"encoding/json"
	"time"
)

// ChatConnectionsListOutputItemsProvidersProvider represents the chat connections list output items providers provider type.
type ChatConnectionsListOutputItemsProvidersProvider struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// CreatedAt - Timestamp when the provider was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatConnectionsListOutputItemsProvidersDeployment represents the chat connections list output items providers deployment type.
type ChatConnectionsListOutputItemsProvidersDeployment struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Deployment ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default deployment
	IsDefault bool `json:"is_default"`
	// Name - Deployment name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatConnectionsListOutputItemsProvidersConfig represents the chat connections list output items providers config type.
type ChatConnectionsListOutputItemsProvidersConfig struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Config ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default config
	IsDefault bool `json:"is_default"`
	// Name - Config name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatConnectionsListOutputItemsProviders represents the chat connections list output items providers type.
type ChatConnectionsListOutputItemsProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat connection provider identifier
	Id string `json:"id"`
	// Status - The chat connection provider status
	Status     string                                             `json:"status"`
	Provider   ChatConnectionsListOutputItemsProvidersProvider    `json:"provider"`
	Deployment *ChatConnectionsListOutputItemsProvidersDeployment `json:"deployment,omitempty"`
	Config     *ChatConnectionsListOutputItemsProvidersConfig     `json:"config,omitempty"`
	// AuthMethodId - The auth method this connection authenticates with, if any
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - The auth credentials this connection authenticates with, if any
	AuthCredentialsId *string `json:"auth_credentials_id,omitempty"`
	// Name - Display name of the provider link
	Name string `json:"name"`
	// Description - Description of the provider link
	Description *string `json:"description,omitempty"`
	// CreatedAt - Timestamp when the provider was linked
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider link was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatConnectionsListOutputItems represents the chat connections list output items type.
type ChatConnectionsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat connection identifier
	Id string `json:"id"`
	// Status - The chat connection status
	Status string `json:"status"`
	// Slug - URL-safe identifier for the chat connection
	Slug string `json:"slug"`
	// Name - Display name of the chat connection
	Name string `json:"name"`
	// Description - Description of the chat connection
	Description *string `json:"description,omitempty"`
	// Metadata - Metadata set on the chat connection
	Metadata map[string]any `json:"metadata"`
	// Providers - The provider linked to this chat connection. Currently always zero or one items.
	Providers []ChatConnectionsListOutputItemsProviders `json:"providers"`
	// CreatedAt - Timestamp when the chat connection was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the chat connection was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatConnectionsListOutputPagination represents the chat connections list output pagination type.
type ChatConnectionsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ChatConnectionsListOutput represents the chat connections list output type.
type ChatConnectionsListOutput struct {
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

// ChatConnectionsListQueryCreatedAt - Filter chat connection creation time by date range
type ChatConnectionsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for chat connection creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for chat connection creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatConnectionsListQueryUpdatedAt - Filter chat connection last update time by date range
type ChatConnectionsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for chat connection last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for chat connection last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatConnectionsListQuery represents the chat connections list query type.
type ChatConnectionsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Search - Search by name
	Search *string `json:"search,omitempty"`
	// Status - Filter by status
	Status *any `json:"status,omitempty"`
	// Id - Filter by chat connection ID(s)
	Id *any `json:"id,omitempty"`
	// CreatedAt - Filter chat connection creation time by date range
	CreatedAt *ChatConnectionsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter chat connection last update time by date range
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
