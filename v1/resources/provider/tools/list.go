package tools

import (
	"encoding/json"
	"time"
)

// ProviderToolsListOutputItemsInputSchema represents the provider tools list output items input schema type.
type ProviderToolsListOutputItemsInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// ProviderToolsListOutputItemsOutputSchema represents the provider tools list output items output schema type.
type ProviderToolsListOutputItemsOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// ProviderToolsListOutputItemsTags represents the provider tools list output items tags type.
type ProviderToolsListOutputItemsTags struct {
	// Destructive - Whether the tool is destructive
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool is read-only
	ReadOnly *bool `json:"read_only,omitempty"`
}

// ProviderToolsListOutputItems represents the provider tools list output items type.
type ProviderToolsListOutputItems struct {
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
	Instructions []string                                  `json:"instructions"`
	InputSchema  *ProviderToolsListOutputItemsInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderToolsListOutputItemsOutputSchema `json:"output_schema,omitempty"`
	Tags         *ProviderToolsListOutputItemsTags         `json:"tags,omitempty"`
	// SpecificationId - Specification ID
	SpecificationId string `json:"specification_id"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderToolsListOutputPagination represents the provider tools list output pagination type.
type ProviderToolsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ProviderToolsListOutput represents the provider tools list output type.
type ProviderToolsListOutput struct {
	Items      []ProviderToolsListOutputItems    `json:"items"`
	Pagination ProviderToolsListOutputPagination `json:"pagination"`
}

// MapProviderToolsListOutputFromJSON deserializes JSON data into a ProviderToolsListOutput.
func MapProviderToolsListOutputFromJSON(data []byte) (*ProviderToolsListOutput, error) {
	var v ProviderToolsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderToolsListOutputToJSON serializes a ProviderToolsListOutput to JSON.
func MapProviderToolsListOutputToJSON(v *ProviderToolsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderToolsListQuery represents the provider tools list query type.
type ProviderToolsListQuery struct {
	Limit             *float64 `json:"limit,omitempty"`
	After             *string  `json:"after,omitempty"`
	Before            *string  `json:"before,omitempty"`
	Cursor            *string  `json:"cursor,omitempty"`
	Order             *string  `json:"order,omitempty"`
	ProviderVersionId string   `json:"provider_version_id"`
}

// MapProviderToolsListQueryFromJSON deserializes JSON data into a ProviderToolsListQuery.
func MapProviderToolsListQueryFromJSON(data []byte) (*ProviderToolsListQuery, error) {
	var v ProviderToolsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderToolsListQueryToJSON serializes a ProviderToolsListQuery to JSON.
func MapProviderToolsListQueryToJSON(v *ProviderToolsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
