package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/token"
)

// TokenEndpoint provides access to a token is the credential a request authenticates with. This resource describes the calling token’s type and scope.
type TokenEndpoint struct {
	client *endpoint.Client
}

// NewTokenEndpoint creates a new TokenEndpoint.
func NewTokenEndpoint(client *endpoint.Client) *TokenEndpoint {
	return &TokenEndpoint{client: client}
}

// Get retrieves metadata about the calling token.
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
