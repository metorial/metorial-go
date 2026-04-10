package providertemplates

import (
	"encoding/json"
	"time"
)

// ProviderTemplatesUpdateOutputToolFilters represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderTemplatesUpdateOutputToolFilters struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderTemplatesUpdateOutput represents the provider templates update output type.
type ProviderTemplatesUpdateOutput struct {
	Object               string         `json:"object"`
	Id                   string         `json:"id"`
	Status               string         `json:"status"`
	Name                 string         `json:"name"`
	Description          *string        `json:"description,omitempty"`
	Metadata             map[string]any `json:"metadata"`
	ProviderDeploymentId string         `json:"provider_deployment_id"`
	// ToolFilters - Tool filter configuration
	ToolFilters ProviderTemplatesUpdateOutputToolFilters `json:"tool_filters"`
	CreatedAt   time.Time                                `json:"created_at"`
	UpdatedAt   time.Time                                `json:"updated_at"`
}

// MapProviderTemplatesUpdateOutputFromJSON deserializes JSON data into a ProviderTemplatesUpdateOutput.
func MapProviderTemplatesUpdateOutputFromJSON(data []byte) (*ProviderTemplatesUpdateOutput, error) {
	var v ProviderTemplatesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderTemplatesUpdateOutputToJSON serializes a ProviderTemplatesUpdateOutput to JSON.
func MapProviderTemplatesUpdateOutputToJSON(v *ProviderTemplatesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderTemplatesUpdateBody represents the provider templates update body type.
type ProviderTemplatesUpdateBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	ToolFilters *any            `json:"tool_filters,omitempty"`
}

// MapProviderTemplatesUpdateBodyFromJSON deserializes JSON data into a ProviderTemplatesUpdateBody.
func MapProviderTemplatesUpdateBodyFromJSON(data []byte) (*ProviderTemplatesUpdateBody, error) {
	var v ProviderTemplatesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderTemplatesUpdateBodyToJSON serializes a ProviderTemplatesUpdateBody to JSON.
func MapProviderTemplatesUpdateBodyToJSON(v *ProviderTemplatesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
