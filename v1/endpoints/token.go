package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/token"
)

// TokenEndpoint provides access to endpoint for retrieving metadata about the token used for authentication. This is useful for clients to understand the type and capabilities of the token they are using, especially since Metorial supports multiple token types with different permission models.
type TokenEndpoint struct {
	client *endpoint.Client
}

// NewTokenEndpoint creates a new TokenEndpoint.
func NewTokenEndpoint(client *endpoint.Client) *TokenEndpoint {
	return &TokenEndpoint{client: client}
}

// Get retrieves metadata and configuration details for a specific token.
func (e *TokenEndpoint) Get() (*token.TokenGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"token"},
	}
	var result token.TokenGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
