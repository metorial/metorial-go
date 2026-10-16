package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chats/commands"
)

// ChatsCommandsEndpoint provides access to chat messages are the individual messages sent within a chat channel.
type ChatsCommandsEndpoint struct {
	client *endpoint.Client
}

// NewChatsCommandsEndpoint creates a new ChatsCommandsEndpoint.
func NewChatsCommandsEndpoint(client *endpoint.Client) *ChatsCommandsEndpoint {
	return &ChatsCommandsEndpoint{client: client}
}

// ChatsCommandsEndpointRespondBody contains the request body for Respond.
type ChatsCommandsEndpointRespondBody struct {
	// ChatEventId - The chat.command.invoked event to respond to. Providers only accept a response within their response window after the command was invoked.
	ChatEventId string `json:"chat_event_id"`
	// Parts - Structured content parts making up the message body
	Parts []any `json:"parts"`
	// AltText - Plain-text fallback for the response
	AltText *string `json:"alt_text,omitempty"`
	// Ephemeral - If true, only the user who invoked the command can see the response. Defaults to the provider behavior. Providers that already acknowledged the command publicly reject it.
	Ephemeral *bool `json:"ephemeral,omitempty"`
}

// Respond replies to a slash command from a chat.command.invoked event, using the provider's native command response.
func (e *ChatsCommandsEndpoint) Respond(instanceId string, chatId string, body *ChatsCommandsEndpointRespondBody) (*commands.ChatsCommandsRespondOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chats", chatId, "commands", "respond"},
		Body: body,
	}
	var result commands.ChatsCommandsRespondOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
