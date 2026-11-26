package magicmcpservers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersToolsOutputItemsInputSchema represents the magic mcp servers tools output items input schema type.
type MagicMcpServersToolsOutputItemsInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersToolsOutputItemsOutputSchema represents the magic mcp servers tools output items output schema type.
type MagicMcpServersToolsOutputItemsOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersToolsOutputItemsTags represents the magic mcp servers tools output items tags type.
type MagicMcpServersToolsOutputItemsTags struct {
	// Destructive - Whether the tool changes or deletes data.
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool only reads data.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// MagicMcpServersToolsOutputItems represents the magic mcp servers tools output items type.
type MagicMcpServersToolsOutputItems struct {
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
	Instructions []string                                     `json:"instructions"`
	InputSchema  *MagicMcpServersToolsOutputItemsInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersToolsOutputItemsOutputSchema `json:"output_schema,omitempty"`
	Tags         *MagicMcpServersToolsOutputItemsTags         `json:"tags,omitempty"`
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

// MagicMcpServersToolsOutput represents the magic mcp servers tools output type.
type MagicMcpServersToolsOutput struct {
	// Object - String representing the object’s type.
	Object string                            `json:"object"`
	Items  []MagicMcpServersToolsOutputItems `json:"items"`
}

// MapMagicMcpServersToolsOutputFromJSON deserializes JSON data into a MagicMcpServersToolsOutput.
func MapMagicMcpServersToolsOutputFromJSON(data []byte) (*MagicMcpServersToolsOutput, error) {
	var v MagicMcpServersToolsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersToolsOutputToJSON serializes a MagicMcpServersToolsOutput to JSON.
func MapMagicMcpServersToolsOutputToJSON(v *MagicMcpServersToolsOutput) ([]byte, error) {
	return json.Marshal(v)
}
