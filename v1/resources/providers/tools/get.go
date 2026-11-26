package tools

import (
	"encoding/json"
	"time"
)

// ProvidersToolsGetOutputInputSchema represents the providers tools get output input schema type.
type ProvidersToolsGetOutputInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// ProvidersToolsGetOutputOutputSchema represents the providers tools get output output schema type.
type ProvidersToolsGetOutputOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// ProvidersToolsGetOutputTags represents the providers tools get output tags type.
type ProvidersToolsGetOutputTags struct {
	// Destructive - Whether the tool changes or deletes data.
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool only reads data.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// ProvidersToolsGetOutput represents the providers tools get output type.
type ProvidersToolsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Key - Key used to call the tool over MCP.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Capabilities - What the tool can do.
	Capabilities map[string]any `json:"capabilities"`
	// Constraints - Rules for using the tool.
	Constraints []string `json:"constraints"`
	// Instructions - Guidance for calling the tool correctly.
	Instructions []string                             `json:"instructions"`
	InputSchema  *ProvidersToolsGetOutputInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProvidersToolsGetOutputOutputSchema `json:"output_schema,omitempty"`
	Tags         *ProvidersToolsGetOutputTags         `json:"tags,omitempty"`
	// ToolCategoryKeys - Keys of the tool categories this tool belongs to
	ToolCategoryKeys []string `json:"tool_category_keys"`
	// SpecificationId - Specification this tool belongs to.
	SpecificationId string `json:"specification_id"`
	// ProviderId - Provider this tool belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProvidersToolsGetOutputFromJSON deserializes JSON data into a ProvidersToolsGetOutput.
func MapProvidersToolsGetOutputFromJSON(data []byte) (*ProvidersToolsGetOutput, error) {
	var v ProvidersToolsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersToolsGetOutputToJSON serializes a ProvidersToolsGetOutput to JSON.
func MapProvidersToolsGetOutputToJSON(v *ProvidersToolsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
