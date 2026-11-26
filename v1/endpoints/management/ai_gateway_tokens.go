package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/aigateway/tokens"
)

// AiGatewayTokensEndpoint provides access to aPI keys for the AI gateway. Each token can be limited to certain models and given its own rate and spend limits.
type AiGatewayTokensEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayTokensEndpoint creates a new AiGatewayTokensEndpoint.
func NewAiGatewayTokensEndpoint(client *endpoint.Client) *AiGatewayTokensEndpoint {
	return &AiGatewayTokensEndpoint{client: client}
}

// AiGatewayTokensEndpointListParams contains optional query parameters for List.
type AiGatewayTokensEndpointListParams struct {
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
	// ConsumerId - Only tokens created for this consumer
	ConsumerId *string `json:"consumer_id,omitempty"`
}

// AiGatewayTokensEndpointCreateBody contains the request body for Create.
type AiGatewayTokensEndpointCreateBody struct {
	// Name - A name for the token
	Name string `json:"name"`
	// ConsumerId - The consumer the token is for. Requests made with it can be filtered by this consumer in inference logs and usage, and virtual model rules can route on it. Cannot be changed afterwards.
	ConsumerId *string `json:"consumer_id,omitempty"`
	// AllowedModels - Models and virtual models the token can call, by slug. An empty list allows every model.
	AllowedModels *[]string `json:"allowed_models,omitempty"`
	// Limits - Limits that reject requests once reached. Replaces the existing limits when updated.
	Limits *[]map[string]any `json:"limits,omitempty"`
	// Metadata - User-defined metadata that virtual model rules can read. Replaces existing metadata when updated.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// ExpiresAt - When the token stops working. Omit for no expiry.
	ExpiresAt *string `json:"expires_at,omitempty"`
}

// AiGatewayTokensEndpointUpdateBody contains the request body for Update.
type AiGatewayTokensEndpointUpdateBody struct {
	// Name - A name for the token
	Name *string `json:"name,omitempty"`
	// AllowedModels - Models and virtual models the token can call, by slug. An empty list allows every model.
	AllowedModels *[]string `json:"allowed_models,omitempty"`
	// Limits - Limits that reject requests once reached. Replaces the existing limits when updated.
	Limits *[]map[string]any `json:"limits,omitempty"`
	// Metadata - User-defined metadata that virtual model rules can read. Replaces existing metadata when updated.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// ExpiresAt - When the token stops working. Pass null to never expire.
	ExpiresAt *nullable.Nullable[string] `json:"expires_at,omitempty"`
}

// List returns a paginated list of tokens, including revoked ones.
func (e *AiGatewayTokensEndpoint) List(instanceId string, params *AiGatewayTokensEndpointListParams) (*tokens.AiGatewayTokensListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "ai-gateway", "tokens"},
		Query: query,
	}
	var result tokens.AiGatewayTokensListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a token by ID. The secret is never returned again.
func (e *AiGatewayTokensEndpoint) Get(instanceId string, tokenId string) (*tokens.AiGatewayTokensGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "ai-gateway", "tokens", tokenId},
	}
	var result tokens.AiGatewayTokensGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a token. The response contains the secret, which is shown only this once.
func (e *AiGatewayTokensEndpoint) Create(instanceId string, body *AiGatewayTokensEndpointCreateBody) (*tokens.AiGatewayTokensCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "ai-gateway", "tokens"},
		Body: body,
	}
	var result tokens.AiGatewayTokensCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates the name, allowed models, limits, metadata or expiry of a token.
func (e *AiGatewayTokensEndpoint) Update(instanceId string, tokenId string, body *AiGatewayTokensEndpointUpdateBody) (*tokens.AiGatewayTokensUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "ai-gateway", "tokens", tokenId},
		Body: body,
	}
	var result tokens.AiGatewayTokensUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Revoke revokes a token. Requests made with it are rejected from then on.
func (e *AiGatewayTokensEndpoint) Revoke(instanceId string, tokenId string) (*tokens.AiGatewayTokensRevokeOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "ai-gateway", "tokens", tokenId, "revoke"},
	}
	var result tokens.AiGatewayTokensRevokeOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
