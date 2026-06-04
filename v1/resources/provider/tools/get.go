package tools

import (
	"encoding/json"
	"time"
)

// ProviderToolsGetOutputInputSchema represents the provider tools get output input schema type.
type ProviderToolsGetOutputInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// ProviderToolsGetOutputOutputSchema represents the provider tools get output output schema type.
type ProviderToolsGetOutputOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// ProviderToolsGetOutputTags represents the provider tools get output tags type.
type ProviderToolsGetOutputTags struct {
	// Destructive - Whether the tool is destructive
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool is read-only
	ReadOnly *bool `json:"read_only,omitempty"`
}

// ProviderToolsGetOutput represents the provider tools get output type.
type ProviderToolsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique tool identifier
	Id string `json:"id"`
	// Key - Tool key
	Key string `json:"key"`
	// Name - Display name of the tool
	Name string `json:"name"`
	// Description - Tool description
	Description *string `json:"description,omitempty"`
	// Capabilities - Tool capabilities
	Capabilities map[string]any `json:"capabilities"`
	// Constraints - Tool constraints
	Constraints []string `json:"constraints"`
	// Instructions - Tool usage instructions
	Instructions []string                            `json:"instructions"`
	InputSchema  *ProviderToolsGetOutputInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderToolsGetOutputOutputSchema `json:"output_schema,omitempty"`
	Tags         *ProviderToolsGetOutputTags         `json:"tags,omitempty"`
	// SpecificationId - Specification ID
	SpecificationId string `json:"specification_id"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderToolsGetOutputFromJSON deserializes JSON data into a ProviderToolsGetOutput.
func MapProviderToolsGetOutputFromJSON(data []byte) (*ProviderToolsGetOutput, error) {
	var v ProviderToolsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderToolsGetOutputToJSON serializes a ProviderToolsGetOutput to JSON.
func MapProviderToolsGetOutputToJSON(v *ProviderToolsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
