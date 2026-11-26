package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chats/channels/members"
)

// ChatsChannelsMembersEndpoint provides access to chat channels are the conversations within a chat, such as Slack channels or Microsoft Teams channels.
type ChatsChannelsMembersEndpoint struct {
	client *endpoint.Client
}

// NewChatsChannelsMembersEndpoint creates a new ChatsChannelsMembersEndpoint.
func NewChatsChannelsMembersEndpoint(client *endpoint.Client) *ChatsChannelsMembersEndpoint {
	return &ChatsChannelsMembersEndpoint{client: client}
}

// ChatsChannelsMembersEndpointListParams contains optional query parameters for List.
type ChatsChannelsMembersEndpointListParams struct {
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
}

// List returns a paginated list of members of a chat channel.
func (e *ChatsChannelsMembersEndpoint) List(chatId string, channelId string, params *ChatsChannelsMembersEndpointListParams) (*members.ChatsChannelsMembersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chats", chatId, "channels", channelId, "members"},
		Query: query,
	}
	var result members.ChatsChannelsMembersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a member of a chat channel.
func (e *ChatsChannelsMembersEndpoint) Get(chatId string, channelId string, userId string) (*members.ChatsChannelsMembersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"chats", chatId, "channels", channelId, "members", userId},
	}
	var result members.ChatsChannelsMembersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
