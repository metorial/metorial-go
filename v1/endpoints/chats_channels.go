package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chats/channels"
)

// ChatsChannelsEndpoint provides access to chat channels are the conversations within a chat, such as Slack channels or Microsoft Teams channels.
type ChatsChannelsEndpoint struct {
	client *endpoint.Client
}

// NewChatsChannelsEndpoint creates a new ChatsChannelsEndpoint.
func NewChatsChannelsEndpoint(client *endpoint.Client) *ChatsChannelsEndpoint {
	return &ChatsChannelsEndpoint{client: client}
}

// ChatsChannelsEndpointListParams contains optional query parameters for List.
type ChatsChannelsEndpointListParams struct {
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
	// WorkspaceId - Only return channels in this workspace.
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// Type - Only return channels of this type.
	Type *string `json:"type,omitempty"`
	// Search - Search by channel name or topic
	Search *string `json:"search,omitempty"`
	// HasAccess - Whether to return channels the chat integration can or cannot access; defaults to true
	HasAccess *bool `json:"has_access,omitempty"`
}

// List returns a paginated list of channels for a chat.
func (e *ChatsChannelsEndpoint) List(chatId string, params *ChatsChannelsEndpointListParams) (*channels.ChatsChannelsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chats", chatId, "channels"},
		Query: query,
	}
	var result channels.ChatsChannelsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a chat channel.
func (e *ChatsChannelsEndpoint) Get(chatId string, channelId string) (*channels.ChatsChannelsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"chats", chatId, "channels", channelId},
	}
	var result channels.ChatsChannelsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
