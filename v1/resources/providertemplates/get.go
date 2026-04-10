package providertemplates

import (
	"encoding/json"
	"time"
)

// ProviderTemplatesGetOutputToolFilters represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderTemplatesGetOutputToolFilters struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderTemplatesGetOutput represents the provider templates get output type.
type ProviderTemplatesGetOutput struct {
	Object               string         `json:"object"`
	Id                   string         `json:"id"`
	Status               string         `json:"status"`
	Name                 string         `json:"name"`
	Description          *string        `json:"description,omitempty"`
	Metadata             map[string]any `json:"metadata"`
	ProviderDeploymentId string         `json:"provider_deployment_id"`
	// ToolFilters - Tool filter configuration
	ToolFilters ProviderTemplatesGetOutputToolFilters `json:"tool_filters"`
	CreatedAt   time.Time                             `json:"created_at"`
	UpdatedAt   time.Time                             `json:"updated_at"`
}

// MapProviderTemplatesGetOutputFromJSON deserializes JSON data into a ProviderTemplatesGetOutput.
func MapProviderTemplatesGetOutputFromJSON(data []byte) (*ProviderTemplatesGetOutput, error) {
	var v ProviderTemplatesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderTemplatesGetOutputToJSON serializes a ProviderTemplatesGetOutput to JSON.
func MapProviderTemplatesGetOutputToJSON(v *ProviderTemplatesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
