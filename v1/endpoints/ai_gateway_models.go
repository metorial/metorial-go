package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/aigateway/models"
)

// AiGatewayModelsEndpoint provides access to the models in the AI gateway catalog, with their capabilities and limits. The catalog is shared by all instances and kept up to date by Metorial.
type AiGatewayModelsEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayModelsEndpoint creates a new AiGatewayModelsEndpoint.
func NewAiGatewayModelsEndpoint(client *endpoint.Client) *AiGatewayModelsEndpoint {
	return &AiGatewayModelsEndpoint{client: client}
}

// AiGatewayModelsEndpointListParams contains optional query parameters for List.
type AiGatewayModelsEndpointListParams struct {
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
	// Lab - Only models made by the lab with this slug
	Lab *string `json:"lab,omitempty"`
	// Kind - Only models of this kind: chat, embedding, reranking, image, video, speech, transcription, realtime, audio, decision or other
	Kind *string `json:"kind,omitempty"`
	// Search - Only models whose slug or name contains this text
	Search *string `json:"search,omitempty"`
}

// AiGatewayModelsEndpointResolveBody contains the request body for Resolve.
type AiGatewayModelsEndpointResolveBody struct {
	// Model - The model string to resolve, as it would be sent to the gateway
	Model string `json:"model"`
}

// List returns a paginated list of the models in the catalog that are not retired.
func (e *AiGatewayModelsEndpoint) List(params *AiGatewayModelsEndpointListParams) (*models.AiGatewayModelsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"ai-gateway", "models"},
		Query: query,
	}
	var result models.AiGatewayModelsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a model by ID or slug. URL-encode the slash of a slug, e.g. anthropic%2Fclaude-sonnet-4.
func (e *AiGatewayModelsEndpoint) Get(modelId string) (*models.AiGatewayModelsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "models", modelId},
	}
	var result models.AiGatewayModelsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Resolve finds the catalog model a model string refers to, the same way the gateway does for requests. Accepts slugs, names providers use for a model, and provider-pinned names like openai:gpt-4o.
func (e *AiGatewayModelsEndpoint) Resolve(body *AiGatewayModelsEndpointResolveBody) (*models.AiGatewayModelsResolveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "models", "resolve"},
		Body: body,
	}
	var result models.AiGatewayModelsResolveOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
