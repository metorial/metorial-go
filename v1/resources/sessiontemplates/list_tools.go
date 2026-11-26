package sessiontemplates

import (
	"encoding/json"
	"time"
)

// SessionTemplatesListToolsOutputItemsInputSchema represents the session templates list tools output items input schema type.
type SessionTemplatesListToolsOutputItemsInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// SessionTemplatesListToolsOutputItemsOutputSchema represents the session templates list tools output items output schema type.
type SessionTemplatesListToolsOutputItemsOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// SessionTemplatesListToolsOutputItemsTags represents the session templates list tools output items tags type.
type SessionTemplatesListToolsOutputItemsTags struct {
	// Destructive - Whether the tool changes or deletes data.
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool only reads data.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// SessionTemplatesListToolsOutputItems represents the session templates list tools output items type.
type SessionTemplatesListToolsOutputItems struct {
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
	Instructions []string                                          `json:"instructions"`
	InputSchema  *SessionTemplatesListToolsOutputItemsInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *SessionTemplatesListToolsOutputItemsOutputSchema `json:"output_schema,omitempty"`
	Tags         *SessionTemplatesListToolsOutputItemsTags         `json:"tags,omitempty"`
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

// SessionTemplatesListToolsOutput represents the session templates list tools output type.
type SessionTemplatesListToolsOutput struct {
	// Object - String representing the object’s type.
	Object string                                 `json:"object"`
	Items  []SessionTemplatesListToolsOutputItems `json:"items"`
}

// MapSessionTemplatesListToolsOutputFromJSON deserializes JSON data into a SessionTemplatesListToolsOutput.
func MapSessionTemplatesListToolsOutputFromJSON(data []byte) (*SessionTemplatesListToolsOutput, error) {
	var v SessionTemplatesListToolsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesListToolsOutputToJSON serializes a SessionTemplatesListToolsOutput to JSON.
func MapSessionTemplatesListToolsOutputToJSON(v *SessionTemplatesListToolsOutput) ([]byte, error) {
	return json.Marshal(v)
}
