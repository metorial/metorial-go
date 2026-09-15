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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// WorkspaceId - Filter to channels in this workspace
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// Type - Filter by channel type
	Type *string `json:"type,omitempty"`
	// Search - Search by channel name or topic
	Search *string `json:"search,omitempty"`
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

// Get retrieves a specific chat channel.
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
