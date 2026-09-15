package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chats/channels/typing"
)

// ChatsChannelsTypingEndpoint provides access to chat channels are the conversations within a chat, such as Slack channels or Microsoft Teams channels.
type ChatsChannelsTypingEndpoint struct {
	client *endpoint.Client
}

// NewChatsChannelsTypingEndpoint creates a new ChatsChannelsTypingEndpoint.
func NewChatsChannelsTypingEndpoint(client *endpoint.Client) *ChatsChannelsTypingEndpoint {
	return &ChatsChannelsTypingEndpoint{client: client}
}

// ChatsChannelsTypingEndpointStartBody contains the request body for Start.
type ChatsChannelsTypingEndpointStartBody struct {
	// ThreadId - Show the typing indicator within this thread
	ThreadId *string `json:"thread_id,omitempty"`
	// Status - Provider-specific typing status text
	Status *string `json:"status,omitempty"`
}

// Start shows a typing indicator in a chat channel.
func (e *ChatsChannelsTypingEndpoint) Start(instanceId string, chatId string, channelId string, body *ChatsChannelsTypingEndpointStartBody) (*typing.ChatsChannelsTypingStartOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chats", chatId, "channels", channelId, "typing"},
		Body: body,
	}
	var result typing.ChatsChannelsTypingStartOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
