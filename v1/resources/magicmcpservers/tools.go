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
	// Destructive - Whether the tool is destructive
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool is read-only
	ReadOnly *bool `json:"read_only,omitempty"`
}

// MagicMcpServersToolsOutputItems represents the magic mcp servers tools output items type.
type MagicMcpServersToolsOutputItems struct {
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
	Instructions []string                                     `json:"instructions"`
	InputSchema  *MagicMcpServersToolsOutputItemsInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersToolsOutputItemsOutputSchema `json:"output_schema,omitempty"`
	Tags         *MagicMcpServersToolsOutputItemsTags         `json:"tags,omitempty"`
	// SpecificationId - Specification ID
	SpecificationId string `json:"specification_id"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersToolsOutput represents the magic mcp servers tools output type.
type MagicMcpServersToolsOutput struct {
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
