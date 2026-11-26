package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/aigateway/virtualmodels"
)

// AiGatewayVirtualModelsEndpoint provides access to virtual models are model names that pick a real model for every request, using rules, optional signals answered by a decision model, a default, and stickiness. Every change to the definition creates a new version.
type AiGatewayVirtualModelsEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayVirtualModelsEndpoint creates a new AiGatewayVirtualModelsEndpoint.
func NewAiGatewayVirtualModelsEndpoint(client *endpoint.Client) *AiGatewayVirtualModelsEndpoint {
	return &AiGatewayVirtualModelsEndpoint{client: client}
}

// AiGatewayVirtualModelsEndpointListParams contains optional query parameters for List.
type AiGatewayVirtualModelsEndpointListParams struct {
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
}

// AiGatewayVirtualModelsEndpointCreateBody contains the request body for Create.
type AiGatewayVirtualModelsEndpointCreateBody struct {
	// Slug - The model name to send in AI gateway requests to use the virtual model. Lowercase letters, digits, ".", "_" and "-", starting with a letter or digit.
	Slug string `json:"slug"`
	// Name - The virtual model's name
	Name string `json:"name"`
	// Description - What the virtual model is for
	Description *string `json:"description,omitempty"`
	// Definition - How the virtual model picks a model for each request
	Definition map[string]any `json:"definition"`
}

// AiGatewayVirtualModelsEndpointCreateFromTemplateBody contains the request body for CreateFromTemplate.
type AiGatewayVirtualModelsEndpointCreateFromTemplateBody struct {
	// Template - The slug of the template to use
	Template string `json:"template"`
	// Slug - The model name to send in AI gateway requests to use the virtual model. Lowercase letters, digits, ".", "_" and "-", starting with a letter or digit.
	Slug string `json:"slug"`
	// Name - The virtual model's name
	Name string `json:"name"`
	// Description - What the virtual model is for. Defaults to the template's description.
	Description *string `json:"description,omitempty"`
	// Parameters - The model for each of the template's parameters, keyed by parameter key
	Parameters map[string]string `json:"parameters"`
	// DecisionConnectionId - The ID of your own provider connection to run the template's decision model on. Omit to use Metorial's, billed per call.
	DecisionConnectionId *string `json:"decision_connection_id,omitempty"`
}

// AiGatewayVirtualModelsEndpointTestDraftBody contains the request body for TestDraft.
type AiGatewayVirtualModelsEndpointTestDraftBody struct {
	// Definition - How the virtual model picks a model for each request
	Definition map[string]any `json:"definition"`
	// TokenId - The AI gateway token to test as, so rules see its consumer and metadata and its allowed models apply. Omit to test without a token's restrictions.
	TokenId *string `json:"token_id,omitempty"`
	// Dialect - The API format of the request: openai_chat (Chat Completions), anthropic_messages (Messages) or open_responses (Responses)
	Dialect string `json:"dialect"`
	// Request - The request body to route, in the format given by dialect
	Request map[string]any `json:"request"`
	// Metadata - Request metadata as sent in the x-metorial-metadata header. Rules read it as metadata.
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// AiGatewayVirtualModelsEndpointUpdateBody contains the request body for Update.
type AiGatewayVirtualModelsEndpointUpdateBody struct {
	// Name - The virtual model's name
	Name *string `json:"name,omitempty"`
	// Description - What the virtual model is for. Pass null to clear it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Definition - How the virtual model picks a model for each request
	Definition *map[string]any `json:"definition,omitempty"`
}

// AiGatewayVirtualModelsEndpointRollbackBody contains the request body for Rollback.
type AiGatewayVirtualModelsEndpointRollbackBody struct {
	// VersionId - The ID of the version to roll back to
	VersionId string `json:"version_id"`
}

// AiGatewayVirtualModelsEndpointTestBody contains the request body for Test.
type AiGatewayVirtualModelsEndpointTestBody struct {
	// Definition - A definition to test instead of the current one, without saving it
	Definition *map[string]any `json:"definition,omitempty"`
	// TokenId - The AI gateway token to test as, so rules see its consumer and metadata and its allowed models apply. Omit to test without a token's restrictions.
	TokenId *string `json:"token_id,omitempty"`
	// Dialect - The API format of the request: openai_chat (Chat Completions), anthropic_messages (Messages) or open_responses (Responses)
	Dialect string `json:"dialect"`
	// Request - The request body to route, in the format given by dialect
	Request map[string]any `json:"request"`
	// Metadata - Request metadata as sent in the x-metorial-metadata header. Rules read it as metadata.
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// AiGatewayVirtualModelsEndpointReplayBody contains the request body for Replay.
type AiGatewayVirtualModelsEndpointReplayBody struct {
	// Definition - A definition to replay instead of the current one, without saving it
	Definition *map[string]any `json:"definition,omitempty"`
	// Limit - How many of the most recent requests to replay. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
}

// List returns a paginated list of virtual models.
func (e *AiGatewayVirtualModelsEndpoint) List(params *AiGatewayVirtualModelsEndpointListParams) (*virtualmodels.AiGatewayVirtualModelsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"ai-gateway", "virtual-models"},
		Query: query,
	}
	var result virtualmodels.AiGatewayVirtualModelsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a virtual model by ID or slug, with its current definition.
func (e *AiGatewayVirtualModelsEndpoint) Get(virtualModelId string) (*virtualmodels.AiGatewayVirtualModelsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "virtual-models", virtualModelId},
	}
	var result virtualmodels.AiGatewayVirtualModelsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a virtual model with its first version. Requests that use its slug as the model are routed by the definition.
func (e *AiGatewayVirtualModelsEndpoint) Create(body *AiGatewayVirtualModelsEndpointCreateBody) (*virtualmodels.AiGatewayVirtualModelsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "virtual-models"},
		Body: body,
	}
	var result virtualmodels.AiGatewayVirtualModelsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateFromTemplate creates a virtual model from a template, filling in the models it asks for. The definition can be changed afterwards like any other.
func (e *AiGatewayVirtualModelsEndpoint) CreateFromTemplate(body *AiGatewayVirtualModelsEndpointCreateFromTemplateBody) (*virtualmodels.AiGatewayVirtualModelsCreateFromTemplateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "virtual-models", "from-template"},
		Body: body,
	}
	var result virtualmodels.AiGatewayVirtualModelsCreateFromTemplateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// TestDraft routes a request through an unsaved definition and shows which rule and models it picks. Nothing is sent to a provider, but the decision model is called and billed when a rule reads signals.
func (e *AiGatewayVirtualModelsEndpoint) TestDraft(body *AiGatewayVirtualModelsEndpointTestDraftBody) (*virtualmodels.AiGatewayVirtualModelsTestDraftOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "virtual-models", "test"},
		Body: body,
	}
	var result virtualmodels.AiGatewayVirtualModelsTestDraftOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates the name, description or definition of a virtual model. A new definition becomes a new version and takes effect right away.
func (e *AiGatewayVirtualModelsEndpoint) Update(virtualModelId string, body *AiGatewayVirtualModelsEndpointUpdateBody) (*virtualmodels.AiGatewayVirtualModelsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "virtual-models", virtualModelId},
		Body: body,
	}
	var result virtualmodels.AiGatewayVirtualModelsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a virtual model. Requests that use its slug are no longer routed by it, and the slug can be used again.
func (e *AiGatewayVirtualModelsEndpoint) Delete(virtualModelId string) (*virtualmodels.AiGatewayVirtualModelsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "virtual-models", virtualModelId},
	}
	var result virtualmodels.AiGatewayVirtualModelsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Rollback makes an earlier version's definition current again, as a new version. The version history is kept.
func (e *AiGatewayVirtualModelsEndpoint) Rollback(virtualModelId string, body *AiGatewayVirtualModelsEndpointRollbackBody) (*virtualmodels.AiGatewayVirtualModelsRollbackOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "virtual-models", virtualModelId, "rollback"},
		Body: body,
	}
	var result virtualmodels.AiGatewayVirtualModelsRollbackOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Test routes a request through the virtual model and shows which rule and models it picks. Nothing is sent to a provider, but the decision model is called and billed when a rule reads signals.
func (e *AiGatewayVirtualModelsEndpoint) Test(virtualModelId string, body *AiGatewayVirtualModelsEndpointTestBody) (*virtualmodels.AiGatewayVirtualModelsTestOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "virtual-models", virtualModelId, "test"},
		Body: body,
	}
	var result virtualmodels.AiGatewayVirtualModelsTestOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Replay routes the virtual model's most recent logged requests again, through the current definition or a draft, and shows which would be routed differently. Only requests whose content was logged are replayed. Logged decision model answers are reused while the signals are unchanged; otherwise the decision model is called and billed.
func (e *AiGatewayVirtualModelsEndpoint) Replay(virtualModelId string, body *AiGatewayVirtualModelsEndpointReplayBody) (*virtualmodels.AiGatewayVirtualModelsReplayOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "virtual-models", virtualModelId, "replay"},
		Body: body,
	}
	var result virtualmodels.AiGatewayVirtualModelsReplayOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
