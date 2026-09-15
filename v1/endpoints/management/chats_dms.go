package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chats/dms"
)

// ChatsDmsEndpoint provides access to open direct message channels with one or more users on a chat.
type ChatsDmsEndpoint struct {
	client *endpoint.Client
}

// NewChatsDmsEndpoint creates a new ChatsDmsEndpoint.
func NewChatsDmsEndpoint(client *endpoint.Client) *ChatsDmsEndpoint {
	return &ChatsDmsEndpoint{client: client}
}

// Open opens (or retrieves) a direct message channel with one or more users.
func (e *ChatsDmsEndpoint) Open(instanceId string, chatId string) (*dms.ChatsDmsOpenOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chats", chatId, "dms"},
	}
	var result dms.ChatsDmsOpenOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
