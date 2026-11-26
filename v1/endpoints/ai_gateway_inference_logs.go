package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/aigateway/inferencelogs"
)

// AiGatewayInferenceLogsEndpoint provides access to a record of every request made through the AI gateway: the model and provider used, tokens, cost, timing and errors. Kept for the configured log retention.
type AiGatewayInferenceLogsEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayInferenceLogsEndpoint creates a new AiGatewayInferenceLogsEndpoint.
func NewAiGatewayInferenceLogsEndpoint(client *endpoint.Client) *AiGatewayInferenceLogsEndpoint {
	return &AiGatewayInferenceLogsEndpoint{client: client}
}

// AiGatewayInferenceLogsEndpointListParams contains optional query parameters for List.
type AiGatewayInferenceLogsEndpointListParams struct {
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
	// TokenId - Only requests made with this AI gateway token
	TokenId *string `json:"token_id,omitempty"`
	// ConsumerId - Only requests made with tokens of this consumer
	ConsumerId *string `json:"consumer_id,omitempty"`
	// Model - Only requests served by this model, by slug
	Model *string `json:"model,omitempty"`
	// VirtualModelId - Only requests routed by this virtual model
	VirtualModelId *string `json:"virtual_model_id,omitempty"`
	// Status - Only requests that ended this way
	Status *string `json:"status,omitempty"`
}

// AiGatewayInferenceLogsEndpointGetParams contains optional query parameters for Get.
type AiGatewayInferenceLogsEndpointGetParams struct {
	// IncludeContent - Return the stored request and response bodies. They exist only for requests made while content logging was enabled, and until their retention ends.
	IncludeContent *bool `json:"include_content,omitempty"`
}

// List returns inference logs, newest first, without request and response content. Pages only go towards older logs: pass the after cursor of the previous page. before cursors are rejected.
func (e *AiGatewayInferenceLogsEndpoint) List(params *AiGatewayInferenceLogsEndpointListParams) (*inferencelogs.AiGatewayInferenceLogsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"ai-gateway", "inference-logs"},
		Query: query,
	}
	var result inferencelogs.AiGatewayInferenceLogsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an inference log by ID. Request and response content is only returned when include_content is set.
func (e *AiGatewayInferenceLogsEndpoint) Get(inferenceLogId string, params *AiGatewayInferenceLogsEndpointGetParams) (*inferencelogs.AiGatewayInferenceLogsGetOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"ai-gateway", "inference-logs", inferenceLogId},
		Query: query,
	}
	var result inferencelogs.AiGatewayInferenceLogsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
