package providertemplates

import (
	"encoding/json"
	"time"
)

// ProviderTemplatesCreateOutput represents the provider templates create output type.
type ProviderTemplatesCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the template.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// IntegrationId - Integration the template was created from. `null` when it wasn’t created from one.
	IntegrationId *string `json:"integration_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderTemplatesCreateOutputFromJSON deserializes JSON data into a ProviderTemplatesCreateOutput.
func MapProviderTemplatesCreateOutputFromJSON(data []byte) (*ProviderTemplatesCreateOutput, error) {
	var v ProviderTemplatesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderTemplatesCreateOutputToJSON serializes a ProviderTemplatesCreateOutput to JSON.
func MapProviderTemplatesCreateOutputToJSON(v *ProviderTemplatesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderTemplatesCreateBody represents the provider templates create body type.
type ProviderTemplatesCreateBody struct {
	// Name - Display name for the provider template.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider template.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationId - Integration the template is created from.
	IntegrationId string `json:"integration_id"`
}

// MapProviderTemplatesCreateBodyFromJSON deserializes JSON data into a ProviderTemplatesCreateBody.
func MapProviderTemplatesCreateBodyFromJSON(data []byte) (*ProviderTemplatesCreateBody, error) {
	var v ProviderTemplatesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderTemplatesCreateBodyToJSON serializes a ProviderTemplatesCreateBody to JSON.
func MapProviderTemplatesCreateBodyToJSON(v *ProviderTemplatesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
