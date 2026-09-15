package instances

import (
	"encoding/json"
	"time"
)

// ChatInstancesListOutputItemsIdentity - The identity this connection is authorized as on the chat provider (e.g. the connected bot or user). Null until this account has been observed on the provider (e.g. through a synced message or channel membership).
type ChatInstancesListOutputItemsIdentity struct {
	// Id - Unique chat author identifier
	Id string `json:"id"`
	// UserId - The provider-side user or app id this connection is authorized as
	UserId string `json:"user_id"`
	// Name - Display name of the authorized identity
	Name string `json:"name"`
	// Username - Username or handle of the authorized identity
	Username string `json:"username"`
	// ProviderType - The kind of identity as classified by the provider, e.g. a human user vs. an app/bot
	ProviderType string `json:"provider_type"`
	// Email - Email address of the authorized identity, if available
	Email *string `json:"email,omitempty"`
	// ImageUrl - Avatar or image URL of the authorized identity, if available
	ImageUrl *string `json:"image_url,omitempty"`
}

// ChatInstancesListOutputItems represents the chat instances list output items type.
type ChatInstancesListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat instance identifier
	Id string `json:"id"`
	// Status - The chat instance status
	Status string `json:"status"`
	// ChatConnectionId - The chat connection this instance belongs to
	ChatConnectionId string `json:"chat_connection_id"`
	// Name - Display name of the chat instance
	Name string `json:"name"`
	// Description - Description of the chat instance
	Description *string `json:"description,omitempty"`
	// Metadata - Metadata set on the chat instance
	Metadata map[string]any `json:"metadata"`
	// Identity - The identity this connection is authorized as on the chat provider (e.g. the connected bot or user). Null until this account has been observed on the provider (e.g. through a synced message or channel membership).
	Identity *ChatInstancesListOutputItemsIdentity `json:"identity,omitempty"`
	// CreatedAt - Timestamp when the chat instance was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the chat instance was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatInstancesListOutputPagination represents the chat instances list output pagination type.
type ChatInstancesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ChatInstancesListOutput represents the chat instances list output type.
type ChatInstancesListOutput struct {
	Items      []ChatInstancesListOutputItems    `json:"items"`
	Pagination ChatInstancesListOutputPagination `json:"pagination"`
}

// MapChatInstancesListOutputFromJSON deserializes JSON data into a ChatInstancesListOutput.
func MapChatInstancesListOutputFromJSON(data []byte) (*ChatInstancesListOutput, error) {
	var v ChatInstancesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatInstancesListOutputToJSON serializes a ChatInstancesListOutput to JSON.
func MapChatInstancesListOutputToJSON(v *ChatInstancesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatInstancesListQueryCreatedAt - Filter chat instance creation time by date range
type ChatInstancesListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for chat instance creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for chat instance creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatInstancesListQueryUpdatedAt - Filter chat instance last update time by date range
type ChatInstancesListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for chat instance last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for chat instance last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatInstancesListQuery represents the chat instances list query type.
type ChatInstancesListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Search - Search by name
	Search *string `json:"search,omitempty"`
	// Status - Filter by status
	Status *any `json:"status,omitempty"`
	// Id - Filter by chat instance ID(s)
	Id *any `json:"id,omitempty"`
	// ChatConnectionId - Filter by chat connection ID(s)
	ChatConnectionId *any `json:"chat_connection_id,omitempty"`
	// CreatedAt - Filter chat instance creation time by date range
	CreatedAt *ChatInstancesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter chat instance last update time by date range
	UpdatedAt *ChatInstancesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapChatInstancesListQueryFromJSON deserializes JSON data into a ChatInstancesListQuery.
func MapChatInstancesListQueryFromJSON(data []byte) (*ChatInstancesListQuery, error) {
	var v ChatInstancesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatInstancesListQueryToJSON serializes a ChatInstancesListQuery to JSON.
func MapChatInstancesListQueryToJSON(v *ChatInstancesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
