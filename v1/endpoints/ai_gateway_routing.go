package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/aigateway/routing"
)

// AiGatewayRoutingEndpoint provides access to shows how the AI gateway routes requests for a model on this instance, without sending one.
type AiGatewayRoutingEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayRoutingEndpoint creates a new AiGatewayRoutingEndpoint.
func NewAiGatewayRoutingEndpoint(client *endpoint.Client) *AiGatewayRoutingEndpoint {
	return &AiGatewayRoutingEndpoint{client: client}
}

// AiGatewayRoutingEndpointPreviewBody contains the request body for Preview.
type AiGatewayRoutingEndpointPreviewBody struct {
	// Model - The model string a request would use
	Model string `json:"model"`
	// RequiredFeatures - Only route to provider connections with these features enabled. zdr: zero data retention. Instances that require zero data retention always include zdr.
	RequiredFeatures *[]string `json:"required_features,omitempty"`
}

// Preview returns the providers a request for the model would be sent to, in the order they are tried, with their pricing. Tokens' allowed models are not applied.
func (e *AiGatewayRoutingEndpoint) Preview(body *AiGatewayRoutingEndpointPreviewBody) (*routing.AiGatewayRoutingPreviewOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "routing", "preview"},
		Body: body,
	}
	var result routing.AiGatewayRoutingPreviewOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
