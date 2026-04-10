package providertemplates

import (
	"encoding/json"
	"time"
)

// ProviderTemplatesDeleteOutputToolFilters represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderTemplatesDeleteOutputToolFilters struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderTemplatesDeleteOutput represents the provider templates delete output type.
type ProviderTemplatesDeleteOutput struct {
	Object               string         `json:"object"`
	Id                   string         `json:"id"`
	Status               string         `json:"status"`
	Name                 string         `json:"name"`
	Description          *string        `json:"description,omitempty"`
	Metadata             map[string]any `json:"metadata"`
	ProviderDeploymentId string         `json:"provider_deployment_id"`
	// ToolFilters - Tool filter configuration
	ToolFilters ProviderTemplatesDeleteOutputToolFilters `json:"tool_filters"`
	CreatedAt   time.Time                                `json:"created_at"`
	UpdatedAt   time.Time                                `json:"updated_at"`
}

// MapProviderTemplatesDeleteOutputFromJSON deserializes JSON data into a ProviderTemplatesDeleteOutput.
func MapProviderTemplatesDeleteOutputFromJSON(data []byte) (*ProviderTemplatesDeleteOutput, error) {
	var v ProviderTemplatesDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderTemplatesDeleteOutputToJSON serializes a ProviderTemplatesDeleteOutput to JSON.
func MapProviderTemplatesDeleteOutputToJSON(v *ProviderTemplatesDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
