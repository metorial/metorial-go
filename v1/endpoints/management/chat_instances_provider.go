package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chat/instances/provider"
)

// ChatInstancesProviderEndpoint provides access to chat instances materialize a chat connection for a specific runtime configuration.
type ChatInstancesProviderEndpoint struct {
	client *endpoint.Client
}

// NewChatInstancesProviderEndpoint creates a new ChatInstancesProviderEndpoint.
func NewChatInstancesProviderEndpoint(client *endpoint.Client) *ChatInstancesProviderEndpoint {
	return &ChatInstancesProviderEndpoint{client: client}
}

// ChatInstancesProviderEndpointSetBody contains the request body for Set.
type ChatInstancesProviderEndpointSetBody struct {
	// ProviderId - The chat provider to use for this instance
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Provider deployment to use
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Provider config to use
	ProviderConfigId *string `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Provider auth config to use
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
}

// Get retrieves the single provider configured for a chat instance.
func (e *ChatInstancesProviderEndpoint) Get(instanceId string, chatInstanceId string) (*provider.ChatInstancesProviderGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chat", "instances", chatInstanceId, "provider"},
	}
	var result provider.ChatInstancesProviderGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Set creates or updates the single provider for a chat instance.
func (e *ChatInstancesProviderEndpoint) Set(instanceId string, chatInstanceId string, body *ChatInstancesProviderEndpointSetBody) (*provider.ChatInstancesProviderSetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chat", "instances", chatInstanceId, "provider"},
		Body: body,
	}
	var result provider.ChatInstancesProviderSetOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AuthenticatedUser retrieves the chat provider account this chat instance is authenticated as.
func (e *ChatInstancesProviderEndpoint) AuthenticatedUser(instanceId string, chatInstanceId string) (*provider.ChatInstancesProviderAuthenticatedUserOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "chat", "instances", chatInstanceId, "provider", "authenticated-user"},
	}
	var result provider.ChatInstancesProviderAuthenticatedUserOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
