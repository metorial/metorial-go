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
	// Destructive - Whether the tool is destructive
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool is read-only
	ReadOnly *bool `json:"read_only,omitempty"`
}

// SessionTemplatesListToolsOutputItems represents the session templates list tools output items type.
type SessionTemplatesListToolsOutputItems struct {
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
	Instructions []string                                          `json:"instructions"`
	InputSchema  *SessionTemplatesListToolsOutputItemsInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *SessionTemplatesListToolsOutputItemsOutputSchema `json:"output_schema,omitempty"`
	Tags         *SessionTemplatesListToolsOutputItemsTags         `json:"tags,omitempty"`
	// SpecificationId - Specification ID
	SpecificationId string `json:"specification_id"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionTemplatesListToolsOutput represents the session templates list tools output type.
type SessionTemplatesListToolsOutput struct {
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
