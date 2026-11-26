package routing

import (
	"encoding/json"
)

// AiGatewayRoutingPreviewOutputRequest represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type AiGatewayRoutingPreviewOutputRequest struct {
	// Type - unified: a plain model name, which routes to your connections first and then to the Metorial provider. metorial: the name was prefixed with metorial:, which only routes to the Metorial provider.
	Type *string `json:"type,omitempty"`
	// Provider - The slug of the pinned provider
	Provider *string `json:"provider,omitempty"`
	// VendorModelId - The provider's id for the model
	VendorModelId *string `json:"vendor_model_id,omitempty"`
	// Name - The model name without any prefix
	Name *string `json:"name,omitempty"`
}

// AiGatewayRoutingPreviewOutputModel - The catalog model the requested string resolved to
type AiGatewayRoutingPreviewOutputModel struct {
	// Slug - The model's slug, e.g. anthropic/claude-sonnet-4
	Slug string `json:"slug"`
	// LabSlug - The slug of the lab that makes the model
	LabSlug string `json:"lab_slug"`
}

// AiGatewayRoutingPreviewOutput represents the ai gateway routing preview output type.
type AiGatewayRoutingPreviewOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Requested - The model string the preview was made for
	Requested string `json:"requested"`
	// Request - How the requested model string was parsed
	Request AiGatewayRoutingPreviewOutputRequest `json:"request"`
	// Via - How the name was matched to a catalog model. slug: it is the model's slug. alias: it is a name a provider, SDK or Metorial uses for the model, e.g. a dated snapshot id. key: it matches such a name once case, separators and date formats are normalized. offering: it pinned a provider and is one of that provider's ids for the model.
	Via string `json:"via"`
	// Model - The catalog model the requested string resolved to
	Model AiGatewayRoutingPreviewOutputModel `json:"model"`
	// Candidates - The providers a request would be sent to, in the order they are tried. The gateway moves on to the next one when a provider fails. Your provider connections come first, then the Metorial provider if it is enabled for the instance. When several Metorial routes share traffic, their order is drawn at random for each request.
	Candidates []any `json:"candidates"`
}

// MapAiGatewayRoutingPreviewOutputFromJSON deserializes JSON data into a AiGatewayRoutingPreviewOutput.
func MapAiGatewayRoutingPreviewOutputFromJSON(data []byte) (*AiGatewayRoutingPreviewOutput, error) {
	var v AiGatewayRoutingPreviewOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayRoutingPreviewOutputToJSON serializes a AiGatewayRoutingPreviewOutput to JSON.
func MapAiGatewayRoutingPreviewOutputToJSON(v *AiGatewayRoutingPreviewOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayRoutingPreviewBody represents the ai gateway routing preview body type.
type AiGatewayRoutingPreviewBody struct {
	// Model - The model string a request would use
	Model string `json:"model"`
	// RequiredFeatures - Only route to provider connections with these features enabled. zdr: zero data retention. Instances that require zero data retention always include zdr.
	RequiredFeatures *[]string `json:"required_features,omitempty"`
}

// MapAiGatewayRoutingPreviewBodyFromJSON deserializes JSON data into a AiGatewayRoutingPreviewBody.
func MapAiGatewayRoutingPreviewBodyFromJSON(data []byte) (*AiGatewayRoutingPreviewBody, error) {
	var v AiGatewayRoutingPreviewBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayRoutingPreviewBodyToJSON serializes a AiGatewayRoutingPreviewBody to JSON.
func MapAiGatewayRoutingPreviewBodyToJSON(v *AiGatewayRoutingPreviewBody) ([]byte, error) {
	return json.Marshal(v)
}
