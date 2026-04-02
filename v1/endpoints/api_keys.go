package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/apikeys"
)

// ApiKeysEndpoint provides access to read and write API key information
type ApiKeysEndpoint struct {
	client *endpoint.Client
}

// NewApiKeysEndpoint creates a new ApiKeysEndpoint.
func NewApiKeysEndpoint(client *endpoint.Client) *ApiKeysEndpoint {
	return &ApiKeysEndpoint{client: client}
}

// ApiKeysEndpointRotateBody contains the request body for Rotate.
type ApiKeysEndpointRotateBody struct {
	CurrentExpiresAt *string `json:"current_expires_at,omitempty"`
}

// Rotate rotate a specific API key
func (e *ApiKeysEndpoint) Rotate(organizationId string, apiKeyId string, body *ApiKeysEndpointRotateBody) (*apikeys.ApiKeysRotateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"dashboard", "organizations", organizationId, "api-keys", apiKeyId, "rotate"},
		Body: body,
	}
	var result apikeys.ApiKeysRotateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Reveal reveal a specific API key
func (e *ApiKeysEndpoint) Reveal(organizationId string, apiKeyId string) (*apikeys.ApiKeysRevealOutput, error) {
	req := &endpoint.Request{
		Path: []string{"dashboard", "organizations", organizationId, "api-keys", apiKeyId, "reveal"},
	}
	var result apikeys.ApiKeysRevealOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
