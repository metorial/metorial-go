package instances

import (
	"encoding/json"
	"time"
)

// ChatInstancesListOutputItemsIdentity - Account this instance is authorized as on the chat provider. `null` until observed.
type ChatInstancesListOutputItemsIdentity struct {
	// Id - Unique ID.
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
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the instance. Archived instances stay listable.
	Status string `json:"status"`
	// ChatConnectionId - The chat connection this instance belongs to
	ChatConnectionId string `json:"chat_connection_id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// Identity - Account this instance is authorized as on the chat provider. `null` until observed.
	Identity *ChatInstancesListOutputItemsIdentity `json:"identity,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatInstancesListOutputPagination represents the chat instances list output pagination type.
type ChatInstancesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ChatInstancesListOutput represents the chat instances list output type.
type ChatInstancesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                            `json:"object"`
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

// ChatInstancesListQueryCreatedAt - Only return results whose chat instance creation time is in this range. Set `gt`, `lt`, or both.
type ChatInstancesListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatInstancesListQueryUpdatedAt - Only return results whose chat instance last update time is in this range. Set `gt`, `lt`, or both.
type ChatInstancesListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatInstancesListQuery represents the chat instances list query type.
type ChatInstancesListQuery struct {
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
	// Status - Only return chat instances with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return chat instances with one of these IDs.
	Id *any `json:"id,omitempty"`
	// ChatConnectionId - Only return chat instances for one of these chat connections.
	ChatConnectionId *any `json:"chat_connection_id,omitempty"`
	// CreatedAt - Only return results whose chat instance creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ChatInstancesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose chat instance last update time is in this range. Set `gt`, `lt`, or both.
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
